package storage

import "gopkg.in/guregu/null.v3"

type ConfirmCode struct {
	BaseModel

	ID null.String `db:"id"`

	PatientID null.String `db:"patient_id"`
	ChannelID null.String `db:"channel_id"`

	TypeCode null.String `db:"type_code"`
	Code     null.String `db:"code"`
}
