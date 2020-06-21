package storage

import (
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ChannelStorage interface {
	StoreChannel(in StoreChannel) (Channel, error)
	GetChannel(in GetChannel) (Channel, error)
	GetChannels(in GetChannels) ([]Channel, error)
	UpdateChannel(in UpdateChannel) (Channel, error)
	DeleteChannel(in DeleteChannel) (Channel, error)
}

var (
	ErrChannelNotFound      = errors.New("not found")
	ErrChannelAlreadyExists = errors.New("already exists")
)

type ChannelType string

const (
	ChannelTypeEmail    ChannelType = "email"
	ChannelTypeLocalSMS ChannelType = "local_sms"
	ChannelTypeMTSSMS   ChannelType = "mts_sms"
	ChannelTypeTelegram ChannelType = "telegram"
	ChannelTypeSlack    ChannelType = "slack"
	ChannelTypeViber    ChannelType = "viber"
)

type Channel struct {
	ID          string      `bson:"uuid"`
	Type        ChannelType `bson:"type" json:"type"`
	AccessToken string      `bson:"access_token" json:"access_token"`

	Params struct {
		From          string `bson:"from,omitempty" json:"from,omitempty"`
		Login         string `bson:"login,omitempty" json:"login,omitempty"`
		Password      string `bson:"password,omitempty" json:"password,omitempty"`
		Port          int64  `bson:"port,omitempty" json:"port,omitempty"`
		Server        string `bson:"server,omitempty" json:"server,omitempty"`
		SSL           bool   `bson:"ssl,omitempty" json:"ssl,omitempty"`
		ModemID       int64  `bson:"modemID,omitempty" json:"modemID,omitempty"`
		Db            string `bson:"db,omitempty" json:"db,omitempty"`
		Token         string `bson:"token,omitempty" json:"token,omitempty"`
		GreetingText  string `bson:"greetingText,omitempty" json:"greetingText,omitempty"`
		AnswerText    string `bson:"answerText,omitempty" json:"answerText,omitempty"`
		AlternateText string `bson:"alternateText,omitempty" json:"alternateText,omitempty"`
		ButtonText    string `bson:"buttonText,omitempty" json:"buttonText,omitempty"`
		Name          string `bson:"name,omitempty" json:"name,omitempty"`
		Avatar        string `bson:"avatar,omitempty" json:"avatar,omitempty"`
		ClientID      string `bson:"clientID,omitempty" json:"clientID,omitempty"`
		Limit         int64  `bson:"limit,omitempty" json:"limit,omitempty"`
		ContentType   string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	} `bson:"params,omitempty" json:"params,omitempty"`
}

func NewChannel() Channel {
	return Channel{
		ID: uuid.NewV4().String(),
	}
}

func (c *Channel) JSONString() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(s)
}

type StoreChannel struct {
	Channel Channel
}

type GetChannel struct {
	ID *string

	Params struct {
		Token *string
	}
}

type GetChannels struct {
	Type        *ChannelType
	AccessToken *string
}

type UpdateChannel struct {
	ID string

	Channel Channel
}

type DeleteChannel struct {
	ID string
}
