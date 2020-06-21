package handlers

import (
	"database/sql"
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"repo.nefrosovet.ru/maximus-platform/patient/api/models"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/employee"
	"repo.nefrosovet.ru/maximus-platform/patient/db/sqlc"
	"repo.nefrosovet.ru/maximus-platform/patient/logger"
)

type EmployeeViewController struct {
	version string
	q       *sqlc.Queries
	l       logger.APIEmployee
}

func NewEmployeeView(version string, q *sqlc.Queries, l logger.APIEntrier) *EmployeeViewController {
	return &EmployeeViewController{
		version: version,
		q:       q,
		l:       l.Employee(),
	}
}

func (ew *EmployeeViewController) Get(params employee.EmployeeViewParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	responseInternalServerError := func(err error) middleware.Responder {
		ew.l.Debug(err)

		payload := new(employee.EmployeeViewInternalServerErrorBody)
		payload.Error500Data = models.Error500Data{
			ErrorData: models.ErrorData{
				BaseData: models.BaseData{Version: &ew.version},
			},
			Errors:  err.Error(),
			Message: &InternalServerError,
		}

		return employee.NewEmployeeViewInternalServerError().WithPayload(payload)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(employee.EmployeeViewNotFoundBody)
		payload.Error404Data = models.Error404Data{
			ErrorData: models.ErrorData{
				BaseData: models.BaseData{Version: &ew.version},
			},
			Message: &EntityNotFoundMessage,
		}

		if err != nil {
			payload.Errors = append(payload.Errors, err)
		}

		return employee.NewEmployeeViewNotFound().WithPayload(payload)
	}

	responseSuccess := func(e sqlc.Employee) middleware.Responder {
		payload := new(employee.EmployeeViewOKBody)
		payload.SuccessData = models.SuccessData{
			BaseData: models.BaseData{Version: &ew.version},
			Message:  &SuccessMessage,
		}
		payload.Data = []*employee.EmployeeDataItem{
			employeeToData(e),
		}

		return employee.NewEmployeeViewOK().WithPayload(payload)
	}

	e, err := ew.q.GetEmployee(ctx, uuid.MustParse(params.EmployeeID.String()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(e)
}

func employeeToData(e sqlc.Employee) *employee.EmployeeDataItem {
	id := strfmt.UUID(e.ID.String())
	photoID := strfmt.UUID(e.PhotoGuid.String())

	return &employee.EmployeeDataItem{
		EmployeeObject: models.EmployeeObject{
			ID:         &id,
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Patronymic: &e.Patronymic.String,
			PhotoID:    &photoID,
		},
	}
}
