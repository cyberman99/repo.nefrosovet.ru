// Package local_sms_db contains the types for schema 'sms'.
package localsmsdb

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Coding is the 'Coding' enum type from schema 'sms'.
type Coding uint16

const (
	// CodingUnicodeNoCompression is the 'Unicode_No_Compression' Coding.
	CodingUnicodeNoCompression = Coding(1)

	// CodingBit is the '8bit' Coding.
	CodingBit = Coding(2)

	// CodingDefaultCompression is the 'Default_Compression' Coding.
	CodingDefaultCompression = Coding(3)

	// CodingUnicodeCompression is the 'Unicode_Compression' Coding.
	CodingUnicodeCompression = Coding(4)
)

// String returns the string value of the Coding.
func (c Coding) String() string {
	var enumVal string

	switch c {
	case CodingUnicodeNoCompression:
		enumVal = "Unicode_No_Compression"

	case CodingBit:
		enumVal = "8bit"

	case CodingDefaultCompression:
		enumVal = "Default_Compression"

	case CodingUnicodeCompression:
		enumVal = "Unicode_Compression"
	}

	return enumVal
}

// MarshalText marshals Coding into text.
func (c Coding) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// UnmarshalText unmarshals Coding from text.
func (c *Coding) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Unicode_No_Compression":
		*c = CodingUnicodeNoCompression

	case "8bit":
		*c = CodingBit

	case "Default_Compression":
		*c = CodingDefaultCompression

	case "Unicode_Compression":
		*c = CodingUnicodeCompression

	default:
		return errors.New("invalid Coding")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Coding.
func (c Coding) Value() (driver.Value, error) {
	return c.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Coding.
func (c *Coding) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Coding")
	}

	return c.UnmarshalText(buf)
}
