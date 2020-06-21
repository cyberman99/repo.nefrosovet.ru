package telegram

import (
    "context"
    "time"

    log "github.com/Sirupsen/logrus"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/utils"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
    _default "repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
)

type Channel struct {
    ID            string
    Token         string
    GreetingText  string
    AnswerText    string
    AlternateText string
    ButtonText    string

    stop    context.CancelFunc
    started bool
    bot     *tgbotapi.BotAPI
}

func (ch *Channel) NormalizeDestination(destination string) (string, error) {
    return utils.NormalizePhoneNumber(destination)
}

func (ch *Channel) Send(destination, data string, _ ...senders.SendOption) error {
    chatID, err := GetChatIDByPhone(destination)
    if err != nil {
        switch err {
        case storage.ErrTelegramContactNotFound:
            return senders.DestinationNotFound(destination)
        default:
            return err
        }
    }
    if ch.bot == nil {
        b, err := tgbotapi.NewBotAPI(ch.Token)
        if err != nil {
            return err
        }
        ch.bot = b
    }
    msg := tgbotapi.NewMessage(chatID, data)
    _, err = ch.bot.Send(msg)
    return err
}

func (ch *Channel) StartContactResolver(ctx context.Context) error {
    if ch.started {
        return nil
    }
    ch.started = true
    ctx, cancelFunc := context.WithCancel(ctx)
    ch.stop = cancelFunc
    go func() {
        u := tgbotapi.NewUpdate(0)
        u.Timeout = 10
        log.WithFields(log.Fields{
            "context":     "CHANNEL",
            "channelType": "TELEGRAM",
            "channelID":   ch.ID,
            "status":      "STARTED",
        }).Info("Telegram contact resolver started")
        for {
            select {
            case <-ctx.Done():
                return
            default:
                updates, err := ch.bot.GetUpdates(u)
                if err != nil {
                    log.WithFields(log.Fields{
                        "context":     "CHANNEL",
                        "channelType": "TELEGRAM",
                        "channelID":   ch.ID,
                        "error":       err,
                    }).Error("Failed to get updates. Retrying in 3 seconds")

                    time.Sleep(time.Second * 3)
                    continue
                }

                for _, update := range updates {

                    u.Offset = update.UpdateID + 1

                    chatID := update.Message.Chat.ID
                    text := update.Message.Text

                    log.WithFields(log.Fields{
                        "context":     "CHANNEL",
                        "channelType": "TELEGRAM",
                        "channelID":   ch.ID,
                        "message":     text,
                        "chatID":      chatID,
                    }).Debug("Got telegram message")

                    contact, err := _default.GetStorage().GetTelegramContact(storage.GetTelegramContact{ChatID: &chatID})
                    if err != nil && err.Error() != "not found" {
                        log.WithFields(log.Fields{
                            "context":     "CHANNEL",
                            "channelType": "TELEGRAM",
                            "channelID":   ch.ID,
                            "error":       err,
                        }).Error("Getting telegram contact by chat ID error")
                    }
                    messageContact := update.Message.Contact
                    var msg tgbotapi.MessageConfig
                    switch {
                    case contact.Phone == "" && messageContact != nil && messageContact.PhoneNumber != "":
                        contact.Phone = utils.FilterNonDigits(messageContact.PhoneNumber)
                        contact.ChatID = chatID
                        contact.Username = update.Message.Chat.UserName
                        _, err := _default.GetStorage().StoreTelegramContact(storage.StoreTelegramContact{
                            TelegramContact: contact,
                        })
                        if err != nil {
                            log.WithFields(log.Fields{
                                "context":     "CHANNEL",
                                "channelType": "TELEGRAM",
                                "action":      "StoreTelegramContact",
                                "channelID":   ch.ID,
                            }).Error(err)
                        }
                        msg = tgbotapi.NewMessage(chatID, ch.AnswerText)
                        msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

                    case messageContact != nil && utils.FilterNonDigits(messageContact.PhoneNumber) == contact.Phone:
                        msg = tgbotapi.NewMessage(chatID, ch.AlternateText)
                        msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
                    case contact.Phone == "":
                        button := tgbotapi.NewKeyboardButtonContact(ch.ButtonText)
                        msg = tgbotapi.NewMessage(chatID, ch.GreetingText)
                        msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(button))
                    default:
                        msg = tgbotapi.NewMessage(chatID, ch.AlternateText)
                    }
                    if _, err := ch.bot.Send(msg); err != nil {
                        log.WithFields(log.Fields{
                            "context":     "CHANNEL",
                            "channelType": "TELEGRAM",
                            "action":      "Send",
                            "channelID":   ch.ID,
                        }).Error(err)
                    }
                }
            }
        }
    }()
    return nil
}

func (ch *Channel) StopContactResolver() {
    if ch.stop != nil {
        ch.stop()
    }
    ch.started = false
}

// GetTgContactByPhone returns chatID by Phone, if found
func GetChatIDByPhone(phone string) (int64, error) {
    contact, err := utils.GetStorage().GetTelegramContact(storage.GetTelegramContact{
        Phone: &phone,
    })
    if err != nil {
        return 0, err
    }

    return contact.ChatID, nil
}
