package mongo

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"gopkg.in/mgo.v2/bson"

	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
)

const (
	BackendsOrderCollectionName = "backends_order"
	BackendsOrderID             = "backends_order"
)

type BackendsOrderStorage struct {
	client *dbMongo.Client
}

func NewBackendsOrderStorage(mongoClient *dbMongo.Client) (storage.BackendsOrderStorage, error) {
	s := &BackendsOrderStorage{client: mongoClient}

	if err := s.ensureBackendsOrderIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *BackendsOrderStorage) Store(in storage.StoreBackendsOrder) (*storage.BackendsOrder, error) {
	collection := s.client.MD.Collection(BackendsOrderCollectionName)

	_, err := collection.InsertOne(nil, bson.M{
		"id":    BackendsOrderID,
		"order": in.Order,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, storage.ErrBackendsOrderAlreadyExists
		}

		return nil, err
	}

	return &storage.BackendsOrder{
		ID:    BackendsOrderID,
		Order: in.Order,
	}, nil
}

func (s *BackendsOrderStorage) Update(in storage.UpdateBackendsOrder) (*storage.BackendsOrder, error) {
	collection := s.client.MD.Collection(BackendsOrderCollectionName)

	res := collection.FindOneAndUpdate(nil, bson.M{"id": BackendsOrderID}, bson.M{"$set": bson.M{
		"order": in.Order,
	}}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		return nil, res.Err()
	}

	var backendsOrder storage.BackendsOrder
	return &backendsOrder, res.Decode(&backendsOrder)
}

func (s *BackendsOrderStorage) Get() (*storage.BackendsOrder, error) {
	collection := s.client.MD.Collection(BackendsOrderCollectionName)

	filter := bson.M{
		"id": BackendsOrderID,
	}

	res := collection.FindOne(nil, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, storage.ErrBackendsOrderNotFound
		}

		return nil, res.Err()
	}

	var backendsOrder storage.BackendsOrder
	return &backendsOrder, res.Decode(&backendsOrder)
}

func (s *BackendsOrderStorage) Delete(in storage.DeleteBackendsOrder) error {
	collection := s.client.MD.Collection(BackendsOrderCollectionName)

	update := bson.M{}
	if len(in.IDs) == 0 {
		update = bson.M{"$set": bson.M{
			"order": storage.DefaultBackendsOrder,
		}}
	} else {
		update = bson.M{"$pull": bson.M{
			"order": bson.M{
				"$in": in.IDs,
			},
		}}
	}

	res := collection.FindOneAndUpdate(nil, bson.M{"id": BackendsOrderID}, update)
	if res.Err() != nil {
		switch {
		case errors.Is(res.Err(), storage.ErrNotFound):
			return storage.ErrBackendsOrderNotFound
		}

		return res.Err()
	}

	return nil
}

func (s *BackendsOrderStorage) ensureBackendsOrderIndex() error {
	collection := s.client.MD.Collection(BackendsOrderCollectionName)

	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key:   "id",
			Value: bsonx.Int32(1),
		}},
		Options: options.Index().
			SetUnique(true).
			SetBackground(true).
			SetSparse(true),
	})
	if err != nil {
		return err
	}

	return err
}
