package handlers

import (
	//errs "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/routes"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

type RouteController struct {
	rtRepo repos.RouteRepository
	rpRepo repos.ReplyRepository
	l      logger.Logger
}

func NewRouteController(
	rtRepo repos.RouteRepository,
	rpRepo repos.ReplyRepository,
	l logger.Logger) *RouteController {
	return &RouteController{
		rtRepo: rtRepo,
		l:      l,
		rpRepo: rpRepo,
	}
}

func (c *RouteController) Post(params routes.RouteCreateParams) middleware.Responder {
	var (
		err     error
		isValid = true
	)

	errs := &routes.RouteCreateBadRequestBodyAO2Errors{}
	errs.Validation = new(routes.RouteCreateBadRequestBodyAO2ErrorsValidation)
	if len(params.Body.Dst) == 0 {
		errs.Validation.Dst = required
		isValid = false
	}
	if params.Body.Src.Topic == nil {
		errs.Validation.Src = required
		errs.Validation.SrcTopic = required
		isValid = false
	}
	if !isValid {
		payload := new(routes.RouteCreateBadRequestBody)
		payload.Errors = errs
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return routes.NewRouteCreateBadRequest().WithPayload(payload)
	}
	if params.Body.ReplyID != nil {
		if *params.Body.ReplyID != "" {
			_, err = c.rpRepo.Get(*params.Body.ReplyID)
			if err == domain.ErrReplyNotFound {
				return replyViewNotFoundResp()
			}
			if err != nil {
				c.l.API().Route().Debug(err)
				return replyViewInternalServerErrorResp(err)
			}
		}
	}

	var rt domain.Route
	rt.ReplyID = params.Body.ReplyID
	rt.Src = domain.Source{params.Body.Src.Payload, params.Body.Src.Topic}
	rt.Dst = make([]domain.Destinations, len(params.Body.Dst))
	for i, d := range params.Body.Dst {
		qos := *d.Qos
		topic := *d.Topic
		rt.Dst[i] = domain.Destinations{byte(qos), topic}
	}

	var res *domain.Route
	res, err = c.rtRepo.Set(rt)
	if err != nil {
		c.l.API().Route().Debug(err)
		payload := new(routes.RouteCreateInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return routes.NewRouteCreateInternalServerError().WithPayload(payload)
	}

	var payload = new(routes.RouteCreateOKBody)

	item := routeToResponse(res)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Route().Info(res.RouteID.String(), "", logger.APICREATED)

	return routes.NewRouteCreateOK().WithPayload(payload)
}

func (c *RouteController) Get(params routes.RouteViewParams) middleware.Responder {
	var err error

	var rt *domain.Route
	rt, err = c.rtRepo.Get(params.RouteID)
	if err == domain.ErrRouteNotFound {
		return routeViewNotFoundResp()
	}

	if err != nil {
		c.l.API().Route().Debug(err)
		return routeViewInternalServerErrorResp(err)
	}

	var payload = new(routes.RouteViewOKBody)
	item := routeToResponse(rt)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	return routes.NewRouteViewOK().WithPayload(payload)
}

func (c *RouteController) Put(params routes.RoutePutParams) middleware.Responder {
	var (
		err     error
		isValid = true
	)

	errs := &routes.RouteCreateBadRequestBodyAO2Errors{}
	errs.Validation = new(routes.RouteCreateBadRequestBodyAO2ErrorsValidation)
	if len(params.Body.Dst) == 0 {
		errs.Validation.Dst = required
		isValid = false
	}
	if params.Body.Src.Topic == nil {
		errs.Validation.Src = required
		errs.Validation.SrcTopic = required
		isValid = false
	}
	if !isValid {
		payload := new(routes.RouteCreateBadRequestBody)
		payload.Errors = errs
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return routes.NewRouteCreateBadRequest().WithPayload(payload)
	}

	if params.Body.ReplyID != nil {
		if *params.Body.ReplyID != "" {
			_, err = c.rpRepo.Get(*params.Body.ReplyID)
			if err == domain.ErrReplyNotFound {
				return replyViewNotFoundResp()
			}
			if err != nil {
				c.l.API().Route().Debug(err)
				return replyViewInternalServerErrorResp(err)
			}
		}
	}

	var rt domain.Route
	rt.ReplyID = params.Body.ReplyID
	rt.Src = domain.Source{params.Body.Src.Payload, params.Body.Src.Topic}
	rt.Dst = make([]domain.Destinations, len(params.Body.Dst))
	for i, d := range params.Body.Dst {
		qos := *d.Qos
		topic := *d.Topic
		rt.Dst[i] = domain.Destinations{byte(qos), topic}
	}

	var res *domain.Route
	res, err = c.rtRepo.Update(
		params.RouteID,
		rt,
	)
	if err == domain.ErrRouteNotFound {
		payload := new(routes.RoutePutNotFoundBody)
		payload.Version = &Version
		payload.Data = append(payload.Data, routes.DataItems0{})
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return routes.NewRoutePutNotFound().WithPayload(payload)
	}
	if err != nil {
		c.l.API().Route().Debug(err)
		payload := new(routes.RoutePutInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Data = append(payload.Data, routes.DataItems0{})
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return routes.NewRoutePutInternalServerError().WithPayload(payload)
	}

	var payload = new(routes.RoutePutOKBody)
	item := routeToResponse(res)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Route().Info(res.RouteID.String(), "", logger.APIEDITED)

	return routes.NewRoutePutOK().WithPayload(payload)
}

func (c *RouteController) Delete(params routes.RouteDeleteParams) middleware.Responder {
	var err error

	_, err = c.rtRepo.Get(params.RouteID)
	if err == domain.ErrRouteNotFound {
		payload := new(routes.RouteDeleteNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return routes.NewRouteDeleteNotFound().WithPayload(payload)
	}

	err = c.rtRepo.Delete(params.RouteID)
	if err != nil {
		c.l.API().Route().Debug(err)
		payload := new(routes.RouteDeleteInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return routes.NewRouteDeleteInternalServerError().WithPayload(payload)
	}

	payload := new(routes.RouteDeleteOKBody)
	payload.Version = &Version
	payload.Data = []interface{}{}
	payload.Message = PayloadSuccessMessage

	c.l.API().Route().Info(params.RouteID.String(), "", logger.APIDELETED)

	return routes.NewRouteDeleteOK().WithPayload(payload)
}

func (c *RouteController) List(params routes.RouteCollectionParams) middleware.Responder {
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

	var rts []domain.Route
	rts, err = c.rtRepo.List(domain.RoutesFilter{
		Limit:  limit,
		Offset: offset,
	})
	if err == domain.ErrRouteNotFound {
		payload := new(routes.RouteCollectionNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return routes.NewRouteCollectionNotFound().WithPayload(payload)
	}

	if err != nil {
		c.l.API().Route().Debug(err)
		payload := new(routes.RouteCollectionInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return routes.NewRouteCollectionInternalServerError().WithPayload(payload)
	}
	var payload = new(routes.RouteCollectionOKBody)

	for _, r := range rts {
		item := routeToResponse(&r)
		payload.Data = append(payload.Data, item)
	}

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage

	return routes.NewRouteCollectionOK().WithPayload(payload)

}

func routeToResponse(p *domain.Route) *routes.DataItems0 {
	created := strfmt.DateTime(p.Created)

	item := new(routes.DataItems0)
	item.ID = p.RouteID.String()
	item.ReplyID = p.ReplyID
	item.Dst = make([]*models.RouteObjectDstItems0, len(p.Dst))
	for i, d := range p.Dst {
		qos := int64(d.Qos)
		topic := d.Topic
		item.Dst[i] = &models.RouteObjectDstItems0{&qos, &topic}
	}

	item.Src = &models.RouteObjectSrc{p.Src.Payload, p.Src.Topic}
	item.Created = created
	return item
}

func routeViewInternalServerErrorResp(err error) *routes.RouteViewInternalServerError {
	payload := new(routes.RouteViewInternalServerErrorBody)
	payload.Errors = err.Error()
	payload.Version = &Version
	payload.Message = *PayloadFailMessage
	return routes.NewRouteViewInternalServerError().WithPayload(payload)
}

func routeViewNotFoundResp() *routes.RouteViewNotFound {
	payload := new(routes.RouteViewNotFoundBody)
	payload.Version = &Version
	payload.Errors = models.Error400DataAO1Errors{}
	payload.Message = *NotFoundMessage
	return routes.NewRouteViewNotFound().WithPayload(payload)
}
