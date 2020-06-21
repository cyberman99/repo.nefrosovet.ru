package _default

import (
	"time"

	"github.com/Sirupsen/logrus"

	dbInflux "repo.nefrosovet.ru/maximus-platform/mailer/db/influx"
	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/influx"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/mongo"
)

type Storage struct {
	storage.AccessTokenStorage
	storage.ChannelStorage
	storage.MessageEventStorage
	storage.TelegramContactStorage
	storage.ViberContactStorage

	Ping func() error
}

var defaultStorage *Storage

func InitStorage(mongoClient *dbMongo.Client, influxClient *dbInflux.Client) {
	defaultStorage = &Storage{
		AccessTokenStorage:     mongo.GetAccessTokenStorage(mongoClient),
		ChannelStorage:         mongo.GetChannelStorage(mongoClient),
		MessageEventStorage:    influx.GetMessageEventStorage(influxClient),
		TelegramContactStorage: mongo.GetTelegramContactStorage(mongoClient),
		ViberContactStorage:    mongo.GetViberContactStorage(mongoClient),

		Ping: func() error {
			err := mongoClient.Ping(nil, nil)
			if err != nil {
				return err
			}

			_, _, err = influxClient.Ping(time.Second * 10)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func GetStorage() *Storage {
	if defaultStorage == nil {
		logrus.Error("default storage isn't initialized")
	}

	return defaultStorage
}
