package telegram

import (
    "context"
    "time"

    log "github.com/Sirupsen/logrus"
    "github.com/pkg/errors"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

type Pool map[string]*Channel

func NewPool() Pool {
    return make(Pool)
}

func (p Pool) AddFromStorage(stor storage.ChannelStorage) error {
    t := storage.ChannelTypeTelegram
    chanCollection, err := stor.GetChannels(storage.GetChannels{
        Type: &t,
    })
    if err != nil {
        return err
    }
    for _, ch := range chanCollection {
        if err := p.Add(&Channel{
            ID:            ch.ID,
            Token:         ch.Params.Token,
            GreetingText:  ch.Params.GreetingText,
            AnswerText:    ch.Params.AnswerText,
            AlternateText: ch.Params.AlternateText,
            ButtonText:    ch.Params.ButtonText,
        }); err != nil {
            return err
        }
    }
    return nil
}

func (p Pool) Add(ch *Channel) error {
    p[ch.ID] = ch
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    var err error
    for {
        select {
        case <-ctx.Done():
            return errors.Wrap(err, "can not connect to telegram")
        default:
            ch.bot, err = tgbotapi.NewBotAPI(ch.Token)
        }
        if err == nil {
            break
        }
    }
    return ch.StartContactResolver(context.Background())
}

func (p Pool) Get(channelID string) (*Channel, bool) {
    ch, found := p[channelID]
    if !found {
        return nil, false
    }
    return ch, true
}
func (p Pool) Delete(channelID string) {
    ch, found := p[channelID]
    if found {
        ch.StopContactResolver()
    }
    delete(p, channelID)
}

func (p Pool) Start() {
    for _, ch := range p {
        if err := ch.StartContactResolver(context.Background()); err != nil {
            log.WithFields(log.Fields{
                "context": "TelegramChannel",
                "action":  "StartContactResolver",
            }).Error(err)
        }
    }
}
func (p Pool) Stop() {
    for _, ch := range p {
        ch.StopContactResolver()
    }
}
