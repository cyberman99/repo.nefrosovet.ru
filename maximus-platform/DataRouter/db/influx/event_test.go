package influx

import (
	"os"
	"strconv"
	"testing"

	guid "github.com/gofrs/uuid"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
)

type EventSuite struct {
	suite.Suite

	r EventRepository
}

var (
	host     = os.Getenv("DATAROUTER_EVENTDB_HOST")
	strPort  = os.Getenv("DATAROUTER_EVENTDB_PORT")
	username = os.Getenv("DATAROUTER_EVENTDB_LOGIN")
	password = os.Getenv("DATAROUTER_EVENTDB_PASSWORD")
	database = os.Getenv("DATAROUTER_EVENTDB_DATABASE")
)

func (s *EventSuite) SetupSuite() {
	port, err := strconv.Atoi(strPort)
	s.Require().NoError(err)

	client, err := ConnectHTTP(
		host,
		port,
		username,
		password,
		database,
		"1h",
	)
	s.Require().NoError(err)

	s.r = NewEventRepo(database, client)
}

func (s *EventSuite) SetupTest() {
	err := s.r.Drop()
	s.Require().NoError(err)
}

func (s *EventSuite) TestStoreEvent() {
	id, err := s.r.StoreEvent(StoreEvent{
		RouteID:          "ROUTE_ID",
		Qos:              5,
		SourceTopic:      "SOURCE_TOPIC",
		DestinationTopic: "DESTINATION_TOPIC",
		TransactionID:    "TRANSACTION_ID",
		ReplyID:          "REPLY_ID",
	})
	s.NoError(err)

	transactionID := "TRANSACTION_ID"
	foundEvent, err := s.r.GetEvent(GetEvent{
		TransactionID: &transactionID,
	})
	s.NoError(err)

	s.Equal(foundEvent.ID, id)
	s.Equal("ROUTE_ID", foundEvent.RouteID)
	s.Equal(uint8(5), foundEvent.Qos)
	s.Equal("SOURCE_TOPIC", foundEvent.SourceTopic)
	s.Equal("DESTINATION_TOPIC", foundEvent.DestinationTopic)
	s.Equal("TRANSACTION_ID", foundEvent.TransactionID)
	s.Equal("REPLY_ID", foundEvent.ReplyID)
}

func (s *EventSuite) TestGetEvents() {
	routeID := "ROUTE_ID-2"

	events, err := s.r.GetEvents(GetEvents{
		RouteID: &routeID,
	})
	s.Error(err)
	s.Equal(domain.ErrEventNotFound, err)

	for i := 0; i < 10; i++ {
		_, err := s.r.StoreEvent(StoreEvent{
			RouteID:          routeID,
			Qos:              6,
			SourceTopic:      "SOURCE_TOPIC-2",
			DestinationTopic: "SOURCE_TOPIC-2",
			TransactionID:    guid.Must(guid.NewV1()).String(),
			ReplyID:          "REPLY_ID-2",
		})
		s.NoError(err)
	}

	events, err = s.r.GetEvents(GetEvents{
		RouteID: &routeID,
	})
	s.NoError(err)

	s.Len(events, 10)
}

func (s *EventSuite) TestGetEvent() {
	transactionID := "unknown"
	_, err := s.r.GetEvent(GetEvent{
		TransactionID: &transactionID,
	})
	s.Error(err)
	s.Equal(domain.ErrEventNotFound, err)

	id, err := s.r.StoreEvent(StoreEvent{
		RouteID:          "ROUTE_ID-3",
		Qos:              0,
		SourceTopic:      "SOURCE_TOPIC-3",
		DestinationTopic: "DESTINATION_TOPIC-3",
		TransactionID:    "TRANSACTION_ID-3",
		ReplyID:          "REPLY_ID-3",
	})
	s.NoError(err)

	transactionID = "TRANSACTION_ID-3"
	event, err := s.r.GetEvent(GetEvent{
		TransactionID: &transactionID,
	})
	s.NoError(err)
	s.Equal(id, event.ID)
}

func TestEvent(t *testing.T) {
	suite.Run(t, new(EventSuite))
}
