package influx

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	dbInflux "repo.nefrosovet.ru/maximus-platform/auth/db/influx"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"time"
)

type EventStorage struct {
	client *dbInflux.Client
}

const (
	eventsTable = "events"

	eventsFieldEventID     = "eventID"
	eventsFieldEventType   = "eventType"
	eventsFieldSourceIP    = "IP"
	eventsFieldStatus      = "status"
	eventsFieldExtraData   = "data"
	eventsFieldEntityLogin = "entityLogin"
	eventsFieldEntityID    = "entityID"
	eventsFieldTime        = "time"
)

// EnsureDatabase - creates DB with data retention policy
func (s *EventStorage) ensureDatabase() {
	queryString := fmt.Sprintf(`CREATE DATABASE "%s"`, s.client.Database)
	query := client.NewQuery(queryString, "", "")
	response, err := s.client.Cln.Query(query)
	if err != nil || response.Error() != nil {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     s.client.Address,
			"database": s.client.Database,
			"status":   "FAILED",
			"errors":    []string{err.Error(), response.Err},
		}).Fatal("Ensuring influxDB database")
	}
}

// alter changes INSERT query to UPDATE on EnsureDatabaseRetentionPolicy function
var alter bool

// ensureDatabaseRetentionPolicy - creates DB retention policy
func (s *EventStorage) ensureDatabaseRetentionPolicy() {
	if s.client.Retention == "" {
		return
	}

	operation := "CREATE"
	if alter {
		operation = "ALTER"
	}
	queryString := fmt.Sprintf(`
		%s RETENTION POLICY "custom"
		ON "%s"
		DURATION %s
		REPLICATION 1 DEFAULT
	`,
		operation,
		s.client.Database,
		s.client.Retention,
	)
	query := client.NewQuery(queryString, "", "")
	response, err := s.client.Cln.Query(query)
	if err != nil || response.Error() != nil {
		if !alter {
			alter = true
			logrus.Debug("ALTER custom database retention policy")
			s.ensureDatabaseRetentionPolicy()
		}

		if err == nil {
			err = response.Error()
		}

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     s.client.Address,
			"database": s.client.Database,
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Ensuring influxDB database retention policy")
	}
}

func NewEventStorage(influxClient *dbInflux.Client) (storage.EventStorage, error) {
	c := &EventStorage{client: influxClient}

	c.ensureDatabase()
	c.ensureDatabaseRetentionPolicy()

	return c, nil
}

func (s *EventStorage) Store(in storage.StoreEvent) error {
	err := s.store(in)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "PostIdentify",
			"error":    err,
		}).Debug("StoreEvent request failed")
	}

	return err
}

func (s *EventStorage) store(in storage.StoreEvent) error {
	eventID := uuid.New().String()

	logrus.WithFields(
		logrus.Fields{
			eventsFieldEventID:     eventID,
			eventsFieldEventType:   in.EventType,
			eventsFieldSourceIP:    in.SourceIP,
			eventsFieldEntityID:    in.EntityID,
			eventsFieldEntityLogin: in.EntityLogin,
			eventsFieldStatus:      in.Status,
		},
	).Info("Auth event")

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  s.client.Database,
		Precision: "s",
	})
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return err
	}

	tags := map[string]string{
		eventsFieldEventID:     eventID,
		eventsFieldEventType:   in.EventType,
		eventsFieldSourceIP:    in.SourceIP,
		eventsFieldEntityID:    in.EntityID,
		eventsFieldEntityLogin: in.EntityLogin,
		eventsFieldStatus:      in.Status,
	}

	fields := map[string]interface{}{
		eventsFieldExtraData: in.Data,
	}

	point, err := client.NewPoint(
		eventsTable,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return err
	}

	bp.AddPoint(point)

	err = s.client.Cln.Write(bp)
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return err
	}

	return nil
}

func (s *EventStorage) Get(in storage.GetEvent) (*storage.Event, error) {
	q := client.Query{
		Command:  fmt.Sprintf("select * from %s where %s = '%s'", eventsTable, eventsFieldEventID, in.ID),
		Database: s.client.Database,
	}

	resp, err := s.client.Cln.Query(q)
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, storage.ErrNotFound
	}

	events := constructEventsResult(resp.Results)

	return events[0], err
}

func (s *EventStorage) GetAll(in storage.GetEvents) ([]*storage.Event, error){
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE true", eventsTable)

	if in.EventType != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldEventType, in.EventType)
	}
	if in.SourceIP != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldSourceIP, in.SourceIP)
	}
	if in.EntityID != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldEntityID, in.EntityID)
	}
	if in.EntityLogin != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldEntityLogin, in.EntityLogin)
	}
	if in.Status != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldStatus, in.Status)
	}
	if in.Limit != 0 {
		queryString = fmt.Sprintf("%s LIMIT %d", queryString, in.Limit)
	}
	if in.Offset != 0 {
		queryString = fmt.Sprintf("%s OFFSET %d", queryString, in.Offset)
	}

	q := client.Query{
		Command:  queryString,
		Database: s.client.Database,
	}

	resp, err := s.client.Cln.Query(q)
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, storage.ErrNotFound
	}

	events := constructEventsResult(resp.Results)

	return events, err
}

func constructEventsResult(results []client.Result) []*storage.Event {
	columns := make(map[string]int)
	var events []*storage.Event
	for _, result := range results {

		for num, column := range result.Series[0].Columns {
			columns[column] = num
		}

		for _, row := range result.Series[0].Values {

			events = append(
				events,
				&storage.Event{
					ID:          stringValue(row[columns[eventsFieldEventID]]),
					Type:        stringValue(row[columns[eventsFieldEventType]]),
					SourceIP:    stringValue(row[columns[eventsFieldSourceIP]]),
					EntityID:    stringValue(row[columns[eventsFieldEntityID]]),
					EntityLogin: stringValue(row[columns[eventsFieldEntityLogin]]),
					Status:      stringValue(row[columns[eventsFieldStatus]]),
					Time:        stringValue(row[columns[eventsFieldTime]]),
				},
			)
		}
	}

	return events
}

func stringValue(value interface{}) string {
	if value == nil {
		return ""
	}

	return value.(string)
}