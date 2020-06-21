package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

const (
	channelCollection = "channels"
)

type ChannelStorage struct {
	client *dbMongo.Client

	storage.ChannelStorage
}

func (s *ChannelStorage) StoreChannel(in storage.StoreChannel) (storage.Channel, error) {
	collection := s.client.MD.Collection(channelCollection)

	_, err := collection.InsertOne(nil, &in.Channel)
	if err != nil {
		switch err := err.(type) {
		case mongo.WriteException:
			if len(err.WriteErrors) >= 1 && err.WriteErrors[0].Code == 11000 {
				return storage.Channel{}, storage.ErrChannelAlreadyExists
			}
		}

		return storage.Channel{}, err
	}

	return in.Channel, nil
}

func (s *ChannelStorage) GetChannel(in storage.GetChannel) (storage.Channel, error) {
	collection := s.client.MD.Collection(channelCollection)

	selector := bson.M{}

	if in.ID != nil {
		selector["uuid"] = *in.ID
	}

	if in.Params.Token != nil {
		selector["params.token"] = *in.Params.Token
	}

	sr := collection.FindOne(nil, selector)
	if sr.Err() != nil {
		return storage.Channel{}, sr.Err()
	}

	var channel storage.Channel
	err := sr.Decode(&channel)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.Channel{}, storage.ErrChannelNotFound
		}

		return storage.Channel{}, err
	}

	return channel, nil
}

func (s *ChannelStorage) GetChannels(in storage.GetChannels) ([]storage.Channel, error) {
	collection := s.client.MD.Collection(channelCollection)

	selector := bson.M{}

	if in.Type != nil {
		selector["type"] = string(*in.Type)
	}

	if in.AccessToken != nil {
		selector["access_token"] = *in.AccessToken
	}

	cur, err := collection.Find(nil, selector)
	if err != nil {
		return nil, err
	}

	defer cur.Close(nil)

	var channels []storage.Channel
	for cur.Next(nil) {
		var channel storage.Channel

		err := cur.Decode(&channel)
		if err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}

	return channels, nil
}

func (s *ChannelStorage) UpdateChannel(in storage.UpdateChannel) (storage.Channel, error) {
	collection := s.client.MD.Collection(channelCollection)

	r := collection.FindOneAndUpdate(nil,
		bson.M{"uuid": in.ID},
		bson.M{"$set": &in.Channel},
	)
	if r.Err() != nil {
		return storage.Channel{}, r.Err()
	}

	var c storage.Channel
	if err := r.Decode(&c); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.Channel{}, storage.ErrChannelNotFound
		}

		return storage.Channel{}, err
	}

	return in.Channel, nil
}

func (s *ChannelStorage) DeleteChannel(in storage.DeleteChannel) (storage.Channel, error) {
	collection := s.client.MD.Collection(channelCollection)

	r := collection.FindOneAndDelete(nil, bson.M{"uuid": in.ID})
	if r.Err() != nil {
		return storage.Channel{}, r.Err()
	}

	var c storage.Channel
	if err := r.Decode(&c); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return storage.Channel{}, storage.ErrChannelNotFound
		}

		return storage.Channel{}, err
	}

	return c, nil
}
