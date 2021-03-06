// Package local_sms_db contains the types for schema 'sms'.
package localsmsdb

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Multipart is the 'MultiPart' enum type from schema 'sms'.
type Multipart uint16

const (
	// MultipartFalse is the 'false' Multipart.
	MultipartFalse = Multipart(1)

	// MultipartTrue is the 'true' Multipart.
	MultipartTrue = Multipart(2)
)

// String returns the string value of the Multipart.
func (m Multipart) String() string {
	var enumVal string

	switch m {
	case MultipartFalse:
		enumVal = "false"

	case MultipartTrue:
		enumVal = "true"
	}

	return enumVal
}

// MarshalText marshals Multipart into text.
func (m Multipart) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText unmarshals Multipart from text.
func (m *Multipart) UnmarshalText(text []byte) error {
	switch string(text) {
	case "false":
		*m = MultipartFalse

	case "true":
		*m = MultipartTrue

	default:
		return errors.New("invalid Multipart")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Multipart.
func (m Multipart) Value() (driver.Value, error) {
	return m.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Multipart.
func (m *Multipart) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Multipart")
	}

	return m.UnmarshalText(buf)
}
