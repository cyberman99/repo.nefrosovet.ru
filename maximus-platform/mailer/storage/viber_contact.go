package storage

import (
	"encoding/json"
	"errors"
)

type ViberContactStorage interface {
	StoreViberContact(in StoreViberContact) (ViberContact, error)
	GetViberContact(in GetViberContact) (ViberContact, error)
	UpdateViberContact(in UpdateViberContact) (ViberContact, error)
}

var (
	ErrViberContactNotFound      = errors.New("not found")
	ErrViberContactAlreadyExists = errors.New("already exists")
)

type ViberContact struct {
	Phone  string `bson:"phone"`
	UserID string `bson:"userID,omitempty" json:"userID,omitempty"`

	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Token string `bson:"token,omitempty" json:"token,omitempty"`
}

func (c *ViberContact) JSONString() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(s)
}

type StoreViberContact struct {
	ViberContact ViberContact
}

type GetViberContact struct {
	Phone  *string
	UserID *string

	Token *string
}

type UpdateViberContact struct {
	Phone  *string
	UserID *string

	ViberContact ViberContact
}
