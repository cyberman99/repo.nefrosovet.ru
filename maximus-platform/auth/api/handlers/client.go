package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/client"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

func GetClients(_ client.GetClientsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(clients []*storage.ClientStorer) middleware.Responder {
		payload := new(client.GetClientsOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		for _, clientItem := range clients {
			item := new(client.DataItems0)
			item.ID = clientItem.ID
			item.Description = &clientItem.Descriptions

			payload.Data = append(payload.Data, item)
		}

		return client.NewGetClientsOK().WithPayload(payload)
	}

	cs := st.GetStorage().ClientStorage
	allClients, err := cs.Get(storage.ClientFilter{})
	if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(allClients)
}

func GetClientsClientID(params client.GetClientsClientIDParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(client.GetClientsClientIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, err.Error())
		payload.Message = &NotFoundMessage
		return client.NewGetClientsClientIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(el *storage.ClientStorer) middleware.Responder {
		payload := new(client.GetClientsClientIDOKBody)
		payload.Version = &Version

		item := new(client.DataItems0)
		item.ID = el.ID
		item.Description = &el.Descriptions

		payload.Data = append(payload.Data, item)
		payload.Message = &PayloadSuccessMessage
		return client.NewGetClientsClientIDOK().WithPayload(payload)
	}

	cs := st.GetStorage().ClientStorage
	clients, err := cs.Get(storage.ClientFilter{ID: &params.ClientID})
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound(err)
	} else if err != nil {
		return responseInternalError(err)
	}
	client := clients[0]

	return responseSuccess(client)
}

func PostClients(params client.PostClientsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(el storage.ClientStorer) middleware.Responder {
		payload := new(client.PostClientsOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		plItem := new(client.DataItems0)
		plItem.ID = el.ID
		plItem.Description = &el.Descriptions

		payload.Data = append(payload.Data, plItem)

		return client.NewPostClientsOK().WithPayload(payload)
	}

	cs := st.GetStorage().ClientStorage
	us := st.GetStorage().UserStorage

	key, err := jwt.HashPassword(*params.Body.Password)
	if err != nil {
		return responseInternalError(err)
	}

	newClient := storage.ClientStorer{
		ID:           uuid.New().String(),
		Descriptions: *params.Body.Description,
		Password:     key,
	}

	err = cs.Store(newClient)
	if err != nil {
		return responseInternalError(err)
	}

	in := storage.StoreUser{
		User: storage.User{
			ID:    newClient.ID,
			Roles: map[string]bool{storage.RoleDefaultClient: true},
		},
	}

	if _, err := us.Store(in); err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(newClient)
}

func PutClientsClientID(params client.PutClientsClientIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(client.PutClientsClientIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, err.Error())
		payload.Message = &NotFoundMessage

		return client.NewPutClientsClientIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(el *storage.ClientStorer) middleware.Responder {
		payload := new(client.PutClientsClientIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		plItem := new(client.DataItems0)
		plItem.ID = el.ID
		plItem.Description = &el.Descriptions

		payload.Data = append(payload.Data, plItem)

		return client.NewPutClientsClientIDOK().WithPayload(payload)
	}

	cs := st.GetStorage().ClientStorage
	key, err := jwt.HashPassword(*params.Body.Password)
	if err != nil {
		return responseInternalServerError(err)
	}
	cl := storage.ClientStorer{
		ID:           params.ClientID,
		Descriptions: *params.Body.Description,
		Password:     key,
	}
	if err := cs.Update(cl.ID, storage.ClientUpdater{
		Descriptions: &cl.Descriptions,
		Password:     &cl.Password,
	}); err != nil && err == storage.ErrNotFound {
		return responseNotFound(err)
	} else if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(&cl)
}

func PatchClientsClientID(params client.PatchClientsClientIDParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		payload := new(client.PatchClientsClientIDInternalServerErrorBody)
		payload.Version = &Version
		payload.Message = &InternalServerErrorMessage
		return client.NewPatchClientsClientIDInternalServerError().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(client.PatchClientsClientIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return client.NewPatchClientsClientIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(clientItem *storage.ClientStorer) middleware.Responder {
		payload := new(client.PatchClientsClientIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		plItem := new(client.DataItems0)
		plItem.ID = clientItem.ID
		plItem.Description = &clientItem.Descriptions

		payload.Data = append(payload.Data, plItem)

		return client.NewPatchClientsClientIDOK().WithPayload(payload)
	}

	us := st.GetStorage().ClientStorage
	clients, err := us.Get(storage.ClientFilter{ID: &params.ClientID})
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	client := clients[0]
	if params.Body.Description != nil {
		client.Descriptions = *params.Body.Description
	}

	if params.Body.Password != nil {
		key, err := jwt.HashPassword(*params.Body.Password)
		if err != nil {
			return responseInternalError(err)
		}
		client.Password = key
	}

	err2 := us.Update(client.ID, storage.ClientUpdater{
		Descriptions: &client.Descriptions,
		Password:     &client.Password,
	})
	if err2 != nil {
		return responseInternalError(err2)
	}

	return responseSuccess(client)
}

func DeleteClientsClientID(params client.DeleteClientsClientIDParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(client.DeleteClientsClientIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, err.Error())
		payload.Message = &NotFoundMessage
		return client.NewDeleteClientsClientIDNotFound().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(client.DeleteClientsClientIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		return client.NewDeleteClientsClientIDOK().WithPayload(payload)
	}

	cs := st.GetStorage().ClientStorage
	err := cs.Delete(params.ClientID)
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound(err)
	} else if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess()
}
