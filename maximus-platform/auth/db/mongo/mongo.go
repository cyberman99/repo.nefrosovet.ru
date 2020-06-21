package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Client struct {
	Database string

	MD *mongo.Database

	*mongo.Client
}

func Connect(host string, port int, username, password, database string) (*Client, error) {
	client := &Client{
		Database: database,
	}

	var err error
	if len(username) > 0 && len(password) > 0 {
		client.Client, err = mongo.NewClient(
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, password, host, port, database),
			),
		)
	} else {
		client.Client, err = mongo.NewClient(
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%d/%s", host, port, database),
			),
		)
	}

	if err != nil {
		return nil, err
	}

	if err = client.Client.Connect(context.Background()); err != nil {
		return nil, err
	}

	client.MD = client.Client.Database(database)

	return client, nil
}
