// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"repo.nefrosovet.ru/go-lms/api-video/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMetaData holds the string denoting the meta_data vertex property in the database.
	FieldMetaData = "meta_data"

	// Table holds the table name of the user in the database.
	Table = "users"
	// SubscriberTable is the table the holds the subscriber relation/edge.
	SubscriberTable = "subscribers"
	// SubscriberInverseTable is the table name for the Subscriber entity.
	// It exists in this package in order to avoid circular dependency with the "subscriber" package.
	SubscriberInverseTable = "subscribers"
	// SubscriberColumn is the table column denoting the subscriber relation/edge.
	SubscriberColumn = "user_subscriber"
	// UseraccountTable is the table the holds the useraccount relation/edge.
	UseraccountTable = "user_accounts"
	// UseraccountInverseTable is the table name for the UserAccount entity.
	// It exists in this package in order to avoid circular dependency with the "useraccount" package.
	UseraccountInverseTable = "user_accounts"
	// UseraccountColumn is the table column denoting the useraccount relation/edge.
	UseraccountColumn = "user_useraccount"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMetaData,
}

var (
	fields = schema.User{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[0].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)
)
