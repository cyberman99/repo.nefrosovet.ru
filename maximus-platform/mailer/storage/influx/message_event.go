package influx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/influxdata/influxdb/client/v2"

	"repo.nefrosovet.ru/maximus-platform/mailer/db/influx"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

const (
	messageEventsTable = "message_events"
)

type MessageEventStorage struct {
	client *influx.Client

	storage.MessageEventStorage
}

func (s *MessageEventStorage) StoreMessageEvent(in storage.StoreMessageEvent) (storage.MessageEvent, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  s.client.Database,
		Precision: "s",
	})
	if err != nil {
		return storage.MessageEvent{}, err
	}

	event := in.MessageEvent

	tags := map[string]string{
		"id":          event.ID,
		"channelID":   event.ChannelID,
		"channelType": event.ChannelType,
		"destination": event.Destination,
		"status":      event.Status,
		"accessToken": event.AccessToken,
	}

	fields := map[string]interface{}{
		"data":    event.Data,
		"created": event.Created,
		"updated": event.Updated,
		"errors":  event.Errors,

		"metaEmailSubject":  event.MetaEmailSubject,
		"metaEmailFrom":     event.MetaEmailFrom,
		"metaSlackDestType": event.MetaSlackDestType,
	}

	point, err := client.NewPoint(
		messageEventsTable,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		logrus.WithError(err).Error("saving event error")
		return storage.MessageEvent{}, err
	}

	bp.AddPoint(point)

	err = s.client.Write(bp)
	if err != nil {
		logrus.WithError(err).Error("saving event error")
		return storage.MessageEvent{}, err
	}

	return in.MessageEvent, nil
}

func (s *MessageEventStorage) GetMessageEvent(in storage.GetMessageEvent) (storage.MessageEvent, error) {
	q := client.Query{
		Command:  fmt.Sprintf(`select * from %s where id = '%s'`, messageEventsTable, in.ID),
		Database: s.client.Database,
	}

	resp, err := s.client.Query(q)
	if err != nil {
		return storage.MessageEvent{}, err
	}

	if resp.Error() != nil {
		return storage.MessageEvent{}, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return storage.MessageEvent{}, storage.ErrMessageEventsNotFound
	}

	return constructMessageEventsByDBResult(resp.Results)[0], err
}

func (s *MessageEventStorage) GetMessageEvents(in storage.GetMessageEvents) ([]storage.MessageEvent, error) {
	queryString := fmt.Sprintf(`SELECT * FROM "%s" WHERE accessToken = '%s'`, messageEventsTable, in.AccessToken)

	if in.ChannelID != nil {
		queryString = fmt.Sprintf("%s AND channelID = '%s'", queryString, *in.ChannelID)
	}
	if in.Status != nil {
		queryString = fmt.Sprintf("%s AND status = '%s'", queryString, *in.Status)
	}
	if in.Destination != nil {
		queryString = fmt.Sprintf("%s AND destination = '%s'", queryString, *in.Destination)
	}

	if in.Limit != nil {
		queryString = fmt.Sprintf("%s LIMIT %d", queryString, *in.Limit)
	}
	if in.Offset != nil {
		queryString = fmt.Sprintf("%s OFFSET %d", queryString, *in.Offset)
	}

	q := client.Query{
		Command:  queryString,
		Database: s.client.Database,
	}

	resp, err := s.client.Query(q)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, storage.ErrMessageEventsNotFound
	}

	return constructMessageEventsByDBResult(resp.Results), nil
}

func (s *MessageEventStorage) GetMessageEventCount(in storage.GetMessageEventCount) (int64, error) {
	var result int64

	q := client.Query{
		Command: fmt.Sprintf(`
			SELECT COUNT("created")
			FROM "%s"
			WHERE time > now() - 30d
				AND "channelID" = '%s'
				AND "status" = 'SENT'
		`, messageEventsTable, in.ChannelID),
		Database: s.client.Database,
	}

	resp, err := s.client.Query(q)
	if err != nil {
		return result, err
	}

	if resp.Error() != nil {
		return result, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return result, nil
	}

	if resp.Results[0].Series[0].Values[0][1] != nil {
		result, _ = resp.Results[0].Series[0].Values[0][1].(json.Number).Int64()
	}

	return result, err
}

func (s *MessageEventStorage) DeleteMessageEvents(in storage.DeleteMessageEvents) error {
	q := client.Query{
		Command:  fmt.Sprintf(`delete from "%s" where channelID = '%s'`, messageEventsTable, in.ChannelID),
		Database: s.client.Database,
	}

	_, err := s.client.Query(q)

	return err
}

func constructMessageEventsByDBResult(results []client.Result) []storage.MessageEvent {
	var events []storage.MessageEvent
	for _, result := range results {
		columns := make(map[string]int)

		for num, column := range result.Series[0].Columns {
			columns[column] = num
		}

		for _, row := range result.Series[0].Values {
			messageEvent := storage.MessageEvent{
				ID:          row[columns["id"]].(string),
				Status:      row[columns["status"]].(string),
				ChannelID:   row[columns["channelID"]].(string),
				ChannelType: row[columns["channelType"]].(string),
				AccessToken: row[columns["accessToken"]].(string),
				Destination: row[columns["destination"]].(string),
				Data:        row[columns["data"]].(string),
				Created:     row[columns["created"]].(string),
				Updated:     row[columns["updated"]].(string),
				Errors:      row[columns["errors"]].(string),

				MetaEmailFrom:     row[columns["metaEmailFrom"]].(string),
				MetaEmailSubject:  row[columns["metaEmailSubject"]].(string),
				MetaSlackDestType: row[columns["metaSlackDestType"]].(string),
			}

			events = append(events, messageEvent)
		}
	}

	return events
}
