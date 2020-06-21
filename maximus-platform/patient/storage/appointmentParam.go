package storage

import "gopkg.in/guregu/null.v3"

type AppointmentParam struct {
	BaseModel

	ID null.String `db:"id"`

	AppointmentID null.String `db:"appointment_id"`

	TypeCode null.String `db:"type_code"`
	Value    null.String `db:"value"`
}
