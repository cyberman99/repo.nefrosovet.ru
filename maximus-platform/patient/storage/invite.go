package storage

import "gopkg.in/guregu/null.v3"

type Invite struct {
	BaseModel

	ID null.String `db:"id"`

	PatientID null.String `db:"patient_id"`
	ChannelID null.String `db:"channel_id"`

	StatusCode null.String `db:"status_code"`
}
