package storage_accessor

import (
	"github.com/sirupsen/logrus"
	dbInflux "repo.nefrosovet.ru/maximus-platform/auth/db/influx"
	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"repo.nefrosovet.ru/maximus-platform/auth/storage/influx"
	"repo.nefrosovet.ru/maximus-platform/auth/storage/mongo"
)

type RootStorage struct {
	storage.AdminPasswordStorage
	storage.BackendStorage
	storage.BackendsOrderStorage
	storage.ClientStorage
	storage.EventStorage
	storage.RoleStorage
	storage.TokenStorage
	storage.UserStorage
}

var defaultStorage *RootStorage

func InitDefaultStorage(mongoClient *dbMongo.Client, influxClient *dbInflux.Client) error {
	as, err := mongo.NewAdminPasswordStorage(mongoClient)
	if err != nil {
		return err
	}
	bos, err := mongo.NewBackendsOrderStorage(mongoClient)
	if err != nil {
		return err
	}
	us, err := mongo.NewUserStorage(mongoClient)
	if err != nil {
		return err
	}
	rs, err := mongo.NewRoleStorage(mongoClient, &us)
	if err != nil {
		return err
	}
	bs, err := mongo.NewBackendStorage(mongoClient)
	if err != nil {
		return err
	}
	cs, err := mongo.NewClientStorage(mongoClient)
	if err != nil {
		return err
	}
	es, err := influx.NewEventStorage(influxClient)
	if err != nil {
		return err
	}
	ts, err := mongo.NewTokenStorage(mongoClient)
	if err != nil {
		return err
	}

	s := RootStorage{
		as,
		bs,
		bos,
		cs,
		es,
		rs,
		ts,
		us,
	}

	defaultStorage = &s

	return nil
}

func GetStorage() *RootStorage {
	if defaultStorage == nil {
		logrus.Error("default storage isn't initialized")
	}

	return defaultStorage
}