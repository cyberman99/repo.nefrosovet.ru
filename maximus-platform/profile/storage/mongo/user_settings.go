package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

const UserSettingsCollectionName = "users_settings"

type UserSettingsStorage struct {
	client *dbMongo.Client

	storage.UserSettingsStorage
}

func newUserSettingsStorage(client *dbMongo.Client) *UserSettingsStorage {
	return &UserSettingsStorage{
		client: client,
	}
}

func (s *UserSettingsStorage) StoreUserSettings(in storage.StoreUserSettings) (storage.UserSettings, error) {
	collection := s.client.Collection(UserSettingsCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	result, err := collection.InsertOne(ctx, in)
	if err != nil {
		return storage.UserSettings{}, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return storage.UserSettings{}, ErrInsertedIDNotFound
	}

	return storage.UserSettings{
		UserID:           storage.UUID(id.String()),
		TwoFAChannelType: in.TwoFAChannelType,
		Locale:           in.Locale,
	}, nil
}

func (s *UserSettingsStorage) UpdateUserSettings(userID storage.UUID, in storage.UpdateUserSettings) (storage.UserSettings, error) {
	collection := s.client.Collection(UserSettingsCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	filter := bson.M{
		"userID": userID,
	}
	update := bson.M{}

	if in.TwoFAChannelType != nil {
		update["2FAChannelType"] = *in.TwoFAChannelType
	}

	if in.Locale != nil {
		update["locale"] = *in.Locale
	}

	result := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": update,
	}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		switch result.Err() {
		case mongo.ErrNoDocuments:
			return storage.UserSettings{}, storage.ErrUserSettingsNotFound
		}

		return storage.UserSettings{}, result.Err()
	}

	var userSettings storage.UserSettings
	if err := result.Decode(&userSettings); err != nil {
		return storage.UserSettings{}, err
	}

	return userSettings, nil
}

func (s *UserSettingsStorage) GetUserSettings(in storage.GetUserSettings) (storage.UserSettings, error) {
	collection := s.client.Collection(UserSettingsCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	filter := bson.M{
		"userID": in.UserID,
	}

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		switch result.Err() {
		case mongo.ErrNoDocuments:
			return storage.UserSettings{}, storage.ErrUserSettingsNotFound
		}

		return storage.UserSettings{}, result.Err()
	}

	var userSettings storage.UserSettings
	if err := result.Decode(&userSettings); err != nil {
		return storage.UserSettings{}, err
	}

	return userSettings, nil
}
