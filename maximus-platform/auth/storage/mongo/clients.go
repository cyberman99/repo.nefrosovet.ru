package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"gopkg.in/mgo.v2/bson"
	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
)

type ClientStorage struct {
	client *dbMongo.Client
}

const (
	ClientsCollectionName = "clients"
)

func NewClientStorage(mongoClient *dbMongo.Client) (storage.ClientStorage, error) {
	s := &ClientStorage{client: mongoClient}

	if err := s.ensureClientsIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *ClientStorage) Store(client storage.ClientStorer) error {
	collection := s.client.MD.Collection(ClientsCollectionName)

	_, err := collection.InsertOne(nil, bson.M{
		"id": client.ID,
		"description": client.Descriptions,
		"password": client.Password,
	})

	return err
}

func convertUpdaterToBson(updater storage.ClientUpdater) (bson.M, error) {
	bUpdater := bson.M{}

	if updater.Descriptions != nil {
		bUpdater["description"] = updater.Descriptions
	}

	if updater.Password != nil {
		bUpdater["password"] = updater.Password
	}

	//TODO

	return bUpdater, nil
}

func (s *ClientStorage) Update(clientId string, updater storage.ClientUpdater) error {
	collection := s.client.MD.Collection(ClientsCollectionName)

	bUpdater, err := convertUpdaterToBson(updater)
	if err != nil {
		return err
	}

	res := collection.FindOneAndUpdate(nil, bson.M{"id": clientId},  bson.M{"$set": bUpdater})
	err = res.Err()

	if err != nil && err == mongo.ErrNoDocuments {
		return storage.ErrNotFound
	}

	return err
}

func (s *ClientStorage) Get(filter storage.ClientFilter) ([]*storage.ClientStorer, error) {
	collection := s.client.MD.Collection(ClientsCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)

	bFilter := bson.M{}
	if filter.ID != nil {
		bFilter["id"] = filter.ID
	}

	cur, err := collection.Find(ctx, bFilter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var clients []*storage.ClientStorer
	for cur.Next(nil) {
		var elem storage.ClientStorer
		err2 := cur.Decode(&elem)
		if err2 != nil {
			return nil, err
		}

		clients = append(clients, &elem)
	}

	if err != nil {
		return nil, err
	}

	if len(clients) == 0 {
		return nil, storage.ErrNotFound
	}

	return clients, nil
}

func (s *ClientStorage) Delete(clientId string) error {
	collection := s.client.MD.Collection(ClientsCollectionName)

	res := collection.FindOneAndDelete(nil, bson.M{"id": clientId})
	err := res.Err()

	if err != nil && err == mongo.ErrNoDocuments {
		return storage.ErrNotFound
	}

	return err
}

func (s *ClientStorage) ensureClientsIndex() error {
	collection := s.client.MD.Collection(ClientsCollectionName)

	//TODO DropDups:   true, now not supported
	options := options.IndexOptions{
		Unique:     ptrFromBool(true),
		Background: ptrFromBool(true),
		Sparse:     ptrFromBool(true),
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		//About what mean 1 at value: https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
		Keys: bsonx.Doc{{ Key: "id", Value: bsonx.Int32(1) }},
		Options: &options,
	})

	if err != nil {
		if v, ok := err.(driver.Error); ok && v.Code == 11000 {
			return nil
		}
	}

	return err
}
