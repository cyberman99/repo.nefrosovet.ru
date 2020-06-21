package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/permissions"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

type ClientPermissionsController struct {
	cliRepo repos.ClientRepository
	l       logger.Logger
}

func NewClientPermissionsController(
	cli repos.ClientRepository,
	l logger.Logger,
) *ClientPermissionsController {
	return &ClientPermissionsController{
		cliRepo: cli,
		l:       l,
	}
}

func (cp *ClientPermissionsController) Post(params permissions.ClientPermissionCreateParams) middleware.Responder {
	var err error

	_, err = cp.cliRepo.Get(params.ClientID)
	if err == domain.ErrClientNotFound {
		payload := new(permissions.ClientPermissionCreateNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return permissions.NewClientPermissionCreateNotFound().WithPayload(payload)
	}

	var perm = new(domain.Permissions)
	perm.Subscribe = make([]domain.Acl, len(params.Body.Subscribe))
	for i, s := range params.Body.Subscribe {
		perm.Subscribe[i] = domain.Acl{s}
	}

	perm.Publish = make([]domain.Acl, len(params.Body.Publish))
	for i, p := range params.Body.Publish {
		perm.Publish[i] = domain.Acl{p}
	}
	perm, err = cp.cliRepo.SetOrReplace(params.ClientID, *perm)
	if err != nil {
		cp.l.API().Client().Debug(err)
		payload := new(permissions.ClientPermissionCreateInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return permissions.NewClientPermissionCreateInternalServerError().WithPayload(payload)
	}

	var payload = new(permissions.ClientPermissionCreateOKBody)
	item := permissionToResponse(perm)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage
	payload.Data = append(payload.Data, item)

	cp.l.API().Client().Info(params.ClientID.String(), "permissions", logger.APIEDITED)

	return permissions.NewClientPermissionCreateOK().WithPayload(payload)
}

func (cp *ClientPermissionsController) Get(params permissions.ClientPermissionViewParams) middleware.Responder {
	var err error

	_, err = cp.cliRepo.Get(params.ClientID)
	if err == domain.ErrClientNotFound {
		payload := new(permissions.ClientPermissionViewNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return permissions.NewClientPermissionViewNotFound().WithPayload(payload)
	}

	var perms *domain.Permissions
	perms, err = cp.cliRepo.GetPermissions(params.ClientID)
	if err == domain.ErrPermissionNotFound {
		payload := new(permissions.ClientPermissionViewNotFoundBody)
		payload.Version = &Version
		payload.Errors = models.Error400DataAO1Errors{}
		payload.Message = *NotFoundMessage
		return permissions.NewClientPermissionViewNotFound().WithPayload(payload)
	}

	if err != nil {
		cp.l.API().Client().Debug(err)
		payload := new(permissions.ClientPermissionViewInternalServerErrorBody)
		payload.Errors = err.Error()
		payload.Version = &Version
		payload.Message = *PayloadFailMessage
		return permissions.NewClientPermissionViewInternalServerError().WithPayload(payload)

	}

	var payload = new(permissions.ClientPermissionViewOKBody)
	item := permissionToResponse(perms)
	payload.Data = append(payload.Data, item)

	payload.Version = &Version
	payload.Message = PayloadSuccessMessage

	return permissions.NewClientPermissionViewOK().WithPayload(payload)
}

func permissionToResponse(d *domain.Permissions) *permissions.DataItems0 {
	item := new(permissions.DataItems0)
	item.Subscribe = make([]string, len(d.Subscribe))
	for i, s := range d.Subscribe {
		item.Subscribe[i] = s.Pattern
	}

	item.Publish = make([]string, len(d.Publish))
	for i, p := range d.Publish {
		item.Publish[i] = p.Pattern
	}
	return item
}
