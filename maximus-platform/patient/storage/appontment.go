package storage

import (
	"gopkg.in/guregu/null.v3"
)

type Appointment struct {
	BaseModel

	ID null.String `db:"id"`

	PatientID null.String `db:"patient_id"`
	ClinicID  null.String `db:"clinic_id"`
	ProgramID null.String `db:"program_id"`
	DoctorID  null.String `db:"doctor_id"`

	TypeCode   null.String `db:"type_code"`
	StatusCode null.String `db:"status_code"`

	Planned   null.Time `db:"planned"`
	Performed null.Time `db:"performed"`
	Duration  null.Int  `db:"duration"`

	Comment null.String `db:"comment"`
}
