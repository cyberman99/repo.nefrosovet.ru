package influx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/influxdata/influxdb/client/v2"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
)

const (
	// Table name
	table = "events"

	// Fields
	fieldID = "id"

	fieldRouteID          = "routeID"
	fieldQos              = "qos"
	fieldSourceTopic      = "sourceTopic"
	fieldDestinationTopic = "destinationTopic"
	fieldTransactionID    = "transactionID"
	fieldReply            = "replyID"

	fieldCreatedAt = "dateTime"
)

type Event struct {
	ID strfmt.UUID `json:"id"`

	RouteID          string `json:"routeID"`
	Qos              byte   `json:"qos"`
	SourceTopic      string `json:"sourceTopic"`
	DestinationTopic string `json:"destinationTopic"`
	TransactionID    string `json:"transactionID"`
	ReplyID          string `json:"replyID"`

	DateTime string `json:"dateTime"`
}

type StoreEvent struct {
	ID               strfmt.UUID `json:"id"`
	RouteID          string      `json:"routeID"`
	Qos              byte        `json:"qos"`
	SourceTopic      string      `json:"sourceTopic"`
	DestinationTopic string      `json:"destinationTopic"`
	TransactionID    string      `json:"transactionID"`
	ReplyID          string      `json:"replyID"`
}

type GetEvents struct {
	RouteID          *string `json:"routeID"`
	Qos              *byte   `json:"qos"`
	SourceTopic      *string `json:"sourceTopic"`
	DestinationTopic *string `json:"destinationTopic"`
	TransactionID    *string `json:"transactionID"`
	ReplyID          *string `json:"replyID"`

	Limit  int64
	Offset int64
}

type GetEvent struct {
	ID            *strfmt.UUID
	TransactionID *string
}

type EventRepository interface {
	StoreEvent(in StoreEvent) (strfmt.UUID, error)
	GetEvents(in GetEvents) ([]Event, error)
	GetEvent(in GetEvent) (Event, error)

	Drop() error
}

type eventRepo struct {
	database string

	c client.Client
}

func NewEventRepo(database string, influxClient client.Client) EventRepository {
	return &eventRepo{
		database: database,
		c:        influxClient,
	}
}

func (s *eventRepo) StoreEvent(in StoreEvent) (strfmt.UUID, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  s.database,
		Precision: "s",
	})
	if err != nil {
		return "", err
	}

	tags := map[string]string{
		fieldDestinationTopic: in.DestinationTopic,
		fieldTransactionID:    in.TransactionID,
	}

	fields := map[string]interface{}{
		fieldID: in.ID.String(),

		fieldRouteID:          in.RouteID,
		fieldQos:              in.Qos,
		fieldSourceTopic:      in.SourceTopic,
		fieldDestinationTopic: in.DestinationTopic,
		fieldTransactionID:    in.TransactionID,
		fieldReply:            in.ReplyID,
	}

	point, err := client.NewPoint(table, tags, fields, time.Now())
	if err != nil {
		return "", err
	}

	bp.AddPoint(point)

	err = s.c.Write(bp)
	if err != nil {
		return "", err
	}

	return in.ID, nil
}

func (s *eventRepo) GetEvents(in GetEvents) ([]Event, error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE true", table)

	if in.RouteID != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldRouteID, *in.RouteID)
	}
	if in.Qos != nil {
		q = fmt.Sprintf("%s AND %s = '%d'", q, fieldQos, *in.Qos)
	}
	if in.SourceTopic != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldSourceTopic, *in.SourceTopic)
	}
	if in.DestinationTopic != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldDestinationTopic, *in.DestinationTopic)
	}
	if in.TransactionID != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldTransactionID, *in.TransactionID)
	}
	if in.ReplyID != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldReply, *in.ReplyID)
	}

	if in.Limit != 0 {
		q = fmt.Sprintf("%s LIMIT %d", q, in.Limit)
	}
	if in.Offset != 0 {
		q = fmt.Sprintf("%s OFFSET %d", q, in.Offset)
	}

	resp, err := s.c.Query(client.Query{
		Command:  q,
		Database: s.database,
	})
	if err != nil {
		return nil, err
	}
	if err = resp.Error(); err != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, domain.ErrEventNotFound
	}

	return parseEvents(resp.Results)

}

func (s *eventRepo) GetEvent(in GetEvent) (Event, error) {
	q := fmt.Sprintf(`SELECT * FROM %s WHERE true`, table)

	if in.ID != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldID, *in.ID)
	}
	if in.TransactionID != nil {
		q = fmt.Sprintf("%s AND %s = '%s'", q, fieldTransactionID, *in.TransactionID)
	}

	resp, err := s.c.Query(client.Query{
		Command:  q,
		Database: s.database,
	})
	if err != nil {
		return Event{}, err
	}
	if resp.Error() != nil {
		return Event{}, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return Event{}, domain.ErrEventNotFound
	}

	events, err := parseEvents(resp.Results)
	if err != nil {
		return Event{}, err
	}

	return events[0], err
}

func (s *eventRepo) Drop() error {
	q := fmt.Sprintf("DROP SERIES FROM %s", table)

	resp, err := s.c.Query(client.Query{
		Command:  q,
		Database: s.database,
	})
	if err != nil {
		return err
	}
	if err := resp.Error(); err != nil {
		return err
	}

	return nil
}

func parseEvents(results []client.Result) ([]Event, error) {
	var events []Event
	columns := make(map[string]int)

	for _, result := range results {
		for k, column := range result.Series[0].Columns {
			columns[column] = k
		}

		for _, row := range result.Series[0].Values {
			qos, err := row[columns[fieldQos]].(json.Number).Int64()
			if err != nil {
				return nil, err
			}

			events = append(
				events,
				Event{
					ID: strfmt.UUID(row[columns[fieldID]].(string)),

					RouteID:          row[columns[fieldRouteID]].(string),
					Qos:              byte(qos),
					SourceTopic:      row[columns[fieldSourceTopic]].(string),
					DestinationTopic: row[columns[fieldDestinationTopic]].(string),
					TransactionID:    row[columns[fieldTransactionID]].(string),
					ReplyID:          row[columns[fieldReply]].(string),

					DateTime: row[columns[fieldCreatedAt]].(string),
				},
			)
		}
	}

	return events, nil
}
