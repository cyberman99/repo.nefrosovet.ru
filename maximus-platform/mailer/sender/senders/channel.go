package senders

import "fmt"

const (
    WrongDestinationMesssage  = "invalid destination format"
    DestanationNotFoundMessage = "destination not found"
)

type Message interface{}

type SendOption interface {
    Apply(to Message) error
}

type Channel interface {
    // NormalizeDestination checks if the given destination is valid
    // and normalizes it if possible
    NormalizeDestination(destination string) (string, error)

    // Send sends message containing data with applied SendOptions to destination.
    //
    // If destination was not found, it returns DestinationNotFound.
    // Other errors comes from sender backend
    Send(destination, data string, opts ...SendOption) error
}

type DestinationNotFound string

func (err DestinationNotFound) Error() string {
    return fmt.Sprintf("%s: %s", DestanationNotFoundMessage, string(err))
}

type DestinationValidationError string

func (err DestinationValidationError) Error() string {
    return fmt.Sprintf("%s: %s", WrongDestinationMesssage, string(err))
}
