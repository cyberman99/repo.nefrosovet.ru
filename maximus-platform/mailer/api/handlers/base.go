package handlers

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/viper"
	botproxyclient "repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/client"
	"repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/client/viber"
	"repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/client/webhook"
	botproxymodels "repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/models"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
)

const (
	PayloadSuccessMessage             = "SUCCESS"
	PayloadInternalServerErrorMessage = "Internal server error"
	PayloadValidationErrorMessage     = "Validation error"
	PayloadNotFoundErrorMessage       = "Entity not found"
	PayloadAccessDeniedErrorMessage   = "Access denied"
)

func wrapGetAccessTokenErrorMessage(err error) string {
	return fmt.Sprintf("checking access token error: %s", err)
}

func wrapGetMessageEventsErrorMessage(err error) string {
	return fmt.Sprintf("get message events error: %s", err)
}

func wrapAccessDeniedErrorMessage(channelID string) string {
	return fmt.Sprintf("access denied to channel (id %s)", channelID)
}

// GetStorage returns default storage implementations
func GetStorage() *_default.Storage {
	return _default.GetStorage()
}

// CreateBotProxyWebHook sends request to bot proxy for webhook create
func CreateBotProxyWebHook(token string) (*viber.PostWebhooksViberOK, error) {
	client := getBotProxyClient()
	params := viber.NewPostWebhooksViberParams().WithBody(&botproxymodels.PostWebhooksViberParamsBody{})
	params.Body.Token = &token

	return client.Viber.PostWebhooksViber(params)
}

// UpdateBotProxyWebHook sends request to bot proxy for webhook update
func UpdateBotProxyWebHook(id, token string) (*viber.PutWebhooksViberWebhookIDOK, error) {
	client := getBotProxyClient()
	params := viber.NewPutWebhooksViberWebhookIDParams().WithBody(&botproxymodels.PutWebhooksViberWebhookIDParamsBody{})
	params.Body.Token = &token
	params.WebhookID = id

	return client.Viber.PutWebhooksViberWebhookID(params)
}

// DeleteBotProxyWebHook sends request to bot proxy for webhook delete
func DeleteBotProxyWebHook(id string) (*webhook.DeleteWebhooksWebhookIDOK, error) {
	client := getBotProxyClient()
	params := webhook.NewDeleteWebhooksWebhookIDParams()
	params.WebhookID = id

	return client.Webhook.DeleteWebhooksWebhookID(params)
}

func getBotProxyClient() *botproxyclient.BotProxy {
	return botproxyclient.NewHTTPClientWithConfig(strfmt.Default, &botproxyclient.TransportConfig{
		Host:     viper.GetString("botProxy.http.host"),
		BasePath: viper.GetString("botProxy.http.path"),
		Schemes:  []string{"http", "https"},
	})
}

func CustomValidationError(field, reason string) map[string]interface{} {
	errors := make(map[string]interface{})
	errors[field] = reason

	validation := make(map[string]interface{})
	validation["validation"] = errors

	return validation
}
