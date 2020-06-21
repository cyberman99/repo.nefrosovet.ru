package slack

import (
    "github.com/nlopes/slack"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

type message struct {
    data     string
    destType DestType
}

type Channel struct {
    username string
    api      *slack.Client
}

func NewChannel(token, username string) *Channel {
    return &Channel{
        username: username,
        api:      slack.New(token),
    }
}

func (c *Channel) NormalizeDestination(dest string) (string, error) {
    return dest, nil
}

func (c *Channel) Send(destination, data string, opts ...senders.SendOption) error {
    m := message{
        data: data,
    }
    for _, opt := range opts {
        if err := opt.Apply(&m); err != nil {
            return err
        }
    }
    if m.destType == ChannelDestType {
        _, _, err := c.api.PostMessage(destination, data, slack.PostMessageParameters{Username: c.username})
        return err
    }

    usersList, err := c.api.GetUsers()
    if err != nil {
        return err
    }
    for _, user := range usersList {
        if user.Name == destination { // TODO: destination should be ID of the user. Use https://godoc.org/github.com/nlopes/slack#Info.GetUserByID
            _, _, chanID, err := c.api.OpenIMChannel(user.ID)
            if err != nil {
                return err
            }
            _, _, err = c.api.PostMessage(chanID, data, slack.PostMessageParameters{Username: c.username})
            return err
        }
    }
    return senders.DestinationNotFound(destination)
}
