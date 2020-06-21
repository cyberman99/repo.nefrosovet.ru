package storage

import "time"

type UserContactStorage interface {
	StoreUserContact(in StoreUserContact) (UserContact, error)
	UpdateUserContact(userID UUID, in UpdateUserContact) (UserContact, error)
	GetUserContacts(in GetUserContacts) ([]UserContact, error)
	GetUserContact(in GetUserContact) (UserContact, error)
}



type StoreUserContact struct {
	ID     string `bson:"_id"`
	UserID UUID   `json:"userID" bson:"userID"`

	TypeCODE string    `json:"typeCODE" bson:"typeCODE"`
	Value    string    `json:"value" bson:"value"`
	Created  time.Time `json:"created" bson:"created"`
}

type UpdateUserContact struct {
	TypeCODE *string
	Value    *string
	Verified *bool
}

type GetUserContacts struct {
	UserID UUID
}

type GetUserContact struct {
	ID    *UUID
	Value *string
}
