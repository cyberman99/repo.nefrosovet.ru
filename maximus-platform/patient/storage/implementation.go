package storage

import (
	"gopkg.in/guregu/null.v3"
)

type Implementation struct {
	BaseModel

	ID null.String `db:"id"`

	AppointmentID null.String `db:"appointment_id"`

	StatusCode null.String `db:"status_code"`
	Performed  null.Time   `db:"performed"`
}
