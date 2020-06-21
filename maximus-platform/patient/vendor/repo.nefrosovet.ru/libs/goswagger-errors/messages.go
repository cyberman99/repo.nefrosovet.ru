package goswagger_errors

import (
	"fmt"
)

const (
	InternalServerErrorMessage = "Internal server error"
	ValidationErrorMessage     = "Validation error"
	NotFoundMessage            = "Entity not found"
	AccessDeniedMessage        = "Access denied"
	UnknownPathMessage         = "Unknown path"
	MethodNotAllowedMessage    = "Method not allowed"
)

func WrapMethodNotAllowedMessage(method string) string {
	if method == "" {
		return MethodNotAllowedMessage
	}

	return fmt.Sprintf("Method %s not allowed", method)
}
