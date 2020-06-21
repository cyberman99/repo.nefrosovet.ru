package slack

import (
    "errors"
    "fmt"

    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

const ChannelDestType DestType = "channel"

type DestType string

func (o DestType) Apply(to senders.Message) error {
    switch m := to.(type) {
    case *message:
        m.destType = o
    default:
        return errors.New(fmt.Sprintf("can not apply DestType to %T", to))
    }
    return nil
}
