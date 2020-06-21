package server

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/apierrors"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

// GET /users/{userID}/contacts
func (s *Server) GetUserContacts(ctx echo.Context, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		panic(err)
	}

	responseNotFound := func(err error) error {
		return s.error(ctx, http.StatusNotFound, err)
	}

	responseSuccess := func(contacts []storage.UserContact) error {
		var userContactsWithIds []api.UserContactWithId
		for _, contact := range contacts {
			userContactsWithIds = append(userContactsWithIds, api.UserContactWithId{
				//ID:          storage.PtrS(string(contact.ID)),
				UserContact: contactToContactParams(contact),
			})
		}

		return ctx.JSON(http.StatusOK, api.UserContacts200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Errors:  nil,
				Message: SuccessMessage,
			},
			Data: &userContactsWithIds,
		})
	}

	contacts, err := s.storage.GetUserContacts(storage.GetUserContacts{
		UserID: storage.UUID(userID),
	})
	if err != nil {
		switch err {
		case storage.ErrUserContactNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(contacts)
}

// POST /users/{userID}/contacts
func (s *Server) PostUserContacts(ctx echo.Context, userID api.UserID) error {
	var params api.PostUserContactsJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERCREATECONTACT, string(userID), logger.AUTHFAIL)

		panic(err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
		s.authLog.Info(logger.AUTHUSERCREATECONTACT, string(userID), logger.AUTHFAIL)

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

	responseSuccess := func(contact storage.UserContact) error {
		s.authLog.Info(logger.AUTHUSERCREATECONTACT, string(userID), logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.UserContacts200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Errors:  nil,
				Message: SuccessMessage,
			},
			Data: &[]api.UserContactWithId{
				{
					//ID:          storage.PtrS(string(contact.ID)),
					UserContact: contactToContactParams(contact),
				},
			},
		})
	}

	validation := map[string]string{}

	if params.TypeCODE == nil {
		validation["typeCODE"] = "required"
	}

	if params.Value == nil {
		validation["value"] = "required"
	}

	if params.Verified == nil {
		validation["verified"] = "required"
	}

	if len(validation) > 0 {
		return responseBadRequest(nil, validation)
	}

	c, err := s.storage.StoreUserContact(storage.StoreUserContact{
		ID:       uuid.New().String(),
		UserID:   storage.UUID(userID),
		TypeCODE: *params.TypeCODE,
		Value:    *params.Value,
	})
	if err != nil {
		switch err {
		case storage.ErrUserContactAlreadyExists:
			return responseBadRequest(err, map[string]string{
				"value": "unique",
			})
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(c)
}

// PATCH /users/{userID}/contacts/{contactID}
func (s *Server) PatchUserContacts(ctx echo.Context, userID api.UserID, contactID api.ContactID) error {
	var params api.PatchUserContactsJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	responseInternalServerError := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATECONTACT, string(userID), logger.AUTHFAIL)

		panic(err)
	}

	responseNotFound := func(err error) error {
		s.authLog.Info(logger.AUTHUSERUPDATECONTACT, string(userID), logger.AUTHFAIL)

		return s.error(ctx, http.StatusNotFound, err)
	}

	responseBadRequest := func(err error, validation map[string]string) error {
		s.authLog.Info(logger.AUTHUSERUPDATECONTACT, string(userID), logger.AUTHFAIL)

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

	responseSuccess := func(contact storage.UserContact) error {
		s.authLog.Info(logger.AUTHUSERUPDATECONTACT, string(userID), logger.AUTHSUCCESS)

		return ctx.JSON(http.StatusOK, api.UserContacts200{
			Success: api.Success{
				Base: api.Base{
					Version: s.Version,
				},
				Errors:  nil,
				Message: SuccessMessage,
			},
			Data: &[]api.UserContactWithId{
				{
					//ID:          storage.PtrS(string(contact.ID)),
					UserContact: contactToContactParams(contact),
				},
			},
		})
	}

	validation := map[string]string{}

	if !strfmt.IsUUID(string(userID)) {
		validation["userID"] = "format"
	}

	if !strfmt.IsUUID(string(contactID)) {
		validation["_id"] = "format"
	}

	if len(validation) > 0 {
		return responseBadRequest(nil, validation)
	}

	in := storage.UpdateUserContact{
		TypeCODE: params.TypeCODE,
		Value:    params.Value,
		Verified: params.Verified,
	}

	c, err := s.storage.UpdateUserContact(storage.UUID(userID), in)
	if err != nil {
		switch err {
		case storage.ErrUserContactAlreadyExists:
			return responseBadRequest(err, nil)
		case storage.ErrUserContactNotFound:
			return responseNotFound(err)
		default:
			return responseInternalServerError(err)
		}
	}

	return responseSuccess(c)
}
