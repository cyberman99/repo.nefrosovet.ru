package viber

import (
    "encoding/json"
    "errors"
    "fmt"
    "regexp"
    "strings"

    log "github.com/Sirupsen/logrus"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    utils2 "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/utils"

    mqtt "github.com/eclipse/paho.mqtt.golang"
    "github.com/mileusna/viber"
    "github.com/spf13/viper"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

var (
    regexpPeekMsgType = regexp.MustCompile("\"type\":\\s*\"([^\"]+)\"")
)

type Channel struct {
    ID string
    ClientID string
    Login    string
    Password string

    GreetingText string
    ButtonText string
    AlternateText string
    AnswerText string

    Bot      *viber.Viber
}

func (c *Channel) NormalizeDestination(dest string) (string, error) {
    return utils2.NormalizePhoneNumber(dest)
}

func (c *Channel) Send(destination, data string, _ ...senders.SendOption) error {
    contact, err := utils2.GetStorage().GetViberContact(storage.GetViberContact{
        Phone: &destination,
        Token: &c.Bot.AppKey,
    })
    if err != nil {
        return err
    }
    _, err = c.Bot.SendTextMessage(contact.UserID, data)
    return err
}

func (c *Channel) StartResolver() error {
    mqtt.ERROR = log.New()
    opts := mqtt.NewClientOptions()
    opts = opts.AddBroker("ssl://" + viper.GetString("botProxy.mq.host") + ":" + viper.GetString("botProxy.mq.port"))
    opts.ClientID = c.ClientID
    opts.Username = c.Login
    opts.Password = c.Password

    mqttClient := mqtt.NewClient(opts)
    token := mqttClient.Connect()
    token.Wait()
    if token.Error() != nil {
        return token.Error()
    }

    go func() {
        token = mqttClient.Subscribe(
            "webhooks/"+c.ClientID,
            0,
            func(client mqtt.Client, msg mqtt.Message) {
                log.WithFields(log.Fields{
                    "context":     "CHANNEL",
                    "channelType": "VIBER",
                    "channelID":   c.ID,
                    "message":     string(msg.Payload()),
                }).Debug("Got message")

                c.HandleEvent(msg.Payload())
            },
        )

        token.Wait()

        if token.Error() != nil {
            log.WithFields(log.Fields{
                "context":     "CHANNEL",
                "channelType": "VIBER",
                "channelID":   c.ID,
                "error":       token.Error(),
            }).Error("Handling viber mq message error")
        }
    }()

    log.WithFields(
        log.Fields{
            "context":     "CHANNEL",
            "channelType": "VIBER",
            "channelID":   c.ID,
            "status":      "STARTED",
        },
    ).Info("Viber contact resolver started")

    return nil
}

type viberEvent struct {
    Event        string          `json:"event"`
    Timestamp    viber.Timestamp `json:"timestamp"`
    MessageToken uint64          `json:"message_token,omitempty"`
    UserID       string          `json:"user_id,omitempty"`

    // failed event
    Descr string `json:"descr,omitempty"`

    // conversation_started event
    Type       string          `json:"type,omitempty"`
    Context    string          `json:"context,omitempty"`
    Subscribed bool            `json:"subscribed,omitempty"`
    User       json.RawMessage `json:"user,omitempty"`

    // message event
    Sender  json.RawMessage `json:"sender,omitempty"`
    Message json.RawMessage `json:"message,omitempty"`
}

type ContactMessage struct {
    viber.TextMessage
    Contact struct {
        Name  string `json:"name"`
        Phone string `json:"phone_number"`
    } `json:"contact"`
}

// HandleEvent - handles viber hook body
func (c *Channel) HandleEvent(body []byte) {
    var e viberEvent
    if err := json.Unmarshal(body, &e); err != nil {
        log.WithFields(log.Fields{
            "context":     "CHANNEL",
            "channelType": "VIBER",
            "channelID":   c.ID,
            "error":       err,
        }).Error("Can't unmarshal event body")

        return
    }

    switch e.Event {
    case "message":
        var u viber.User
        if err := json.Unmarshal(e.Sender, &u); err != nil {
            log.WithFields(log.Fields{
                "context":     "CHANNEL",
                "channelType": "VIBER",
                "channelID":   c.ID,
                "error":       err,
            }).Error("Can't unmarshal message event body")

            return
        }

        contact, err := utils2.GetStorage().GetViberContact(storage.GetViberContact{
            UserID: &u.ID,
        })
        if err != nil && err != storage.ErrViberContactNotFound {
            log.WithFields(log.Fields{
                "context":     "CHANNEL",
                "channelType": "VIBER",
                "channelID":   c.ID,
                "error":       err,
            }).Error("Getting viber contact by user ID error")
        }

        msgType := getViberMessageType(e.Message)
        switch msgType {
        case "text":
            var m viber.TextMessage
            if err = json.Unmarshal(e.Message, &m); err != nil {
                log.WithFields(log.Fields{
                    "context":     "CHANNEL",
                    "channelType": "VIBER",
                    "channelID":   c.ID,
                    "error":       err,
                }).Error("Can't unmarshal message text body")

                return
            }

            if contact.Phone == "" {
                msg := c.Bot.NewTextMessage(c.GreetingText)
                msg.MinAPIVersion = 3
                keyboard := c.Bot.NewKeyboard("#FFFFFF", false)
                button := c.Bot.NewButton(6, 1, "share-phone", "phone", c.ButtonText, "")
                button.SetBgColor("#DDDDDD")

                keyboard.AddButton(button)
                msg.SetKeyboard(keyboard)
                _, err = c.Bot.SendMessage(u.ID, msg)
                if err != nil {
                    log.WithFields(log.Fields{
                        "context":     "CHANNEL",
                        "channelType": "VIBER",
                        "channelID":   c.ID,
                        "error":       err,
                    }).Error("Can't send contacts request to viber user")
                }

                return
            }

            _, err = c.Bot.SendTextMessage(u.ID, c.AlternateText)
            if err != nil {
                log.WithFields(log.Fields{
                    "context":     "CHANNEL",
                    "channelType": "VIBER",
                    "channelID":   c.ID,
                    "error":       err,
                }).Error("Can't send default answer to viber user")
            }

        case "contact":
            var contactMsg ContactMessage
            if err = json.Unmarshal(e.Message, &contactMsg); err != nil {
                log.WithFields(log.Fields{
                    "context":     "CHANNEL",
                    "channelType": "VIBER",
                    "channelID":   c.ID,
                    "error":       err,
                }).Error("Can't unmarshal user contact body")

                return
            }

            if contact.Phone == "" && contactMsg.Contact.Phone != "" {
                contact.Phone = fmt.Sprintf("+%s", utils2.FilterNonDigits(contactMsg.Contact.Phone))
                contact.UserID = u.ID
                contact.Name = contactMsg.Contact.Name
                contact.Token = c.Bot.AppKey

                utils2.GetStorage().StoreViberContact(storage.StoreViberContact{ViberContact: contact})

                _, err = c.Bot.SendTextMessage(u.ID, c.AnswerText)
                if err != nil {
                    log.WithFields(log.Fields{
                        "context":     "CHANNEL",
                        "channelType": "VIBER",
                        "channelID":   c.ID,
                        "error":       err,
                    }).Error("Can't send AnswerText to viber user")
                }

                return
            }

            _, err = c.Bot.SendTextMessage(u.ID, c.AlternateText)
            if err != nil {
                log.WithFields(log.Fields{
                    "context":     "CHANNEL",
                    "channelType": "VIBER",
                    "channelID":   c.ID,
                    "error":       err,
                }).Error("Can't send AlternateText to viber user")
            }

        default:
            return
        }
    default:
        return
    }
}

func (c *Channel) StopResolver() error {
    if c.Bot == nil {
        return errors.New("channel bot is not configured")
    }
    opts := mqtt.NewClientOptions()
    opts = opts.AddBroker("ssl://" + viper.GetString("botProxy.mq.host") + ":" + viper.GetString("botProxy.mq.port"))
    opts.ClientID = c.ClientID
    opts.Username = c.Login
    opts.Password = c.Password
    mqttClient := mqtt.NewClient(opts)
    token := mqttClient.Connect()
    token.Wait()
    if token.Error() != nil {
        return token.Error()
    }

    token = mqttClient.Unsubscribe(
        "webhooks/" + c.ClientID,
    )
    token.Wait()
    if token.Error() != nil {
        log.WithFields(log.Fields{
            "context":     "CHANNEL",
            "channelType": "VIBER",
            "channelID":   c.ID,
            "error":       token.Error(),
        }).Error("Handling viber mq message error")
    }

    c.Bot = nil

    log.WithFields(
        log.Fields{
            "context":     "CHANNEL",
            "channelType": "VIBER",
            "channelID":   c.ID,
            "status":      "STOPPED",
        },
    ).Info("Contact resolver stopped")

    return nil
}

// getViberMessageType uses regexp to determin message type for unmarshaling
func getViberMessageType(b []byte) string {
    matches := regexpPeekMsgType.FindAllSubmatch(b, -1)
    if len(matches) == 0 {
        return ""
    }

    return strings.ToLower(string(matches[0][1]))
}
