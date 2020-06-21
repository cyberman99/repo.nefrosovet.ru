package mongod

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Storer interface {
	Connect(context.Context) error
	Status() error
	ServerStatus() (status *ServerStatus, err error)
	Database() *mongo.Database
	Close() error
	Session(fn func()) error
	Collection(string) *mongo.Collection
	MQTTCredentials() map[string]interface{}
}

type mgo struct {
	dbName  string
	mqCreds map[string]interface{}

	cli  *mongo.Database
	opts *options.ClientOptions
	ctx  context.Context
}

var storage Storer

func GetStorage() Storer {
	if storage == nil {
		log.Fatal("Mongodb session is not found")
	}

	return storage
}

type ServerStatus struct {
	Host           string    `bson:"host"`
	Version        string    `bson:"version"`
	Process        string    `bson:"process"`
	Pid            int64     `bson:"pid"`
	Uptime         int64     `bson:"uptime"`
	UptimeMillis   int64     `bson:"uptimeMillis"`
	UptimeEstimate int64     `bson:"uptimeEstimate"`
	LocalTime      time.Time `bson:"localTime"`
}

func NewCli(
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
	return &m, err
}

func (m *mgo) Connect(ctx context.Context) (err error) {
	m.ctx = ctx
	err = m.cli.Client().Connect(m.ctx)
	if err != nil {
		return err
	}
	err = m.cli.Client().Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}

	storage = m
	return nil
}

func (m *mgo) MQTTCredentials() map[string]interface{} {
	return m.mqCreds
}

func (m *mgo) Status() error {
	return m.cli.Client().Ping(context.Background(), readpref.Primary())

}

func (m *mgo) ServerStatus() (_ *ServerStatus, err error) {
	var result *mongo.SingleResult
	result = m.cli.RunCommand(m.ctx, bson.D{{"serverStatus", 1}})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var status = new(ServerStatus)
	err = result.Decode(status)
	if err != nil {
		return nil, err
	}

	return status, nil
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

func (m *mgo) Session(fn func()) (err error) {
	var sess mongo.Session
	client := m.Database().Client()
	if sess, err = client.StartSession(); err != nil {
		return err
	}
	if err = sess.StartTransaction(); err != nil {
		return err
	}
	if err = mongo.WithSession(m.ctx, sess, func(sc mongo.SessionContext) error {
		//if result, err = fn(); err != nil {
		//	return err
		//} TODO

		if err = sess.AbortTransaction(sc); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	sess.EndSession(m.ctx)
	return nil
}

func (m *mgo) Collection(collName string) *mongo.Collection {
	return m.cli.Collection(collName)
}
