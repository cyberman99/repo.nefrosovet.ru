package mongo

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
)

const (
	AdminPasswordCollectionName = "admin_password"
	AdminPasswordID             = "admin"
)

type AdminPasswordStorage struct {
	client *dbMongo.Client
}

func NewAdminPasswordStorage(mongoClient *dbMongo.Client) (storage.AdminPasswordStorage, error) {
	s := &AdminPasswordStorage{client: mongoClient}

	if err := s.ensurePasswordIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *AdminPasswordStorage) Store(in storage.StoreAdminPassword) (*storage.AdminPassword, error) {
	collection := s.client.MD.Collection(AdminPasswordCollectionName)

	_, err := collection.InsertOne(nil, bson.M{
		"id":   AdminPasswordID,
		"hash": in.Hash,
	})
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate"):
			return nil, storage.ErrAdminPasswordAlreadyExists
		}

		return nil, err
	}

	return nil, err
}

func (s *AdminPasswordStorage) Update(in storage.UpdateAdminPassword) (*storage.AdminPassword, error) {
	collection := s.client.MD.Collection(AdminPasswordCollectionName)

	res := collection.FindOneAndUpdate(nil, bson.M{"id": AdminPasswordID}, bson.M{"$set": bson.M{
		"hash": in.Hash,
	}}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		switch {
		case errors.Is(res.Err(), mongo.ErrNoDocuments):
			return nil, storage.ErrAdminPasswordNotFound
		}

		return nil, res.Err()
	}

	var adminPassword storage.AdminPassword
	return &adminPassword, res.Decode(&adminPassword)
}

func (s *AdminPasswordStorage) Get() ([]*storage.AdminPassword, error) {
	collection := s.client.MD.Collection(AdminPasswordCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	cur, err := collection.Find(nil, bson.M{"id": AdminPasswordID})
	if err != nil {
		return nil, err
	}

	var adminPasswords []*storage.AdminPassword
	for cur.Next(ctx) {
		var adminPassword storage.AdminPassword
		if err := cur.Decode(&adminPassword); err != nil {
			return nil, err
		}

		adminPasswords = append(adminPasswords, &adminPassword)
	}

	return adminPasswords, err
}

func (s *AdminPasswordStorage) ensurePasswordIndex() error {
	collection := s.client.MD.Collection(AdminPasswordCollectionName)

	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key:   "id",
			Value: bsonx.Int32(1),
		}},
		Options: options.Index().
			SetUnique(true).
			SetBackground(true).
			SetSparse(true),
	}, options.CreateIndexes().SetMaxTime(10*time.Second))
	return err
}
