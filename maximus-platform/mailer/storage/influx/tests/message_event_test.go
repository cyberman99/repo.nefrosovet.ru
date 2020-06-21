package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	dbInflux "repo.nefrosovet.ru/maximus-platform/mailer/db/influx"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/influx"
)

var (
	MessageEventStorage storage.MessageEventStorage
)

func TestMain(m *testing.M) {
	host := flag.String("c.eventDB.host", "127.0.0.1", "EventDB host")
	port := flag.Int("c.eventDB.port", 8086, "EventDB port")
	username := flag.String("c.eventDB.login", "", "EventDB login")
	password := flag.String("c.eventDB.password", "", "EventDB password")
	database := flag.String("c.eventDB.database", "mailer", "EventDB database name")
	retention := flag.String("c.eventDB.retention", "", "EventDB retention policy")
	flag.Parse()

	client, err := dbInflux.Connect(*host, *port, *username, *password, *database, *retention)
	if err != nil {
		panic(err)
	}

	err = influx.Ensure(client)
	if err != nil {
		panic(err)
	}

	MessageEventStorage = influx.GetMessageEventStorage(client)

	os.Exit(m.Run())
}

type MessageEventSuite struct {
	s storage.MessageEventStorage

	suite.Suite
}

func (s *MessageEventSuite) SetupSuite() {
	s.s = MessageEventStorage
}

func (s *MessageEventSuite) Test1StoreMessageEvent() {
	_, err := s.s.StoreMessageEvent(storage.StoreMessageEvent{
		MessageEvent: storage.MessageEvent{
			ID:                "id",
			Created:           "created",
			Updated:           "updated",
			Status:            "SENT",
			Errors:            "errors",
			ChannelID:         "channelID",
			ChannelType:       "channelType",
			AccessToken:       "accessToken",
			Destination:       "destination",
			Data:              "data",
			MetaEmailSubject:  "metaEmailSubject",
			MetaEmailFrom:     "metaEmailFrom",
			MetaSlackDestType: "metaSlackDestType",
		},
	})
	s.NoError(err)
}

func (s *MessageEventSuite) Test2GetMessageEvents() {
	channelID := "channelID"
	status := "SENT"
	destination := "destination"

	messageEvents, err := s.s.GetMessageEvents(storage.GetMessageEvents{
		AccessToken: "accessToken",
		ChannelID:   &channelID,
		Status:      &status,
		Destination: &destination,
	})
	s.NoError(err)

	s.Len(messageEvents, 1)
}

func (s *MessageEventSuite) Test3GetMessageEventCount() {
	count, err := s.s.GetMessageEventCount(storage.GetMessageEventCount{
		ChannelID: "channelID",
	})
	s.NoError(err)

	s.Equal(int64(1), count)
}

func (s *MessageEventSuite) Test4DeleteMessageEvents() {
	err := s.s.DeleteMessageEvents(storage.DeleteMessageEvents{
		ChannelID: "channelID",
	})
	s.NoError(err)
}

func TestMessageEvents(t *testing.T) {
	suite.Run(t, new(MessageEventSuite))
}
