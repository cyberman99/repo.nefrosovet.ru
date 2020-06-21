package storage

import (
	"encoding/json"
	"errors"
)

type MessageEventStorage interface {
	StoreMessageEvent(in StoreMessageEvent) (MessageEvent, error)
	GetMessageEvent(in GetMessageEvent) (MessageEvent, error)
	GetMessageEvents(in GetMessageEvents) ([]MessageEvent, error)
	GetMessageEventCount(in GetMessageEventCount) (int64, error)
	DeleteMessageEvents(in DeleteMessageEvents) error
}

var (
	ErrMessageEventsNotFound = errors.New("not found")
)

type MessageEvent struct {
	ID          string
	Created     string
	Updated     string
	Status      string
	Errors      string
	ChannelID   string
	ChannelType string
	AccessToken string
	Destination string
	Data        string

	MetaEmailSubject  string
	MetaEmailFrom     string
	MetaSlackDestType string
}

func (e *MessageEvent) JSONString() string {
	s, err := json.Marshal(e)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(s)
}

type StoreMessageEvent struct {
	MessageEvent MessageEvent
}

type GetMessageEvent struct {
	ID string
}

type GetMessageEvents struct {
	AccessToken string
	ChannelID   *string
	Status      *string
	Destination *string

	Limit  *int64
	Offset *int64
}

type GetMessageEventCount struct {
	ChannelID string
}

type DeleteMessageEvents struct {
	ChannelID string
}
