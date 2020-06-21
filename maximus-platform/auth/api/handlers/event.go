package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/event"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

// GetEvents - GET /events
func GetEvents(_ event.GetEventsParams) middleware.Responder {
	responseNotFound := func() middleware.Responder {
		//TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
		}).Error("Not found")

		panic(nil)
	}

	responseInternalError := func(err error) middleware.Responder {
		//TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(events []*storage.Event) middleware.Responder {
		payload := new(event.GetEventsOKBody)
		payload.Version = &Version

		for _, eventItem := range events {
			item := new(event.DataItems0)
			item.ID = &eventItem.ID
			item.Created = &eventItem.Time
			item.IP = eventItem.SourceIP
			item.Type = &eventItem.Type
			item.EntityID = &eventItem.EntityID
			item.EntityLogin = &eventItem.EntityLogin
			item.Status = &eventItem.Status

			payload.Data = append(payload.Data, item)
		}

		payload.Message = &PayloadSuccessMessage
		return event.NewGetEventsOK().WithPayload(payload)
	}

	es := st.GetStorage().EventStorage
	events, err := es.GetAll(storage.GetEvents{})
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(events)
}

// GetEventsEventID - GET /events
func GetEventsEventID(params event.GetEventsEventIDParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		//TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(event.GetEventsEventIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return event.NewGetEventsEventIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(eventItem *storage.Event) middleware.Responder {
		payload := new(event.GetEventsEventIDOKBody)
		payload.Version = &Version

		item := new(event.DataItems0)
		item.ID = &eventItem.ID
		item.Created = &eventItem.Time
		item.IP = eventItem.SourceIP
		item.Type = &eventItem.Type
		item.EntityID = &eventItem.EntityID
		item.EntityLogin = &eventItem.EntityLogin
		item.Status = &eventItem.Status

		payload.Data = append(payload.Data, item)

		payload.Message = &PayloadSuccessMessage
		return event.NewGetEventsEventIDOK().WithPayload(payload)
	}

	es := st.GetStorage().EventStorage
	// params can be used later
	eventItem, err := es.Get(storage.GetEvent{ID: params.EventID})
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(eventItem)
}
