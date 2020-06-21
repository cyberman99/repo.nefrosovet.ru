package repos

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
)

type ClientRepository interface {
	Set(client domain.Client) (*domain.Client, error)
	Get(uuid strfmt.UUID) (*domain.Client, error)
	Update(strfmt.UUID, domain.Client) (*domain.Client, error)
	Delete(uuid strfmt.UUID) error
	List(domain.ClientsFilter) ([]domain.Client, error)
	SetOrReplace(clientID strfmt.UUID, filter domain.Permissions) (resp *domain.Permissions, err error)
	GetPermissions(clientID strfmt.UUID) (_ *domain.Permissions, err error)
}

type clientRepo struct {
	ctx     context.Context
	db      *mongo.Collection
	mqCreds map[string]interface{}
}

func NewClientRepo(store mongod.Storer, mqCreds map[string]interface{}) ClientRepository {
	cr := clientRepo{
		context.Background(),
		store.Collection(domain.ClientCollectionName),
		mqCreds,
	}
	cr.initExpirationIndex()
	cr.initUniqueIndex()
	cr.storeSystemClient()
	return &cr
}

func (c *clientRepo) storeSystemClient() { // dirty devops hack
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("can't store system client. Improper config flags", r)
		}
	}()
	var (
		err      error
		passhash []byte
		login    string
	)
	if c.mqCreds == nil {
		return
	}

	passhash, err = bcrypt.GenerateFromPassword(
		[]byte(c.mqCreds["password"].(string)), 10,
	)
	if err != nil {
		log.Fatal(err)
	}
	login = c.mqCreds["login"].(string)

	_, err = c.db.UpdateOne(
		c.ctx,
		bson.M{"client_id": c.mqCreds["subClientID"].(string)},
		bson.M{"$set": domain.Client{
			Username: login,
			Passhash: string(passhash),
			System:   true},
		}, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
	}

	subscribe := make([]domain.Acl, len(c.mqCreds["subscribe"].([]string)))
	for i, s := range c.mqCreds["subscribe"].([]string) {
		subscribe[i] = domain.Acl{s}
	}

	publish := make([]domain.Acl, len(c.mqCreds["publish"].([]string)))
	for i, s := range c.mqCreds["publish"].([]string) {
		publish[i] = domain.Acl{s}
	}

	_, err = c.SetOrReplace(
		strfmt.UUID(c.mqCreds["subClientID"].(string)),
		domain.Permissions{
			make([]domain.Acl, 0),
			subscribe,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.db.UpdateOne(
		c.ctx,
		bson.M{"client_id": c.mqCreds["pubClientID"].(string)},
		bson.M{"$set": domain.Client{
			Username: login,
			Passhash: string(passhash),
			System:   true},
		}, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
	}
	_, err = c.SetOrReplace(
		strfmt.UUID(c.mqCreds["pubClientID"].(string)),
		domain.Permissions{
			publish,
			make([]domain.Acl, 0),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *clientRepo) initExpirationIndex() {
	var (
		index     mongo.IndexModel
		doc       bsonx.Doc
		err       error
		indexName = "client_autoprune"
		cur       *mongo.Cursor
	)

	cur, err = c.db.Indexes().List(c.ctx, &options.ListIndexesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(c.ctx) {
		err = cur.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}
		for _, elem := range doc {
			if elem.Value.String() == indexName {
				return
			}
		}
	}

	keys := bsonx.Doc{{Key: "expired", Value: bsonx.Int32(int32(1))}}
	index.Keys = keys
	index.Options = new(options.IndexOptions)
	index.Options.SetExpireAfterSeconds(0)
	index.Options.SetName(indexName)
	if _, err = c.db.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal(err)
	}
}

func (c *clientRepo) initUniqueIndex() {
	var (
		noErr            = true
		index            mongo.IndexModel
		doc              bsonx.Doc
		err              error
		indexReplaceName = "unique_client"
		cur              *mongo.Cursor
	)
	defer func() {
		if !noErr {
			c.db.Indexes().DropAll(c.ctx) // rollback
		}
	}()

	cur, err = c.db.Indexes().List(c.ctx, &options.ListIndexesOptions{})
	if err != nil {
		log.Fatal(err)
		noErr = false
	}
	for cur.Next(c.ctx) {
		err = cur.Decode(&doc)
		if err != nil {
			log.Fatal(err)
			noErr = false
		}
		for _, elem := range doc {
			if elem.Value.String() == indexReplaceName {
				return
			}
		}
	}
	keys := bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(int32(1))}}
	index.Keys = keys
	index.Options = new(options.IndexOptions)
	index.Options.SetName(indexReplaceName)
	index.Options.SetUnique(true)
	if _, err := c.db.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal(err)
	}
}

func (c *clientRepo) Set(filter domain.Client) (resp *domain.Client, err error) {
	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)
	now := time.Now()

	filter.Created, err = time.Parse(time.RFC3339, now.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	if filter.TTL != nil {
		if *filter.TTL != 0 {
			expiration := now.Add(time.Duration(*filter.TTL))
			filter.Expired = &expiration
		} else {
			filter.TTL = nil
		}
	}

	if filter.ClientID == "" {
		var guid strfmt.UUID
		err = guid.Scan(uuid.Must(uuid.NewV1()).String())
		if err != nil {
			return nil, err
		}
		filter.ClientID = guid.String()
	}

	filter.Subscribe = make([]domain.Acl, 0)
	filter.Publish = make([]domain.Acl, 0)

	resp = &filter

	_, err = c.db.InsertOne(ctx, filter)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, domain.ErrClientAlreadyExists
		}
		return nil, err
	}

	return resp, nil
}

func (c *clientRepo) Get(guid strfmt.UUID) (_ *domain.Client, err error) {
	var result domain.Client

	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)
	if err := c.db.FindOne(
		ctx,
		bson.M{
			"client_id": guid.String(),
			"system":    false,
		},
	).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrClientNotFound
		}

		return nil, err
	}
	return &result, nil

}

func (c *clientRepo) Update(
	id strfmt.UUID,
	update domain.Client,
) (resp *domain.Client, err error) {
	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)

	tm, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	update.Updated = &tm

	opts := new(options.FindOneAndUpdateOptions)
	opts.SetUpsert(false)
	opts.SetReturnDocument(options.After)

	err = c.db.FindOneAndUpdate(
		ctx,
		bson.M{"client_id": id.String()},
		bson.M{"$set": update},
		opts).
		Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrClientNotFound
		}
		return nil, err
	}

	return resp, nil
}

func (c *clientRepo) Delete(uuid strfmt.UUID) (err error) {
	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)

	err = c.db.FindOneAndDelete(
		ctx,
		bson.M{"client_id": uuid.String()},
		&options.FindOneAndDeleteOptions{},
	).Err()

	if err == mongo.ErrNoDocuments {
		return domain.ErrClientNotFound
	}

	return err
}

func (c *clientRepo) List(filter domain.ClientsFilter) (_ []domain.Client, err error) {
	var clies = make([]domain.Client, 0)
	ctx, _ := context.WithTimeout(c.ctx, 5*time.Second)

	opts := new(options.FindOptions)
	opts.SetLimit(filter.Limit)
	opts.SetSkip(filter.Offset)

	doc := bson.M{"system": false}
	if filter.Username != "" {
		doc = bson.M{
			"username": filter.Username,
			"system":   false,
		}
	}
	cur, err := c.db.Find(ctx, doc, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.Client
		if err = cur.Decode(&result); err != nil {
			return nil, err
		}
		clies = append(clies, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	if len(clies) == 0 {
		return clies, domain.ErrClientNotFound
	}

	return clies, nil
}

func (c *clientRepo) SetOrReplace(clientID strfmt.UUID, filter domain.Permissions) (resp *domain.Permissions, err error) {
	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)

	err = c.db.FindOneAndUpdate(
		ctx,
		bson.M{"client_id": clientID.String()},
		bson.M{"$set": filter},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clientRepo) GetPermissions(clientID strfmt.UUID) (_ *domain.Permissions, err error) {
	ctx, _ := context.WithTimeout(c.ctx, 1*time.Second)

	var perms *domain.Permissions
	err = c.db.FindOne(ctx, bson.M{"client_id": clientID.String()}).Decode(&perms)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrPermissionNotFound
	}
	if err != nil {
		return nil, err
	}

	return perms, nil
}
