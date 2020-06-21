package jwt

import (
	"gopkg.in/hlandau/passlib.v1"
)

// VerifyPassword validates password and regenerate hash
func VerifyPassword(password, hash string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Verify(password, hash)
}

// HashPassword returns password hash
func HashPassword(password string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Hash(password)
}
