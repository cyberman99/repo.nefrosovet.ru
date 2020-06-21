package storage

import (
	"fmt"
)

var (
	ErrAdminPasswordNotFound = fmt.Errorf("admin password: %w", ErrNotFound)
	ErrAdminPasswordAlreadyExists = fmt.Errorf("admin password: %w", ErrAlreadyExists)
)

type AdminPasswordStorage interface {
	Store(in StoreAdminPassword) (*AdminPassword, error)
	Update(in UpdateAdminPassword) (*AdminPassword, error)
	Get() ([]*AdminPassword, error)
}

type AdminPassword struct {
	ID   string `json:"id" bson:"id"`
	Hash string `json:"hash" bson:"hash"`
}

type StoreAdminPassword struct {
	Hash string `bson:"hash"`
}

type UpdateAdminPassword struct {
	Hash string `bson:"hash"`
}
