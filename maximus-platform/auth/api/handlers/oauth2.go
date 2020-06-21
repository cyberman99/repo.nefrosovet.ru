package handlers

import (
	"fmt"

	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	xOAuth2 "golang.org/x/oauth2"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/auth"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/oauth2"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/oauth2/esia"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	NoSuchBackendIDErrorMessage = "There is no such backend '%v'"
)

// AuthGetOAuth2BackendID - GET /oauth2/{BackendID}
func AuthGetOAuth2BackendID(params auth.GetOauth2BackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalServerError(err)
	} else if len(backends) == 0 {
		return responseInternalServerError(storage.ErrBackendNotFound)
	}

	payload := new(auth.GetOauth2BackendIDOKBody)
	switch backends[0].Provider {
	case storage.BackendOAuth2ProviderGoogle:
		p := oauth2.GoogleConfig.AuthCodeURL("", xOAuth2.AccessTypeOffline)
		// p := getOAuth2GooglePath(params.RedirectURI, "") // TODO: provide useful state
		payload.OAuth2Path = &p

		return auth.NewGetOauth2BackendIDOK().WithPayload(payload)
	case storage.BackendOAuth2ProviderESIA:
		timestamp := esia.Timestamp()
		state := esia.State()

		clientSecret, err := esia.Secret(backends[0].ClientID, backends[0].ClientSecret, oauth2.EsiaConfig.Scopes, timestamp, state)
		if err != nil {
			return responseInternalServerError(err)
		}

		url := oauth2.EsiaConfig.AuthCodeURL(
			state,
			xOAuth2.AccessTypeOffline,
			xOAuth2.SetAuthURLParam("client_id", backends[0].ClientID),
			xOAuth2.SetAuthURLParam("client_secret", clientSecret),
			xOAuth2.SetAuthURLParam("timestamp", timestamp),
			xOAuth2.SetAuthURLParam("redirect_uri", params.RedirectURI),
		)
		payload.OAuth2Path = &url

		return auth.NewGetOauth2BackendIDOK().WithPayload(payload)
	default:
		payload := new(auth.GetOauth2BackendIDBadRequestBody)
		payload.Version = &Version
		payload.Message = &PayloadFailMessage
		payload.Errors = &auth.GetOauth2BackendIDBadRequestBodyAO1Errors{
			Core: fmt.Sprintf(NoSuchBackendIDErrorMessage, params.BackendID),
		}

		return auth.NewGetOauth2BackendIDBadRequest().WithPayload(payload)
	}
}

// POST: /oauth2/{backendID}/consent
func AuthPostOAuth2Consent(params auth.PostOauth2BackendIDConsentParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(jwt *jwt.JWT) middleware.Responder {
		payload := new(auth.PostOauth2BackendIDConsentOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(auth.DataItems0)
		item.AccessToken = &jwt.AccessToken
		item.RefreshToken = &jwt.RefreshToken

		payload.Data = append(payload.Data, item)

		return auth.NewPostOauth2BackendIDConsentOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalServerError(err)
	} else if len(backends) == 0 {
		return responseInternalServerError(storage.ErrBackendNotFound)
	}

	sourceIP := getSourceIP(params.HTTPRequest)

	res := authentication.Auth(&oauth2.Credentials{
		Code:        params.Body.AuthorizationCode,
		RedirectURI: params.Body.RedirectURI,
		Backend:     backends[0],
	})
	es := st.GetStorage().EventStorage
	if res.Err() != nil {
		err := es.Store(storage.StoreEvent{
			EventType: "OAUTH2",
			SourceIP:  sourceIP,
			EntityID:  res.EntityID,
			Status:    PayloadFailMessage,
		})

		return responseInternalServerError(err)
	}

	es.Store(storage.StoreEvent{
		EventType: "OAUTH2",
		SourceIP:  sourceIP,
		EntityID:  res.EntityID,
		Status:    PayloadSuccessMessage,
	})

	return responseSuccess(res.JWT)
}
