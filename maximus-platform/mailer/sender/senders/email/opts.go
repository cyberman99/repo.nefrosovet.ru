package email

import (
    "errors"
    "fmt"

    "github.com/jordan-wright/email"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

type Subject string

func (o Subject) Apply(to senders.Message) error {
    switch m := to.(type) {
    case *email.Email:
        m.Subject = string(o)
    default:
        return errors.New(fmt.Sprintf("can not apply Subject to %T", to))
    }
    return nil
}

type From string

func (o From) Apply(to senders.Message) error {
    switch m := to.(type) {
    case *email.Email:
        m.From = string(o)
    default:
        return errors.New(fmt.Sprintf("can not apply Subject to %T", to))
    }
    return nil
}
