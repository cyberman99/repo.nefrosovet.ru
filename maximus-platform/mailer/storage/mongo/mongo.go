package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

func GetAccessTokenStorage(mongoClient *dbMongo.Client) storage.AccessTokenStorage {
	return &AccessTokenStorage{client: mongoClient}
}

func GetChannelStorage(mongoClient *dbMongo.Client) storage.ChannelStorage {
	return &ChannelStorage{client: mongoClient}
}

func GetTelegramContactStorage(mongoClient *dbMongo.Client) storage.TelegramContactStorage {
	return &TelegramContactStorage{client: mongoClient}
}

func GetViberContactStorage(mongoClient *dbMongo.Client) storage.ViberContactStorage {
	return &ViberContactStorage{client: mongoClient}
}

func Ensure(mongoClient *dbMongo.Client) error {
	indexes := map[string][]string{
		channelCollection:         {"uuid"},
		accessTokenCollection:     {"token"},
		telegramContactCollection: {"phone"},
		viberContactCollection:    {"phone", "token"},
	}

	for collectionName, key := range indexes {
		keys := bsonx.Doc{}
		for _, v := range key {
			keys = keys.Append(v, bsonx.Int32(1))
		}

		_, err := mongoClient.MD.Collection(collectionName).Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: keys,
				Options: options.Index().
					SetUnique(true).
					SetBackground(true).
					SetSparse(true),
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
