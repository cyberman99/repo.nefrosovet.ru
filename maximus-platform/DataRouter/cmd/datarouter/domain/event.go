package domain

import (
	"errors"
)

type Event struct {
	ID               string `json:"ID"`
	DateTime         string `json:"dateTime"`
	RouteID          string `json:"routeID"`
	Qos              byte
	SourceTopic      string `json:"sourceTopic"`
	DestinationTopic string `json:"destinationTopic"`
	TransactionID    string `json:"transactionID"`
	Reply            bool   `json:"reply"` // maybe IsReplied?
}

type LanModel struct {
	Payload       interface{} `json:"payload"`
	TransactionID string      `json:"transactionID"`
}

var (
	ErrEventNotFound = errors.New("event not found")
)
