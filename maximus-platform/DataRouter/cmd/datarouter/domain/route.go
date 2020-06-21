package domain

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"time"
)

const RouteCollectionName = "routes"

type Route struct {
	RouteID strfmt.UUID    `bson:"_id,omitempty"`
	ReplyID *strfmt.UUID   `bson:"reply_id,omitempty"`
	Dst     []Destinations `bson:"dst,omitempty"`
	Src     Source         `bson:"src,omitempty"`
	Created time.Time      `bson:"created,omitempty"`
	Updated *time.Time     `bson:"updated,omitempty"`
}

type Destinations struct {
	Qos   byte   `json:"qos" bson:"qos,omitempty"`
	Topic string `json:"topic" bson:"topic,omitempty"`
}

type Source struct {
	Payload interface{} `bson:"payload,omitempty"`
	Topic   interface{} `bson:"topic,omitempty"`
}

type RoutesFilter struct {
	ReplyID *strfmt.UUID
	Limit   int64
	Offset  int64
}

var (
	ErrRouteNotFound      = errors.New("route not found")
	ErrRouteAlreadyExists = errors.New("route already exist")
)
