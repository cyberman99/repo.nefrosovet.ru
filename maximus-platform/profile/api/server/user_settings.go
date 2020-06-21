package server

import (
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/apierrors"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
)

// GET /users/{userID}/settings
func (s *Server) GetUserSettings(ctx echo.Context, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		panic(err)
	}

	responseNotFound := func(err error) error {
		return s.error(ctx, http.StatusNotFound, err)
	}

	responseSuccess := func(settings storage.UserSettings) error {
		return ctx.JSON(http.StatusOK, api.UserSettings200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]api.SettingsParams{
				settingsToSettingsParams(settings),
			},
		})
	}

	settings, err := s.storage.GetUserSettings(storage.GetUserSettings{
		UserID: storage.UUID(userID),
	})
	if err != nil {
		switch err {
		case storage.ErrUserSettingsNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(settings)
}

// PATCH /users/{userID}/settings
func (s *Server) PatchUserSettings(ctx echo.Context, userID api.UserID) error {
	var params api.PatchUserSettingsJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATESETTINGS, string(userID), logger.AUTHFAIL)

		panic(err)
	}

	responseNotFound := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATESETTINGS, string(userID), logger.AUTHFAIL)

		return s.error(ctx, http.StatusNotFound, err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
		s.authLog.Info(logger.AUTHUSERUPDATESETTINGS, string(userID), logger.AUTHFAIL)

		response := apierrors.ValidationErrorResponse{
			Response: apierrors.Response{
				Version: s.Version,
				Message: apierrors.ValidationErrorMessage,
			},
			Errors: &apierrors.ValidationError{
				Validation: validation,
			},
		}

		if err != nil {
			response.Errors.Core = err.Error()
		}

		return ctx.JSON(http.StatusBadRequest, response)
	}

	responseSuccess := func(settings storage.UserSettings) error {
		s.authLog.Info(logger.AUTHUSERUPDATESETTINGS, string(userID), logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.UserSettings200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]api.SettingsParams{
				settingsToSettingsParams(settings),
			},
		})
	}

	if !strfmt.IsUUID(string(userID)) {
		return responseBadRequest(nil, map[string]string{
			"userID": "format",
		})
	}

	in := storage.UpdateUserSettings{
		TwoFAChannelType: params.N2FAChannelType,
		Locale:           params.Locale,
	}

	settings, err := s.storage.UpdateUserSettings(storage.UUID(userID), in)
	if err != nil {
		switch err {
		case storage.ErrUserSettingsNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(settings)
}

func settingsToSettingsParams(s storage.UserSettings) api.SettingsParams {
	params := api.SettingsParams{
		N2FAChannelType: &s.TwoFAChannelType,
		Locale:          &s.Locale,
	}

	return params
}
