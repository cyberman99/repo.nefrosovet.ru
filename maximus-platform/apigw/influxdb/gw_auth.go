package influxdb

import (
	"errors"
	"time"

	"fmt"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
)

type Event struct {
	ID       string
	Status   string
	IP       string
	PolicyID string
	Endpoint string
	Method   string
	Path     string
	Roles    string
}

const (
	eventsTable = "events"

	eventsFieldEventID  = "eventID"
	eventsFieldStatus   = "status"
	eventsFieldSourceIP = "IP"
	eventsFieldPolicyID = "policyID"
	eventsFieldEndpoint = "endpoint"
	eventsFieldMethod   = "method"
	eventsFieldPath     = "path"
	eventsFieldRoles    = "roles"
)

// EventsParams is params struct
type EventsParams struct {
	ID       string
	Status   string
	IP       string
	Endpoint string
	Method   string
	Path     string
	Roles    string
	PolicyID string

	Limit  int64
	Offset int64
}

// LogEvent - logs event
func LogEvent(params EventsParams) {
	eventID := uuid.New().String()

	logrus.WithFields(
		logrus.Fields{
			eventsFieldEventID:  eventID,
			eventsFieldStatus:   params.Status,
			eventsFieldSourceIP: params.IP,
			eventsFieldPolicyID: params.PolicyID,
			eventsFieldEndpoint: params.Endpoint,
			eventsFieldMethod:   params.Method,
			eventsFieldPath:     params.Path,
			eventsFieldRoles:    params.Roles,
		},
	).Info("APIgw authorization event")

	if cln == nil {
		logrus.Warn("No InfluxDB connection")

		return
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Database,
		Precision: "s",
	})
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return
	}

	tags := map[string]string{
		eventsFieldEventID:  eventID,
		eventsFieldStatus:   params.Status,
		eventsFieldSourceIP: params.IP,
		eventsFieldPolicyID: params.PolicyID,
		eventsFieldEndpoint: params.Endpoint,
		eventsFieldMethod:   params.Method,
		eventsFieldPath:     params.Path,
	}

	fields := map[string]interface{}{
		eventsFieldRoles: params.Roles,
	}

	point, err := client.NewPoint(
		eventsTable,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return
	}

	bp.AddPoint(point)

	err = cln.Write(bp)
	if err != nil {
		logrus.WithError(err).Error("saving event error")

		return
	}

	return
}

// GetEventsParams is params struct
type GetEventsParams struct {
	ID       string
	Status   string
	IP       string
	PolicyID string
	Endpoint string
	Method   string
	Path     string

	Limit  int64
	Offset int64
}

// GetEvents - returns events log
func GetEvents(params GetEventsParams) ([]*Event, error) {
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE true", eventsTable)

	if params.ID != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldEventID, params.ID)
	}
	if params.Status != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldStatus, params.Status)
	}
	if params.IP != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldSourceIP, params.IP)
	}
	if params.PolicyID != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldPolicyID, params.PolicyID)
	}
	if params.Endpoint != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldEndpoint, params.Endpoint)
	}
	if params.Method != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldMethod, params.Method)
	}
	if params.Path != "" {
		queryString = fmt.Sprintf("%s AND %s = '%s'", queryString, eventsFieldPath, params.Path)
	}
	if params.Limit != 0 {
		queryString = fmt.Sprintf("%s LIMIT %d", queryString, params.Limit)
	}
	if params.Offset != 0 {
		queryString = fmt.Sprintf("%s OFFSET %d", queryString, params.Offset)
	}

	q := client.Query{
		Command:  queryString,
		Database: Database,
	}

	resp, err := cln.Query(q)
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, errors.New("events not found")
	}

	events := constructEventsResult(resp.Results)

	return events, err
}

// GetEventByID - returns event by ID
func GetEventByID(id string) (*Event, error) {
	q := client.Query{
		Command:  fmt.Sprintf("select * from %s where %s = '%s'", eventsTable, eventsFieldEventID, id),
		Database: Database,
	}

	resp, err := cln.Query(q)
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if len(resp.Results) == 0 || len(resp.Results[0].Series) == 0 {
		return nil, errors.New("events not found")
	}

	events := constructEventsResult(resp.Results)

	return events[0], err
}

func constructEventsResult(results []client.Result) []*Event {
	columns := make(map[string]int)
	var events []*Event
	for _, result := range results {

		for num, column := range result.Series[0].Columns {
			columns[column] = num
		}

		for _, row := range result.Series[0].Values {

			events = append(
				events,
				&Event{
					ID:       stringValue(row[columns[eventsFieldEventID]]),
					Status:   stringValue(row[columns[eventsFieldStatus]]),
					IP:       stringValue(row[columns[eventsFieldSourceIP]]),
					PolicyID: stringValue(row[columns[eventsFieldPolicyID]]),
					Endpoint: stringValue(row[columns[eventsFieldEndpoint]]),
					Method:   stringValue(row[columns[eventsFieldMethod]]),
					Path:     stringValue(row[columns[eventsFieldPath]]),
					Roles:    stringValue(row[columns[eventsFieldRoles]]),
				},
			)
		}
	}

	return events
}
