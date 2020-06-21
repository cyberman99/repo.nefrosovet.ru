package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"time"
)

type TokenStorage struct {
	client *dbMongo.Client
}

const (
	// TokenCollectionName - is document collection name
	TokenCollectionName = "tokens"
)

func NewTokenStorage(mongoClient *dbMongo.Client) (storage.TokenStorage, error) {
	s := &TokenStorage{client: mongoClient}

	if err := s.ensureTokensIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

// Insert inserts token to DB
func (s *TokenStorage) Store(token storage.Token) error {
	if token.Refresh == "" {
		return errors.New("RefreshToken is mandatory for DB save")
	}

	collection := s.client.MD.Collection(TokenCollectionName)

	_, err := collection.InsertOne(nil, bson.M{
		"refresh": token.Refresh,
		"access": token.Access,
		"username": token.Username,
		"parent": token.Parent,
		"expired": token.Expired,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Token.Insert",
			"error":    err,
		}).Debug("mongo request failed")
	}

	return err
}

// Update rewrites token on DB
func (s *TokenStorage) Update(token storage.Token) error {
	if token.Refresh == "" {
		return errors.New("RefreshToken is mandatory for DB save")
	}

	collection := s.client.MD.Collection(TokenCollectionName)
	_, err := collection.UpdateOne(nil, bson.M{"refresh": token.Refresh}, bson.M{
		"access": token.Access,
		"username": token.Username,
		"parent": token.Parent,
		"expired": token.Expired,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Token.Update",
			"error":    err,
		}).Debug("mongo request failed")
	}

	return err
}

// Delete removes token from DB
func (s *TokenStorage) Delete(refresh string) error {
	if refresh == "" {
		return errors.New("RefreshToken is mandatory for DB delete")
	}

	collection := s.client.MD.Collection(TokenCollectionName)

	_, err := collection.DeleteOne(nil, bson.M{"refresh": refresh})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Token.Delete",
			"error":    err,
		}).Debug("mongo request failed")
	}

	return err
}

// GetTokenByRefresh returns tokens by RefreshToken
func (s *TokenStorage) GetByRefresh(refresh string) (*storage.Token, error) {
	if refresh == "" {
		return nil, errors.New("RefreshToken is mandatory")
	}

	collection := s.client.MD.Collection(TokenCollectionName)

	res := collection.FindOne(nil, bson.M{"refresh": refresh})
	if res.Err() != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetTokenByRefresh",
			"error":    res.Err(),
		}).Debug("mongo request failed")

		return nil, res.Err()
	}

	var token storage.Token
	err := res.Decode(&token)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetTokenByRefresh",
			"error":    res,
		}).Debug("decode doc error")
	}

	return &token, nil
}

// GetTokenByParent returns tokens by RefreshToken
func (s *TokenStorage) GetByParent(parent string) (*storage.Token, error) {
	collection := s.client.MD.Collection(TokenCollectionName)

	res := collection.FindOne(nil, bson.M{"parent": parent})
	if res.Err() != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetTokenByParent",
			"error":    res.Err(),
		}).Debug("mongo request failed")

		return nil, res.Err()
	}

	var token storage.Token
	err := res.Decode(&token)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetTokenByParent",
			"error":    err,
		}).Debug("decode doc error")
	}

	return &token, err
}

// EnsureTokensIndex works with mongo index
func (s *TokenStorage) ensureTokensIndex() error {
	collection := s.client.MD.Collection(TokenCollectionName)

	isEnsured, err := needEnsureTokensIndex(collection)
	if err != nil {
		return err
	}

	if *isEnsured {
		duration := int32(time.Second)
		name := "token_auto_prune"
		options := options.IndexOptions{
			Name: &name,
			ExpireAfterSeconds: &duration,
		}
		_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			//About what mean 1 at value: https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
			Keys: bsonx.Doc{{ Key: "expired", Value: bsonx.Int32(1) }},
			Options: &options,
		})

		if err != nil {
			if v, ok := err.(driver.Error); ok && v.Code == 11000 {
				return nil
			}
		}

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"function": "UserStorage.EnsureRolesIndex",
				"error":    err,
			}).Debug("mongo request failed")
		}

		return err
	}

	return nil
}

func needEnsureTokensIndex(collection *mongo.Collection) (*bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)

	cur, err := collection.Indexes().List(ctx)
	if err != nil {
		logrus.WithError(err).Debug()
		return nil, errors.New(fmt.Sprintf("cat't get tokens collection indexes. err: %s", err))
	}
	defer cur.Close(ctx)

	var isEnsured bool
	for cur.Next(nil) {
		docs := bsonx.Doc{}
		err2 := cur.Decode(&docs)
		if err2 != nil {
			logrus.WithFields(logrus.Fields{
				"function": "needEnsureTokensIndex",
				"error": err2,
			}).Debug("decode doc error")
		}

		for _, doc := range docs {
			if doc.Key ==  "token_auto_prune" {
				isEnsured = false
				return &isEnsured, nil
			}
		}
	}
	isEnsured = true
	return &isEnsured, nil
}
