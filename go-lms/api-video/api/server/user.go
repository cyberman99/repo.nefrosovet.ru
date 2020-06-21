package server

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/go-lms/api-video/api"
	"repo.nefrosovet.ru/go-lms/api-video/ent"
	"repo.nefrosovet.ru/go-lms/api-video/ent/subscriber"
	"repo.nefrosovet.ru/go-lms/api-video/ent/user"
	"repo.nefrosovet.ru/go-lms/api-video/helpers"
	"repo.nefrosovet.ru/go-lms/api-video/libs"
)

const (
	passwordLenStore = 40
	passwordLenShow  = 20
	timeoutDuration  = 2 * time.Second
)

var (
	re       = regexp.MustCompile(`[<]*(?:(?:sip\:)?)((\w+)@([^>]*))`)
	metaData = []byte("[]")
)

// Получение списка юзеров// (GET /api/v1/user)
func (s *Server) UserIndex(c echo.Context) error {
	ctx, _ := context.WithTimeout(c.Request().Context(), timeoutDuration)

	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(users []*ent.User) error {
		response := &api.UserIndexResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   ptrB(false),
				Message: nil,
			},
		}

		var data []api.UsersData
		for _, u := range users {
			kamailioUser, err := u.QuerySubscriber().Only(ctx)
			if err != nil {
				return responseInternalServerError(err)
			}
			userData := userToUsersData(u.ID, kamailioUser.Username, kamailioUser.Domain)
			data = append(data, userData)
		}

		response.Data = &data

		return c.JSON(http.StatusOK, response)
	}

	users, err := s.ent.User.Query().All(ctx)
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(users)
}

// Создание юзера// (POST /api/v1/user)
func (s *Server) UserStore(c echo.Context, params api.UserStoreParams) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseBadRequest := func(validation map[string]string) error {
		response := &api.UserStoreResponse400{
			ApiError: api.ApiError{
				Code:  ptrI(http.StatusBadRequest),
				Data:  &[]interface{}{nil},
				Error: ptrB(true),
			},
		}

		var message struct {
			api.UserValidationData
		}
		if required, ok := validation["required"]; ok {
			message.Required = &required
		}
		if sipValid, ok := validation["sipValid"]; ok {
			message.SipValid = &sipValid
		}

		response.Message = &message
		return c.JSON(http.StatusOK, response)
	}

	responseSuccess := func(userID int) error {
		var data struct {
			api.UserCreateData
		}
		data.UserId = userID
		return c.JSON(http.StatusOK, &api.UserStoreResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   nil,
				Message: ptrS("OK"),
			},
			ApiUserStoreObject: api.ApiUserStoreObject{Data: &data},
		})
	}

	validation, err := libs.ValidateSip(params.Sip)
	if err != nil {
		return responseInternalServerError(err)
	}

	if len(*validation) != 0 {
		return responseBadRequest(*validation)
	}

	password := helpers.RandomString(passwordLenStore, helpers.Alphanumeric)

	ctx, _ := context.WithTimeout(c.Request().Context(), timeoutDuration)

	match := re.FindStringSubmatch(string(*params.Sip))
	username := match[2]
	domain := match[3]

	kamailioUser, err := s.ent.Subscriber.Query().
		Where(
			subscriber.Username(username),
			subscriber.Domain(domain),
		).
		First(ctx)

	var u *ent.User
	if ent.IsNotFound(err) {
		err := s.ent.WithTx(ctx, func(tx *ent.Tx) (err error) {
			u, err = tx.User.
				Create().
				SetMetaData(string(metaData)).
				Save(ctx)
			if err != nil {
				return err
			}

			kamailioUser, err = tx.Subscriber.
				Create().
				SetUsername(username).
				SetDomain(domain).
				SetHa1(helpers.GetMD5Hash(helpers.GetHa1(username, domain, password))).
				SetHa1b(helpers.GetMD5Hash(helpers.GetHa1b(username, domain, password))).
				SetUser(u).
				Save(ctx)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return responseInternalServerError(err)
		}
	} else {
		u, err = kamailioUser.QueryUser().Only(ctx)
		if err != nil {
			return responseInternalServerError(err)
		}
	}

	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(u.ID)
}

// Данные юзера// (GET /api/v1/user/{userID})
func (s *Server) UserShow(c echo.Context, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(username, password string) error {
		response := &api.UserShowResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   ptrB(false),
				Message: nil,
			},
			ApiUserShowObject: api.ApiUserShowObject{Data: &api.UserData{
				Password: password,
				Username: username,
			}},
		}

		return c.JSON(http.StatusOK, response)
	}

	password := helpers.RandomString(passwordLenShow, helpers.Alphanumeric)

	ctx, _ := context.WithTimeout(c.Request().Context(), timeoutDuration)

	sub, err := s.ent.User.
		Query().
		Where(user.ID(int(userID))).
		QuerySubscriber().
		First(ctx)
	if err != nil {
		return responseInternalServerError(err)
	}

	_, err = sub.
		Update().
		SetHa1(helpers.GetMD5Hash(helpers.GetHa1(sub.Username, sub.Domain, password))).
		SetHa1b(helpers.GetMD5Hash(helpers.GetHa1b(sub.Username, sub.Domain, password))).
		Save(ctx)
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(sub.Username, password)
}

func userToUsersData(userID int, username, domain string) api.UsersData {
	return api.UsersData{
		Sip:    helpers.GetSip(username, domain),
		UserId: userID,
	}
}
