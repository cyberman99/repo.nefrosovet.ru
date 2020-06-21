package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
)

const (
	BackendsCollectionName = "backends"
)

type BackendStorage struct {
	client *dbMongo.Client
}

func NewBackendStorage(mongoClient *dbMongo.Client) (storage.BackendStorage, error) {
	s := &BackendStorage{
		client: mongoClient,
	}

	if err := s.ensureBackendsIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *BackendStorage) Store(in storage.StoreBackend) (*storage.Backend, error) {
	collection := s.client.MD.Collection(BackendsCollectionName)

	if in.Groups == nil {
		in.Groups = map[string]string{}
	}

	_, err := collection.InsertOne(nil, in)
	if err != nil {
		return nil, err
	}

	return &in.Backend, nil
}

func (s *BackendStorage) Update(id string, in storage.UpdateBackend) (*storage.Backend, error) {
	collection := s.client.MD.Collection(BackendsCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	filter := bson.M{
		"id": id,
	}

	set := bson.M{}
	unset := bson.M{}

	if in.ID != nil {
		set["id"] = *in.ID
	}

	if in.Type != nil {
		set["type"] = *in.Type
	}

	if in.Description != nil {
		set["description"] = *in.Description
	}

	if in.Sync != nil {
		set["sync"] = *in.Sync
	}

	if in.Attributes != nil {
		set["attributes"] = *in.Attributes
	}

	if in.Host != nil {
		set["host"] = *in.Host
	}

	if in.Port != nil {
		set["port"] = *in.Port
	}

	if in.Cipher != nil {
		set["cipher"] = *in.Cipher
	}

	if in.SSL != nil {
		set["ssl"] = *in.SSL
	}

	if in.BindDN != nil {
		set["bindDN"] = *in.BindDN
	}

	if in.BaseDN != nil {
		set["baseDN"] = *in.BaseDN
	}

	if in.Filter != nil {
		set["filter"] = *in.Filter
	}

	if in.Password != nil {
		set["password"] = *in.Password
	}

	if in.Groups != nil {
		if len(in.Groups) == 0 {
			set["groups"] = in.Groups
		} else {
			for k, v := range in.Groups {
				if v != "" {
					set["groups."+k] = v
				} else {
					unset["groups."+k] = ""
				}
			}
		}
	}

	if in.ClientID != nil {
		set["clientID"] = *in.ClientID
	}

	if in.ClientSecret != nil {
		set["clientSecret"] = *in.ClientSecret
	}

	if in.Provider != nil {
		set["provider"] = *in.Provider
	}

	update := bson.M{}
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(unset) > 0 {
		update["$unset"] = unset
	}

	res := collection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().
			SetReturnDocument(options.After),
	)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var backend storage.Backend
	return &backend, res.Decode(&backend)
}

func (s *BackendStorage) Get(in storage.GetBackend) ([]*storage.Backend, error) {
	collection := s.client.MD.Collection(BackendsCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	filter := bson.M{}

	if in.ID != nil {
		filter["id"] = *in.ID
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var backends []*storage.Backend
	for cur.Next(nil) {
		var backend storage.Backend
		if err := cur.Decode(&backend); err != nil {
			return nil, err
		}

		backends = append(backends, &backend)
	}

	return backends, err
}

func (s *BackendStorage) Delete(id string) error {
	collection := s.client.MD.Collection(BackendsCollectionName)

	_, err := collection.DeleteOne(nil, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}

func (s *BackendStorage) ensureBackendsIndex() error {
	collection := s.client.MD.Collection(BackendsCollectionName)

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
