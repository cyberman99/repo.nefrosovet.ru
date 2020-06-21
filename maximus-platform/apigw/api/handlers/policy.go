package handlers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
	"repo.nefrosovet.ru/maximus-platform/apigw/api"
	"repo.nefrosovet.ru/maximus-platform/apigw/mongodb"
)

// convertPolicy returns api.PolicyObjectWithId object from mongodb.Policy
func convertPolicy(policy *mongodb.Policy) (item api.PolicyObjectWithId) {
	item.ID = policy.ID
	item.Description = policy.Description
	item.Method = policy.Method
	item.Path = policy.Path
	item.BackendHost = policy.BackendHost
	item.BackendPath = policy.BackendPath
	item.Resource = policy.Resource
	item.Roles = policy.Roles
	item.QuerystringParams = policy.QueryStringParams
	item.HeadersToPass = policy.HeadersToPass
	item.KeyCache = policy.KeyCache
	item.Cache = policy.Cache

	return
}

// getValidationErrors getValidation fields from error
func getValidationErrors(err error) map[string]interface{} {
	validation := make(map[string]interface{})
	if httpError, ok := err.(*echo.HTTPError); ok {
		if unmarshalError, ok := httpError.Internal.(*json.UnmarshalTypeError); ok {
			validation[unmarshalError.Field] = unmarshalError.Type.String()
		}
	}
	return validation
}

// PolicyCollection Коллекция политик (GET /policies)
func (s *Server) PolicyCollection(ctx echo.Context) error {
	responseInternalServerError := func() error {
		payload500 := api.BaseResponse500{}
		payload500.Version = s.Version
		payload500.Message = InternalServerErrorMessage
		return ctx.JSON(http.StatusInternalServerError, payload500)
	}

	responseSuccess := func(policyCollection []*mongodb.Policy) error {
		payload200 := api.PolicyResponse200{}
		payload200.Version = s.Version
		payload200.Message = PayloadSuccessMessage
		for _, policy := range policyCollection {
			item := convertPolicy(policy)
			payload200.Data = append(payload200.Data, item)
		}
		return ctx.JSON(http.StatusOK, payload200)
	}

	policyCollection, err := s.repo.GetPolicies()
	if err != nil && err != mongo.ErrNoDocuments {
		return responseInternalServerError()
	}

	return responseSuccess(policyCollection)
}

// PolicyCreate Создание политики (POST /policies)
func (s *Server) PolicyCreate(ctx echo.Context) error {
	responseInternalServerError := func() error {
		payload500 := api.BaseResponse500{}
		payload500.Version = s.Version
		payload500.Message = InternalServerErrorMessage
		return ctx.JSON(http.StatusInternalServerError, payload500)
	}

	responseBadRequest := func(msg string, validation *map[string]interface{}) error {
		payload400 := api.BaseResponse400{}
		payload400.Version = s.Version
		payload400.Message = msg
		payload400.Errors.Validation = validation

		return ctx.JSON(http.StatusBadRequest, payload400)
	}

	responseSuccess := func(policy *mongodb.Policy) error {
		payload200 := api.PolicyResponse200{}
		payload200.Version = Version
		payload200.Message = PayloadSuccessMessage
		item := convertPolicy(policy)
		payload200.Data = append(payload200.Data, item)
		return ctx.JSON(http.StatusOK, payload200)
	}

	params := api.PolicyCreateJSONRequestBody{}
	if err := ctx.Bind(&params); err != nil {
		validation := getValidationErrors(err)
		return responseBadRequest(PayloadValidationErrorMessage, &validation)
	}

	validation := make(map[string]interface{})
	if params.Description == "" {
		validation["description"] = "required"
	}
	if params.Resource == "" {
		validation["resource"] = "required"
	}
	if params.Method == "" {
		validation["method"] = "required"
	}
	if params.Path == "" {
		validation["path"] = "required"
	}
	if params.BackendPath == "" {
		validation["backendPath"] = "required"
	}
	if len(validation) > 0 {
		return responseBadRequest(PayloadValidationErrorMessage, &validation)
	}

	existedPolicy, err := s.repo.GetPolicyByResourceMethodPath(params.Resource, params.Method, params.Path)
	if err != nil && err != mongo.ErrNoDocuments {
		return responseInternalServerError()
	}
	if existedPolicy != nil {
		return responseBadRequest(PayloadPolicyAlreadyExisted, nil)
	}
	policy := mongodb.Policy{}
	policy.ID = params.ID
	policy.Description = params.Description
	policy.Method = params.Method
	policy.Path = params.Path
	policy.BackendHost = params.BackendHost
	policy.BackendPath = params.BackendPath
	policy.Resource = params.Resource
	policy.Roles = params.Roles
	policy.QueryStringParams = params.QuerystringParams
	policy.HeadersToPass = params.HeadersToPass
	policy.KeyCache = params.KeyCache
	policy.Cache = params.Cache

	policy.ID, err = s.repo.Insert(policy)
	if err != nil {
		return responseInternalServerError()
	}

	return responseSuccess(&policy)
}

// PolicyDelete Удаление политики (DELETE /policies/{policyID})
func (s *Server) PolicyDelete(ctx echo.Context, policyID api.PolicyID) error {
	responseInternalServerError := func() error {
		payload500 := api.BaseResponse500{}
		payload500.Version = s.Version
		payload500.Message = InternalServerErrorMessage
		return ctx.JSON(http.StatusInternalServerError, payload500)
	}

	responseNotFound := func() error {
		payload404 := api.BaseResponse404{}
		payload404.Version = s.Version
		payload404.Message = NotFoundMessage
		return ctx.JSON(http.StatusNotFound, payload404)
	}

	responseSuccess := func() error {
		payload200 := api.BaseResponse200{}
		payload200.Version = s.Version
		payload200.Message = PayloadSuccessMessage
		return ctx.JSON(http.StatusOK, payload200)
	}

	policy, err := s.repo.GetPolicyByID(string(policyID))
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return responseNotFound()
		}
		return responseInternalServerError()
	}

	err = s.repo.Delete(policy.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return responseNotFound()
		}
		return responseInternalServerError()
	}

	return responseSuccess()
}

// PolicyView Информация о политике (GET /policies/{policyID})
func (s *Server) PolicyView(ctx echo.Context, policyID api.PolicyID) error {
	responseInternalServerError := func() error {
		payload500 := api.BaseResponse500{}
		payload500.Version = s.Version
		payload500.Message = InternalServerErrorMessage
		return ctx.JSON(http.StatusInternalServerError, payload500)
	}

	responseNotFound := func() error {
		payload404 := api.BaseResponse404{}
		payload404.Version = s.Version
		payload404.Message = NotFoundMessage
		return ctx.JSON(http.StatusNotFound, payload404)
	}

	responseSuccess := func(policy *mongodb.Policy) error {
		payload200 := api.PolicyResponse200{}
		payload200.Version = s.Version
		payload200.Message = PayloadSuccessMessage
		item := convertPolicy(policy)
		payload200.Data = append(payload200.Data, item)
		return ctx.JSON(http.StatusOK, payload200)
	}

	policy, err := s.repo.GetPolicyByID(string(policyID))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return responseNotFound()
		}
		return responseInternalServerError()
	}
	return responseSuccess(policy)
}

// PolicyPatch Изменение политики (PATCH /policies/{policyID})
func (s *Server) PolicyPatch(ctx echo.Context, policyID api.PolicyID) error {
	responseInternalServerError := func() error {
		payload500 := api.Error500Data{}
		payload500.Version = s.Version
		payload500.Message = InternalServerErrorMessage
		return ctx.JSON(http.StatusInternalServerError, payload500)
	}

	responseNotFound := func() error {
		payload404 := api.BaseResponse404{}
		payload404.Version = s.Version
		payload404.Message = NotFoundMessage
		return ctx.JSON(http.StatusNotFound, payload404)
	}

	responseBadRequest := func(msg string, validation *map[string]interface{}) error {
		payload400 := api.BaseResponse400{}
		payload400.Version = s.Version
		payload400.Message = msg
		payload400.Errors.Validation = validation
		return ctx.JSON(http.StatusBadRequest, payload400)
	}

	responseSuccess := func(policy *mongodb.Policy) error {
		payload200 := api.PolicyResponse200{}
		payload200.Version = s.Version
		payload200.Message = PayloadSuccessMessage
		item := convertPolicy(policy)
		payload200.Data = append(payload200.Data, item)
		return ctx.JSON(http.StatusOK, payload200)
	}

	policy, err := s.repo.GetPolicyByID(string(policyID))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return responseNotFound()
		}
		return responseInternalServerError()
	}

	params := api.PolicyPatchJSONRequestBody{}
	if err := ctx.Bind(&params); err != nil {
		validation := getValidationErrors(err)
		return responseBadRequest(PayloadValidationErrorMessage, &validation)
	}

	if params.Description != nil {
		policy.Description = *params.Description
	}
	if params.Method != nil {
		policy.Method = *params.Method
	}
	if params.Path != nil {
		policy.Path = *params.Path
	}
	if params.Resource != nil {
		policy.Resource = *params.Resource
	}
	if params.BackendHost != nil {
		policy.BackendHost = *params.BackendHost
	}
	if params.BackendPath != nil {
		policy.BackendPath = *params.BackendPath
	}
	if params.Roles != nil {
		policy.Roles = *params.Roles
	}
	if params.QuerystringParams != nil {
		policy.QueryStringParams = *params.QuerystringParams
	}
	if params.HeadersToPass != nil {
		policy.HeadersToPass = *params.HeadersToPass
	}
	if params.KeyCache != nil {
		policy.KeyCache = *params.KeyCache
	}
	if params.Cache != nil {
		policy.Cache = *params.Cache
	}

	err = s.repo.Update(*policy)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return responseNotFound()
		}
		return responseInternalServerError()
	}

	return responseSuccess(policy)
}
