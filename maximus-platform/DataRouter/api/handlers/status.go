package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/status"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
)

type StatusViewController struct {
	db mongod.Storer
}

// NewStatusView is event controller constructor.
func NewStatusView(store mongod.Storer) *StatusViewController {
	return &StatusViewController{
		db: store,
	}
}

func (sw *StatusViewController) Get(params status.StatusViewParams) middleware.Responder {
	var err error
	var ctx = context.Background()

	if err = sw.db.Status(); err != nil {
		payload := new(status.StatusViewInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return status.NewStatusViewInternalServerError().WithPayload(payload)
	}

	coll := sw.db.Collection("admin")

	_, err = coll.InsertOne(ctx, bson.M{"status": "ok"})
	if err != nil {
		payload := new(status.StatusViewInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return status.NewStatusViewInternalServerError().WithPayload(payload)
	}

	err = coll.FindOneAndDelete(ctx, bson.M{"status": "ok"}).Err()
	if err != nil {
		payload := new(status.StatusViewInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return status.NewStatusViewInternalServerError().WithPayload(payload)
	}

	payload := new(status.StatusViewOKBody)
	payload.Version = &Version
	payload.Data = []interface{}{}
	payload.Message = PayloadSuccessMessage
	return status.NewStatusViewOK().WithPayload(payload)
}
