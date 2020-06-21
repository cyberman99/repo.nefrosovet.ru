package storage

import (
	"errors"
)

var (
	ErrBadInput = errors.New("bad input")

	// User
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")

	// User settings
	ErrUserSettingsNotFound = errors.New("user settings not found")

	// User contacts
	ErrUserContactNotFound = errors.New("user contacts not found")
	ErrUserContactAlreadyVerified = errors.New("user contact is already verified")
	ErrUserContactAlreadyExists = errors.New("user contact already exists")
)
