package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"golang.org/x/crypto/bcrypt"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/clients"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

type ClientController struct {
	cliRepo repos.ClientRepository
	l       logger.Logger
}

func NewClientController(
	cRepo repos.ClientRepository,
	l logger.Logger) *ClientController {
	return &ClientController{
		cliRepo: cRepo,
		l:       l,
	}
}

func (c *ClientController) Post(params clients.ClientCreateParams) middleware.Responder {
	var err error

	errs := &clients.ClientCreateBadRequestBodyAO2Errors{
		Validation: &clients.ClientCreateBadRequestBodyAO2ErrorsValidation{},
	}
	if params.Body.Username == "" {
		errs.Validation.Username = required
	}
	if params.Body.Password == nil {
		errs.Validation.Password = required
	} else if *params.Body.Password == "" {
		errs.Validation.Password = min
	}
	if errs.Validation.Password != "" || errs.Validation.Username != "" {
		payload := new(clients.ClientCreateBadRequestBody)
		payload.Errors = errs
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return clients.NewClientCreateBadRequest().WithPayload(payload)
	}

	var passhash []byte
	passhash, err = bcrypt.GenerateFromPassword([]byte(*params.Body.Password), 10)
	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientCreateInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientCreateInternalServerError().WithPayload(payload)
	}

	var cli *domain.Client
	cli, err = c.cliRepo.Set(domain.Client{
		ClientID: params.Body.ID.String(),
		TTL:      &params.Body.TTL,
		Username: params.Body.Username,
		Passhash: string(passhash),
	})
	if err == domain.ErrClientAlreadyExists {
		payload := new(clients.ClientCreateBadRequestBody)
		payload.Errors = &clients.ClientCreateBadRequestBodyAO2Errors{
			Validation: &clients.ClientCreateBadRequestBodyAO2ErrorsValidation{
				ID: unique,
			},
		}
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return clients.NewClientCreateBadRequest().WithPayload(payload)
	}
	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientCreateInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientCreateInternalServerError().WithPayload(payload)
	}

	var payload = new(clients.ClientCreateOKBody)
	item := clientToResponse(cli)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Client().Info(cli.ClientID, "", logger.APICREATED)
	return clients.NewClientCreateOK().WithPayload(payload)
}

func (c *ClientController) Get(params clients.ClientViewParams) middleware.Responder {
	var err error

	var pat *domain.Client
	pat, err = c.cliRepo.Get(params.ClientID)
	if err == domain.ErrClientNotFound {
		return clientViewNotFoundResp()
	}
	if err != nil {
		c.l.API().Client().Debug(err)
		return clientViewInternalServerErrorResp(err)
	}

	var payload = new(clients.ClientViewOKBody)
	item := clientToResponse(pat)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	return clients.NewClientViewOK().WithPayload(payload)
}

func (c *ClientController) Patch(params clients.ClientPatchParams) middleware.Responder {
	var err error

	errs := &clients.ClientCreateBadRequestBodyAO2Errors{
		Validation: &clients.ClientCreateBadRequestBodyAO2ErrorsValidation{},
	}
	if params.Body.Username == "" {
		errs.Validation.Username = required
	}
	if params.Body.Password == "" {
		errs.Validation.Password = min
	}
	if errs.Validation.Password != "" || errs.Validation.Username != "" {
		payload := new(clients.ClientCreateBadRequestBody)
		payload.Errors = errs
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		return clients.NewClientCreateBadRequest().WithPayload(payload)
	}

	var passhash []byte
	passhash, err = bcrypt.GenerateFromPassword([]byte(params.Body.Password), 10)
	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientPatchInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientPatchInternalServerError().WithPayload(payload)
	}

	var cli *domain.Client
	cli, err = c.cliRepo.Update(
		params.ClientID,
		domain.Client{
			TTL:      &params.Body.TTL,
			Username: params.Body.Username,
			Passhash: string(passhash),
		},
	)
	if err == domain.ErrClientNotFound {
		payload := new(clients.ClientPatchNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return clients.NewClientPatchNotFound().WithPayload(payload)
	}
	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientPatchInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientPatchInternalServerError().WithPayload(payload)
	}

	var payload = new(clients.ClientPatchOKBody)
	item := clientToResponse(cli)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	c.l.API().Client().Info(cli.ClientID, "", logger.APIEDITED)

	return clients.NewClientPatchOK().WithPayload(payload)
}

func (c *ClientController) Delete(params clients.ClientDeleteParams) middleware.Responder {
	var err error

	_, err = c.cliRepo.Get(params.ClientID)
	if err == domain.ErrClientNotFound {
		payload := new(clients.ClientDeleteNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return clients.NewClientDeleteNotFound().WithPayload(payload)
	}

	err = c.cliRepo.Delete(params.ClientID)
	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientDeleteInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientDeleteInternalServerError().WithPayload(payload)
	}

	payload := new(clients.ClientDeleteOKBody)
	payload.Version = &Version
	payload.Data = []interface{}{}
	payload.Message = PayloadSuccessMessage

	c.l.API().Client().Info(params.ClientID.String(), "", logger.APIDELETED)

	return clients.NewClientDeleteOK().WithPayload(payload)
}

func (c *ClientController) List(params clients.ClientCollectionParams) middleware.Responder {
	var (
		username string
		err      error
		limit    int64
		offset   int64
	)

	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Offset != nil {
		offset = *params.Offset
	}

	if params.Username != nil {
		username = *params.Username
	}

	var clies []domain.Client
	clies, err = c.cliRepo.List(domain.ClientsFilter{
		Limit:    limit,
		Offset:   offset,
		Username: username,
	})
	if err == domain.ErrClientNotFound {
		payload := new(clients.ClientCollectionNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return clients.NewClientCollectionNotFound().WithPayload(payload)
	}

	if err != nil {
		c.l.API().Client().Debug(err)
		payload := new(clients.ClientCollectionInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return clients.NewClientCollectionInternalServerError().WithPayload(payload)
	}
	var payload = new(clients.ClientCollectionOKBody)

	for _, c := range clies {
		item := clientToResponse(&c)
		payload.Data = append(payload.Data, item)
	}

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage

	return clients.NewClientCollectionOK().WithPayload(payload)

}

func clientToResponse(p *domain.Client) *clients.DataItems0 {
	var expired *strfmt.DateTime
	if p.Expired != nil {
		tm := strfmt.DateTime(*p.Expired)
		expired = &tm
	}

	created := strfmt.DateTime(p.Created)

	item := new(clients.DataItems0)
	item.ID = strfmt.UUID(p.ClientID)
	item.Created = &created
	item.Username = p.Username
	item.TTL = p.TTL
	item.Expired = expired
	return item
}

func clientViewInternalServerErrorResp(err error) *clients.ClientViewInternalServerError {
	payload := new(clients.ClientViewInternalServerErrorBody)
	payload.Errors = err.Error()
	payload.Version = &Version
	payload.Message = *PayloadFailMessage
	return clients.NewClientViewInternalServerError().WithPayload(payload)
}

func clientViewNotFoundResp() *clients.ClientViewNotFound {
	payload := new(clients.ClientViewNotFoundBody)
	payload.Version = &Version
	payload.Errors = models.Error400DataAO1Errors{}
	payload.Message = *NotFoundMessage
	return clients.NewClientViewNotFound().WithPayload(payload)
}
