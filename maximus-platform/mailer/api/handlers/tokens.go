package handlers

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/manage"
	"repo.nefrosovet.ru/maximus-platform/mailer/sender"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

// GetTokensHandler - GET /tokens
func GetTokensHandler(params manage.GetTokensParams) middleware.Responder {
	responseForbidden := func(message string) middleware.Responder {
		payload := new(manage.GetTokensForbiddenBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewGetTokensForbidden().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(manage.GetTokensTokenIDNotFoundBody)
		payload.Version = sender.Version
		payload.Message = PayloadNotFoundErrorMessage

		return manage.NewGetTokensTokenIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(tokens []storage.AccessToken) middleware.Responder {
		payload := new(manage.GetTokensOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message
		payload.Errors = new([]error)
		payload.Data = []*manage.DataItems0{}

		for i := range tokens {
			token := tokens[i]

			item := new(manage.DataItems0)
			item.ID = token.Token
			item.Description = &token.Description

			payload.Data = append(payload.Data, item)
		}

		return manage.NewGetTokensOK().WithPayload(payload)
	}

	if params.MasterToken != viper.GetString("masterToken") {
		return responseForbidden("access denied")
	}

	tokens, err := GetStorage().GetAccessTokens(storage.GetAccessTokens{})
	if err != nil {
		return responseNotFound()
	}

	return responseSuccess(tokens)
}

// PostTokensHandler - POST /tokens
func PostTokensHandler(params manage.PostTokensParams) middleware.Responder {
	responseForbidden := func(message string) middleware.Responder {
		payload := new(manage.PostTokensForbiddenBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewPostTokensForbidden().WithPayload(payload)
	}

	responseBadRequest := func(message string) middleware.Responder {
		payload := new(manage.PostTokensBadRequestBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewPostTokensBadRequest().WithPayload(payload)
	}

	responseSuccess := func(item *manage.DataItems0) middleware.Responder {
		payload := new(manage.PostTokensOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message
		payload.Errors = new([]error)

		payload.Data = append(payload.Data, item)

		return manage.NewPostTokensOK().WithPayload(payload)
	}

	if params.MasterToken != viper.GetString("masterToken") {
		return responseForbidden("access denied")
	}

	token := storage.NewAccessToken()
	token.Description = *params.Body.Description

	token, err := GetStorage().StoreAccessToken(storage.StoreAccessToken{AccessToken: token})
	if err != nil {
		return responseBadRequest(fmt.Sprintf("inserting access token error: %s", err.Error()))
	}

	log.WithFields(log.Fields{
		"context": "API",
		"tokenID": token.Token,
		"status":  "CREATED",
	}).Info("Access token created")

	log.Debug(token.JSONString())

	item := new(manage.DataItems0)
	item.ID = token.Token
	item.Description = &token.Description

	return responseSuccess(item)
}

// GetTokensTokenIDHandler - GET /tokens/<ID>
func GetTokensTokenIDHandler(params manage.GetTokensTokenIDParams) middleware.Responder {
	responseForbidden := func(message string) middleware.Responder {
		payload := new(manage.GetTokensTokenIDForbiddenBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewGetTokensTokenIDForbidden().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(manage.GetTokensTokenIDNotFoundBody)
		payload.Version = sender.Version
		message := PayloadNotFoundErrorMessage
		payload.Message = message

		return manage.NewGetTokensTokenIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(token storage.AccessToken) middleware.Responder {
		payload := new(manage.GetTokensOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message
		payload.Errors = new([]error)

		item := new(manage.DataItems0)
		item.ID = token.Token
		item.Description = &token.Description

		payload.Data = append(payload.Data, item)

		return manage.NewGetTokensOK().WithPayload(payload)
	}

	if params.MasterToken != viper.GetString("masterToken") {
		return responseForbidden("access denied")
	}

	token, err := GetStorage().GetAccessToken(storage.GetAccessToken{Token: params.TokenID})
	if err != nil {
		return responseNotFound()
	}

	return responseSuccess(token)
}

// PutTokensTokenIDHandler - PUT /tokens/<ID>
func PutTokensTokenIDHandler(params manage.PutTokensTokenIDParams) middleware.Responder {
	responseForbidden := func(message string) middleware.Responder {
		payload := new(manage.PutTokensTokenIDForbiddenBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewPutTokensTokenIDForbidden().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(manage.GetTokensTokenIDNotFoundBody)
		payload.Version = sender.Version
		payload.Message = PayloadNotFoundErrorMessage

		return manage.NewGetTokensTokenIDNotFound().WithPayload(payload)
	}

	responseBadRequest := func(message string) middleware.Responder {
		payload := new(manage.PutTokensTokenIDBadRequestBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewPutTokensTokenIDBadRequest().WithPayload(payload)
	}

	responseSuccess := func(item *manage.DataItems0) middleware.Responder {
		payload := new(manage.PutTokensTokenIDOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message
		payload.Errors = new([]error)

		payload.Data = append(payload.Data, item)

		return manage.NewPutTokensTokenIDOK().WithPayload(payload)
	}

	if params.MasterToken != viper.GetString("masterToken") {
		return responseForbidden("access denied")
	}

	token, err := GetStorage().GetAccessToken(storage.GetAccessToken{Token: params.TokenID})
	if err != nil {
		return responseNotFound()
	}

	token.Description = *params.Body.Description

	token, err = GetStorage().UpdateAccessToken(storage.UpdateAccessToken{
		Token:       token.Token,
		AccessToken: token,
	})
	if err != nil {
		return responseBadRequest(fmt.Sprintf("updating access token error: %s", err.Error()))
	}

	log.WithFields(log.Fields{
		"context": "API",
		"tokenID": token.Token,
		"status":  "EDITED",
	}).Info("Access token edited")

	log.Debug(token.JSONString())

	item := new(manage.DataItems0)
	item.ID = token.Token
	item.Description = &token.Description

	return responseSuccess(item)
}

// DeleteTokensTokenIDHandler - DELETE /tokens/<ID>
func DeleteTokensTokenIDHandler(params manage.DeleteTokensTokenIDParams) middleware.Responder {
	responseForbidden := func(message string) middleware.Responder {
		payload := new(manage.DeleteTokensTokenIDForbiddenBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewDeleteTokensTokenIDForbidden().WithPayload(payload)
	}

	responseNotFound := func(message string) middleware.Responder {
		payload := new(manage.DeleteTokensTokenIDNotFoundBody)
		payload.Version = sender.Version
		payload.Message = message

		return manage.NewDeleteTokensTokenIDNotFound().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(manage.DeleteTokensTokenIDOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message
		payload.Errors = new([]error)

		payload.Data = make([]*manage.DataItems0, 0)

		return manage.NewDeleteTokensTokenIDOK().WithPayload(payload)
	}

	if params.MasterToken != viper.GetString("masterToken") {
		return responseForbidden("access denied")
	}

	token, err := GetStorage().GetAccessToken(storage.GetAccessToken{Token: params.TokenID})
	if err != nil {
		return responseNotFound(PayloadNotFoundErrorMessage)
	}

	channels, err := GetStorage().GetChannels(storage.GetChannels{
		AccessToken: &params.TokenID,
	})
	if err != nil {
		return responseForbidden(fmt.Sprintf("deleting access token channels error: %s", err.Error()))
	}

	for _, channel := range channels {
		err = GetStorage().DeleteMessageEvents(storage.DeleteMessageEvents{ChannelID: channel.ID})
		if err != nil {
			return responseForbidden(fmt.Sprintf("deleting channels message events error: %s", err.Error()))
		}

		_, err = GetStorage().DeleteChannel(storage.DeleteChannel{ID: channel.ID})
		if err != nil {
			return responseForbidden(fmt.Sprintf("deleting access token channels error: %s", err.Error()))
		}
	}

	_, err = GetStorage().DeleteAccessToken(storage.DeleteAccessToken{Token: token.Token})
	if err != nil {
		return responseNotFound("Deleting token error: " + err.Error())
	}

	log.WithFields(log.Fields{
		"context": "API",
		"tokenID": token.Token,
		"status":  "DELETED",
	}).Info("Access token deleted")

	return responseSuccess()
}
