package sender

import (
    "errors"
    "fmt"
    "time"

    log "github.com/Sirupsen/logrus"
    viberbot "github.com/mileusna/viber"
    uuid "github.com/satori/go.uuid"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/email"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/localsms"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/mtssms"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/slack"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/telegram"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/viber"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
)

const (
    StatusQueue    = "QUEUE"
    StatusSent     = "SENT"
    StatusError    = "ERROR"
    Status500Error = "500"
)

var (
    Version   string
    TgPool    telegram.Pool
    ViberPool viber.Pool
)

type Meta struct {
    AccessToken   *string
    EmailFrom     *string
    EmailSubject  *string
    SlackDestType *string
}

// Result used for Send() returns
type Result struct {
    ID          string `json:"ID"`
    Created     string `json:"created,omitempty"`
    Updated     string `json:"updated,omitempty"`
    Status      string `json:"status,omitempty"`
    Errors      string `json:"errors,omitempty"`
    ChannelID   string `json:"channelID,omitempty"`
    Destination string `json:"destination,omitempty"`
    Data        string `json:"data,omitempty"`

    Meta struct {
        EmailSubject  string `json:"emailSubject,omitempty"`
        EmailFrom     string `json:"emailFrom,omitempty"`
        SlackDestType string `json:"slackDestType,omitempty"`
    } `json:"meta,omitempty"`
}

// Send sends message with data to destination via channel, res is always not nil
func Send(channel *storage.Channel, destination, data string, meta *Meta) (res *Result, err error) {
    defer func() {
        res = &Result{
            ID:          uuid.NewV4().String(),
            Created:     time.Now().Format("2006-01-02T15:04:05-07:00"),
            ChannelID:   channel.ID,
            Destination: destination,
            Data:        data,
            Meta: struct {
                EmailSubject  string `json:"emailSubject,omitempty"`
                EmailFrom     string `json:"emailFrom,omitempty"`
                SlackDestType string `json:"slackDestType,omitempty"`
            }{},
        }
        if err != nil {
            switch err.(type) {
            case senders.DestinationValidationError, senders.DestinationNotFound:
                res.Status = StatusError
            default:
                res.Status = Status500Error
            }
            res.Errors = err.Error()
        } else {
            res.Status = StatusSent
        }
        if meta != nil {
            if meta.EmailSubject != nil {
                res.Meta.EmailSubject = *meta.EmailSubject
            }
            if meta.EmailFrom != nil {
                res.Meta.EmailFrom = *meta.EmailFrom
            }
            if meta.SlackDestType != nil {
                res.Meta.SlackDestType = *meta.SlackDestType
            }
        }
        go func() {
            var errStr string
            if err != nil {
                errStr = err.Error()
            }
            msgEvent := storage.MessageEvent{
                ID:                res.ID,
                Created:           res.Created,
                Status:            res.Status,
                Errors:            errStr,
                ChannelID:         res.ChannelID,
                ChannelType:       string(channel.Type),
                Destination:       res.Destination,
                Data:              res.Data,
                MetaEmailSubject:  res.Meta.EmailSubject,
                MetaEmailFrom:     res.Meta.EmailFrom,
                MetaSlackDestType: res.Meta.SlackDestType,
            }
            if meta != nil && meta.AccessToken != nil {
                msgEvent.AccessToken = *meta.AccessToken
            }
            _, err := _default.GetStorage().StoreMessageEvent(storage.StoreMessageEvent{
                MessageEvent: msgEvent,
            })
            if err != nil {
                log.WithFields(log.Fields{
                    "context": "EVENTDB",
                    "action":  "StoreMessageEvent",
                }).Error(err)
            }
        }()
    }()
    var (
        ch   senders.Channel
        opts []senders.SendOption
    )
    switch channel.Type {
    case storage.ChannelTypeTelegram:
        tgCh, exists := TgPool.Get(channel.ID)
        if !exists {
            tgCh := &telegram.Channel{
                ID:            channel.ID,
                Token:         channel.Params.Token,
                GreetingText:  channel.Params.GreetingText,
                AnswerText:    channel.Params.AnswerText,
                AlternateText: channel.Params.AlternateText,
                ButtonText:    channel.Params.ButtonText,
            }
            TgPool.Add(tgCh)
        }
        ch = tgCh
    case storage.ChannelTypeSlack:
        ch = slack.NewChannel(channel.Params.Token, channel.Params.Name)
        if meta != nil && meta.SlackDestType != nil {
            opts = append(opts, slack.DestType(*meta.SlackDestType))
        }
    case storage.ChannelTypeViber:
        vibCh, exists := ViberPool.GetByID(channel.ID)
        if !exists {
            vibCh := &viber.Channel{
                ID:            channel.ID,
                ClientID:      channel.Params.ClientID,
                Login:         channel.Params.ClientID,
                Password:      channel.Params.ClientID,
                GreetingText:  channel.Params.ClientID,
                ButtonText:    channel.Params.ClientID,
                AlternateText: channel.Params.ClientID,
                AnswerText:    channel.Params.ClientID,
                Bot:           viberbot.New(channel.Params.Token, channel.Params.Name, channel.Params.Avatar),
            }
            ViberPool.Add(vibCh)
        }
        ch = vibCh
    case storage.ChannelTypeEmail:
        ch = &email.Channel{
            ID:          channel.ID,
            From:        channel.Params.From,
            Login:       channel.Params.Login,
            Password:    channel.Params.Password,
            Port:        channel.Params.Port,
            Server:      channel.Params.Server,
            SSL:         channel.Params.SSL,
            ContentType: channel.Params.ContentType,
        }
        if meta != nil {
            if meta.EmailSubject != nil {
                opts = append(opts,
                    email.Subject(*meta.EmailSubject))
            }
            if meta.EmailFrom != nil {
                opts = append(opts,
                    email.From(*meta.EmailFrom))
            }
        }
    case storage.ChannelTypeLocalSMS:
        ch = &localsms.Channel{
            ID:       channel.ID,
            Login:    channel.Params.Login,
            Password: channel.Params.Password,
            Port:     channel.Params.Port,
            Server:   channel.Params.Server,
            ModemID:  channel.Params.ModemID,
            Db:       channel.Params.Db,
        }
        opts = append(opts, localsms.AccessToken(channel.AccessToken))
    case storage.ChannelTypeMTSSMS:
        ch = &mtssms.Channel{
            ID:       channel.ID,
            Login:    channel.Params.Login,
            Password: channel.Params.Password,
            From:     channel.Params.From,
        }
    default:
        err = errors.New(fmt.Sprintf("Unsupported channel type %q", channel.Type))
        return
    }
    dest, err := ch.NormalizeDestination(destination)
    if err != nil {
        return
    }
    err = ch.Send(dest, data, opts...)
    return
}
