package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Storer interface {
	Status() error
	Database() *mongo.Database
	Close() error
	Collection(string) *mongo.Collection
}

type mgo struct {
	dbName string

	cli  *mongo.Database
	opts *options.ClientOptions
	ctx  context.Context
}

func New(
	host string,
	port int,
	login string,
	password string,
	database string,
) (_ Storer, err error) {
	var m mgo
	m.opts = options.Client()
	m.dbName = database

	m.opts.SetDirect(true)
	m.opts.SetConnectTimeout(2 * time.Second)
	m.opts.SetSocketTimeout(2 * time.Second)
	m.opts.SetMaxConnIdleTime(1 * time.Second)
	m.opts.SetMaxPoolSize(10)

	uri := fmt.Sprintf("mongodb://@%s:%d", host, port)

	if login != "" && password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", login, password, host, port, database)
	}
	m.opts.ApplyURI(uri)

	var cli *mongo.Client
	cli, err = mongo.NewClient(m.opts)
	if err != nil {
		return nil, err
	}

	m.cli = cli.Database(m.dbName)
	err = m.cli.Client().Connect(m.ctx)
	if err != nil {
		return nil, err
	}
	err = m.cli.Client().Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"context":  "CORE",
		"resource": "configDB",
		"addr":     host + ":" + strconv.Itoa(port),
		"status":   "CONNECTED",
	}).Info("Connection to mongoDB")
	return &m, err
}

func (m *mgo) Status() error {
	return m.cli.Client().Ping(context.Background(), readpref.Primary())
}

func (m *mgo) Database() *mongo.Database {
	return m.cli
}

func (m *mgo) Close() (err error) {
	err = m.cli.Client().Disconnect(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *mgo) Collection(collName string) *mongo.Collection {
	return m.cli.Collection(collName)
}
