package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/go-lms/api-video/api"
	"repo.nefrosovet.ru/go-lms/api-video/ent"
	"repo.nefrosovet.ru/go-lms/api-video/ent/webinaruser"
)

const (
	webinarUserDestroyMessage = `Subscribers delete from webinar`
)

// Получение списка юзеров вебинара
// (GET /api/v1/webinar/{webinarID}/user)
func (s *Server) WebinarUserIndex(c echo.Context, webinarID api.WebinarID) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(users []*ent.WebinarUser) error {
		response := &api.WebinarUserIndexResponse200{
			ApiMessage: api.ApiMessage{
				Code:    nil,
				Error:   nil,
				Message: nil,
			},
			ApiWebinarUserIndexObject: api.ApiWebinarUserIndexObject{},
		}

		var data []api.WebinarUserJsonData
		for _, user := range users {
			webinarUserJsonData, err := webinarUserToWebinarUserJsonData(user)
			if err != nil {
				return responseInternalServerError(err)
			}

			data = append(data, webinarUserJsonData)
		}

		response.ApiWebinarUserIndexObject = api.ApiWebinarUserIndexObject{
			Data: &data,
		}

		return c.JSON(http.StatusOK, response)
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), time.Second*2)
	users, err := s.ent.WebinarUser.Query().Where(
		webinaruser.WebinarIDEQ(int(webinarID)),
	).All(ctx)
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(users)
}

// Создание юзера вебинара
// (POST /api/v1/webinar/{webinarID}/user)
func (s *Server) WebinarUserStore(c echo.Context, webinarID api.WebinarID, params api.WebinarUserStoreParams) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(user *ent.WebinarUser) error {
		return c.JSON(http.StatusOK, &api.RootResponse200{
			ApiMessage: api.ApiMessage{
				Code:    nil,
				Error:   nil,
				Message: nil,
			},
			ApiDataStringObject: api.ApiDataStringObject{
				Data: ptrS("OK!"),
			},
		})
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), time.Second*2)
	user, err := s.ent.WebinarUser.Create().
		SetWebinarID(int(webinarID)).
		SetUserID(int(*params.UserId)).
		Save(ctx)
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(user)
}

// Удаление юзера вебинара
// (DELETE /api/v1/webinar/{webinarID}/user/{userID})
func (s *Server) WebinarUserDestroy(c echo.Context, webinarID api.WebinarID, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseNotFound := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusNotFound),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(user *ent.WebinarUser) error {
		data := struct {
			api.WebinarUserDestroyData
		}{
			WebinarUserDestroyData: api.WebinarUserDestroyData{
				Status:       webinarUserDestroyMessage,
				SubscriberId: user.ID,
				UserId:       user.UserID,
				Webinar:      user.WebinarID,
			},
		}

		return c.JSON(http.StatusOK, &api.WebinarUserDestroyResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   ptrB(false),
				Message: ptrS(""),
			},
			ApiWebinarUserDestroyObject: api.ApiWebinarUserDestroyObject{
				Data: &data,
			},
		})
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), time.Second*2)

	var user *ent.WebinarUser
	err := s.ent.WithTx(ctx, func(tx *ent.Tx) (err error) {
		user, err = tx.WebinarUser.Query().Where(
			webinaruser.IDEQ(int(userID)),
			webinaruser.WebinarIDEQ(int(webinarID)),
		).First(ctx)
		if err != nil {
			return err
		}

		_, err = tx.WebinarUser.Delete().
			Where(
				webinaruser.IDEQ(int(userID)),
				webinaruser.WebinarIDEQ(int(webinarID)),
			).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if ent.IsNotFound(err) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(user)
}

// Получение юзера вебинара
// (GET /api/v1/webinar/{webinarID}/user/{userID})
func (s *Server) WebinarUserShow(c echo.Context, webinarID api.WebinarID, userID api.UserID) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseNotFound := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusNotFound),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(user *ent.WebinarUser) error {
		webinarUserJsonData, err := webinarUserToWebinarUserJsonData(user)
		if err != nil {
			return responseInternalServerError(err)
		}

		return c.JSON(http.StatusOK, &api.RootResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   ptrB(false),
				Message: ptrS(""),
			},
			ApiDataStringObject: api.ApiDataStringObject{
				Data: ptrS(string(webinarUserJsonData)),
			},
		})
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), time.Second*2)

	user, err := s.ent.WebinarUser.Query().Where(
		webinaruser.IDEQ(int(userID)),
		webinaruser.WebinarIDEQ(int(webinarID)),
	).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(user)
}

// Изменение юзера вебинара
// (PATCH /api/v1/webinar/{webinarID}/user/{userID})
func (s *Server) WebinarUserUpdate(c echo.Context, webinarID api.WebinarID, userID api.UserID, params api.WebinarUserUpdateParams) error {
	responseInternalServerError := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusInternalServerError),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseNotFound := func(err error) error {
		return c.JSON(http.StatusOK, &api.ApiError{
			Code:  ptrI(http.StatusNotFound),
			Data:  &[]interface{}{nil},
			Error: ptrB(true),
		})
	}

	responseSuccess := func(user *ent.WebinarUser) error {
		webinarUserJsonData, err := webinarUserToWebinarUserJsonData(user)
		if err != nil {
			return responseInternalServerError(err)
		}

		return c.JSON(http.StatusOK, &api.RootResponse200{
			ApiMessage: api.ApiMessage{
				Code:    ptrI(http.StatusOK),
				Error:   ptrB(false),
				Message: ptrS(""),
			},
			ApiDataStringObject: api.ApiDataStringObject{
				Data: ptrS(string(webinarUserJsonData)),
			},
		})
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), time.Second*2)

	var user *ent.WebinarUser
	err := s.ent.WithTx(ctx, func(tx *ent.Tx) (err error) {
		user, err = tx.WebinarUser.Query().Where(
			webinaruser.IDEQ(int(userID)),
			webinaruser.WebinarIDEQ(int(webinarID)),
		).First(ctx)
		if err != nil {
			return err
		}

		err = tx.WebinarUser.Update().
			Where(
				webinaruser.IDEQ(int(userID)),
				webinaruser.WebinarIDEQ(int(webinarID)),
			).
			SetMic(int16(params.Mic)).
			SetSound(int16(params.Sound)).
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if ent.IsNotFound(err) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(user)
}

func webinarUserToWebinarUserJsonData(user *ent.WebinarUser) (api.WebinarUserJsonData, error) {
	bb, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	return api.WebinarUserJsonData(string(bb)), nil
}
