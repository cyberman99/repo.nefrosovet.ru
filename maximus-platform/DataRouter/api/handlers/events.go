package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/events"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

type EventsController struct {
	eventsRepo influx.EventRepository
	l          logger.Logger
}

func NewEventsController(
	eRepo influx.EventRepository,
	l logger.Logger) *EventsController {
	return &EventsController{
		eventsRepo: eRepo,
		l:          l,
	}
}

func (c *EventsController) List(params events.EventCollectionParams) middleware.Responder {
	InternalServerError := func() middleware.Responder {
		payload := new(events.EventCollectionInternalServerErrorBody)
		payload.Version = &Version
		payload.Message = *InternalServerErrorMessage
		payload.Data = []interface{}{}
		return events.NewEventCollectionInternalServerError().WithPayload(payload)
	}

	ViewOK := func(objs []influx.Event) middleware.Responder {
		payload := new(events.EventCollectionOKBody)
		payload.Message = PayloadSuccessMessage
		payload.Version = &Version
		payload.Data = []*events.DataItems0{}
		for _, obj := range objs {
			payload.Data = append(payload.Data, eventObjToItem(obj))
		}

		return events.NewEventCollectionOK().WithPayload(payload)
	}

	var in = influx.GetEvents{
		RouteID:          (*string)(params.RouteID),
		TransactionID:    (*string)(params.TransactionID),
		SourceTopic:      params.SrcTopic,
		DestinationTopic: params.DstTopic,
	}
	if params.ReplyID != nil {
		str := params.ReplyID.String()
		in.ReplyID = &str
	}

	if params.Offset != nil {
		in.Offset = *params.Offset
	}
	if params.Limit != nil {
		in.Limit = *params.Limit
	}

	objs, err := c.eventsRepo.GetEvents(in)
	if err == domain.ErrEventNotFound {
		c.l.Debug(err)
		payload := new(events.EventCollectionNotFoundBody)
		payload.Version = &Version
		payload.Message = *NotFoundMessage
		payload.Data = []interface{}{}
		return events.NewEventCollectionNotFound().WithPayload(payload)
	}
	if err != nil {
		c.l.Debug(err, " internal")
		return InternalServerError()
	}

	return ViewOK(objs)
}

func (c *EventsController) Get(params events.EventViewParams) middleware.Responder {
	InternalServerError := func() middleware.Responder {
		payload := new(events.EventViewInternalServerErrorBody)
		payload.Version = &Version
		payload.Message = *InternalServerErrorMessage

		return events.NewEventViewInternalServerError().WithPayload(payload)
	}

	ViewOK := func(obj influx.Event) middleware.Responder {
		payload := new(events.EventViewOKBody)
		payload.Message = PayloadSuccessMessage
		payload.Version = &Version
		payload.Data = []*events.DataItems0{eventObjToItem(obj)}

		return events.NewEventViewOK().WithPayload(payload)
	}

	NotFound := func() middleware.Responder {
		payload := new(events.EventViewNotFoundBody)
		payload.Message = *NotFoundMessage

		return events.NewEventViewNotFound().WithPayload(payload)
	}

	obj, err := c.eventsRepo.GetEvent(influx.GetEvent{ID: &params.EventID})
	if err != nil && err == domain.ErrEventNotFound {
		c.l.Debug(err)
		return NotFound()
	}
	if err != nil {
		c.l.Debug(err)
		return InternalServerError()
	}

	return ViewOK(obj)
}

func eventObjToItem(model influx.Event) *events.DataItems0 {
	date, _ := strfmt.ParseDateTime(model.DateTime)
	var id *strfmt.UUID
	if model.ReplyID != "" {
		formatted := strfmt.UUID(model.ReplyID)
		id = &formatted
	}

	item := events.DataItems0{
		ID: model.ID.String(),
		EventObject: models.EventObject{
			ReplyID:       id,
			Date:          date,
			DstTopic:      model.DestinationTopic,
			RouteID:       (strfmt.UUID)(model.RouteID),
			SrcTopic:      model.SourceTopic,
			TransactionID: (strfmt.UUID)(model.TransactionID),
		},
	}

	return &item
}
