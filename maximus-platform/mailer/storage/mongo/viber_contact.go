package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

const (
	viberContactCollection = "viber_contacts"
)

type ViberContactStorage struct {
	client *dbMongo.Client

	storage.ViberContactStorage
}

func (s *ViberContactStorage) StoreViberContact(in storage.StoreViberContact) (storage.ViberContact, error) {
	collection := s.client.MD.Collection(viberContactCollection)

	_, err := collection.InsertOne(nil, &in.ViberContact)
	if err != nil {
		switch err := err.(type) {
		case mongo.WriteException:
			if len(err.WriteErrors) >= 1 && err.WriteErrors[0].Code == 11000 {
				return storage.ViberContact{}, storage.ErrViberContactAlreadyExists
			}
		}

		return storage.ViberContact{}, err
	}

	return in.ViberContact, nil
}

func (s *ViberContactStorage) GetViberContact(in storage.GetViberContact) (storage.ViberContact, error) {
	collection := s.client.MD.Collection(viberContactCollection)

	selector := bson.M{}

	if in.Phone != nil {
		selector["phone"] = *in.Phone
	}

	if in.UserID != nil {
		selector["userID"] = *in.UserID
	}

	if in.Token != nil {
		selector["token"] = *in.Token
	}

	r := collection.FindOne(nil, selector)
	if r.Err() != nil {
		return storage.ViberContact{}, r.Err()
	}

	var viberContact storage.ViberContact
	err := r.Decode(&viberContact)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.ViberContact{}, storage.ErrViberContactNotFound
		}

		return storage.ViberContact{}, err
	}

	return viberContact, nil
}

func (s *ViberContactStorage) UpdateViberContact(in storage.UpdateViberContact) (storage.ViberContact, error) {
	collection := s.client.MD.Collection(viberContactCollection)

	selector := bson.M{}

	if in.Phone != nil {
		selector["phone"] = *in.Phone
	}

	if in.UserID != nil {
		selector["userID"] = *in.UserID
	}

	r := collection.FindOneAndUpdate(nil,
		selector,
		bson.M{"$set": &in.ViberContact},
	)
	if r.Err() != nil {
		return storage.ViberContact{}, r.Err()
	}

	var vc storage.ViberContact
	if err := r.Decode(&vc); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.ViberContact{}, storage.ErrViberContactNotFound
		}

		return storage.ViberContact{}, err
	}

	return in.ViberContact, nil
}
