package utils

import (
    "regexp"

    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

const (
    WrongDestinationFormat = "invalid destination format"
)

var (
    phoneNumberRegex = regexp.MustCompile(`^\+(\d+)$`)
    digitsRegex      = regexp.MustCompile(`[^\d]`)
)

// NormalizePhoneNumber checks the string for compliance with the
// international standard for writing a phone number: +12345678900
// and returns it without plus
func NormalizePhoneNumber(phone string) (string, error) {
    if !phoneNumberRegex.MatchString(phone) {
        return "", senders.DestinationValidationError(phone)
    }
    return phoneNumberRegex.ReplaceAllString(phone, "$1"), nil
}

func FilterNonDigits(source string) string {
    return digitsRegex.ReplaceAllLiteralString(source, "")
}
