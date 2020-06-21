package handlers

import (
    log "github.com/Sirupsen/logrus"
    "github.com/pkg/errors"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/models"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"

    "github.com/go-openapi/runtime/middleware"

    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/messages"
)

func insertVersion(data *models.BaseData) {
    data.Version = sender.Version
}

// PostSendHandler - POST /send
func PostSendHandler(params messages.PostSendParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(messages.PostSendForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return messages.NewPostSendForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string, item *messages.DataItems0, err error) middleware.Responder {
        payload := new(messages.PostSendBadRequestBody)
        insertVersion(&payload.BaseData)
        payload.Errors = err.Error()
        payload.Message = message

        payload.Data = []*messages.DataItems0{item}
        return messages.NewPostSendBadRequest().WithPayload(payload)
    }

    responseInternalServerError := func(item *messages.DataItems0, err error) middleware.Responder {
        payload := new(messages.PostSendInternalServerErrorBody)
        insertVersion(&payload.BaseData)
        payload.Message = PayloadInternalServerErrorMessage
        payload.Data = []*messages.DataItems0{item}

        return messages.NewPostSendInternalServerError().WithPayload(payload)
    }

    responseSuccess := func(item *messages.DataItems0) middleware.Responder {
        payload := new(messages.PostSendOKBody)
        insertVersion(&payload.BaseData)
        message := PayloadSuccessMessage
        payload.Message = &message
        payload.Data = append(payload.Data, item)

        return messages.NewPostSendOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }
    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: params.Body.ChannelID,
    })
    if err != nil || accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }
    meta := &sender.Meta{
        AccessToken: &params.AccessToken,
    }
    if params.Body.Meta != nil {
        meta.EmailFrom = &params.Body.Meta.EmailFrom
        meta.EmailSubject = &params.Body.Meta.EmailSubject
        meta.SlackDestType = &params.Body.Meta.SLACKDestType
    }

    res, err := sender.Send(&channel, *params.Body.Destination, *params.Body.Data, meta)
    if res == nil {
        // res should not be not nil, as it has to be returned in response,
        // even if there was an error while sending message
        panic(errors.Wrap(err, "nil result was returned with error"))
    }
    item :=  &messages.DataItems0{
        ID:        res.ID,
        ChannelID: res.ChannelID,
        Created:   res.Created,
        Errors:    res.Errors,
        Status:    res.Status,
        MessageObject: models.MessageObject{
            ChannelID:   params.Body.ChannelID,
            Data:        params.Body.Data,
            Destination: params.Body.Destination,
            Meta:        params.Body.Meta,
        },
    }
    if err != nil {
        switch err.(type) {
        case senders.DestinationValidationError:
            return responseBadRequest(PayloadValidationErrorMessage, item, errors.Wrap(err, "bad destination format"))
        case senders.DestinationNotFound:
            return responseBadRequest(PayloadValidationErrorMessage, item, errors.New("user: not found"))
        default:
            log.WithFields(log.Fields{
                "context":   "SEND",
                "channelID": channel.ID,
            }).Error(err)
            return responseInternalServerError(item, errors.New("Error while sending message"))
        }
    }
    return responseSuccess(item)
}

// GetMessagesHandler - GET /messages
func GetMessagesHandler(params messages.GetMessagesParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(messages.GetMessagesForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message
        payload.Data = make([]messages.DataItems0, 0)

        return messages.NewGetMessagesForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(messages.PostSendBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message
        payload.Data = make([]messages.DataItems0, 0)

        return messages.NewPostSendBadRequest().WithPayload(payload)
    }

    responseSuccess := func(items []*messages.DataItems0) middleware.Responder {
        payload := new(messages.GetMessagesOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message
        payload.Data = append(payload.Data, items...)

        return messages.NewGetMessagesOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{Token: params.AccessToken})
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    messageEvents, err := GetStorage().GetMessageEvents(storage.GetMessageEvents{
        AccessToken: accessToken.Token,
        ChannelID:   params.ChannelID,
        Status:      params.Status,
        Destination: params.Destination,
        Limit:       params.Limit,
        Offset:      params.Offset,
    })
    if err != nil && err != storage.ErrMessageEventsNotFound {
        return responseBadRequest(wrapGetMessageEventsErrorMessage(err))
    }

    var items []*messages.DataItems0
    for i := range messageEvents {
        // We need local copy of event, cuz we use pointers
        sendEvent := messageEvents[i]

        item := new(messages.DataItems0)
        item.ID = sendEvent.ID
        item.Created = sendEvent.Created
        item.Destination = &sendEvent.Destination
        item.MessageObject.ChannelID = &sendEvent.ChannelID
        item.Data = &sendEvent.Data
        item.Status = sendEvent.Status
        if sendEvent.Errors != "" {
            item.Errors = sendEvent.Errors
        }

        switch storage.ChannelType(sendEvent.ChannelType) {
        default:
            break
        case storage.ChannelTypeEmail:
            item.Meta = new(models.MessageObjectMeta)

            item.Meta.EmailFrom = sendEvent.MetaEmailFrom
            item.Meta.EmailSubject = sendEvent.MetaEmailSubject
        case storage.ChannelTypeSlack:
            item.Meta = new(models.MessageObjectMeta)

            item.Meta.SLACKDestType = sendEvent.MetaSlackDestType
        }

        items = append(items, item)

    }

    return responseSuccess(items)
}

// GetMessagesMessageIDHandler - GET /messages/<ID>
func GetMessagesMessageIDHandler(params messages.GetMessagesMessageIDParams) middleware.Responder {
    responseIDNotFound := func() middleware.Responder {
        payload := new(messages.GetMessagesMessageIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return messages.NewGetMessagesMessageIDNotFound().WithPayload(payload)
    }

    responseSuccess := func(item *messages.DataItems0) middleware.Responder {
        payload := new(messages.GetMessagesOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return messages.NewGetMessagesOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{Token: params.AccessToken})
    if err != nil {
        return responseIDNotFound()
    }

    sendEvent, err := GetStorage().GetMessageEvent(storage.GetMessageEvent{
        ID: params.MessageID,
    })
    if err != nil || sendEvent.AccessToken != accessToken.Token {
        return responseIDNotFound()
    }

    item := new(messages.DataItems0)
    item.ID = sendEvent.ID
    item.Created = sendEvent.Created
    item.Destination = &sendEvent.Destination
    item.MessageObject.ChannelID = &sendEvent.ChannelID
    item.Data = &sendEvent.Data
    item.Status = sendEvent.Status
    if sendEvent.Errors != "" {
        item.Errors = sendEvent.Errors
    }

    switch storage.ChannelType(sendEvent.ChannelType) {
    case storage.ChannelTypeEmail:
        item.Meta = new(models.MessageObjectMeta)

        item.Meta.EmailFrom = sendEvent.MetaEmailFrom
        item.Meta.EmailSubject = sendEvent.MetaEmailSubject
    case storage.ChannelTypeSlack:
        item.Meta = new(models.MessageObjectMeta)

        item.Meta.SLACKDestType = sendEvent.MetaSlackDestType
    }

    return responseSuccess(item)
}
