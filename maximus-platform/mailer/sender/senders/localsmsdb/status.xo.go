// Package local_sms_db contains the types for schema 'sms'.
package localsmsdb

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Status is the 'Status' enum type from schema 'sms'.
type Status uint16

const (
	// StatusSendingok is the 'SendingOK' Status.
	StatusSendingok = Status(1)

	// StatusSendingoknoreport is the 'SendingOKNoReport' Status.
	StatusSendingoknoreport = Status(2)

	// StatusSendingerror is the 'SendingError' Status.
	StatusSendingerror = Status(3)

	// StatusDeliveryok is the 'DeliveryOK' Status.
	StatusDeliveryok = Status(4)

	// StatusDeliveryfailed is the 'DeliveryFailed' Status.
	StatusDeliveryfailed = Status(5)

	// StatusDeliverypending is the 'DeliveryPending' Status.
	StatusDeliverypending = Status(6)

	// StatusDeliveryunknown is the 'DeliveryUnknown' Status.
	StatusDeliveryunknown = Status(7)

	// StatusError is the 'Error' Status.
	StatusError = Status(8)
)

// String returns the string value of the Status.
func (s Status) String() string {
	var enumVal string

	switch s {
	case StatusSendingok:
		enumVal = "SendingOK"

	case StatusSendingoknoreport:
		enumVal = "SendingOKNoReport"

	case StatusSendingerror:
		enumVal = "SendingError"

	case StatusDeliveryok:
		enumVal = "DeliveryOK"

	case StatusDeliveryfailed:
		enumVal = "DeliveryFailed"

	case StatusDeliverypending:
		enumVal = "DeliveryPending"

	case StatusDeliveryunknown:
		enumVal = "DeliveryUnknown"

	case StatusError:
		enumVal = "Error"
	}

	return enumVal
}

// MarshalText marshals Status into text.
func (s Status) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText unmarshals Status from text.
func (s *Status) UnmarshalText(text []byte) error {
	switch string(text) {
	case "SendingOK":
		*s = StatusSendingok

	case "SendingOKNoReport":
		*s = StatusSendingoknoreport

	case "SendingError":
		*s = StatusSendingerror

	case "DeliveryOK":
		*s = StatusDeliveryok

	case "DeliveryFailed":
		*s = StatusDeliveryfailed

	case "DeliveryPending":
		*s = StatusDeliverypending

	case "DeliveryUnknown":
		*s = StatusDeliveryunknown

	case "Error":
		*s = StatusError

	default:
		return errors.New("invalid Status")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Status.
func (s Status) Value() (driver.Value, error) {
	return s.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Status.
func (s *Status) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Status")
	}

	return s.UnmarshalText(buf)
}