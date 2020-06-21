package domain

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"time"
)

const ReplyCollectionName = "reply_replaces"

type Reply struct {
	ReplyID     strfmt.UUID `bson:"_id,omitempty"`
	Description *string     `bson:"description,omitempty"`
	Regex       string      `json:"regex" bson:"regex,omitempty"`
	Replace     string      `json:"replace" bson:"replace,omitempty"`
	Created     time.Time   `bson:"created,omitempty"`
	Updated     *time.Time  `bson:"updated,omitempty"`
}

type RepliesFilter struct {
	Limit  int64
	Offset int64
}

var (
	ErrReplyNotFound      = errors.New("reply not found")
	ErrReplyAlreadyExists = errors.New("reply already exist")
)
