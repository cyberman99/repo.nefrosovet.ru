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
	UserCollectionName = "users"
)

type UserStorage struct {
	client *dbMongo.Client
}

func NewUserStorage(mongoClient *dbMongo.Client) (storage.UserStorage, error) {
	userStorage := UserStorage{
		client: mongoClient,
	}

	if err := userStorage.ensure(); err != nil {
		return nil, err
	}

	return &userStorage, nil
}

func (s *UserStorage) Store(in storage.StoreUser) (*storage.User, error) {
	collection := s.client.MD.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	_, err := collection.InsertOne(ctx, in)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, storage.ErrUserNotFound
		case strings.Contains(err.Error(), "duplicate"):
			return nil, storage.ErrUserAlreadyExists
		}

		return nil, err
	}

	return &in.User, nil
}

func (s *UserStorage) Update(id string, in storage.UpdateUser) (*storage.User, error) {
	collection := s.client.MD.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	filter := bson.M{
		"id": id,
	}

	set := bson.M{}
	unset := bson.M{}

	if in.BackendEntryIDs != nil {
		for k, v := range in.BackendEntryIDs {
			if v != "" {
				set["backendEntryIDs."+k] = v
			} else {
				unset["backendEntryIDs."+k] = ""
			}
		}
	}

	if in.Roles != nil {
		if len(in.Roles) == 0 {
			set["roles"] = in.Roles
		} else {
			for k, v := range in.Roles {
				if v {
					set["roles."+k] = v
				} else {
					unset["roles."+k] = ""
				}
			}
		}
	}

	update := bson.M{}
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(unset) > 0 {
		update["$unset"] = unset
	}

	res := collection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		switch {
		case errors.Is(res.Err(), mongo.ErrNoDocuments):
			return nil, storage.ErrUserNotFound
		}

		return nil, res.Err()
	}

	var user storage.User
	return &user, res.Decode(&user)
}

func (s *UserStorage) Get(in storage.GetUser) ([]*storage.User, error) {
	collection := s.client.MD.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	filter := bson.M{}

	if in.ID != nil {
		filter["id"] = *in.ID
	}

	if in.RoleID != nil {
		filter["roles."+*in.RoleID] = true
	}

	if in.BackendID != nil && in.BackendEntryID != nil {
		filter["backendEntryIDs."+*in.BackendID] = *in.BackendEntryID
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []*storage.User
	for cur.Next(ctx) {
		var user storage.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (s *UserStorage) ensure() error {
	collection := s.client.MD.Collection(UserCollectionName)
	ctx := context.Background()

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
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
