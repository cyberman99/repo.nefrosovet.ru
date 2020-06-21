package viber

import (
    log "github.com/Sirupsen/logrus"
    "github.com/mileusna/viber"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

type Pool map[string]*Channel

func NewPool() Pool {
    return make(Pool)
}

func (p Pool) Add(c *Channel) error {
    p[c.ID] = c
    return c.StartResolver()
}

// GetByID returns channel with given ID and if it was found
func (p Pool) GetByID(id string) (*Channel, bool) {
    ch, found := p[id]
    return ch, found
}

func (p Pool) Delete(id string) {
    c, found := p[id]
    if found {
        _ = c.StopResolver()
    }
    delete(p, id)
}

func (p *Pool) AddFromStorage(stor storage.ChannelStorage) error {
    t := storage.ChannelTypeViber
    channels, err := stor.GetChannels(storage.GetChannels{
        Type: &t,
    })
    if err != nil {
        return err
    }
    for _, ch := range channels {
        if err := p.Add(&Channel{
            ID:            ch.ID,
            ClientID:      ch.Params.ClientID,
            Login:         ch.Params.Login,
            Password:      ch.Params.Password,
            GreetingText:  ch.Params.GreetingText,
            ButtonText:    ch.Params.ButtonText,
            AlternateText: ch.Params.AlternateText,
            AnswerText:    ch.Params.AnswerText,
            Bot:           viber.New(ch.Params.Token, ch.Params.Name, ch.Params.Avatar),
        }); err != nil {
            return err
        }
    }
    return nil
}

func (p Pool) Start() {
    for _, ch := range p {
        go func(c *Channel) {
            if err := c.StartResolver(); err != nil {
                log.WithFields(log.Fields{
                    "context": "ViberChannel",
                    "ID":      c.ID,
                    "action":  "StartResolver",
                    "status":  "ERROR",
                }).Error(err)
            }
        }(ch)
    }
}

func (p Pool) Stop() {
    for _, ch := range p {
        go func(c *Channel) {
            if err := c.StopResolver(); err != nil {
                log.WithFields(log.Fields{
                    "context": "ViberChannel",
                    "ID":      c.ID,
                    "action":  "StopResolver",
                    "status":  "ERROR",
                }).Error(err)
            }
        }(ch)
    }
}
