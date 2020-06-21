package handlers

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/role"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/index"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var roleAlreadyApplied = "Role already applied"
var userDoesNotHaveRole = "User does not have the specified role"
var canNotDeleteAdminRole = "Can't delete ADMIN role from admin"

// GetRoles - GET /roles
func GetRoles(_ role.GetRolesParams) middleware.Responder {
	responseNotFound := func() middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
		}).Error("Not found")

		panic(nil)
	}

	responseInternalError := func(err error) middleware.Responder {
		payload := new(role.GetRolesInternalServerErrorBody)
		payload.Version = &Version
		payload.Errors = err.Error()
		payload.Message = &PayloadSuccessMessage

		return role.NewGetRolesInternalServerError().WithPayload(payload)
	}

	responseSuccess := func(roles []*storage.Role) middleware.Responder {
		payload := new(role.GetRolesOKBody)
		payload.Version = &Version

		for _, roleItem := range roles {
			item := new(role.DataItems0)
			item.ID = &roleItem.ID
			item.Description = &roleItem.Description
			isDefault := storage.IsDefaultRole(roleItem.ID)
			item.Default = &isDefault

			payload.Data = append(payload.Data, item)
		}

		payload.Message = &PayloadSuccessMessage
		return role.NewGetRolesOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	roles, err := rs.GetAll()
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(roles)
}

// GetRolesRoleID - GET /roles/<roleID>
func GetRolesRoleID(params role.GetRolesRoleIDParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(role.GetRolesRoleIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return role.NewGetRolesRoleIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(el *storage.Role) middleware.Responder {
		payload := new(role.GetRolesRoleIDOKBody)
		payload.Version = &Version

		item := new(role.DataItems0)
		item.ID = &el.ID
		item.Description = &el.Description
		isDefault := storage.IsDefaultRole(el.ID)
		item.Default = &isDefault

		payload.Data = append(payload.Data, item)
		payload.Message = &PayloadSuccessMessage
		return role.NewGetRolesRoleIDOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	roleItem, err := rs.Get(params.RoleID)
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(roleItem)
}

func PostRoles(params role.PostRolesParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseBadRequest := func() middleware.Responder {
		payload := new(role.PostRolesRoleIDUsersUserIDBadRequestBody)
		payload.Version = &Version
		payload.Message = &PayloadValidationErrorMessage

		validation := make(map[string]interface{})
		validation["ID"] = "unique"

		payload.Errors = new(role.PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors)
		payload.Errors.Validation = validation
		return role.NewPostRolesRoleIDUsersUserIDBadRequest().WithPayload(payload)
	}

	responseSuccess := func(roleItem storage.Role) middleware.Responder {
		payload := new(role.PostRolesOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(role.DataItems0)
		item.ID = &roleItem.ID
		item.Description = &roleItem.Description
		isDefault := storage.IsDefaultRole(roleItem.ID)
		item.Default = &isDefault
		payload.Data = append(payload.Data, item)

		return role.NewPostRolesOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	_, err := rs.Get(*params.Body.ID)
	if err == nil {
		return responseBadRequest()
	}

	roleItem := storage.Role{
		ID:          *params.Body.ID,
		Description: *params.Body.Description,
	}
	err2 := rs.Store(roleItem)
	if err2 != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(roleItem)
}

// PutRolesRoleID - PUT /roles/<roleID>
func PutRolesRoleID(params role.PutRolesRoleIDParams) middleware.Responder {
	responseNotFound := func() middleware.Responder {
		payload := new(role.PutRolesRoleIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return role.NewPutRolesRoleIDNotFound().WithPayload(payload)
	}

	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(roleItem *storage.Role) middleware.Responder {
		payload := new(role.PutRolesRoleIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		item := new(role.DataItems0)
		item.ID = &roleItem.ID
		item.Description = &roleItem.Description
		isDefault := storage.IsDefaultRole(roleItem.ID)
		item.Default = &isDefault
		payload.Data = append(payload.Data, item)
		return role.NewPutRolesRoleIDOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	roleItem := &storage.Role{
		ID:          params.RoleID,
		Description: params.Body.Description,
	}

	err := rs.Update(*roleItem)
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(roleItem)
}

func DeleteRolesRoleID(params role.DeleteRolesRoleIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(role.DeleteRolesRoleIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return role.NewDeleteRolesRoleIDNotFound().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(role.DeleteRolesRoleIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		payload.Data = nil
		return role.NewDeleteRolesRoleIDOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	err := rs.Delete(params.RoleID)
	if err != nil && err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess()
}

// GetRolesRoleIDUsers - GET /roles/<roleID>/users
func GetRolesRoleIDUsers(params role.GetRolesRoleIDUsersParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(role.GetRolesRoleIDUsersNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return role.NewGetRolesRoleIDUsersNotFound().WithPayload(payload)
	}

	responseSuccess := func(users []*storage.User) middleware.Responder {
		payload := new(role.GetRolesRoleIDUsersOKBody)
		payload.Version = &Version

		for _, user := range users {
			payload.Data = append(payload.Data, user.ID)
		}

		payload.Message = &PayloadSuccessMessage
		return role.NewGetRolesRoleIDUsersOK().WithPayload(payload)
	}

	rs := st.GetStorage().RoleStorage
	us := st.GetStorage().UserStorage
	_, err := rs.Get(params.RoleID)
	if err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalError(err)
	}

	users, err := us.Get(storage.GetUser{
		RoleID: &params.RoleID,
	})
	if err != nil {
		return responseInternalError(err)
	}

	if len(users) == 0 {
		return responseNotFound()
	}

	return responseSuccess(users)
}

// PostRolesRoleIDUsersUserID - POST /roles/<roleID>/users/<userID>
func PostRolesRoleIDUsersUserID(params role.PostRolesRoleIDUsersUserIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err *error) middleware.Responder {
		payload := new(role.PostRolesRoleIDUsersUserIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		if err != nil {
			payload.Errors = append(payload.Errors, (*err).Error())
		}

		return role.NewPostRolesRoleIDUsersUserIDNotFound().WithPayload(payload)
	}

	responseBadRequest := func() middleware.Responder {
		payload := new(role.PostRolesRoleIDUsersUserIDBadRequestBody)
		payload.Version = &Version
		payload.Message = &roleAlreadyApplied
		return role.NewPostRolesRoleIDUsersUserIDBadRequest().WithPayload(payload)
	}

	responseSuccess := func(userId string) middleware.Responder {
		payload := new(role.PostRolesRoleIDUsersUserIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		payload.Data = append(payload.Data, userId)

		return role.NewPostRolesRoleIDUsersUserIDOK().WithPayload(payload)
	}

	us := st.GetStorage().UserStorage
	rs := st.GetStorage().RoleStorage

	users, err := us.Get(storage.GetUser{
		ID: &params.UserID,
	})
	if err != nil {
		return responseInternalServerError(err)
	} else if len(users) == 0 {
		// If user not found in local db, we try do search request to Index.
		// If user exist at Index we duplicate user at local db
		searchOk, err := index.Search(params.UserID)
		if err != nil {
			return responseNotFound(nil)
		}

		if searchOk == nil {
			return responseNotFound(nil)
		}

		data := searchOk.Payload.Data[0]

		if user, err := us.Store(storage.StoreUser{
			User: storage.User{
				ID: data.GUID,
			},
		}); err != nil {
			return responseInternalServerError(err)
		} else {
			users = append(users, user)
		}
	}
	user := users[0]

	userRole, err := rs.Get(params.RoleID)
	if err != nil {
		return responseNotFound(&err)
	}

	if user.Roles[userRole.ID] {
		return responseBadRequest()
	}

	if user.Roles == nil {
		user.Roles = make(map[string]bool)
	}

	user.Roles[userRole.ID] = true

	if _, err := us.Update(user.ID, storage.UpdateUser{
		Roles:           user.Roles,
		BackendEntryIDs: user.BackendEntryIDs,
	}); err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(user.ID)
}

// DeleteRolesRoleIDUsersUserID - DELETE /roles/<roleID>/users/<userID>
func DeleteRolesRoleIDUsersUserID(params role.DeleteRolesRoleIDUsersUserIDParams) middleware.Responder {
	responceNotFound := func(err error) middleware.Responder {
		payload := new(role.DeleteRolesRoleIDUsersUserIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, err.Error())
		payload.Message = &NotFoundMessage

		return role.NewDeleteRolesRoleIDUsersUserIDNotFound().WithPayload(payload)
	}

	responseUserNotHaveRole := func() middleware.Responder {
		payload := new(role.DeleteRolesRoleIDUsersUserIDBadRequestBody)
		payload.Version = &Version
		payload.Message = &userDoesNotHaveRole

		return role.NewDeleteRolesRoleIDUsersUserIDBadRequest().WithPayload(payload)
	}

	responseCannotDeleteAdminRole := func() middleware.Responder {
		payload := new(role.DeleteRolesRoleIDUsersUserIDBadRequestBody)
		payload.Version = &Version
		payload.Errors = new(role.DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors)
		payload.Message = &canNotDeleteAdminRole

		return role.NewDeleteRolesRoleIDUsersUserIDBadRequest().WithPayload(payload)
	}

	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(role.DeleteRolesRoleIDUsersUserIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		return role.NewDeleteRolesRoleIDUsersUserIDOK().WithPayload(payload)
	}

	us := st.GetStorage().UserStorage
	rs := st.GetStorage().RoleStorage

	users, err := us.Get(storage.GetUser{
		ID: &params.UserID,
	})
	if err != nil {
		return responseInternalError(err)
	}

	// FIXME:
	if len(users) == 0 {
		return responceNotFound(errors.New("user not found"))
	}

	userRole, err := rs.Get(params.RoleID)
	if err != nil && err == storage.ErrNotFound {
		return responceNotFound(err)
	} else if err != nil {
		return responseInternalError(err)
	}

	if !users[0].Roles[userRole.ID] {
		return responseUserNotHaveRole()
	}

	if users[0].ID == "admin" && userRole.ID == "ADMIN" {
		return responseCannotDeleteAdminRole()
	}

	if _, err := us.Update(params.UserID, storage.UpdateUser{
		Roles: map[string]bool{
			userRole.ID: false,
		},
	}); err != nil {
		return responseInternalError(err)
	}

	return responseSuccess()
}
