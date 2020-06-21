package localsms

import (
    "errors"
    "fmt"

    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

type AccessToken string

func (o AccessToken) Apply(to senders.Message) error {
    switch m := to.(type) {
    case *message:
        m.CreatorID = string(o)
    default:
        return errors.New(fmt.Sprintf("can not apply AccessToken to %T", to))
    }
    return nil
}
