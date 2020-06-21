package broker

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"

	"github.com/diegoholiveira/jsonlogic"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

var (
	ErrInvalidRule        = errors.New("invalid json logic rule")
	ErrInvalidResultType  = errors.New("result must be convertible to bool type")
	ErrParamsRuleMismatch = errors.New("params do not match with the rules")
	ErrReplyDoesntExist   = errors.New("reply wasn't found in config db")
	ErrNoTxnID            = errors.New("transactionID not found")
)

type Handler struct {
	l         logger.Logger
	cli       Messenger
	routeRepo repos.RouteRepository
	replyRepo repos.ReplyRepository
	eventRepo influx.EventRepository
}

func NewHandler(l logger.Logger, db mongod.Storer, infCli influx.Influxer, cli Messenger) *Handler {
	h := &Handler{
		l,
		cli,
		repos.NewRouteRepo(db),
		repos.NewReplyRepo(db),
		influx.NewEventRepo(infCli.DBName(), infCli),
	}
	return h
}

func (h *Handler) RouteMessage(msg mqtt.Message) {
	h.l.Debug("received msg with ID:", msg.MessageID())
	h.l.Debug("and with Payload:", string(msg.Payload()))

	var (
		err    error
		body   = new(MessageBody)
		events []influx.Event
	)

	if err = body.DoUnmarshal(msg.Payload()); err != nil {
		h.failLog(err)
		h.l.Debug("incorrect message json:", err)
		return
	}
	if body.TransactionID == "" {
		h.l.Debug(ErrNoTxnID)
		return
	}

	events, err = h.findUnreplied(body.TransactionID)
	if err != nil {
		h.failLog(err)
		h.l.Debug("retrieving event error:", err)
		return
	}

	if len(events) != 0 {
		for _, event := range events {
			var topic, replyID string
			topic, replyID, err = h.buildReplyDirection(event, msg.Topic())
			if err != nil || topic == "" {
				h.l.Debug("building reply error:", err)
				continue
			}

			h.cli.AsyncPublish(topic, msg.Qos(), msg.Payload())

			newReplyEventID := h.createEventID()
			h.l.Event().Info(newReplyEventID.String(), body.TransactionID, event.RouteID, replyID, "",
				logger.EVENTREPLY, logger.EVENTPASS)

			go func() {
				_, err = h.eventRepo.StoreEvent(
					influx.StoreEvent{
						newReplyEventID,
						event.RouteID,
						msg.Qos(),
						msg.Topic(),
						topic,
						body.TransactionID,
						replyID,
					},
				)
				if err != nil {
					h.l.Debugf("storing event error: %v. TransactionID: %s", err, body.TransactionID)
				}
			}()
		}
		return
	}

	var (
		routeFound bool
		wg         = new(sync.WaitGroup)
	)

	var routeList []domain.Route
	routeList, err = h.routeRepo.List(domain.RoutesFilter{})
	if err != nil {
		h.l.Debug("routes not found error:", err)
		h.failLog(err)
		return
	}

	var topicName = map[string]interface{}{"name": msg.Topic()}

	for _, rt := range routeList {
		var (
			errLogicTopic   error
			errLogicPayload error
		)
		wg.Add(1)
		go func() {
			defer wg.Done()
			errLogicTopic = doJSONLogic(rt.Src.Topic, topicName)
		}()
		errLogicPayload = doJSONLogic(rt.Src.Payload, body.Payload)
		wg.Wait()

		if errLogicTopic != nil || errLogicPayload != nil {
			h.l.Debug("json logic topic:", errLogicTopic)
			h.l.Debug(rt.Src.Topic, " | ", topicName)
			h.l.Debug("json logic payload:", errLogicPayload)
			h.l.Event().Error("", body.TransactionID, "Logic Error in Topic or Payload",
				logger.EVENTREQUEST, logger.EVENTFAILED)
			continue
		}
		routeFound = true

		for _, dst := range rt.Dst {
			newEventID := h.createEventID()
			h.l.Event().Debug(rt.Src.Topic, dst.Topic, string(msg.Payload()))

			_, err := h.eventRepo.StoreEvent(
				influx.StoreEvent{
					newEventID,
					rt.RouteID.String(),
					dst.Qos,
					msg.Topic(),
					dst.Topic,
					body.TransactionID,
					"",
				},
			)
			if err != nil {
				h.l.Debugf("storing event error: %v. TransactionID: %s", err, body.TransactionID)
			}

			h.cli.AsyncPublish(dst.Topic, dst.Qos, msg.Payload())

			h.l.Event().Info(newEventID.String(), body.TransactionID, rt.RouteID.String(), "",
				"", logger.EVENTREQUEST, logger.EVENTPASS)
		}
	}
	if !routeFound {
		h.failLog(domain.ErrRouteNotFound)
	}
}

func (h *Handler) createEventID() (newEventID strfmt.UUID) {
	err := newEventID.Scan(uuid.Must(uuid.NewV1()).String())
	if err != nil {
		h.l.Event().Error(newEventID.String(), "", err.Error(), "", logger.EVENTFAILED)
	}
	return
}

func (h *Handler) failLog(err error) {
	h.l.Event().Error("", "", err.Error(), logger.EVENTREQUEST, logger.EVENTFAILED)
}

func (h *Handler) buildReplyDirection(event influx.Event, msgTopic string) (_, replyID string, err error) {
	var route *domain.Route
	route, err = h.routeRepo.Get(strfmt.UUID(event.RouteID))
	if err != nil {
		return "", "", err
	}
	if route.ReplyID == nil {
		return "", "", ErrReplyDoesntExist
	}

	var reply *domain.Reply
	reply, err = h.replyRepo.Get(*route.ReplyID)
	if err != nil {
		return "", "", err
	}
	h.l.Debugln("found reply id: ", reply.ReplyID, "with regex:", reply.Regex, "and replace:", reply.Replace)

	xp, err := regexp.Compile(reply.Regex)
	if err != nil {
		return "", "", err
	}
	replaced := xp.ReplaceAllString(msgTopic, reply.Replace)
	if replaced != event.DestinationTopic {
		h.l.Debugln("regexp replacement error. replaced string: ", replaced,
			". destination topic: ", event.DestinationTopic)
		return "", "", nil
	}

	return xp.ReplaceAllString(event.SourceTopic, reply.Replace), reply.ReplyID.String(), nil
}

func (h *Handler) findUnreplied(txID string) (events []influx.Event, err error) {
	events, err = h.eventRepo.GetEvents(
		influx.GetEvents{
			RouteID:          nil,
			Qos:              nil,
			SourceTopic:      nil,
			DestinationTopic: nil,
			TransactionID:    &txID,
			ReplyID:          nil,
			Limit:            0,
			Offset:           0,
		},
	)
	if err != nil && err != domain.ErrEventNotFound {
		return nil, err
	}
	if len(events) == 0 {
		return nil, nil
	}

	var nonRepliedEvents = make([]influx.Event, 0)
	for _, ev := range events {
		h.l.Debugf("tx # %s found event: %#v\n", txID, ev)
		if ev.ReplyID == "" {
			nonRepliedEvents = append(nonRepliedEvents, ev)
		}
	}

	if len(nonRepliedEvents) == 0 {
		return nil, nil

	}
	return nonRepliedEvents, nil
}

func doJSONLogic(rule interface{}, data interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	if rule == nil || reflect.DeepEqual(rule, make(map[string]interface{})) {
		return nil
	}

	if !jsonlogic.IsValid(rule) {
		return errors.Wrap(ErrInvalidRule, fmt.Sprintf("%#v", rule))
	}

	var (
		result interface{}
		ok     bool
		match  bool
	)

	err = jsonlogic.Apply(rule, data, &result)
	if err != nil {
		return err
	}

	// result must have only bool type
	match, ok = result.(bool)
	if !ok {
		return ErrInvalidResultType
	}
	if !match {
		return ErrParamsRuleMismatch
	}
	return nil
}
