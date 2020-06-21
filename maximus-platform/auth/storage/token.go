package storage

import "time"

type TokenStorage interface {
	Store(token Token) error
	Update(token Token) error
	Delete(refresh string) error
	GetByRefresh(refresh string) (*Token, error)
	GetByParent(refresh string) (*Token, error)
}

// Token is a struct with information about access Access
type Token struct {
	Refresh  string     `bson:"refresh" json:"refresh"`
	Access   string     `bson:"access" json:"access"`
	Username string     `bson:"username" json:"username"`
	Parent   string     `bson:"parent,omitempty" json:"parent,omitempty"`
	Expired  *time.Time `bson:"expired" json:"expired"`
}
