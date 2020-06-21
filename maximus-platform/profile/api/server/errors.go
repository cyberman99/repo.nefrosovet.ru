package server

import (
	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/apierrors"
)

func (s *Server) error(ctx echo.Context, httpCode int, err error) error {
	var errors []interface{}

	if err != nil {
		errors = append(errors, err.Error())
	}

	switch httpCode {
	case 400: // BadRequest
		return ctx.JSON(httpCode, api.Error400{
			Error: api.Error{
				Base: api.Base{
					Version: s.Version,
				},
			},
			Errors:  errors,
			Message: apierrors.ValidationErrorMessage,
		})
	case 401: // Unauthorized
		return ctx.JSON(httpCode, api.Error401{
			Error: api.Error{
				Base: api.Base{
					Version: s.Version,
				},
			},
			Errors:  errors,
			Message: apierrors.AccessDeniedMessage,
		})
	case 404: // NotFound
		return ctx.JSON(httpCode, api.Error404{
			Error: api.Error{
				Base: api.Base{
					Version: s.Version,
				},
			},
			Errors:  errors,
			Message: apierrors.NotFoundMessage,
		})
	case 405: // MethodNotAllowed
		return ctx.JSON(httpCode, api.Error405{
			Error: api.Error{
				Base: api.Base{
					Version: s.Version,
				},
			},
			Errors:  &errors,
			Message: &apierrors.MethodNotAllowedMessage,
		})
	}

	return ctx.JSON(httpCode, api.Error500{
		Error: api.Error{
			Base: api.Base{
				Version: s.Version,
			},
		},
		Errors:  errors,
		Message: apierrors.InternalServerErrorMessage,
	})
}
