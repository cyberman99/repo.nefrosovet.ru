package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/service"
	"repo.nefrosovet.ru/maximus-platform/mailer/sender"
)

// GetStatusHandler - GET /status
func GetStatusHandler(_ service.GetStatusParams) middleware.Responder {
	responseInternalServerError := func() middleware.Responder {
		payload := new(service.GetStatusInternalServerErrorBody)
		payload.Version = sender.Version
		payload.Message = PayloadInternalServerErrorMessage

		return service.NewGetStatusInternalServerError().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(service.GetStatusOKBody)
		payload.Version = sender.Version
		message := PayloadSuccessMessage
		payload.Message = &message

		return service.NewGetStatusOK().WithPayload(payload)
	}

	if err := GetStorage().Ping(); err != nil {
		return responseInternalServerError()
	}

	return responseSuccess()
}
