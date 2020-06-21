package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
)

type RouteRepository interface {
	Set(client domain.Route) (*domain.Route, error)
	Get(uuid strfmt.UUID) (*domain.Route, error)
	Update(strfmt.UUID, domain.Route) (*domain.Route, error)
	UnsetReplyID(uuid strfmt.UUID) error
	Delete(uuid strfmt.UUID) error
	List(domain.RoutesFilter) ([]domain.Route, error)
}

type routeRepo struct {
	ctx context.Context
	db  *mongo.Collection
}

func NewRouteRepo(store mongod.Storer) RouteRepository {
	return &routeRepo{
		context.Background(),
		store.Collection(domain.RouteCollectionName),
	}
}

func (p *routeRepo) Set(filter domain.Route) (resp *domain.Route, err error) {
	ctx, _ := context.WithTimeout(p.ctx, 1*time.Second)

	filter.Created, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	var guid strfmt.UUID
	err = guid.Scan(uuid.Must(uuid.NewV1()).String())
	if err != nil {
		return nil, err
	}
	filter.RouteID = guid

	resp = &filter

	res, err := p.db.InsertOne(ctx, filter)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, domain.ErrRouteAlreadyExists
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

	resp.RouteID = id
	return resp, nil
}

func (p *routeRepo) Get(id strfmt.UUID) (_ *domain.Route, err error) {
	var result domain.Route
	ctx, _ := context.WithTimeout(p.ctx, 1*time.Second)
	if err := p.db.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrRouteNotFound
		}
		return nil, err
	}
	result.Src.Payload, err = fromDocToMap(result.Src.Payload)
	if err != nil {
		return nil, err
	}
	result.Src.Topic, err = fromDocToMap(result.Src.Topic)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func (p *routeRepo) Update(
	id strfmt.UUID,
	update domain.Route,
) (resp *domain.Route, err error) {
	ctx, _ := context.WithTimeout(p.ctx, 2*time.Second)

	tm, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	update.Updated = &tm

	opts := new(options.FindOneAndUpdateOptions)
	opts.SetUpsert(false)
	opts.SetReturnDocument(options.After)

	err = p.db.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": update}, opts).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrRouteNotFound
		}
		return nil, err
	}

	resp.Src.Payload, err = fromDocToMap(resp.Src.Payload)
	if err != nil {
		return nil, err
	}
	resp.Src.Topic, err = fromDocToMap(resp.Src.Topic)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *routeRepo) UnsetReplyID(id strfmt.UUID) (err error) {
	ctx, _ := context.WithTimeout(p.ctx, 2*time.Second)

	err = p.db.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$unset": bson.M{"reply_id": nil}}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrRouteNotFound
		}
		return err
	}

	return nil
}

func (p *routeRepo) Delete(id strfmt.UUID) error {
	ctx, _ := context.WithTimeout(p.ctx, 1*time.Second)

	err := p.db.FindOneAndDelete(ctx, bson.M{"_id": id}, &options.FindOneAndDeleteOptions{}).Err()

	if err == mongo.ErrNoDocuments {
		return domain.ErrClientNotFound
	}

	return err
}

func (p *routeRepo) List(filter domain.RoutesFilter) (_ []domain.Route, err error) {
	var rs = make([]domain.Route, 0)
	ctx, _ := context.WithTimeout(p.ctx, 5*time.Second)

	opts := new(options.FindOptions)
	opts.SetLimit(filter.Limit)
	opts.SetSkip(filter.Offset)

	doc := bson.M{}
	if filter.ReplyID != nil {
		doc = bson.M{"reply_id": *filter.ReplyID}
	}

	cur, err := p.db.Find(ctx, doc, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.Route
		if err = cur.Decode(&result); err != nil {
			return nil, err
		}

		result.Src.Payload, err = fromDocToMap(result.Src.Payload)
		if err != nil {
			return nil, err
		}
		result.Src.Topic, err = fromDocToMap(result.Src.Topic)
		if err != nil {
			return nil, err
		}

		rs = append(rs, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	if len(rs) == 0 {
		return rs, domain.ErrRouteNotFound
	}

	return rs, nil
}

func fromDocToMap(doc interface{}) (map[string]interface{}, error) {
	switch d := doc.(type) {
	case nil:
		return map[string]interface{}{}, nil
	case primitive.D:
		return mToMap(d.Map())
	case primitive.A:
		var bigD primitive.D
		for _, v := range d {
			bigD = append(bigD, v.(primitive.D)...)
		}
		return mToMap(bigD.Map())
	}
	return mToMap(doc.(bson.M))
}

func mToMap(m primitive.M) (result map[string]interface{}, err error) {
	var bt []byte
	bt, err = bson.MarshalExtJSON(m, false, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bt, &result)
	return
}
