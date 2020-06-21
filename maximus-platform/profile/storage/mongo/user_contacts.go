package mongo

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

const UserContactsCollectionName = "users_contacts"

type UserContactStorage struct {
	client *dbMongo.Client

	storage.UserContactStorage
}

func newUserContactStorage(client *dbMongo.Client) *UserContactStorage {
	return &UserContactStorage{
		client: client,
	}
}

func (s *UserContactStorage) StoreUserContact(in storage.StoreUserContact) (storage.UserContact, error) {
	var err error

	collection := s.client.Collection(UserContactsCollectionName)
	ctx, _ := context.WithTimeout(background, 1*time.Second)

	in.Created, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return storage.UserContact{}, err
	}

	_, err = collection.InsertOne(ctx, in)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return storage.UserContact{}, storage.ErrUserContactAlreadyExists
		}

		return storage.UserContact{}, err
	}

	return storage.UserContact{
		TypeCODE: in.TypeCODE,
		Value:    in.Value,
		Created:  in.Created,
	}, nil
}

func (s *UserContactStorage) UpdateUserContact(userID storage.UUID, in storage.UpdateUserContact) (storage.UserContact, error) {
	collection := s.client.Collection(UserContactsCollectionName)
	ctx, _ := context.WithTimeout(background, 1*time.Second)

	filter := bson.M{
		"userID":   userID,
		"typeCODE": in.TypeCODE,
		"value":    *in.Value,
	}

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		switch result.Err() {
		case mongo.ErrNoDocuments:
			return storage.UserContact{}, storage.ErrUserContactNotFound
		default:
			return storage.UserContact{}, result.Err()
		}
	}

	var userContact storage.UserContact

	if err := result.Decode(&userContact); err != nil {
		return storage.UserContact{}, err
	}

	if userContact.Verified != nil {
		return storage.UserContact{}, storage.ErrUserContactAlreadyVerified
	}

	update := bson.M{}

	if in.TypeCODE != nil {
		update["typeCODE"] = *in.TypeCODE
	}

	if in.Value != nil {
		update["value"] = *in.Value
	}

	if in.Verified != nil && *in.Verified {
		verified, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			return storage.UserContact{}, err
		}
		update["verified"] = &verified
	}

	result = collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": update,
	}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		switch {
		//case errors.Is(result.Err(), mongo.ErrNoDocuments):
		//	return storage.UserContact{}, storage.ErrUserContactNotFound
		case strings.Contains(result.Err().Error(), "duplicate"):
			return storage.UserContact{}, storage.ErrUserContactAlreadyExists
		default:
			return storage.UserContact{}, result.Err()
		}
	}

	if err := result.Decode(&userContact); err != nil {
		return storage.UserContact{}, err
	}

	return userContact, nil
}

func (s *UserContactStorage) GetUserContacts(in storage.GetUserContacts) ([]storage.UserContact, error) {
	//collection := s.client.Collection(UserContactsCollectionName)
	//ctx, _ := context.WithTimeout(background, 1*time.Second)

	var contacts []storage.UserContact

	//lastEmailContact := getLastContact("EMAIL", user, ctx)
	//contacts = append(contacts, lastEmailContact)
	//
	//lastMobileContact:= getLastContact("MOBILE", user, ctx)
	//contacts = append(contacts, lastMobileContact)

	return contacts, nil
}

func (s *UserContactStorage) GetUserContact(in storage.GetUserContact) (storage.UserContact, error) {
	collection := s.client.Collection(UserContactsCollectionName)
	ctx, _ := context.WithTimeout(background, 1*time.Second)

	filter := bson.M{}

	if in.ID != nil {
		filter["_id"] = *in.ID
	}

	if in.Value != nil {
		filter["value"] = *in.Value
	}

	if len(filter) == 0 {
		return storage.UserContact{}, storage.ErrBadInput
	}

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		switch result.Err() {
		case mongo.ErrNoDocuments:
			return storage.UserContact{}, storage.ErrUserContactNotFound
		default:
			return storage.UserContact{}, result.Err()
		}
	}

	var contact storage.UserContact
	if err := result.Decode(&contact); err != nil {
		return storage.UserContact{}, err
	}

	return contact, nil
}
