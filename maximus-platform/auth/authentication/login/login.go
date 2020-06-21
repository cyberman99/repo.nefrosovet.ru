package login

import (
	"errors"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/admin"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/index"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/ldap"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	ErrUnknownBackendType = errors.New("unknown backend type")
	ErrInAuth             = errors.New("auth error")
)

type Credentials struct {
	Login    string
	Password string

	SmartCardNumber string
}

type Result struct {
	EntityID    string
	EntityLogin string

	JWT *jwt.JWT

	Error error
}

func Auth(credentials *Credentials) *Result {
	if credentials.Login == "admin" {
		res := admin.Auth(&admin.Credentials{
			Password: credentials.Password,
		})
		if res.Error != nil {
			return &Result{
				Error: res.Error,
			}
		}

		tokens, err := jwt.GenerateTokens("admin", "", []*storage.Role{})
		if err != nil {
			return &Result{
				EntityLogin: credentials.Login,
				Error:       err,
			}
		}

		return &Result{
			EntityLogin: credentials.Login,
			JWT:         tokens,
		}
	}

	bos := st.GetStorage().BackendsOrderStorage

	backendsOrder, err := bos.Get()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "authentication",
			"status":   "FAILED",
		}).Error("getting backends order failed")
		logrus.Debug(err)

		return &Result{
			EntityLogin: credentials.Login,
			Error:       err,
		}
	}

	logrus.WithFields(logrus.Fields{
		"context":       "CORE",
		"resource":      "Login",
		"function":      "Auth",
		"backendsOrder": backendsOrder,
	}).Debug("Auth by login order got")

	var tokens *jwt.JWT
	for _, backendID := range backendsOrder.Order {
		err = nil

		if backendID == "index" {
			res := index.Auth(&index.Credentials{
				Login:    credentials.Login,
				Password: credentials.Password,

				SmartCardNumber: credentials.SmartCardNumber,
			})
			if res.Error != nil {
				err = res.Error

				continue
			} else {
				tokens, err = jwt.GenerateTokens(res.EntityID, "", []*storage.Role{})
			}

			continue
		}

		bs := st.GetStorage().BackendStorage

		backends, err := bs.Get(storage.GetBackend{
			ID: &backendID,
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "authentication",
				"status":   "FAILED",
			}).Error("getting backend by ID failed")
			logrus.Debug(err)

			return &Result{
				EntityLogin: credentials.Login,
				Error:       err,
			}
		} else if len(backends) == 0 {
			return &Result{
				EntityLogin: credentials.Login,
				Error:       storage.ErrBackendNotFound,
			}
		}

		switch backends[0].Type {
		default:
			logrus.WithFields(logrus.Fields{
				"context":    "CORE",
				"eventType":  "AuthByLogin",
				"entityType": "backend",
				"status":     "FAILED",
			}).Error("unknown backend type")
			logrus.Debug(backends[0])

			return &Result{
				Error: ErrUnknownBackendType,
			}
		case storage.BackendTypeOAuth2:
			continue
		case storage.BackendTypeLDAP:
			res := ldap.Auth(&ldap.Credentials{
				Login:    credentials.Login,
				Password: credentials.Password,

				Backend: backends[0],
			})
			if res.Error != nil {
				err = res.Error

				continue
			} else {
				tokens, err = jwt.GenerateTokens(res.EntityID, "", res.TempRoles)
			}
		}

		if tokens != nil {
			break
		}
	}

	if tokens == nil || tokens.UserID == "" {
		if err == nil {
			err = ErrInAuth
		}

		return &Result{
			EntityLogin: credentials.Login,
			Error:       err,
		}
	}

	return &Result{
		EntityID:    tokens.UserID,
		EntityLogin: credentials.Login,
		JWT:         tokens,
	}
}
