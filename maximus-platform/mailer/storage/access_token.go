package storage

import (
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type AccessTokenStorage interface {
	StoreAccessToken(in StoreAccessToken) (AccessToken, error)
	GetAccessToken(in GetAccessToken) (AccessToken, error)
	GetAccessTokens(in GetAccessTokens) ([]AccessToken, error)
	UpdateAccessToken(in UpdateAccessToken) (AccessToken, error)
	DeleteAccessToken(in DeleteAccessToken) (AccessToken, error)
}

var (
	ErrAccessTokenNotFound      = errors.New("not found")
	ErrAccessTokenAlreadyExists = errors.New("already exists")
)

type AccessToken struct {
	Token       string `bson:"token" json:"token"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

func NewAccessToken() AccessToken {
	return AccessToken{
		Token: uuid.NewV4().String(),
	}
}

func (t *AccessToken) JSONString() string {
	s, err := json.Marshal(t)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(s)
}

type StoreAccessToken struct {
	AccessToken
}

type GetAccessToken struct {
	Token string
}

type GetAccessTokens struct {
	Limit  *int
	Offset *int
}

type UpdateAccessToken struct {
	Token string

	AccessToken AccessToken
}

type DeleteAccessToken struct {
	Token string
}
