// Code generated by entc, DO NOT EDIT.

package webinaruser

import (
	"fmt"

	"repo.nefrosovet.ru/go-lms/api-video/ent/schema"
)

const (
	// Label holds the string label denoting the webinaruser type in the database.
	Label = "webinar_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id vertex property in the database.
	FieldUserID = "user_id"
	// FieldWebinarID holds the string denoting the webinar_id vertex property in the database.
	FieldWebinarID = "webinar_id"
	// FieldStatus holds the string denoting the status vertex property in the database.
	FieldStatus = "status"
	// FieldMedoozeID holds the string denoting the medooze_id vertex property in the database.
	FieldMedoozeID = "medooze_id"
	// FieldOldMedoozeID holds the string denoting the old_medooze_id vertex property in the database.
	FieldOldMedoozeID = "old_medooze_id"
	// FieldMic holds the string denoting the mic vertex property in the database.
	FieldMic = "mic"
	// FieldSound holds the string denoting the sound vertex property in the database.
	FieldSound = "sound"

	// Table holds the table name of the webinaruser in the database.
	Table = "webinar_users"
)

// Columns holds all SQL columns for webinaruser fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldWebinarID,
	FieldStatus,
	FieldMedoozeID,
	FieldOldMedoozeID,
	FieldMic,
	FieldSound,
}

var (
	fields = schema.WebinarUser{}.Fields()

	// descMic is the schema descriptor for mic field.
	descMic = fields[5].Descriptor()
	// DefaultMic holds the default value on creation for the mic field.
	DefaultMic = descMic.Default.(int16)

	// descSound is the schema descriptor for sound field.
	descSound = fields[6].Descriptor()
	// DefaultSound holds the default value on creation for the sound field.
	DefaultSound = descSound.Default.(int16)
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusWAIT    Status = "WAIT"
	StatusOFFLINE Status = "OFFLINE"
	StatusONLINE  Status = "ONLINE"
	StatusBANNED  Status = "BANNED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusWAIT, StatusOFFLINE, StatusONLINE, StatusBANNED:
		return nil
	default:
		return fmt.Errorf("webinaruser: invalid enum value for status field: %q", s)
	}
}
