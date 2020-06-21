package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

const (
	telegramContactCollection = "telegram_contacts"
)

type TelegramContactStorage struct {
	client *dbMongo.Client

	storage.TelegramContactStorage
}

func (s *TelegramContactStorage) StoreTelegramContact(in storage.StoreTelegramContact) (storage.TelegramContact, error) {
	collection := s.client.MD.Collection(telegramContactCollection)

	_, err := collection.InsertOne(nil, &in.TelegramContact)
	if err != nil {
		switch err := err.(type) {
		case mongo.WriteException:
			if len(err.WriteErrors) >= 1 && err.WriteErrors[0].Code == 11000 {
				return storage.TelegramContact{}, storage.ErrTelegramContactAlreadyExists
			}
		}

		return storage.TelegramContact{}, err
	}

	return in.TelegramContact, nil
}

func (s *TelegramContactStorage) GetTelegramContact(in storage.GetTelegramContact) (storage.TelegramContact, error) {
	collection := s.client.MD.Collection(telegramContactCollection)

	selector := bson.M{}

	if in.Phone != nil {
		selector["phone"] = *in.Phone
	}

	if in.ChatID != nil {
		selector["chatID"] = *in.ChatID
	}

	r := collection.FindOne(nil, selector)
	if r.Err() != nil {
		return storage.TelegramContact{}, r.Err()
	}

	var telegramContact storage.TelegramContact
	err := r.Decode(&telegramContact)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.TelegramContact{}, storage.ErrTelegramContactNotFound
		}

		return storage.TelegramContact{}, err
	}

	return telegramContact, nil
}

func (s *TelegramContactStorage) UpdateTelegramContact(in storage.UpdateTelegramContact) (storage.TelegramContact, error) {
	collection := s.client.MD.Collection(telegramContactCollection)

	selector := bson.M{}

	if in.Phone != nil {
		selector["phone"] = *in.Phone
	}

	if in.ChatID != nil {
		selector["chatID"] = *in.ChatID
	}

	r := collection.FindOneAndUpdate(nil,
		selector,
		bson.M{"$set": &in.TelegramContact},
	)
	if r.Err() != nil {
		return storage.TelegramContact{}, r.Err()
	}

	var tc storage.TelegramContact
	if err := r.Decode(&tc); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.TelegramContact{}, storage.ErrTelegramContactNotFound
		}

		return storage.TelegramContact{}, err
	}

	return in.TelegramContact, nil
}
