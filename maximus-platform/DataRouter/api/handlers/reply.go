package handlers

import (
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	//errs "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/replies"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

type ReplyController struct {
	rtRepo repos.RouteRepository
	rpRepo repos.ReplyRepository
	l      logger.Logger
}

func NewReplyController(
	rpRepo repos.ReplyRepository,
	rtRepo repos.RouteRepository,
	l logger.Logger) *ReplyController {
	return &ReplyController{
		rpRepo: rpRepo,
		rtRepo: rtRepo,
		l:      l,
	}
}

func (c *ReplyController) Post(params replies.ReplyCreateParams) middleware.Responder {
	var err error

	if *params.Body.Replace == "" || *params.Body.Regex == "" {
		payload := new(replies.ReplyCreateBadRequestBody)
		payload.Errors = &replies.ReplyCreateBadRequestBodyAO2Errors{
			Validation: &replies.ReplyCreateBadRequestBodyAO2ErrorsValidation{
				Regex:   required,
				Replace: required,
			},
		}
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return replies.NewReplyCreateBadRequest().WithPayload(payload)
	}

	var rp *domain.Reply
	rp, err = c.rpRepo.Set(domain.Reply{
		Description: &params.Body.Description,
		Regex:       *params.Body.Regex,
		Replace:     *params.Body.Replace,
	})
	if err == domain.ErrReplyAlreadyExists {
		payload := new(replies.ReplyCreateBadRequestBody)
		payload.Errors = &replies.ReplyCreateBadRequestBodyAO2Errors{
			Validation: &replies.ReplyCreateBadRequestBodyAO2ErrorsValidation{
				Regex:   unique,
				Replace: unique,
			},
		}
		payload.Data = []interface{}{}
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return replies.NewReplyCreateBadRequest().WithPayload(payload)
	}
	if err != nil {
		c.l.API().Reply().Debug(err)
		payload := new(replies.ReplyCreateInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return replies.NewReplyCreateInternalServerError().WithPayload(payload)
	}

	var payload = new(replies.ReplyCreateOKBody)
	item := replyToResponse(rp)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Reply().Info(rp.ReplyID.String(), "", logger.APICREATED)

	return replies.NewReplyCreateOK().WithPayload(payload)
}

func (c *ReplyController) Get(params replies.ReplyViewParams) middleware.Responder {
	var err error

	var rp *domain.Reply
	rp, err = c.rpRepo.Get(params.ReplyID)
	if err == domain.ErrReplyNotFound {
		return replyViewNotFoundResp()
	}
	if err != nil {
		c.l.API().Reply().Debug(err)
		return replyViewInternalServerErrorResp(err)
	}

	var payload = new(replies.ReplyViewOKBody)
	item := replyToResponse(rp)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	return replies.NewReplyViewOK().WithPayload(payload)
}

func (c *ReplyController) Put(params replies.ReplyPutParams) middleware.Responder {
	var err error

	if *params.Body.Replace == "" || *params.Body.Regex == "" {
		payload := new(replies.ReplyPutBadRequestBody)
		payload.Errors = &replies.ReplyPutBadRequestBodyAO2Errors{
			Validation: &replies.ReplyPutBadRequestBodyAO2ErrorsValidation{
				Regex:   required,
				Replace: required,
			},
		}
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return replies.NewReplyPutBadRequest().WithPayload(payload)
	}

	var rp *domain.Reply
	rp, err = c.rpRepo.Update(
		params.ReplyID,
		domain.Reply{
			Description: &params.Body.Description,
			Regex:       *params.Body.Regex,
			Replace:     *params.Body.Replace,
		},
	)
	if err == domain.ErrReplyNotFound {
		payload := new(replies.ReplyPutNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return replies.NewReplyPutNotFound().WithPayload(payload)
	}
	if err != nil {
		c.l.API().Reply().Debug(err)
		payload := new(replies.ReplyPutInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return replies.NewReplyPutInternalServerError().WithPayload(payload)
	}

	var payload = new(replies.ReplyPutOKBody)
	item := replyToResponse(rp)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Reply().Info(rp.ReplyID.String(), "", logger.APIEDITED)

	return replies.NewReplyPutOK().WithPayload(payload)
}

func (c *ReplyController) Delete(params replies.ReplyDeleteParams) middleware.Responder {
	var err error

	_, err = c.rpRepo.Get(params.ReplyID)
	if err == domain.ErrReplyNotFound {
		payload := new(replies.ReplyDeleteNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return replies.NewReplyDeleteNotFound().WithPayload(payload)
	}

	var routes []domain.Route
	routes, err = c.rtRepo.List(domain.RoutesFilter{ReplyID: &params.ReplyID})
	if err != nil {
		c.l.API().Reply().Debug("can't delete reply id from routes", err)
		payload := new(replies.ReplyDeleteInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return replies.NewReplyDeleteInternalServerError().WithPayload(payload)
	}

	for _, rt := range routes {
		if err = c.rtRepo.UnsetReplyID(rt.RouteID); err != nil {
			c.l.API().Reply().Debug("can't delete reply id from routes", err)
			payload := new(replies.ReplyDeleteInternalServerErrorBody)
			payload.Errors = err.Error()
			payload.Version = &Version
			payload.Message = *PayloadFailMessage
			return replies.NewReplyDeleteInternalServerError().WithPayload(payload)
		}
	}

	err = c.rpRepo.Delete(params.ReplyID)
	if err != nil {
		c.l.API().Reply().Debug(err)
		payload := new(replies.ReplyDeleteInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return replies.NewReplyDeleteInternalServerError().WithPayload(payload)
	}

	payload := new(replies.ReplyDeleteOKBody)
	payload.Version = &Version
	payload.Data = []interface{}{}
	payload.Message = PayloadSuccessMessage

	c.l.API().Reply().Info(params.ReplyID.String(), "", logger.APIDELETED)

	return replies.NewReplyDeleteOK().WithPayload(payload)
}

func (c *ReplyController) List(params replies.ReplyCollectionParams) middleware.Responder {
	var (
		err    error
		limit  int64
		offset int64
	)

	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Offset != nil {
		offset = *params.Offset
	}

	var reps []domain.Reply
	reps, err = c.rpRepo.List(domain.RepliesFilter{
		Limit:  limit,
		Offset: offset,
	})
	if err == domain.ErrRouteNotFound {
		payload := new(replies.ReplyCollectionNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return replies.NewReplyCollectionNotFound().WithPayload(payload)
	}

	if err != nil {
		c.l.API().Reply().Debug(err)
		payload := new(replies.ReplyCollectionInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return replies.NewReplyCollectionInternalServerError().WithPayload(payload)
	}
	var payload = new(replies.ReplyCollectionOKBody)

	for _, r := range reps {
		item := replyToResponse(&r)
		payload.Data = append(payload.Data, item)
	}

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage

	return replies.NewReplyCollectionOK().WithPayload(payload)

}

func replyToResponse(p *domain.Reply) *replies.DataItems0 {
	created := strfmt.DateTime(p.Created)

	item := new(replies.DataItems0)
	item.ID = p.ReplyID.String()
	item.Description = *p.Description
	item.Regex = &p.Regex
	item.Replace = &p.Replace
	item.Created = created
	return item
}

func replyViewInternalServerErrorResp(err error) *replies.ReplyViewInternalServerError {
	payload := new(replies.ReplyViewInternalServerErrorBody)
	payload.Errors = err.Error()
	payload.Version = &Version
	payload.Message = *PayloadFailMessage
	return replies.NewReplyViewInternalServerError().WithPayload(payload)
}

func replyViewNotFoundResp() *replies.ReplyViewNotFound {
	payload := new(replies.ReplyViewNotFoundBody)
	payload.Version = &Version
	payload.Errors = models.Error400DataAO1Errors{}
	payload.Message = *NotFoundMessage
	return replies.NewReplyViewNotFound().WithPayload(payload)
}
