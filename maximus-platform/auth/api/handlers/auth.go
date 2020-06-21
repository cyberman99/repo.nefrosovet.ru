package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/auth"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/client"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

// PostAuthClient - POST /auth/client
func PostAuthClient(params auth.PostClientParams) middleware.Responder {
	responseClientUnauthorized := func() middleware.Responder  {
		payload := new(auth.PostClientUnauthorizedBody)
		payload.Version = &Version
		payload.Message = &PayloadAuthFailureMessage

		return auth.NewPostClientUnauthorized().WithPayload(payload)
	}

	responseSuccess := func(jwt *jwt.JWT) middleware.Responder  {
		payload := new(auth.PostClientOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(auth.DataItems0)
		item.AccessToken = &jwt.AccessToken
		item.RefreshToken = &jwt.RefreshToken

		payload.Data = append(payload.Data, item)

		return auth.NewPostClientOK().WithPayload(payload)
	}

	sourceIP := getSourceIP(params.HTTPRequest)

	res := authentication.Auth(&client.Credentials{
		Login:    *params.Body.ClientID,
		Password: *params.Body.Password,
	})
	es := st.GetStorage().EventStorage
	if res.Err() != nil {
		es.Store(storage.StoreEvent{
			EventType:   "CLIENT_LOGIN",
			SourceIP:    sourceIP,
			EntityID:    res.EntityID,
			EntityLogin: res.EntityLogin,
			Status:      PayloadFailMessage,
			Data:        res.Err().Error(),
		})

		return responseClientUnauthorized()
	}

	es.Store(storage.StoreEvent{
		EventType:   "CLIENT_LOGIN",
		SourceIP:    sourceIP,
		EntityID:    res.EntityID,
		EntityLogin: res.EntityLogin,
		Status:      PayloadSuccessMessage,
	})

	return responseSuccess(res.JWT)
}

// PostUser - POST /user
func PostUser(params auth.PostUserParams) middleware.Responder {
	responseUserUnauthorized := func() middleware.Responder  {
		payload := new(auth.PostUserUnauthorizedBody)
		payload.Version = &Version
		payload.Message = &PayloadAuthFailureMessage

		return auth.NewPostUserUnauthorized().WithPayload(payload)
	}

	responseSuccess := func(jwt *jwt.JWT) middleware.Responder {
		payload := new(auth.PostUserOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(auth.DataItems0)
		item.AccessToken = &jwt.AccessToken
		item.RefreshToken = &jwt.RefreshToken

		payload.Data = append(payload.Data, item)

		return auth.NewPostUserOK().WithPayload(payload)
	}

	sourceIP := getSourceIP(params.HTTPRequest)

	res := authentication.Auth(&login.Credentials{
		Login:    *params.Body.Login,
		Password: *params.Body.Password,
	})
	es := st.GetStorage().EventStorage
	if res.Err() != nil {
		es.Store(storage.StoreEvent{
			EventType:   "USER_LOGIN",
			SourceIP:    sourceIP,
			EntityID:    res.EntityID,
			EntityLogin: res.EntityLogin,
			Status:      PayloadFailMessage,
		})

		return responseUserUnauthorized()
	}

	es.Store(storage.StoreEvent{
		EventType:   "USER_LOGIN",
		SourceIP:    sourceIP,
		EntityID:    res.EntityID,
		EntityLogin: res.EntityLogin,
		Status:      PayloadSuccessMessage,
	})

	return responseSuccess(res.JWT)
}

// PostRefresh - POST /auth/refresh
func PostRefresh(params auth.PostRefreshParams) middleware.Responder {
	responseRefreshUnauthorized  := func() middleware.Responder {
		payload := new(auth.PostRefreshUnauthorizedBody)
		payload.Version = &Version
		payload.Message = &PayloadAuthFailureMessage

		return auth.NewPostRefreshUnauthorized().WithPayload(payload)
	}

	responseSuccess := func(jwt *jwt.JWT) middleware.Responder {
		payload := new(auth.PostRefreshOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(auth.DataItems0)
		item.AccessToken = &jwt.AccessToken
		item.RefreshToken = &jwt.RefreshToken

		payload.Data = append(payload.Data, item)

		return auth.NewPostRefreshOK().WithPayload(payload)
	}

	sourceIP := getSourceIP(params.HTTPRequest)

	res := authentication.Refresh(*params.Body.RefreshToken)
	es := st.GetStorage().EventStorage
	if res.Err() != nil {
		es.Store(storage.StoreEvent{
			EventType:   "REFRESH_TOKEN",
			SourceIP:    sourceIP,
			EntityID:    res.EntityID,
			EntityLogin: res.EntityLogin,
			Status:      PayloadFailMessage,
			Data:        res.Err().Error(),
		})

		return responseRefreshUnauthorized()
	}

	es.Store(storage.StoreEvent{
		EventType:   "REFRESH_TOKEN",
		SourceIP:    sourceIP,
		EntityID:    res.EntityID,
		EntityLogin: res.EntityLogin,
		Status:      PayloadSuccessMessage,
	})

	return responseSuccess(res.JWT)
}

// PostIdentify - POST /identify
func PostIdentify(params auth.PostIdentifyParams) middleware.Responder {
	responseIdentifyUnauthorized := func() middleware.Responder {
		payload := new(auth.PostIdentifyUnauthorizedBody)
		payload.Version = &Version
		payload.Message = &PayloadAuthFailureMessage

		return auth.NewPostIdentifyUnauthorized().WithPayload(payload)
	}

	responseSuccess := func(jwt *jwt.JWT) middleware.Responder {
		payload := new(auth.PostIdentifyOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(auth.DataItems0)
		item.AccessToken = &jwt.AccessToken
		item.RefreshToken = &jwt.RefreshToken

		payload.Data = append(payload.Data, item)

		return auth.NewPostIdentifyOK().WithPayload(payload)
	}

	sourceIP := getSourceIP(params.HTTPRequest)

	res := authentication.Auth(&login.Credentials{
		SmartCardNumber: *params.Body.CardNumber,
	})
	es := st.GetStorage().EventStorage
	if res.Err() != nil {
		es.Store(storage.StoreEvent{
			EventType:   "DRIVER_LOGIN",
			SourceIP:    sourceIP,
			EntityID:    res.EntityID,
			EntityLogin: res.EntityLogin,
			Status:      PayloadFailMessage,
			Data:        res.Err().Error(),
		})

		return responseIdentifyUnauthorized()
	}

	es.Store(storage.StoreEvent{
		EventType:   "DRIVER_LOGIN",
		SourceIP:    sourceIP,
		EntityID:    res.EntityID,
		EntityLogin: res.EntityLogin,
		Status:      PayloadSuccessMessage,
	})

	return responseSuccess(res.JWT)
}
