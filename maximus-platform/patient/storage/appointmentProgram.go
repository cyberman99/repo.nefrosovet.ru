package storage

import (
	"gopkg.in/guregu/null.v3"
)

type AppointmentProgram struct {
	ID null.String

	DoctorID null.String `db:"doctor_id"`

	TypeCode   null.String `db:"type_code"`
	StatusCode null.String `db:"status_code"`

	Begin null.Time `db:"begin"`
	End   null.Time `db:"end"`

	Comment     null.String `db:"comment"`
	Periodicity null.String `db:"periodicity"`
}
