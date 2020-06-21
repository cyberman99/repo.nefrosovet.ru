package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

const (
	accessTokenCollection = "access_tokens"
)

type AccessTokenStorage struct {
	client *dbMongo.Client

	storage.AccessTokenStorage
}

func (s *AccessTokenStorage) StoreAccessToken(in storage.StoreAccessToken) (storage.AccessToken, error) {
	collection := s.client.MD.Collection(accessTokenCollection)

	_, err := collection.InsertOne(nil, &in.AccessToken)
	if err != nil {
		switch err := err.(type) {
		case mongo.WriteException:
			if len(err.WriteErrors) >= 1 && err.WriteErrors[0].Code == 11000 {
				return storage.AccessToken{}, storage.ErrTelegramContactAlreadyExists
			}
		}

		return storage.AccessToken{}, err
	}

	return in.AccessToken, nil
}

func (s *AccessTokenStorage) GetAccessToken(in storage.GetAccessToken) (storage.AccessToken, error) {
	collection := s.client.MD.Collection(accessTokenCollection)

	selector := bson.M{
		"token": in.Token,
	}

	r := collection.FindOne(nil, selector)
	if r.Err() != nil {
		return storage.AccessToken{}, r.Err()
	}

	var accessToken storage.AccessToken
	if err := r.Decode(&accessToken); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.AccessToken{}, storage.ErrAccessTokenNotFound
		}

		return storage.AccessToken{}, err
	}

	return accessToken, nil
}

func (s *AccessTokenStorage) GetAccessTokens(in storage.GetAccessTokens) ([]storage.AccessToken, error) {
	collection := s.client.MD.Collection(accessTokenCollection)

	opts := options.Find()

	if in.Limit != nil {
		opts.SetLimit(int64(*in.Limit))
	}

	if in.Offset != nil {
		opts.SetSkip(int64(*in.Offset))
	}

	cur, err := collection.Find(nil, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	defer cur.Close(nil)

	var accessTokens []storage.AccessToken
	for cur.Next(nil) {
		var accessToken storage.AccessToken

		err := cur.Decode(&accessToken)
		if err != nil {
			return nil, err
		}

		accessTokens = append(accessTokens, accessToken)
	}

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	return accessTokens, nil
}

func (s *AccessTokenStorage) UpdateAccessToken(in storage.UpdateAccessToken) (storage.AccessToken, error) {
	collection := s.client.MD.Collection(accessTokenCollection)

	r := collection.FindOneAndUpdate(nil,
		bson.M{"token": in.Token},
		bson.M{"$set": &in.AccessToken},
	)
	if r.Err() != nil {
		return storage.AccessToken{}, r.Err()
	}

	var ac storage.AccessToken
	if err := r.Decode(&ac); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.AccessToken{}, storage.ErrAccessTokenNotFound
		}

		return storage.AccessToken{}, err
	}

	return in.AccessToken, nil
}

func (s *AccessTokenStorage) DeleteAccessToken(in storage.DeleteAccessToken) (storage.AccessToken, error) {
	collection := s.client.MD.Collection(accessTokenCollection)

	r := collection.FindOneAndDelete(nil, bson.M{"token": in.Token})
	if r.Err() != nil {
		return storage.AccessToken{}, r.Err()
	}

	var ac storage.AccessToken
	if err := r.Decode(&ac); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.AccessToken{}, storage.ErrAccessTokenNotFound
		}

		return storage.AccessToken{}, err
	}

	return ac, nil
}
