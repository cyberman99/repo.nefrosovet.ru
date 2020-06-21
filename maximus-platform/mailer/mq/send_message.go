package mq

import (
    "encoding/json"

    log "github.com/Sirupsen/logrus"
    "github.com/pkg/errors"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

func (m *MQ) handleSendMessage(message *IncomingMessage) {
    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &message.Payload.Body.ChannelID,
    })
    if err == nil && channel.AccessToken != message.Payload.Params.AccessToken {
        err = errors.New("access denied")
    }
    if err != nil {
        log.WithFields(log.Fields{
            "context":     "MQ",
            "error":       "can't get channel",
            "channelID":   message.Payload.Body.ChannelID,
            "accessToken": message.Payload.Params.AccessToken,
        }).Error(err)
        m.PublishSendErrorMessage(message.TransactionID, err.Error())
        return
    }

    var response MessageSendResult
    response.TransactionID = message.TransactionID
    response.Payload.Body.Version = sender.Version
    res, err := sender.Send(&channel, message.Payload.Body.Destination, message.Payload.Body.Data, nil)
    if err != nil {
        switch err.(type) {
        case senders.DestinationValidationError, senders.DestinationNotFound:
            response.Payload.Body.Status.Code = 200
            response.Payload.Body.Status.Message = "SUCCESS"
        default:
            response.Payload.Body.Status.Code = 500
            response.Payload.Body.Status.Message = "Internal server error"
        }
    }
    if res == nil {
        // res should not be not nil, as it has to be returned in response,
        // even if there was an error while sending message
        panic(errors.Wrap(err, "nil result was returned with error"))
    }
    response.Payload.Data = append(response.Payload.Data, *res)

    msg, err := json.Marshal(&response)
    if err != nil {
        log.WithFields(log.Fields{
            "context": "MQ",
            "error":   err,
        }).Error("Can't marshal sending message responce")

        return
    }

    log.WithFields(log.Fields{
        "context":       "MQ",
        "action":        "REPLY",
        "transactionID": response.TransactionID,
    }).Info("Sending MQ answer")
    log.WithFields(log.Fields{
        "message": string(msg),
    }).Debug()

    m.PublishMessage(msg)

}
