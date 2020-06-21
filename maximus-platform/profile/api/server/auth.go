package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

// POST /authorize
func (s *Server) Authorize(ctx echo.Context) error {
	var params api.AuthorizeJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERLOGIN, params.Login, logger.AUTHFAIL)
		panic(err)
	}

	responseUnauthorized := func(err error) error {
		s.authLog.Info(logger.AUTHUSERLOGIN, params.Login, logger.AUTHFAIL)
		return s.error(ctx, http.StatusUnauthorized, err)
	}

	responseSuccess := func(userID string) error {
		s.authLog.Info(logger.AUTHUSERLOGIN, params.Login, logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.Authorize200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]struct {
				UserID string `json:"userID"`
			}{{
				UserID: userID,
			}},
		})
	}

	user, err := s.storage.GetUser(storage.GetUser{Value: &params.Login})
	if err != nil {
		switch {
		//case errors.Is(err, storage.ErrUserNotFound):
		//	return responseUnauthorized(err)
		}

		return responseInternalServerError(err)
	}

	newHash, err := VerifyPassword(params.Password, user.PasswordHash)
	if err != nil {
		return responseUnauthorized(err)
	}

	if newHash != "" {
		user, err = s.storage.UpdateUser(user.ID, storage.UpdateUser{
			PasswordHash: &newHash,
		})
		if err != nil {
			return responseInternalServerError(err)
		}
	}



	return responseSuccess(string(user.ID))
}
