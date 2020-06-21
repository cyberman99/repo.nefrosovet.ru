package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/admin"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var settingAdminPasswordError = "setting admin password error"
var gettingAdminPasswordError = "getting admin password error"
var storingAdminPasswordError = "storing admin password error"

// PostAdmin - POST /admin
func PostAdmin(params admin.PostAdminParams) middleware.Responder {
	responseBadRequest := func(err error, msg string) middleware.Responder {
		payload := new(admin.PostAdminBadRequestBody)
		payload.Version = &Version
		payload.Errors = new(admin.PostAdminBadRequestBodyAO1Errors)
		payload.Errors.Core = err.Error()
		payload.Message = &msg
		return admin.NewPostAdminBadRequest().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(admin.PostAdminOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		return admin.NewPostAdminOK().WithPayload(payload)
	}

	key, err := jwt.HashPassword(*params.Body.Password)
	if err != nil {
		return responseBadRequest(err, settingAdminPasswordError)
	}

	ps := st.GetStorage().AdminPasswordStorage
	_, err = ps.Get()
	if err != nil {
		return responseBadRequest(err, gettingAdminPasswordError)
	}

	_, err = ps.Update(storage.UpdateAdminPassword{
		Hash: key,
	})
	if err != nil {
		return responseBadRequest(err, storingAdminPasswordError)
	}

	return responseSuccess()
}
