package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Config *Config

	*mongo.Database
	ctx context.Context
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func Connect(config *Config, ctx context.Context) (*Client, error) {
	var client *mongo.Client
	var err error
	if len(config.Username) > 0 && len(config.Password) > 0 {
		client, err = mongo.NewClient(
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database),
			),
		)
	} else {
		client, err = mongo.NewClient(
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%d/%s", config.Host, config.Port, config.Database),
			),
		)
	}

	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return &Client{
		Config: config,

		Database: client.Database(config.Database),

		ctx: ctx,
	}, nil
}

func (c *Client) Session(fn func(sc mongo.SessionContext) error) (err error) {
	var sess mongo.Session
	if sess, err = c.Client().StartSession(); err != nil {
		return err
	}
	defer sess.EndSession(c.ctx)

	if err = sess.StartTransaction(); err != nil {
		return err
	}

	if err = mongo.WithSession(c.ctx, sess, func(sc mongo.SessionContext) (err error) {
		if err = fn(sc); err != nil {
			sess.AbortTransaction(sc)
			return err
		}

		if err = sess.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}
