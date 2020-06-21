package handlers

import (
	"errors"
	"fmt"
	"sort"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/auth/authentication"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/backend"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	ErrWrongBackendType = errors.New("wrong backend type")
)

// GetBackends - GET /backends
func GetBackends(_ backend.GetBackendsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		payload := new(backend.GetBackendsInternalServerErrorBody)
		payload.Version = &Version
		payload.Errors = err.Error()
		payload.Message = &PayloadSuccessMessage

		return backend.NewGetBackendsInternalServerError().WithPayload(payload)
	}

	responseSuccess := func(backends []*storage.Backend) middleware.Responder {
		payload := new(backend.GetBackendsOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		for _, clientItem := range backends {
			payload.Data = append(payload.Data, clientItem)
		}
		return backend.NewGetBackendsOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{})
	if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(backends)
}

// GetBackendsBackendID - GET /backends/{backendID}
func GetBackendsBackendID(params backend.GetBackendsBackendIDParams) middleware.Responder {
	responseNotFound := func() middleware.Responder {
		payload := new(backend.GetBackendsBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = []interface{}{NotFoundMessage}
		payload.Message = &NotFoundMessage

		return backend.NewGetBackendsBackendIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(item *storage.Backend) middleware.Responder {
		payload := new(backend.GetBackendsBackendIDOKBody)
		payload.Version = &Version
		payload.Data = append(payload.Data, item)
		payload.Message = &PayloadSuccessMessage

		return backend.NewGetBackendsBackendIDOK().WithPayload(payload)
	}

	responseInternalError := func(err error) middleware.Responder {
		payload := new(backend.GetBackendsBackendIDInternalServerErrorBody)
		payload.Version = &Version
		payload.Errors = err.Error()
		payload.Message = &PayloadSuccessMessage

		return backend.NewGetBackendsBackendIDInternalServerError().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalError(err)
	} else if len(backends) == 0 {
		return responseNotFound()
	}

	return responseSuccess(backends[0])
}

// GetBackendsBackendIDTest - GET /backends/{backendID}/test
func GetBackendsBackendIDTest(params backend.GetBackendsBackendIDTestParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(backend.GetBackendsBackendIDTestNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, err.Error())
		payload.Message = &NotFoundMessage
		return backend.NewGetBackendsBackendIDTestNotFound().WithPayload(payload)
	}

	responseBadRequest := func() middleware.Responder {
		payload := new(backend.GetBackendsBackendIDTestBadRequestBody)
		payload.Version = &Version
		payload.Message = &PayloadFailMessage
		item := new(backend.TestDataItem)
		item.Status = "FAILED"
		text := "Protocol failure"
		item.Error = &text
		payload.Data = append(payload.Data, item)

		return backend.NewGetBackendsBackendIDTestBadRequest().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(backend.GetBackendsBackendIDTestOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		item := new(backend.TestDataItem)
		item.Status = PayloadSuccessMessage
		payload.Data = append(payload.Data, item)

		return backend.NewGetBackendsBackendIDTestOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalError(err)
	} else if len(backends) == 0 {
		return responseNotFound(storage.ErrNotFound)
	}

	err = authentication.TestBackend(backends[0])
	if err != nil {
		return responseBadRequest()
	}

	return responseSuccess()
}

// PostBackendsBackendIDGroups - POST /backends/{backendID}/groups
func PostBackendsBackendIDGroups(params backend.PostBackendsBackendIDGroupsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(backend.PostBackendsBackendIDGroupsNotFoundBody)
		payload.Version = &Version
		payload.Errors = []interface{}{err.Error()}
		payload.Message = &NotFoundMessage
		return backend.NewPostBackendsBackendIDGroupsNotFound().WithPayload(payload)
	}

	responseSuccess := func(groups map[string]string) middleware.Responder {
		payload := new(backend.PostBackendsBackendIDGroupsOKBody)
		payload.Version = &Version
		for group, roleID := range groups {
			itemGroup := group
			itemRoleID := roleID

			item := new(backend.DataItems0)
			item.Group = &itemGroup
			item.RoleID = &itemRoleID

			payload.Data = append(payload.Data, item)
		}
		payload.Message = &PayloadSuccessMessage
		return backend.NewPostBackendsBackendIDGroupsOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalError(err)
	} else if len(backends) == 0 {
		return responseNotFound(storage.ErrNotFound)
	}

	if _, err := bs.Update(backends[0].ID, storage.UpdateBackend{
		Groups: map[string]string{},
	}); err != nil {
		return responseInternalError(err)
	}

	var stored *storage.Backend
	var errs []error
	for _, entry := range params.Body {
		_, err := st.GetStorage().RoleStorage.Get(*entry.RoleID)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errs = append(errs, fmt.Errorf("role %s not found", *entry.RoleID))
				continue
			}

			return responseInternalError(err)
		}

		stored, err = bs.Update(backends[0].ID, storage.UpdateBackend{
			Groups: map[string]string{
				*entry.Group: *entry.RoleID,
			},
		})
		if err != nil {
			return responseInternalError(err)
		}
	}

	if len(errs) != 0 {
		payload := new(backend.PostBackendsBackendIDGroupsNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		for _, err := range errs {
			payload.Errors = append(payload.Errors, err)
		}

		return backend.NewPostBackendsBackendIDGroupsNotFound().WithPayload(payload)
	}

	return responseSuccess(stored.Groups)
}

// DeleteBackendsBackendIDGroups - DELETE /backends/{backendID}/groups
func DeleteBackendsBackendIDGroups(params backend.DeleteBackendsBackendIDGroupsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(backend.DeleteBackendsBackendIDGroupsNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return backend.NewDeleteBackendsBackendIDGroupsNotFound().WithPayload(payload)
	}

	responseSuccess := func(groups map[string]string) middleware.Responder {
		payload := new(backend.DeleteBackendsBackendIDGroupsOKBody)
		payload.Version = &Version

		for group, roleID := range groups {
			itemGroup := group
			itemRoleID := roleID

			item := new(backend.DataItems0)
			item.Group = &itemGroup
			item.RoleID = &itemRoleID

			payload.Data = append(payload.Data, item)
		}
		payload.Message = &PayloadSuccessMessage
		return backend.NewDeleteBackendsBackendIDGroupsOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalError(err)
	} else if len(backends) == 0 {
		return responseNotFound()
	}

	stored, err := bs.Update(backends[0].ID, storage.UpdateBackend{
		Groups: map[string]string{},
	})
	if err != nil {
		return responseInternalError(err)
	}

	return responseSuccess(stored.Groups)
}

// GetBackendsBackendIDGroups - GET /backends/{backendID}/groups
func GetBackendsBackendIDGroups(params backend.GetBackendsBackendIDGroupsParams) middleware.Responder {
	responseInternalError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(backend.GetBackendsBackendIDGroupsNotFoundBody)
		payload.Version = &Version
		payload.Errors = []interface{}{err.Error()}
		payload.Message = &NotFoundMessage
		return backend.NewGetBackendsBackendIDGroupsNotFound().WithPayload(payload)
	}

	responseSuccess := func(groups map[string]string) middleware.Responder {
		payload := new(backend.GetBackendsBackendIDGroupsOKBody)
		payload.Version = &Version
		for group, roleID := range groups {
			itemGroup := group
			itemRoleID := roleID

			item := new(backend.DataItems0)
			item.Group = &itemGroup
			item.RoleID = &itemRoleID

			payload.Data = append(payload.Data, item)
		}
		payload.Message = &PayloadSuccessMessage

		return backend.NewGetBackendsBackendIDGroupsOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	backends, err := bs.Get(storage.GetBackend{
		ID: &params.BackendID,
	})
	if err != nil {
		return responseInternalError(err)
	} else if len(backends) == 0 {
		return responseNotFound(storage.ErrNotFound)
	}

	return responseSuccess(backends[0].Groups)
}

// PostBackendsLDAP - POST /backends/ldap
func PostBackendsLDAP(params backend.PostBackendsLdapParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(backendItem storage.Backend) middleware.Responder {
		payload := new(backend.PostBackendsLdapOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.LdapDataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Host = &backendItem.Host
		item.Port = &backendItem.Port
		item.Cipher = &backendItem.Cipher
		item.BindDN = &backendItem.BindDN
		item.BaseDN = &backendItem.BaseDN
		item.Filter = &backendItem.Filter
		item.Sync = &backendItem.Sync

		item.Attributes.BackendAttributeParams = backendItem.Attributes

		payload.Data = append(payload.Data, item)

		return backend.NewPostBackendsLdapOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage
	bos := st.GetStorage().BackendsOrderStorage

	var backendItem storage.Backend
	backendItem.ID = uuid.New().String()
	backendItem.Type = storage.BackendTypeLDAP
	backendItem.Description = *params.Body.Description
	backendItem.Host = *params.Body.Host
	backendItem.Port = *params.Body.Port
	backendItem.Cipher = *params.Body.Cipher
	backendItem.BindDN = *params.Body.BindDN
	backendItem.BaseDN = *params.Body.BaseDN
	backendItem.Filter = *params.Body.Filter
	backendItem.Password = *params.Body.Password
	backendItem.Sync = *params.Body.Sync
	backendItem.Attributes = params.Body.Attributes.BackendAttributeParams

	stored, err := bs.Store(storage.StoreBackend{
		Backend: backendItem,
	})
	if err != nil {
		return responseInternalServerError(err)
	}

	// Append backend to auth flow.
	backendsOrder, err := bos.Get()
	if err != nil {
		return responseInternalServerError(err)
	}

	backendsOrder.Order = append(backendsOrder.Order, backendItem.ID)

	_, err = bos.Update(storage.UpdateBackendsOrder{
		Order: backendsOrder.Order,
	})
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(*stored)
}

// PutBackendsLdapBackendID - PUT /backends/ldap/{backendID}
func PutBackendsLdapBackendID(params backend.PutBackendsLdapBackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(backend.PutBackendsLdapBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return backend.NewPutBackendsLdapBackendIDNotFound().WithPayload(payload)
	}

	/*responseValidationError := func() middleware.Responder {
		// TODO: the response use logically incorrect type (NotFound)
		payload := new(backend.PutBackendsLdapBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, "wrong type")
		payload.Message = &NotFoundMessage
		return backend.NewPutBackendsLdapBackendIDNotFound().WithPayload(payload)
	}*/

	responseSuccess := func(backendItem *storage.Backend) middleware.Responder {
		payload := new(backend.PutBackendsLdapBackendIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.LdapDataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Host = &backendItem.Host
		item.Port = &backendItem.Port
		item.Cipher = &backendItem.Cipher
		item.BindDN = &backendItem.BindDN
		item.BaseDN = &backendItem.BaseDN
		item.Filter = &backendItem.Filter
		item.Sync = &backendItem.Sync

		item.Attributes.BackendAttributeParams = backendItem.Attributes

		payload.Data = append(payload.Data, item)

		return backend.NewPutBackendsLdapBackendIDOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	stored, err := bs.Update(params.BackendID, storage.UpdateBackend{
		Description: params.Body.Description,
		Sync:        params.Body.Sync,
		Attributes:  &params.Body.Attributes.BackendAttributeParams,
		Host:        params.Body.Host,
		Port:        params.Body.Port,
		Cipher:      params.Body.Cipher,
		BindDN:      params.Body.BindDN,
		BaseDN:      params.Body.BaseDN,
		Filter:      params.Body.Filter,
		Password:    params.Body.Password,
	})
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return responseNotFound()
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(stored)
}

// PatchBackendsLdapBackendID - PATCH /backends/ldap/{backendID}
func PatchBackendsLdapBackendID(params backend.PatchBackendsLdapBackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(backend.PatchBackendsLdapBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return backend.NewPatchBackendsLdapBackendIDNotFound().WithPayload(payload)
	}

	/*responseValidationError := func() middleware.Responder {
		// TODO: the response use logically incorrect type (NotFound)
		payload := new(backend.PatchBackendsLdapBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Errors = append(payload.Errors, "wrong type")
		payload.Message = &NotFoundMessage
		return backend.NewPatchBackendsLdapBackendIDNotFound().WithPayload(payload)
	}*/

	responseSuccess := func(backendItem *storage.Backend) middleware.Responder {
		payload := new(backend.PatchBackendsLdapBackendIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.LdapDataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Host = &backendItem.Host
		item.Port = &backendItem.Port
		item.Cipher = &backendItem.Cipher
		item.BindDN = &backendItem.BindDN
		item.BaseDN = &backendItem.BaseDN
		item.Filter = &backendItem.Filter
		item.Sync = &backendItem.Sync

		item.Attributes.BackendAttributeParams = backendItem.Attributes

		payload.Data = append(payload.Data, item)

		return backend.NewPatchBackendsLdapBackendIDOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	stored, err := bs.Update(params.BackendID, storage.UpdateBackend{
		Description: params.Body.Description,
		Sync:        params.Body.Sync,
		Host:        params.Body.Host,
		Port:        params.Body.Port,
		Cipher:      params.Body.Cipher,
		BindDN:      params.Body.BindDN,
		BaseDN:      params.Body.BaseDN,
		Filter:      params.Body.Filter,
		Password:    params.Body.Password,
	})
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return responseNotFound()
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(stored)
}

// PostBackendsOAuth2 - POST /backends/oauth2
func PostBackendsOAuth2(params backend.PostBackendsOauth2Params) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseSuccess := func(backendItem *storage.Backend) middleware.Responder {
		payload := new(backend.PostBackendsOauth2OKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.Oauth2DataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Sync = &backendItem.Sync

		item.ClientID = &backendItem.ClientID
		item.ClientSecret = &backendItem.ClientSecret
		item.Provider = &backendItem.Provider

		item.Attributes.BackendAttributeParams = backendItem.Attributes

		payload.Data = append(payload.Data, item)

		return backend.NewPostBackendsOauth2OK().WithPayload(payload)
	}

	backendItem := new(storage.Backend)
	backendItem.Type = storage.BackendTypeOAuth2
	backendItem.Description = *params.Body.Description
	backendItem.Sync = *params.Body.Sync
	backendItem.Attributes = params.Body.Attributes.BackendAttributeParams
	backendItem.ClientID = *params.Body.ClientID
	backendItem.ClientSecret = *params.Body.ClientSecret
	backendItem.Provider = *params.Body.Provider

	// Custom Backend ID
	if params.Body.ID != "" {
		backendItem.ID = params.Body.ID
	} else {
		backendItem.ID = uuid.New().String()
	}

	bs := st.GetStorage().BackendStorage

	stored, err := bs.Store(storage.StoreBackend{
		Backend: *backendItem,
	})
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(stored)
}

// PutBackendsOAuth2BackendID - PUT /backends/oauth2/{backendID}
func PutBackendsOAuth2BackendID(params backend.PutBackendsOauth2BackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(backend.PutBackendsOauth2BackendIDNotFoundBody)
		payload.Errors = append(payload.Errors, err)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return backend.NewPutBackendsOauth2BackendIDNotFound().WithPayload(payload)
	}

	responseSuccess := func(backendItem *storage.Backend) middleware.Responder {
		payload := new(backend.PutBackendsOauth2BackendIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.Oauth2DataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Sync = &backendItem.Sync
		item.Attributes.BackendAttributeParams = backendItem.Attributes

		item.ClientID = &backendItem.ClientID
		item.ClientSecret = &backendItem.ClientSecret
		item.Provider = &backendItem.Provider

		payload.Data = append(payload.Data, item)

		return backend.NewPutBackendsOauth2BackendIDOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	in := storage.UpdateBackend{
		Description:  params.Body.Description,
		Sync:         params.Body.Sync,
		Attributes:   &params.Body.Attributes.BackendAttributeParams,
		ClientID:     params.Body.ClientID,
		ClientSecret: params.Body.ClientSecret,
		Provider:     params.Body.Provider,
	}

	if params.Body.ID != "" {
		in.ID = &params.Body.ID
	}

	stored, err := bs.Update(params.BackendID, in)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(stored)
}

// PatchBackendsOAuth2BackendID - PATCH /backends/oauth2/{backendID}
func PatchBackendsOAuth2BackendID(params backend.PatchBackendsOauth2BackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func(err error) middleware.Responder {
		payload := new(backend.PatchBackendsOauth2BackendIDNotFoundBody)
		payload.Errors = append(payload.Errors, err)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return backend.NewPatchBackendsOauth2BackendIDNotFound().WithPayload(payload)
	}

	/*responseValidationError := func(err error) middleware.Responder {
		// TODO: the response use logically incorrect type (NotFound)
		payload := new(backend.PatchBackendsOauth2BackendIDNotFoundBody)
		payload.Errors = append(payload.Errors, err)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return backend.NewPatchBackendsOauth2BackendIDNotFound().WithPayload(payload)
	}*/

	responseSuccess := func(backendItem *storage.Backend) middleware.Responder {
		payload := new(backend.PatchBackendsOauth2BackendIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		item := new(backend.Oauth2DataItem)
		item.ID = backendItem.ID
		item.Description = &backendItem.Description
		item.Sync = &backendItem.Sync
		item.Attributes.BackendAttributeParams = backendItem.Attributes

		item.ClientID = &backendItem.ClientID
		item.ClientSecret = &backendItem.ClientSecret
		item.Provider = &backendItem.Provider

		payload.Data = append(payload.Data, item)

		return backend.NewPatchBackendsOauth2BackendIDOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage

	in := storage.UpdateBackend{
		Description:  params.Body.Description,
		Sync:         params.Body.Sync,
		ClientID:     params.Body.ClientID,
		ClientSecret: params.Body.ClientSecret,
		Provider:     params.Body.Provider,
	}

	if params.Body.ID != "" {
		in.ID = &params.Body.ID
	}

	stored, err := bs.Update(params.BackendID, in)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return responseNotFound(err)
		}

		return responseInternalServerError(err)
	}

	return responseSuccess(stored)
}

// DeleteBackendsBackendID - DELETE /backends/{backendID}
func DeleteBackendsBackendID(params backend.DeleteBackendsBackendIDParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"error":   err,
			"context": "API",
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(backend.DeleteBackendsBackendIDNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage
		return backend.NewDeleteBackendsBackendIDNotFound().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(backend.DeleteBackendsBackendIDOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		return backend.NewDeleteBackendsBackendIDOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendStorage
	bos := st.GetStorage().BackendsOrderStorage

	if err := bos.Delete(storage.DeleteBackendsOrder{IDs: []string{
		params.BackendID,
	}}); err != nil {
		return responseInternalServerError(err)
	}

	err := bs.Delete(params.BackendID)
	if err == storage.ErrNotFound {
		return responseNotFound()
	} else if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess()
}

// GetFlow - GET /flow
func GetFlow(_ backend.GetFlowParams) middleware.Responder {
	internalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"error":   err,
			"context": "API",
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	successResponse := func(backendsOrder []string) middleware.Responder {
		payload := new(backend.GetFlowOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		for i, item := range backendsOrder {
			orderID := item
			orderPosition := int64(i + 1)

			flowDataItem := new(backend.FlowDataItem)
			flowDataItem.BackendID = &orderID
			flowDataItem.Order = &orderPosition

			payload.Data = append(payload.Data, flowDataItem)
		}

		return backend.NewGetFlowOK().WithPayload(payload)
	}

	bs := st.GetStorage().BackendsOrderStorage
	backendsOrder, err := bs.Get()
	if err != nil {
		return internalServerError(err)
	}

	return successResponse(backendsOrder.Order)
}

// PostFlow - POST /flow
func PostFlow(params backend.PostFlowParams) middleware.Responder {
	responseInternalServerError := func(err error) middleware.Responder {
		// TODO: the specification does not determine the response that has arisen
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   err,
		}).Error(InternalServerErrorMessage)

		panic(err)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(backend.PostFlowNotFoundBody)
		payload.Version = &Version
		payload.Message = &NotFoundMessage

		return backend.NewPostFlowNotFound().WithPayload(payload)
	}

	responseSuccess := func(backendsOrder []string) middleware.Responder {
		payload := new(backend.PostFlowOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		for i, item := range backendsOrder {
			orderID := item
			orderPosition := int64(i + 1)

			flowDataItem := new(backend.FlowDataItem)
			flowDataItem.BackendID = &orderID
			flowDataItem.Order = &orderPosition

			payload.Data = append(payload.Data, flowDataItem)
		}

		return backend.NewPostFlowOK().WithPayload(payload)
	}

	bos := st.GetStorage().BackendsOrderStorage
	bs := st.GetStorage().BackendStorage

	backends := params.Body
	sort.Sort(backendsByOrder(backends))

	backendsOrder, err := bos.Get()
	if err != nil {
		return responseInternalServerError(err)
	}

	for _, item := range backends {
		if *item.BackendID != "index" {
			// Check backend availability
			backends, err := bs.Get(storage.GetBackend{
				ID: item.BackendID,
			})
			if err != nil {
				return responseInternalServerError(err)
			} else if len(backends) == 0 {
				return responseNotFound()
			}
		}

		backendsOrder.Order = append(backendsOrder.Order, *item.BackendID)
	}

	backendsOrder, err = bos.Update(storage.UpdateBackendsOrder{
		Order: backendsOrder.Order,
	})
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(backendsOrder.Order)
}

/* POST /flow backends sorter */
type backendsByOrder []*backend.PostFlowParamsBodyItems0

func (s backendsByOrder) Len() int {
	return len(s)
}
func (s backendsByOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s backendsByOrder) Less(i, j int) bool {
	return *s[i].Order < *s[j].Order
}
