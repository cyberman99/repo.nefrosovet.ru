package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

var (
	background = context.Background()
)

func New(client *dbMongo.Client) (*storage.Storage, error) {
	if err := ensure(client); err != nil {
		return nil, err
	}

	return &storage.Storage{
		UserStorage:         newUserStorage(client),
		UserSettingsStorage: newUserSettingsStorage(client),
		UserContactStorage: newUserContactStorage(client),
	}, nil
}

func ensure(client *dbMongo.Client) error {
	_, err := client.Collection(UserSettingsCollectionName).Indexes().CreateOne(background, mongo.IndexModel{
		Keys: bson.M{
			"userID": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	_, err = client.Collection(UserContactsCollectionName).Indexes().CreateOne(background, mongo.IndexModel{
		Keys:    bson.M{
			"value": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	return nil
}
