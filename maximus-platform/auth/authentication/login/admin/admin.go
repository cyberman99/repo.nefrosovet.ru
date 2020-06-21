package admin

import (
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

type Credentials struct {
	Password string
}

type Result struct {
	Error error
}

func Auth(credentials *Credentials) *Result {
	ps := st.GetStorage().AdminPasswordStorage

	pass, err := ps.Get()
	if err != nil {
		return &Result{
			Error: err,
		}
	} else if len(pass) == 0 {
		return &Result{
			Error: storage.ErrAdminPasswordNotFound,
		}
	}

	newHash, err := jwt.VerifyPassword(credentials.Password, pass[0].Hash)
	if err != nil {
		return &Result{
			Error: err,
		}
	}

	if newHash != "" {
		_, err = ps.Update(storage.UpdateAdminPassword{
			Hash: newHash,
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":    "CORE",
				"eventType":  "configDB save",
				"entityType": "adminPassword",
				"status":     "FAILED",
			}).Error("save password hash failed")
			logrus.Debug(err)
		}
	}

	return &Result{
		Error: err,
	}
}
