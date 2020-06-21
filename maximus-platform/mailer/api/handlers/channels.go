package handlers

import (
    "fmt"

    log "github.com/Sirupsen/logrus"
    "github.com/go-openapi/runtime/middleware"
    viberbot "github.com/mileusna/viber"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/telegram"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/viber"

    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/channels"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

// GetChannelsHandler - GET /channels
func GetChannelsHandler(params channels.GetChannelsParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.GetChannelsForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewGetChannelsForbidden().WithPayload(payload)
    }

    responseSuccess := func(channelsCollection []storage.Channel) middleware.Responder {
        payload := new(channels.GetChannelsOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        for i := range channelsCollection {
            // We need local copy of channel, cuz we use pointers
            channel := channelsCollection[i]

            payload.Data = append(payload.Data, channelToItem(&channel))
        }

        return channels.NewGetChannelsOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channelsCollection, err := GetStorage().GetChannels(storage.GetChannels{
        AccessToken: &accessToken.Token,
    })
    if err != nil {
        return responseForbidden(fmt.Sprintf("get channels error: %s", err))
    }

    return responseSuccess(channelsCollection)
}

// GetChannelsChannelIDHandler - GET /channels/<ID>
func GetChannelsChannelIDHandler(params channels.GetChannelsChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.GetChannelsChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewGetChannelsChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.GetChannelsChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewGetChannelsChannelIDNotFound().WithPayload(payload)
    }

    responseSuccess := func(channel storage.Channel) middleware.Responder {
        payload := new(channels.GetChannelsChannelIDOKBody)
        payload.Version = sender.Version

        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, channelToItem(&channel))

        return channels.NewGetChannelsChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    return responseSuccess(channel)
}

// PostChannelsEmailHandler - POST /channels/email
func PostChannelsEmailHandler(params channels.PostChannelsEmailParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsEmailForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message
        return channels.NewPostChannelsEmailForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsEmailBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message
        return channels.NewPostChannelsEmailBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemEmail) middleware.Responder {
        payload := new(channels.PostChannelsEmailOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsEmailOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel := storage.NewChannel()

    channel.Type = storage.ChannelTypeEmail
    channel.AccessToken = accessToken.Token
    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.Port = *params.Body.Port
    channel.Params.Server = *params.Body.Server
    channel.Params.SSL = *params.Body.Ssl

    if params.Body.ContentType != nil && *params.Body.ContentType != "" {
        channel.Params.ContentType = *params.Body.ContentType
    } else {
        channel.Params.ContentType = "text/html; charset=UTF-8"
    }

    if params.Body.From != nil {
        channel.Params.From = *params.Body.From
    }

    channel, err = GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("inserting channel error: %s", err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemEmail)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.Port = &channel.Params.Port
    item.Server = &channel.Params.Server
    item.Ssl = &channel.Params.SSL
    item.ContentType = &channel.Params.ContentType

    if channel.Params.From != "" {
        item.From = &channel.Params.From
    }

    return responseSuccess(item)
}

// PostChannelsLocalSmsHandler - POST /channels/local_sms
func PostChannelsLocalSmsHandler(params channels.PostChannelsLocalSmsParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsLocalSmsForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsLocalSmsForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsLocalSmsBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsLocalSmsBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemLocalSms) middleware.Responder {
        payload := new(channels.PostChannelsLocalSmsOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsLocalSmsOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel := storage.NewChannel()

    channel.Type = storage.ChannelTypeLocalSMS
    channel.AccessToken = accessToken.Token
    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.Port = *params.Body.Port
    channel.Params.Server = *params.Body.Server
    channel.Params.ModemID = *params.Body.ModemID
    channel.Params.Db = *params.Body.Db
    channel.Params.Limit = *params.Body.Limit

    channel, err = GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("store channel error: %s", err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemLocalSms)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.Port = &channel.Params.Port
    item.Server = &channel.Params.Server
    item.ModemID = &channel.Params.ModemID
    item.Db = &channel.Params.Db
    item.Limit = &channel.Params.Limit

    return responseSuccess(item)
}

// PostChannelsMtsSmsHandler - POST /channels/mts_sms
func PostChannelsMtsSmsHandler(params channels.PostChannelsMtsSmsParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsMtsSmsForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsMtsSmsForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsMtsSmsBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsMtsSmsBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemMtsSms) middleware.Responder {
        payload := new(channels.PostChannelsMtsSmsOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsMtsSmsOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel := storage.NewChannel()

    channel.Type = storage.ChannelTypeMTSSMS
    channel.AccessToken = accessToken.Token
    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.From = *params.Body.From
    channel.Params.Limit = *params.Body.Limit

    channel, err = GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("inserting channel error: %s", err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemMtsSms)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.From = &channel.Params.From
    item.Limit = &channel.Params.Limit

    return responseSuccess(item)
}

// PostChannelsSLACKHandler - POST /channels/slack
func PostChannelsSLACKHandler(params channels.PostChannelsSLACKParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsSLACKForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsSLACKForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsSLACKBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsSLACKBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemSLACK) middleware.Responder {
        payload := new(channels.PostChannelsSLACKOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsSLACKOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel := storage.NewChannel()

    channel.Type = storage.ChannelTypeSlack
    channel.AccessToken = accessToken.Token
    channel.Params.Token = *params.Body.Token
    channel.Params.Name = *params.Body.Name

    channel, err = GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("inserting channel error: %s", err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemSLACK)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.Name = &channel.Params.Name

    return responseSuccess(item)
}

// PostChannelsTelegramHandler - POST /channels/telegram
func PostChannelsTelegramHandler(params channels.PostChannelsTelegramParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsTelegramForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsTelegramForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string, errors interface{}) middleware.Responder {
        payload := new(channels.PostChannelsTelegramBadRequestBody)
        payload.Version = sender.Version
        if errors != nil {
            payload.Errors = errors
        }

        payload.Message = message
        return channels.NewPostChannelsTelegramBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemTelegram) middleware.Responder {
        payload := new(channels.PostChannelsTelegramOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsTelegramOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        Params: struct {
            Token *string
        }{
            Token: params.Body.Token,
        },
    })
    if err == nil || err != storage.ErrChannelNotFound {
        return responseBadRequest(PayloadValidationErrorMessage, CustomValidationError("token", "unique"))
    }

    channel = storage.NewChannel()

    channel.Type = storage.ChannelTypeTelegram
    channel.AccessToken = accessToken.Token
    channel.Params.Token = *params.Body.Token
    channel.Params.GreetingText = *params.Body.GreetingText
    channel.Params.AnswerText = *params.Body.AnswerText
    channel.Params.AlternateText = *params.Body.AlternateText
    channel.Params.ButtonText = *params.Body.ButtonText

    channel, err = GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("inserting channel error: %s", err.Error()), nil)
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    if err := sender.TgPool.Add(&telegram.Channel{
        ID:            channel.ID,
        Token:         channel.Params.Token,
        GreetingText:  channel.Params.GreetingText,
        AnswerText:    channel.Params.AnswerText,
        AlternateText: channel.Params.AlternateText,
        ButtonText:    channel.Params.ButtonText,
    }); err != nil {
        return responseBadRequest(fmt.Sprintf("can not add telegram channel: %s", err.Error()), nil)
    }

    item := new(channels.DataItemTelegram)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.GreetingText = &channel.Params.GreetingText
    item.AnswerText = &channel.Params.AnswerText
    item.AlternateText = &channel.Params.AlternateText
    item.ButtonText = &channel.Params.ButtonText

    return responseSuccess(item)
}

// PostChannelsViberHandler - POST /channels/viber
func PostChannelsViberHandler(params channels.PostChannelsViberParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsViberForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsViberForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsViberBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsViberBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemViber) middleware.Responder {
        payload := new(channels.PostChannelsViberOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPostChannelsViberOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        Params: struct {
            Token *string
        }{
            Token: params.Body.Token,
        },
    })
    if err == nil || err != storage.ErrChannelNotFound {
        payload := new(channels.PostChannelsViberBadRequestBody)
        payload.Version = sender.Version
        payload.Errors = CustomValidationError("token", "unique")

        payload.Message = fmt.Sprintf("Validation error")
        return channels.NewPostChannelsViberBadRequest().WithPayload(payload)
    }

    // Create viber webhook on Bot Proxy
    botProxyResponse, err := CreateBotProxyWebHook(*params.Body.Token)
    if err != nil || *botProxyResponse.Payload.Message != PayloadSuccessMessage {
        return responseBadRequest(fmt.Sprintf("post bot-proxy webhook error: %s", err.Error()))
    }

    // Save channel to DB
    channel = storage.NewChannel()

    channel.Type = storage.ChannelTypeViber
    channel.AccessToken = accessToken.Token
    channel.Params.Token = *params.Body.Token
    channel.Params.Name = *params.Body.BotName
    channel.Params.Avatar = *params.Body.BotAvatar
    channel.Params.GreetingText = *params.Body.GreetingText
    channel.Params.AnswerText = *params.Body.AnswerText
    channel.Params.AlternateText = *params.Body.AlternateText
    channel.Params.ButtonText = *params.Body.ButtonText
    channel.Params.ClientID = botProxyResponse.Payload.Data[0].ID
    channel.Params.Login = botProxyResponse.Payload.Data[0].Login
    channel.Params.Password = botProxyResponse.Payload.Data[0].Password

    ch, err := GetStorage().StoreChannel(storage.StoreChannel{Channel: channel})
    if err != nil {
        return responseBadRequest(fmt.Sprintf("inserting channel error: %s", err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "CREATED",
    }).Info("Channel created")

    log.Debug(channel.JSONString())

    if err := sender.ViberPool.Add(&viber.Channel{
        ID:            ch.ID,
        ClientID:      ch.Params.ClientID,
        Login:         ch.Params.Login,
        Password:      ch.Params.Password,
        GreetingText:  ch.Params.GreetingText,
        ButtonText:    ch.Params.ButtonText,
        AlternateText: ch.Params.AlternateText,
        AnswerText:    ch.Params.AnswerText,
        Bot:           viberbot.New(ch.Params.Token, ch.Params.Name, ch.Params.Avatar),
    }); err != nil {
        return responseBadRequest(fmt.Sprintf("can not add viber bot: %s", err.Error()))
    }

    item := new(channels.DataItemViber)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.BotName = &channel.Params.Name
    item.BotAvatar = &channel.Params.Avatar
    item.GreetingText = &channel.Params.GreetingText
    item.AnswerText = &channel.Params.AnswerText
    item.AlternateText = &channel.Params.AlternateText
    item.ButtonText = &channel.Params.ButtonText

    return responseSuccess(item)
}

// PutChannelsEmailChannelIDHandler - PUT /channels/email/<ID>
func PutChannelsEmailChannelIDHandler(params channels.PutChannelsEmailChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsEmailChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message
        return channels.NewPutChannelsEmailChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsEmailChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewPutChannelsEmailChannelIDNotFound().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsEmailChannelIDBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message
        return channels.NewPutChannelsEmailChannelIDBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemEmail) middleware.Responder {
        payload := new(channels.PutChannelsEmailChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsEmailChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.Port = *params.Body.Port
    channel.Params.Server = *params.Body.Server
    channel.Params.SSL = *params.Body.Ssl
    if params.Body.ContentType != nil && *params.Body.ContentType != "" {
        channel.Params.ContentType = *params.Body.ContentType
    } else {
        channel.Params.ContentType = "text/html; charset=UTF-8"
    }

    if params.Body.From != nil {
        channel.Params.From = *params.Body.From
    }

    channel, err = GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemEmail)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.Port = &channel.Params.Port
    item.Server = &channel.Params.Server
    item.Ssl = &channel.Params.SSL
    item.ContentType = &channel.Params.ContentType

    if channel.Params.From != "" {
        item.From = &channel.Params.From
    }

    return responseSuccess(item)
}

// PutChannelsLocalSmsChannelIDHandler - PUT /channels/local_sms/<ID>
func PutChannelsLocalSmsChannelIDHandler(params channels.PutChannelsLocalSmsChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsLocalSmsChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsLocalSmsChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsLocalSmsChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewPutChannelsLocalSmsChannelIDNotFound().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsLocalSmsChannelIDBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsLocalSmsChannelIDBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemLocalSms) middleware.Responder {
        payload := new(channels.PutChannelsLocalSmsChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsLocalSmsChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.Port = *params.Body.Port
    channel.Params.Server = *params.Body.Server
    channel.Params.ModemID = *params.Body.ModemID
    channel.Params.Db = *params.Body.Db
    channel.Params.Limit = *params.Body.Limit

    channel, err = GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemLocalSms)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.Port = &channel.Params.Port
    item.Server = &channel.Params.Server
    item.ModemID = &channel.Params.ModemID
    item.Db = &channel.Params.Db
    item.Limit = &channel.Params.Limit

    return responseSuccess(item)
}

// PutChannelsMtsSmsChannelIDHandler - PUT /channels/mts_sms/<ID>
func PutChannelsMtsSmsChannelIDHandler(params channels.PutChannelsMtsSmsChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsMtsSmsChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsMtsSmsChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsMtsSmsChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewPutChannelsMtsSmsChannelIDNotFound().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsMtsSmsChannelIDBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsMtsSmsChannelIDBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemMtsSms) middleware.Responder {
        payload := new(channels.PutChannelsMtsSmsChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsMtsSmsChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Login = *params.Body.Login
    channel.Params.Password = *params.Body.Password
    channel.Params.From = *params.Body.From
    channel.Params.Limit = *params.Body.Limit

    channel, err = GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemMtsSms)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Login = &channel.Params.Login
    item.Password = &channel.Params.Password
    item.From = &channel.Params.From
    item.Limit = &channel.Params.Limit

    return responseSuccess(item)
}

// PutChannelsSLACKChannelIDHandler - PUT /channels/slack/<ID>
func PutChannelsSLACKChannelIDHandler(params channels.PutChannelsSLACKChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsSLACKChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsSLACKChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsSLACKChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewPutChannelsSLACKChannelIDNotFound().WithPayload(payload)
    }

    responseBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsSLACKChannelIDBadRequestBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsSLACKChannelIDBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemSLACK) middleware.Responder {
        payload := new(channels.PutChannelsSLACKChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsSLACKChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Token = *params.Body.Token
    channel.Params.Name = *params.Body.Name

    channel, err = GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())

    item := new(channels.DataItemSLACK)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.Name = &channel.Params.Name

    return responseSuccess(item)
}

// PutChannelsTelegramChannelIDHandler - PUT /channels/telegram/<ID>
func PutChannelsTelegramChannelIDHandler(params channels.PutChannelsTelegramChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsTelegramChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsTelegramChannelIDForbidden().WithPayload(payload)
    }

    responseBadRequest := func(message string, errors interface{}) middleware.Responder {
        payload := new(channels.PostChannelsTelegramBadRequestBody)
        payload.Version = sender.Version
        if errors != nil {
            payload.Errors = errors
        }

        payload.Message = message
        return channels.NewPostChannelsTelegramBadRequest().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsTelegramChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage

        return channels.NewPutChannelsTelegramChannelIDNotFound().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemTelegram) middleware.Responder {
        payload := new(channels.PutChannelsTelegramChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsTelegramChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        Params: struct {
            Token *string
        }{
            Token: params.Body.Token,
        },
    })
    if err == nil && channel.ID != "" && channel.ID != params.ChannelID {
        return responseBadRequest(PayloadValidationErrorMessage, CustomValidationError("token", "unique"))
    }

    channel, err = GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Token = *params.Body.Token
    channel.Params.GreetingText = *params.Body.GreetingText
    channel.Params.AnswerText = *params.Body.AnswerText
    channel.Params.AlternateText = *params.Body.AlternateText
    channel.Params.ButtonText = *params.Body.ButtonText

    ch, err := GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()), nil)
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())

    sender.TgPool.Delete(ch.ID)
    if err := sender.TgPool.Add(&telegram.Channel{
        ID:            ch.ID,
        Token:         ch.Params.Token,
        GreetingText:  ch.Params.GreetingText,
        AnswerText:    ch.Params.AnswerText,
        AlternateText: ch.Params.AlternateText,
        ButtonText:    ch.Params.ButtonText,
    }); err != nil {
        return responseBadRequest(fmt.Sprintf("can not add telegram channel: %s", err.Error()), nil)
    }

    item := new(channels.DataItemTelegram)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.GreetingText = &channel.Params.GreetingText
    item.AnswerText = &channel.Params.AnswerText
    item.AlternateText = &channel.Params.AlternateText
    item.ButtonText = &channel.Params.ButtonText

    return responseSuccess(item)
}

// PutChannelsViberChannelIDHandler - PUT /channels/viber/<ID>
func PutChannelsViberChannelIDHandler(params channels.PutChannelsViberChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsViberChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPutChannelsViberChannelIDForbidden().WithPayload(payload)
    }

    responseViberBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PostChannelsViberBadRequestBody) // TODO !!!!! POST???
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewPostChannelsViberBadRequest().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.PutChannelsViberChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = PayloadNotFoundErrorMessage
        return channels.NewPutChannelsViberChannelIDNotFound().WithPayload(payload)
    }

    responseChannelIDBadRequest := func(message string) middleware.Responder {
        payload := new(channels.PutChannelsViberChannelIDBadRequestBody)
        payload.Version = sender.Version

        payload.Message = message
        return channels.NewPutChannelsViberChannelIDBadRequest().WithPayload(payload)
    }

    responseSuccess := func(item *channels.DataItemViber) middleware.Responder {
        payload := new(channels.PutChannelsViberChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        payload.Data = append(payload.Data, item)

        return channels.NewPutChannelsViberChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        Params: struct {
            Token *string
        }{
            Token: params.Body.Token,
        },
    })
    if channel.ID != "" && channel.ID != params.ChannelID {
        return responseViberBadRequest(fmt.Sprintf("channel with token %s already existed", *params.Body.Token))
    }

    channel, err = GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    if err != nil {
        return responseNotFound()
    }

    if accessToken.Token != channel.AccessToken {
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    channel.Params.Name = *params.Body.BotName
    channel.Params.Avatar = *params.Body.BotAvatar
    channel.Params.GreetingText = *params.Body.GreetingText
    channel.Params.AnswerText = *params.Body.AnswerText
    channel.Params.AlternateText = *params.Body.AlternateText
    channel.Params.ButtonText = *params.Body.ButtonText

    if channel.Params.Token != *params.Body.Token {
        channel.Params.Token = *params.Body.Token

        // Update viber webhook on Bot Proxy
        botProxyResponse, err := UpdateBotProxyWebHook(channel.Params.ClientID, channel.Params.Token)
        if err != nil || *botProxyResponse.Payload.Message != PayloadSuccessMessage {
            return responseChannelIDBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
        }
    }

    ch, err := GetStorage().UpdateChannel(storage.UpdateChannel{
        ID:      channel.ID,
        Channel: channel,
    })
    if err != nil {
        return responseChannelIDBadRequest(fmt.Sprintf("updating channel %s error: %s", params.ChannelID, err.Error()))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "EDITED",
    }).Info("Channel edited")

    log.Debug(channel.JSONString())
    sender.ViberPool.Delete(channel.ID)
    if err := sender.ViberPool.Add(&viber.Channel{
        ID:            ch.ID,
        ClientID:      ch.Params.ClientID,
        Login:         ch.Params.Login,
        Password:      ch.Params.Password,
        GreetingText:  ch.Params.GreetingText,
        ButtonText:    ch.Params.ButtonText,
        AlternateText: ch.Params.AlternateText,
        AnswerText:    ch.Params.AnswerText,
        Bot:           viberbot.New(ch.Params.Token, ch.Params.Name, ch.Params.Avatar),
    }); err != nil {
        return responseViberBadRequest(fmt.Sprintf("can not add viber bot: %s", err.Error()))
    }

    item := new(channels.DataItemViber)
    item.ID = channel.ID
    item.Type = string(channel.Type)
    item.Token = &channel.Params.Token
    item.BotName = &channel.Params.Name
    item.BotAvatar = &channel.Params.Avatar
    item.GreetingText = &channel.Params.GreetingText
    item.AnswerText = &channel.Params.AnswerText
    item.AlternateText = &channel.Params.AlternateText
    item.ButtonText = &channel.Params.ButtonText

    return responseSuccess(item)
}

// DeleteChannelsChannelIDHandler - DELETE /channels/<ID>
func DeleteChannelsChannelIDHandler(params channels.DeleteChannelsChannelIDParams) middleware.Responder {
    responseForbidden := func(message string) middleware.Responder {
        payload := new(channels.DeleteChannelsChannelIDForbiddenBody)
        payload.Version = sender.Version
        payload.Message = message

        return channels.NewDeleteChannelsChannelIDForbidden().WithPayload(payload)
    }

    responseNotFound := func() middleware.Responder {
        payload := new(channels.DeleteChannelsChannelIDNotFoundBody)
        payload.Version = sender.Version
        payload.Message = fmt.Sprintf("Entity not found")

        return channels.NewDeleteChannelsChannelIDNotFound().WithPayload(payload)
    }

    responseSuccess := func() middleware.Responder {
        payload := new(channels.DeleteChannelsChannelIDOKBody)
        payload.Version = sender.Version
        message := PayloadSuccessMessage
        payload.Message = &message

        return channels.NewDeleteChannelsChannelIDOK().WithPayload(payload)
    }

    accessToken, err := GetStorage().GetAccessToken(storage.GetAccessToken{
        Token: params.AccessToken,
    })
    if err != nil {
        return responseForbidden(wrapGetAccessTokenErrorMessage(err))
    }

    channel, err := GetStorage().GetChannel(storage.GetChannel{
        ID: &params.ChannelID,
    })
    switch {
    case err != nil:
        return responseNotFound()
    case accessToken.Token != channel.AccessToken:
        return responseForbidden(wrapAccessDeniedErrorMessage(channel.ID))
    }

    err = GetStorage().DeleteMessageEvents(storage.DeleteMessageEvents{ChannelID: channel.ID})
    if err != nil {
        return responseForbidden(fmt.Sprintf("error deleting %s channels message events: %s", params.ChannelID, err))
    }

    log.WithFields(log.Fields{
        "context":   "API",
        "channelID": channel.ID,
        "status":    "DELETED",
    }).Info("Channel deleted")

    switch channel.Type {
    case storage.ChannelTypeViber:
        // Delete viber webhook on bot proxy
        DeleteBotProxyWebHook(channel.Params.ClientID)
    }

    channel, err = GetStorage().DeleteChannel(storage.DeleteChannel{ID: channel.ID})
    if err != nil {
        return responseForbidden(fmt.Sprintf("error deleting %s channel: %s", params.ChannelID, err.Error()))
    }

    switch channel.Type {
    case storage.ChannelTypeTelegram:
        sender.TgPool.Delete(channel.ID)
    case storage.ChannelTypeViber:
        sender.ViberPool.Delete(channel.ID)
    }

    return responseSuccess()
}

// GetChannelsItem is a struct for GetChannels methods
type GetChannelsItem struct {
    ID   string `bson:"uuid"`
    Type string `bson:"type" json:"type"`

    From          *string `bson:"from,omitempty" json:"from,omitempty"`
    Login         *string `bson:"login,omitempty" json:"login,omitempty"`
    Password      *string `bson:"password,omitempty" json:"password,omitempty"`
    Port          *int64  `bson:"port,omitempty" json:"port,omitempty"`
    Server        *string `bson:"server,omitempty" json:"server,omitempty"`
    SSL           *bool   `bson:"ssl,omitempty" json:"ssl,omitempty"`
    ModemID       *int64  `bson:"modemID,omitempty" json:"modemID,omitempty"`
    Db            *string `bson:"db,omitempty" json:"db,omitempty"`
    Token         *string `bson:"token,omitempty" json:"token,omitempty"`
    GreetingText  *string `bson:"greetingText,omitempty" json:"greetingText,omitempty"`
    AnswerText    *string `bson:"answerText,omitempty" json:"answerText,omitempty"`
    AlternateText *string `bson:"alternateText,omitempty" json:"alternateText,omitempty"`
    ButtonText    *string `bson:"buttonText,omitempty" json:"buttonText,omitempty"`
    BotName       *string `bson:"botName,omitempty" json:"botName,omitempty"`
    BotAvatar     *string `bson:"botAvatar,omitempty" json:"botAvatar,omitempty"`
    Name          *string `bson:"name,omitempty" json:"name,omitempty"`
    Limit         *int64  `bson:"limit,omitempty" json:"limit,omitempty"`
    ContentType   *string `bson:"contentType,omitempty" json:"contentType,omitempty"`
}

func channelToItem(channel *storage.Channel) *GetChannelsItem {
    item := new(GetChannelsItem)

    item.ID = channel.ID
    item.Type = string(channel.Type)

    switch channel.Type {
    case storage.ChannelTypeEmail:
        item.Login = &channel.Params.Login
        item.Password = &channel.Params.Password
        item.Port = &channel.Params.Port
        item.Server = &channel.Params.Server
        item.SSL = &channel.Params.SSL
        item.ContentType = &channel.Params.ContentType

        if channel.Params.From != "" {
            item.From = &channel.Params.From
        }
    case storage.ChannelTypeLocalSMS:
        item.Login = &channel.Params.Login
        item.Password = &channel.Params.Password
        item.Port = &channel.Params.Port
        item.Server = &channel.Params.Server
        item.ModemID = &channel.Params.ModemID
        item.Db = &channel.Params.Db
        item.Limit = &channel.Params.Limit
    case storage.ChannelTypeMTSSMS:
        item.From = &channel.Params.From
        item.Login = &channel.Params.Login
        item.Password = &channel.Params.Password
        item.Limit = &channel.Params.Limit
    case storage.ChannelTypeSlack:
        item.Token = &channel.Params.Token
        item.Name = &channel.Params.Name
    case storage.ChannelTypeTelegram:
        item.Token = &channel.Params.Token
        item.GreetingText = &channel.Params.GreetingText
        item.AnswerText = &channel.Params.AnswerText
        item.AlternateText = &channel.Params.AlternateText
        item.ButtonText = &channel.Params.ButtonText
    case storage.ChannelTypeViber:
        item.Token = &channel.Params.Token
        item.GreetingText = &channel.Params.GreetingText
        item.AnswerText = &channel.Params.AnswerText
        item.AlternateText = &channel.Params.AlternateText
        item.ButtonText = &channel.Params.ButtonText
        item.BotName = &channel.Params.Name
        item.BotAvatar = &channel.Params.Avatar
    }

    return item
}
