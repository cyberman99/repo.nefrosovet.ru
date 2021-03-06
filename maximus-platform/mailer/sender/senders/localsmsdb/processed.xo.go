// Package local_sms_db contains the types for schema 'sms'.
package localsmsdb

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Processed is the 'Processed' enum type from schema 'sms'.
type Processed uint16

const (
	// ProcessedFalse is the 'false' Processed.
	ProcessedFalse = Processed(1)

	// ProcessedTrue is the 'true' Processed.
	ProcessedTrue = Processed(2)
)

// String returns the string value of the Processed.
func (p Processed) String() string {
	var enumVal string

	switch p {
	case ProcessedFalse:
		enumVal = "false"

	case ProcessedTrue:
		enumVal = "true"
	}

	return enumVal
}

// MarshalText marshals Processed into text.
func (p Processed) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// UnmarshalText unmarshals Processed from text.
func (p *Processed) UnmarshalText(text []byte) error {
	switch string(text) {
	case "false":
		*p = ProcessedFalse

	case "true":
		*p = ProcessedTrue

	default:
		return errors.New("invalid Processed")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Processed.
func (p Processed) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Processed.
func (p *Processed) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Processed")
	}

	return p.UnmarshalText(buf)
}
