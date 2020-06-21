package repos

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
)

type ReplyRepository interface {
	Set(client domain.Reply) (*domain.Reply, error)
	Get(uuid strfmt.UUID) (*domain.Reply, error)
	Update(strfmt.UUID, domain.Reply) (resp *domain.Reply, err error)
	Delete(uuid strfmt.UUID) error
	List(domain.RepliesFilter) ([]domain.Reply, error)
}

type replyRepo struct {
	ctx context.Context
	db  *mongo.Collection
}

func NewReplyRepo(store mongod.Storer) ReplyRepository {
	rr := &replyRepo{
		context.Background(),
		store.Collection(domain.ReplyCollectionName),
	}
	rr.initIndex()
	return rr
}

func (r *replyRepo) initIndex() {
	var (
		noErr            = true
		index            mongo.IndexModel
		doc              bsonx.Doc
		err              error
		indexReplaceName = "unique_regex"
		cur              *mongo.Cursor
	)
	defer func() {
		if !noErr {
			r.db.Indexes().DropAll(r.ctx)
		}
	}()

	cur, err = r.db.Indexes().List(r.ctx, &options.ListIndexesOptions{})
	if err != nil {
		log.Fatal(err)
		noErr = false
	}
	for cur.Next(r.ctx) {
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
	keys := bsonx.Doc{{Key: "regex", Value: bsonx.Int32(int32(1))}}
	index.Keys = keys
	index.Options = new(options.IndexOptions)
	index.Options.SetName(indexReplaceName)
	index.Options.SetUnique(true)
	if _, err := r.db.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal(err)
	}
}

func (r *replyRepo) Set(filter domain.Reply) (resp *domain.Reply, err error) {
	ctx, _ := context.WithTimeout(r.ctx, 1*time.Second)

	filter.Created, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	var guid strfmt.UUID
	err = guid.Scan(uuid.Must(uuid.NewV1()).String())
	if err != nil {
		return nil, err
	}
	filter.ReplyID = guid

	resp = &filter

	res, err := r.db.InsertOne(ctx, filter)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, domain.ErrReplyAlreadyExists
		}
		return nil, err
	}

	var id strfmt.UUID
	switch res.InsertedID.(type) {
	case string:
		err = id.Scan(res.InsertedID)
	case primitive.ObjectID:
		err = id.Scan(
			fmt.Sprintf("%q", res.InsertedID.(primitive.ObjectID).Hex()),
		)
	case primitive.D:
		err = id.Scan(res.InsertedID.(primitive.D).Map()["data"])
	}
	if err != nil {
		return nil, err
	}

	resp.ReplyID = id

	return resp, nil
}

func (r *replyRepo) Get(uuid strfmt.UUID) (_ *domain.Reply, err error) {
	var result domain.Reply
	ctx, _ := context.WithTimeout(r.ctx, 1*time.Second)
	if err := r.db.FindOne(ctx, bson.M{"_id": uuid}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrReplyNotFound
		}

		return nil, err
	}
	return &result, nil
}

func (r *replyRepo) Update(
	id strfmt.UUID,
	update domain.Reply,
) (resp *domain.Reply, err error) {
	ctx, _ := context.WithTimeout(r.ctx, 1*time.Second)

	tm, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	update.Updated = &tm

	opts := new(options.FindOneAndUpdateOptions)
	opts.SetUpsert(false)
	opts.SetReturnDocument(options.After)

	err = r.db.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": update}, opts).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrReplyNotFound
		}
		return nil, err
	}

	return resp, nil
}

func (r *replyRepo) Delete(uuid strfmt.UUID) error {
	ctx, _ := context.WithTimeout(r.ctx, 1*time.Second)

	err := r.db.FindOneAndDelete(ctx, bson.M{"_id": uuid}, &options.FindOneAndDeleteOptions{}).Err()

	if err == mongo.ErrNoDocuments {
		return domain.ErrReplyNotFound
	}

	return err
}

func (r *replyRepo) List(filter domain.RepliesFilter) (_ []domain.Reply, err error) {
	var replies = make([]domain.Reply, 0)
	ctx, _ := context.WithTimeout(r.ctx, 5*time.Second)

	opts := new(options.FindOptions)
	opts.SetLimit(filter.Limit)
	opts.SetSkip(filter.Offset)
	cur, err := r.db.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.Reply
		if err = cur.Decode(&result); err != nil {
			return nil, err
		}
		replies = append(replies, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	if len(replies) == 0 {
		return replies, domain.ErrReplyNotFound
	}

	return replies, nil
}
