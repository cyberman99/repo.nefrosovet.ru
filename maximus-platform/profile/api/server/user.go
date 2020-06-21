package server

import (
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/apierrors"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

// POST /users
func (s *Server) PostUsers(ctx echo.Context) error {
	var params api.PostUsersJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	contactValue := (*params.Contacts)[0].Value

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERCREATE, *contactValue, logger.AUTHFAIL)

		panic(err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
		s.authLog.Info(logger.AUTHUSERCREATE, *contactValue, logger.AUTHFAIL)

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

	responseSuccess := func(user storage.User) error {
		s.authLog.Info(logger.AUTHUSERCREATE, *contactValue, logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.User200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]api.UserParamsWithId{{
				ID:         storage.PtrS(string(user.ID)),
				UserParams: userToUserParams(user),
			}},
		})
	}

	validation := map[string]string{}

	if params.FirstName == nil {
		validation["firstName"] = "required"
	}

	if params.LastName == nil {
		validation["lastName"] = "required"
	}

	if params.MiddleName == nil {
		validation["middleName"] = "required"
	}

	if params.Password == nil {
		validation["password"] = "required"
	}

	if (*params.Contacts)[0].TypeCODE == nil {
		validation["typeCODE"] = "required"
	}

	if (*params.Contacts)[0].Value == nil {
		validation["value"] = "required"
	}

	if len(validation) > 0 {
		return responseBadRequest(nil, validation)
	}

	passwordHash, err := HashPassword(*params.Password)
	if err != nil {
		return responseInternalServerError(err)
	}

	var u storage.User

	var contacts []storage.UserContact
	for _, c := range *params.Contacts {
		contacts = append(contacts, storage.UserContact{
			TypeCODE: *c.TypeCODE,
			Value:    *c.Value,
		})
	}

	u, err = s.storage.StoreUser(storage.StoreUser{
		ID:           uuid.New().String(),
		PasswordHash: passwordHash,
		FirstName:    *params.FirstName,
		LastName:     *params.LastName,
		MiddleName:   *params.MiddleName,
		Contacts:     contacts,
	})
	if err != nil {
		switch err {
		case storage.ErrUserContactAlreadyExists:
			return responseBadRequest(err, map[string]string{
				"value": "unique",
			})
		}
		return responseInternalServerError(err)
	}

	_, err = s.storage.StoreUserSettings(storage.StoreUserSettings{
		UserID:           u.ID,
		TwoFAChannelType: "",
		Locale:           "RUS",
	})
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(u)
}

// GET /users
func (s *Server) GetUsers(ctx echo.Context) error {
	responseInternalServerError := func(err error) error {
		panic(err)
	}

	responseSuccess := func(users []storage.User) error {
		var userParamsWithIDs []api.UserParamsWithId
		for _, user := range users {
			userParamsWithIDs = append(userParamsWithIDs, api.UserParamsWithId{
				ID:         storage.PtrS(string(user.ID)),
				UserParams: userToUserParams(user),
			})
		}

		return ctx.JSON(http.StatusOK, api.User200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Errors:  nil,
				Message: SuccessMessage,
			},
			Data: &userParamsWithIDs,
		})
	}

	users, err := s.storage.GetUsers()
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(users)
}

// GET /users/{userID}
func (s *Server) GetUser(ctx echo.Context, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		panic(err)
	}

	responseNotFound := func(err error) error {
		return s.error(ctx, http.StatusNotFound, err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
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

	responseSuccess := func(user storage.User) error {
		return ctx.JSON(http.StatusOK, api.User200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]api.UserParamsWithId{{
				ID:         storage.PtrS(string(user.ID)),
				UserParams: userToUserParams(user),
			}},
		})
	}

	if !strfmt.IsUUID(string(userID)) {
		return responseBadRequest(nil, map[string]string{
			"userID": "format",
		})
	}

	u, err := s.storage.GetUser(storage.GetUser{
		ID: storage.PtrUID(storage.UUID(userID)),
	})
	if err != nil {
		switch err {
		case storage.ErrUserNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(u)
}

// PATCH /users/{userID}
func (s *Server) PatchUser(ctx echo.Context, userID api.UserID) error {
	var params api.PatchUserJSONRequestBody
	var err error
	if err = ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATE, string(userID), logger.AUTHFAIL)

		panic(err)
	}

	responseNotFound := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATE, string(userID), logger.AUTHFAIL)

		return s.error(ctx, http.StatusNotFound, err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
		s.authLog.Info(logger.AUTHUSERUPDATE, string(userID), logger.AUTHFAIL)

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

	responseSuccess := func(user storage.User) error {
		s.authLog.Info(logger.AUTHUSERUPDATE, string(user.ID), logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.User200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Message: SuccessMessage,
			},
			Data: &[]api.UserParamsWithId{{
				ID:         storage.PtrS(string(user.ID)),
				UserParams: userToUserParams(user),
			}},
		})
	}

	if !strfmt.IsUUID(string(userID)) {
		return responseBadRequest(nil, map[string]string{
			"userID": "format",
		})
	}

	var contacts []storage.UserContact
	for _, c := range *params.Contacts {
		var verified time.Time
		if c.Verified != nil && *c.Verified {
			verified, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			if err != nil {
				return responseInternalServerError(err)
			}
		}

		contacts = append(contacts, storage.UserContact{
			TypeCODE: *c.TypeCODE,
			Value:    *c.Value,
			Verified: &verified,
		})
	}

	in := storage.UpdateUser{
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		MiddleName: params.MiddleName,
		Contacts:   &contacts,
	}

	if params.Password != nil {
		passwordHash, err := HashPassword(*params.Password)
		if err != nil {
			return responseInternalServerError(err)
		}

		in.PasswordHash = &passwordHash
	}

	u, err := s.storage.UpdateUser(storage.UUID(userID), in)
	if err != nil {
		switch err {
		case storage.ErrUserAlreadyExists:
			return responseBadRequest(err, nil)
		case storage.ErrUserNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(u)
}

func userToUserParams(user storage.User) api.UserParams {
	var contacts []api.UserContact
	for _, c := range user.Contacts {
		contacts = append(contacts, contactToContactParams(c))
	}

	params := api.UserParams{
		Contacts:   &contacts,
		FirstName:  &user.FirstName,
		LastName:   &user.LastName,
		MiddleName: &user.MiddleName,
	}

	return params
}

func contactToContactParams(c storage.UserContact) api.UserContact {
	isVerified := c.Verified != nil
	contact := api.UserContact{
		TypeCODE: &c.TypeCODE,
		Value:    &c.Value,
		Verified: &isVerified,
	}

	return contact
}
