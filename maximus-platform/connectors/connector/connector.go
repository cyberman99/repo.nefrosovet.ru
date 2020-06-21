package connector

import "time"

type (
	ConnectorType string
	EventStatus   string
)

type Connector interface {
	Listen() (<-chan ConnectorMessage, <-chan error)
	Close()
	ConnectorType() ConnectorType
}

// declarations
type (
	ConnectorState string

	ConnectorMessage struct {
		Event     ConnectorState `json:"event"`
		Connector ConnectorData  `json:"connector"`
		Data      ButtonData     `json:"data"`
	}

	ConnectorData struct {
		Type ConnectorType `json:"type"`
		ID   string        `json:"ID"`
	}

	ButtonData struct {
		Date       time.Time   `json:"date"`
		StatusCode EventStatus `json:"statusCode"`
	}
)

func BuildMessage(
	connType ConnectorType,
	status EventStatus,
	state ConnectorState) ConnectorMessage {
	msg := ConnectorMessage{
		Event: state,
		Connector: ConnectorData{
			connType,
			"",
		},
		Data: ButtonData{
			time.Now(),
			status,
		},
	}
	return msg
}
