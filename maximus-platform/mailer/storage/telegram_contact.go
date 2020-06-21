package storage

import (
	"encoding/json"
	"errors"
)

type TelegramContactStorage interface {
	StoreTelegramContact(in StoreTelegramContact) (TelegramContact, error)
	GetTelegramContact(in GetTelegramContact) (TelegramContact, error)
	UpdateTelegramContact(in UpdateTelegramContact) (TelegramContact, error)
}

var (
	ErrTelegramContactNotFound      = errors.New("not found")
	ErrTelegramContactAlreadyExists = errors.New("already exists")
)

type TelegramContact struct {
	Phone  string `bson:"phone"`
	ChatID int64  `bson:"chatID,omitempty" json:"chatID,omitempty"`

	Username string `bson:"username,omitempty" json:"username,omitempty"`
}

func (c *TelegramContact) JSONString() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(s)
}

type StoreTelegramContact struct {
	TelegramContact TelegramContact
}

type GetTelegramContact struct {
	Phone  *string
	ChatID *int64
}

type UpdateTelegramContact struct {
	Phone  *string
	ChatID *int64

	TelegramContact TelegramContact
}
