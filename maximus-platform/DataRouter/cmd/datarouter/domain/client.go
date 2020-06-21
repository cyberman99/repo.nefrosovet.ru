package domain

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"time"
)

const ClientCollectionName = "vmq_acl_auth"

type Client struct {
	ClientID    string     `bson:"client_id,omitempty"`
	Passhash    string     `bson:"passhash,omitempty"`
	Username    string     `bson:"username,omitempty"`
	Created     time.Time  `bson:"created,omitempty"`
	Updated     *time.Time `bson:"updated,omitempty"`
	TTL         *int64     `bson:"ttl,omitempty"`
	Expired     *time.Time `bson:"expired,omitempty"`
	System      bool       `bson:"system"`
	Mountpoint  string     `bson:"mountpoint"`
	Permissions `bson:"inline"`
}

type ClientsFilter struct {
	Limit    int64
	Offset   int64
	Username string `bson:"username,omitempty"`
}

type Permissions struct {
	Publish   []Acl `bson:"publish_acl"`
	Subscribe []Acl `bson:"subscribe_acl"`
}

type Acl struct {
	Pattern string `bson:"pattern" json:"pattern"`
}

type PermissionsFilter struct {
	Limit    int64
	Offset   int64
	ClientID strfmt.UUID `bson:"client_id"`
}

var (
	ErrClientNotFound      = errors.New("client not found")
	ErrClientAlreadyExists = errors.New("client already exist")
	ErrPermissionNotFound  = errors.New("permission not found")
)
