package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/suite"

	//benchtool "repo.nefrosovet.ru/libs/mqtt-benchmark/tool"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

const targetMPS = 100000

type Dsts []string

var (
	testInfo = make(chan string)

	messagesWg = sync.WaitGroup{}

	sourceTopicTest   interface{}
	sourceTopicDiff   interface{}
	sourcePayloadTest interface{}
	sourcePayloadDiff interface{}
	rp                *domain.Reply
	routeIDs          = make([]strfmt.UUID, 0)

	receivedMessages, sentMessages = 0, 4

	txnID1, txnID2, txnID3 strfmt.UUID

	_ = json.Unmarshal([]byte(`{"and": [{"==": [{"var": "name"}, "testServices/test/OUT"]}]}`), &sourceTopicTest)
	_ = json.Unmarshal([]byte(`{"and": [{"==": [{"var": "name"}, "testServices/diff/OUT"]}]}`), &sourceTopicDiff)

	_ = json.Unmarshal([]byte(`{ "and" : [
	  {"<" : [ { "var" : "temp" }, 110 ]},
	  {"==" : [ { "var" : "pie.filling" }, "apple" ] }
	] }`), &sourcePayloadTest)
	_ = json.Unmarshal([]byte(`{ "and" : [
	  {">" : [ { "var" : "temp" }, 120 ]},
	  {"==" : [ { "var" : "pie.filling" }, "orange" ] }
	] }`), &sourcePayloadDiff)

	dstTest1 = Dsts{"testServices/test1/IN"}
	dstTest2 = Dsts{"testServices/test2-1/IN", "testServices/test2-2/IN", "testServices/test2-3/IN"}
	dstDiff  = Dsts{"testServices/diff/IN"}
	dstReply = Dsts{"testServices/reply/IN"}

	_ = txnID1.Scan(uuid.Must(uuid.NewV4()).String())
	_ = txnID2.Scan(uuid.Must(uuid.NewV4()).String())
	_ = txnID3.Scan(uuid.Must(uuid.NewV4()).String())
)

type HandlerTestSuite struct {
	suite.Suite

	db      mongod.Storer
	cli     Messenger
	h       *Handler
	replyID strfmt.UUID
	infCli  influx.Influxer
}

func (s *HandlerTestSuite) SetupSuite() {
	subClientId := os.Getenv("DATAROUTER_MQ_SUBCLIENTID")
	pubClientId := os.Getenv("DATAROUTER_MQ_PUBCLIENTID")

	mqLogin := os.Getenv("DATAROUTER_MQ_LOGIN")
	mqPass := os.Getenv("DATAROUTER_MQ_PASSWORD")
	mqHost := os.Getenv("DATAROUTER_MQ_HOST")
	mqStringPort := os.Getenv("DATAROUTER_MQ_PORT")
	logLevel := "info"

	mqPort, err := strconv.Atoi(mqStringPort)
	if err != nil {
		s.Suite.T().Fatal(err)
	}

	infHost := os.Getenv("DATAROUTER_EVENTDB_HOST")
	infStringPort := os.Getenv("DATAROUTER_EVENTDB_PORT")
	infLogin := os.Getenv("DATAROUTER_EVENTDB_LOGIN")
	infPassword := os.Getenv("DATAROUTER_EVENTDB_PASSWORD")
	infDatabase := os.Getenv("DATAROUTER_EVENTDB_DATABASE")
	infRetention := "1h"

	infPort, err := strconv.Atoi(infStringPort)
	if err != nil {
		s.Suite.T().Fatal(err)
	}

	l := logger.New(10, 1000, logLevel, "", "")

	db := mongoConnect(l)

	s.db = db

	infCli, err := influx.ConnectHTTP(
		infHost,
		infPort,
		infLogin,
		infPassword,
		infDatabase,
		infRetention,
	)
	if err != nil {
		l.Core().Fatal(logger.COREEVENTDB, infHost, infStringPort, err.Error(), logger.COREFAILED)
	}

	s.infCli = infCli

	l.Core().Info(logger.COREEVENTDB, infHost, infStringPort, "", logger.CORECONNECTED)

	mqtt.ERROR = l

	br, err := NewMQTTClient(
		l,
		subClientId,
		pubClientId,
		mqHost,
		mqPort,
		mqLogin,
		mqPass,
	)
	if err != nil {
		l.Core().Fatal(logger.COREMQ, mqHost, mqStringPort, err.Error(), logger.COREFAILED)
	}

	s.cli = br

	h := NewHandler(l, db, infCli, br)
	s.h = h

	s.SetReply()

	s.SetRoute(sourceTopicTest, sourcePayloadTest, dstTest1)
	s.SetRoute(sourceTopicDiff, sourcePayloadDiff, dstDiff)
	s.SetRoute(sourceTopicTest, sourcePayloadDiff, dstTest2)
	s.SetRoute(sourceTopicTest, sourcePayloadTest, dstReply)
}

func (s *HandlerTestSuite) SetReply() {
	replyList, err := s.h.replyRepo.List(domain.RepliesFilter{})
	if err != nil {
		s.h.l.API().Reply().Debug(err)
		if err != domain.ErrReplyNotFound {
			return
		}
	}

	if len(replyList) == 0 {
		description := "Переправлять ответы на /IN"
		rp, err = s.h.replyRepo.Set(domain.Reply{
			Description: &description,
			Regex:       "testServices/(.*)/OUT",
			Replace:     "testServices/$1/IN",
		})
		if err != nil {
			s.h.l.API().Reply().Debug(err)
			return
		}
	} else {
		rp = &replyList[0]
	}

	s.replyID = rp.ReplyID
}

func (s *HandlerTestSuite) SetRoute(srcTopic interface{}, srcPayload interface{}, dstTopics Dsts) {
	dsts := make([]domain.Destinations, 0)
	for _, dstTopic := range dstTopics {
		dst := domain.Destinations{
			Qos:   0,
			Topic: dstTopic,
		}

		dsts = append(dsts, dst)
	}

	route, err := s.h.routeRepo.Set(domain.Route{
		ReplyID: &s.replyID,
		Dst:     dsts,
		Src: domain.Source{
			Payload: srcPayload,
			Topic:   srcTopic,
		},
		Created: time.Now(),
		Updated: nil,
	})
	if err != nil {
		s.h.l.API().Route().Debug(err)
		return
	}
	routeIDs = append(routeIDs, route.RouteID)
}

func (s *HandlerTestSuite) Test1Subscribe() {
	var err error

	// Sum up all /IN/ topics
	inTopics := append(dstTest1, dstTest2...)
	inTopics = append(inTopics, dstDiff...)
	inTopics = append(inTopics, dstReply...)

	// Subscribe on OUT topics
	err = s.h.cli.Subscribe([]string{"testServices/test/OUT", "testServices/diff/OUT"}, func(msg mqtt.Message) {
		defer messagesWg.Done()

		receivedMessages++

		s.h.l.Debugln("Message | Received topic:", msg.Topic())
		s.h.l.Debugln("Message | Received payload:", string(msg.Payload()))

		s.Contains([][]byte{
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 100, "pie" : { "filling" : "apple" } }}`,
				txnID1,
			)),
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
				txnID2,
			)),
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
				txnID3,
			)),
		}, msg.Payload())

		s.Contains([]string{
			"testServices/test/OUT",
			"testServices/diff/OUT",
		}, msg.Topic())

		s.h.RouteMessage(msg)
	})
	if err != nil {
		s.h.l.Fatal(err.Error())
	}

	err = s.h.cli.Subscribe(inTopics, func(msg mqtt.Message) {
		s.h.l.Debugln("Reply Message | Received topic:", msg.Topic())
		s.h.l.Debugln("Reply Message | Received payload:", string(msg.Payload()))

		s.Contains(inTopics, msg.Topic())

		s.Contains([][]byte{
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 100, "pie" : { "filling" : "apple" } }}`,
				txnID1,
			)),
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
				txnID2,
			)),
			[]byte(fmt.Sprintf(
				`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
				txnID3,
			)),
		}, msg.Payload())
	})
	if err != nil {
		s.h.l.Fatal(err.Error())
	}
	s.h.l.Core().Debug("Subscribe to all topics")
}

func (s *HandlerTestSuite) Test2Publish() {
	messagesWg.Add(sentMessages)

	s.h.cli.Publish("testServices/test/OUT", 0, []byte(fmt.Sprintf(
		`{"transactionID": "%s", "payload":{ "temp" : 100, "pie" : { "filling" : "apple" } }}`,
		txnID1,
	)))

	s.h.cli.Publish("testServices/diff/OUT", 0, []byte(fmt.Sprintf(
		`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
		txnID2,
	)))

	s.h.cli.Publish("testServices/test/OUT", 0, []byte(fmt.Sprintf(
		`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
		txnID3,
	)))

	s.h.cli.Publish("testServices/diff/OUT", 0, []byte(fmt.Sprintf(
		`{"transactionID": "%s", "payload":{ "temp" : 125, "pie" : { "filling" : "orange" } }}`,
		txnID2,
	)))

	s.h.l.Debug("Test2Publish passed")

	messagesWg.Wait()
}

//func (s *HandlerTestSuite) TestBenchmark() {
//	execOpts := benchtool.ExecOptions{}
//	execOpts.Broker = os.Getenv("DATAROUTER_MQ_HOST") + ":" + os.Getenv("DATAROUTER_MQ_PORT")
//	execOpts.Qos = byte(2)
//	execOpts.Topic = "testServices/test/OUT"
//	execOpts.Username = os.Getenv("DATAROUTER_MQ_LOGIN")
//	execOpts.Password = os.Getenv("DATAROUTER_MQ_PASSWORD")
//	execOpts.ClientNum = 100
//	execOpts.Count = 10
//	execOpts.UseDefaultHandler = true
//	execOpts.PreTime = 0
//	execOpts.IntervalTime = 0
//	execOpts.TargetMPS = targetMPS
//	err := benchtool.Execute(
//		benchtool.PublishAllClient,
//		execOpts,
//		`{"transactionID": "bench", "payload":{ "temp" : 100, "pie" : { "filling" : "apple" } }}`,
//	)
//	s.NoError(err, targetMPS)
//}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func mongoConnect(lg logger.Logger) (db mongod.Storer) {
	var err error

	dbHost := os.Getenv("DATAROUTER_CONFIGDB_HOST")
	dbStringPort := os.Getenv("DATAROUTER_CONFIGDB_PORT")
	dbDatabase := os.Getenv("DATAROUTER_CONFIGDB_DATABASE")

	dbPort, err := strconv.Atoi(dbStringPort)
	if err != nil {
		lg.Fatal(err)
	}

	db, err = mongod.NewCli(
		dbHost,
		dbPort,
		"",
		"",
		dbDatabase,
	)
	if err != nil {
		lg.Core().Fatal(logger.CORECONFIGDB, dbHost, dbStringPort, err.Error(), logger.COREFAILED)
	}

	err = db.Connect(context.Background())
	if err != nil {
		lg.Core().Fatal(logger.CORECONFIGDB, dbHost, dbStringPort, err.Error(), logger.COREFAILED)
	}
	lg.Core().Info(logger.CORECONFIGDB, dbHost, dbStringPort, "", logger.CORECONNECTED)

	return
}
