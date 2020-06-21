package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
)

type RoleStorage struct {
	client *dbMongo.Client
	userStorage *storage.UserStorage
}

// RoleCollectionName - is document collection name
const RoleCollectionName = "roles"

func NewRoleStorage(mongoClient *dbMongo.Client, userStorage *storage.UserStorage) (storage.RoleStorage, error) {
	s := &RoleStorage{client: mongoClient, userStorage: userStorage}

	if err := s.ensureRolesIndex(); err != nil {
		return nil, err
	}

	if err := s.ensureDefaultRoles(); err != nil {
		return nil, err
	}

	return s, nil
}

// Insert inserts role to DB
func (s *RoleStorage) Store(role storage.Role) error {
	collection := s.client.MD.Collection(RoleCollectionName)

	_, err := collection.InsertOne(nil, bson.M{
		"id": role.ID,
		"description": role.Description,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Role.Insert",
			"error":    err,
		}).Debug("mongo request failed")
	}

	return err
}

// Update rewrites role on DB
func (s *RoleStorage) Update(role storage.Role) error {
	collection := s.client.MD.Collection(RoleCollectionName)
	res := collection.FindOneAndUpdate(nil, bson.M{"id": role.ID,},  bson.M{"$set": bson.M{
		"id": role.ID,
		"description": role.Description,
	}})
	err := res.Err()

	if err != nil && err == mongo.ErrNoDocuments {
		return storage.ErrNotFound
	}
	
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Role.UpdateRole",
			"error":    err,
		}).Debug("mongo request failed")
	}

	return err
}

// GetRole returns role from DB
func (s *RoleStorage) Get(id string) (*storage.Role, error) {
	collection := s.client.MD.Collection(RoleCollectionName)

	res := collection.FindOne(nil, bson.M{"id": id})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, storage.ErrNotFound
	}

	if res.Err() != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetRole",
			"error":    res.Err(),
		}).Debug("mongo request failed")

		return nil, res.Err()
	}

	var doc storage.Role
	err := res.Decode(&doc)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetRole",
			"error":    err,
		}).Debug("decode doc error")
	}

	return &doc, nil
}

// GetRoles returns all roles from DB
func (s *RoleStorage) GetAll() ([]*storage.Role, error) {
	collection := s.client.MD.Collection(RoleCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetRoles",
			"error":    err,
		}).Debug("mongo request failed")
	}
	defer cur.Close(ctx)

	var roles []*storage.Role
	for cur.Next(nil) {
		var elem storage.Role
		err2 := cur.Decode(&elem)
		if err2 != nil {
			logrus.WithFields(logrus.Fields{
				"function": "GetRoles",
				"error":    err2,
			}).Debug("decode doc error")
		}

		roles = append(roles, &elem)
	}

	return roles, err
}

//us storage.UserStorage,
func (s *RoleStorage) Delete(id string) error {
	if storage.IsDefaultRole(id) {
		return errors.New("can't delete default role")
	}

	collection := s.client.MD.Collection(RoleCollectionName)
	res := collection.FindOneAndDelete(nil, bson.M{"id": id})
	err := res.Err()

	if err != nil && err == mongo.ErrNoDocuments {
		return storage.ErrNotFound
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Role.Delete",
			"error":    err,
		}).Debug("mongo request failed")

		return err
	}

	return nil
}

func ptrFromBool(v bool) *bool {
	return &v
}

// EnsureRolesIndex works with mongo index
func (s *RoleStorage) ensureRolesIndex() error {
	collection := s.client.MD.Collection(RoleCollectionName)

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

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Role.EnsureRolesIndex",
			"error":    err,
		}).Debug("mongo request failed")
	}
	return err
}

// EnsureDefaultRoles creates default roles on DB if not exist.
func (s *RoleStorage) ensureDefaultRoles() error {
	for _, roleID := range []string{storage.RoleDefaultAdmin, storage.RoleDefaultEmployee, storage.RoleDefaultPatient} {
		_, err := s.Get(roleID)
		if err != nil && err == storage.ErrNotFound {
			role := storage.Role{
				ID: roleID,
				Description: fmt.Sprintf("Default %s role", roleID),
			}
			err = s.Store(role)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
