package storage

import (
	"gopkg.in/guregu/null.v3"
)

type BaseModel struct {
	Created null.Time `db:"created"`
	Updated null.Time `db:"updated"`
}

const (
	GUID                     = "guid"
	PatientGUID              = "PatientGUID"
	FirstName                = "FirstName"
	LastName                 = "LastName"
	Patronymic               = "Patronymic"
	BirthDate                = "BirthDate"
	VaTypeCode               = "VaTypeCode"
	VaSideCode               = "VaSideCode"
	LocationCode             = "LocationCode"
	PositionCode             = "PositionCode"
	PhotoGUID                = "PhotoGUID"
	Date                     = "Date"
	Comment                  = "Comment"
	DeviationCodes           = "DeviationCodes"
	Photos                   = "PhotoIDs"
	FromStarted              = "FromStarted"
	ToStarted                = "ToStarted"
	ClassCode                = "ClassCode"
	TypeCode                 = "TypeCode"
	ProcedureGUID            = "ProcedureID"
	StatusCode               = "StatusCode"
	Begin                    = "Begin"
	End                      = "End"
	PatientStatusCode        = "PatientStatusCode"
	NurseGUID                = "NurseGUID"
	DoctorGUID               = "DoctorGUID"
	TreatmentEpisodeTypeCode = "TreatmentEpisodeTypeCode"
	Duration                 = "Duration"
	ServiceTypeCode          = "ServiceTypeCode"
	DryWeight                = "DryWeight"
	WristBandLocationCode    = "WristbandLocationCode"
	MachineTypeCode          = "MachineTypeCode"
	MachineStatusCode        = "MachineStatusCode"
	AcTypeCode               = "AcTypeCode"
	AcInjectionTypeCode      = "AcInjectionTypeCode"
	AcInjectionStatusCode    = "AcInjectionStatusCode"
	Abnormal                 = "Abnormal"
)
