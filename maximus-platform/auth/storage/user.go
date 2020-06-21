package storage

import (
	"fmt"
)

var (
	ErrUserNotFound      = fmt.Errorf("user: %w", ErrNotFound)
	ErrUserAlreadyExists = fmt.Errorf("user: %w", ErrAlreadyExists)
)

type UserStorage interface {
	Store(in StoreUser) (*User, error)
	Update(id string, in UpdateUser) (*User, error)
	Get(in GetUser) ([]*User, error)
}

type User struct {
	ID string `bson:"id"`
	Roles map[string]bool `bson:"roles,omitempty"`
	BackendEntryIDs map[string]string `bson:"backendEntryIDs,omitempty"`
}

type StoreUser struct {
	User `bson:",inline"`
}

type UpdateUser struct {
	Roles           map[string]bool   `bson:"roles,omitempty"`
	BackendEntryIDs map[string]string `bson:"backendEntryIDs,omitempty"`
}

type GetUser struct {
	ID             *string
	RoleID         *string
	BackendID      *string
	BackendEntryID *string
}
