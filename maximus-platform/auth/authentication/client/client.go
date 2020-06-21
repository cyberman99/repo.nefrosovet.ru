package client

import (
	"errors"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/sirupsen/logrus"
	"gopkg.in/hlandau/passlib.v1"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
)

var (
	ErrNoClientLoginGiven = errors.New("no client login given")
)

type Credentials struct {
	Login    string
	Password string
}

type Result struct {
	EntityID string
	JWT      *jwt.JWT

	Error error
}

func Auth(credentials *Credentials) *Result {
	if credentials.Login == "" {
		return &Result{
			Error: ErrNoClientLoginGiven,
		}
	}

	cs := st.GetStorage().ClientStorage
	clients, err := cs.Get(storage.ClientFilter{ID: &credentials.Login})
	if err != nil {
		return &Result{
			Error: err,
		}
	}
	client := clients[0]

	newHash, err := VerifyPassword(credentials.Password, client.Password)
	if err != nil {
		return &Result{
			EntityID: client.ID,
			Error:    err,
		}
	}

	if newHash != "" {
		err := cs.Update(client.ID, storage.ClientUpdater{
			Descriptions: &client.Descriptions,
			Password: &newHash,
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":    "CORE",
				"eventType":  "configDB save",
				"entityType": "client",
				"status":     "FAILED",
			}).Error("save client password hash failed")
			logrus.Debug(err)
		}
	}

	tokens, err := jwt.GenerateTokens(client.ID, "", []*storage.Role{})
	if err != nil {
		return &Result{
			EntityID: client.ID,
			Error:    err,
		}
	}

	return &Result{
		EntityID: client.ID,
		JWT:      tokens,
	}
}

// HashPassword returns password hash
func HashPassword(password string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Hash(password)
}

// VerifyPassword validates password and regenerate hash
func VerifyPassword(password, hash string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Verify(password, hash)
}
