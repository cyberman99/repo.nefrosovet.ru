package storage

import "time"

type UserStorage interface {
	StoreUser(in StoreUser) (User, error)
	UpdateUser(id UUID, in UpdateUser) (User, error)
	GetUser(in GetUser) (User, error)
	GetUsers() ([]User, error)
}

type User struct {
	ID UUID `json:"id" bson:"_id"`

	PasswordHash string `json:"-" bson:"passwordHash"`

	FirstName  string `json:"firstName" bson:"firstName"`
	LastName   string `json:"lastName" bson:"lastName"`
	MiddleName string `json:"middleName" bson:"middleName"`

	Contacts []UserContact `json:"contacts" bson:"contacts"`
}

type StoreUser struct {
	ID           string        `bson:"_id"`
	PasswordHash string        `bson:"passwordHash"`
	FirstName    string        `bson:"firstName"`
	LastName     string        `bson:"lastName"`
	MiddleName   string        `bson:"middleName"`
	Contacts     []UserContact `json:"contacts" bson:"contacts"`
}

type UpdateUser struct {
	PasswordHash *string
	FirstName    *string
	LastName     *string
	MiddleName   *string
	Contacts     *[]UserContact `json:"contacts" bson:"contacts"`
}

type GetUser struct {
	ID    *UUID
	Value *string
}

type UserContact struct {
	TypeCODE string     `json:"typeCODE" bson:"typeCODE"`
	Value    string     `json:"value" bson:"value"`
	Verified *time.Time `json:"verified" bson:"verified,omitempty"`
	Created  time.Time  `json:"created" bson:"created,omitempty"`
}
