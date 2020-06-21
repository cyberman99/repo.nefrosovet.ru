package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
)

const (
	EchoContextKey = "oapi-codegen/echo-context"
	UserDataKey    = "oapi-codegen/user-data"
)

var (
	errRouteMissingSwagger = errors.New("route is missing OpenAPI specification")
)

func OAPIRequestValidator(swagger *openapi3.Swagger) echo.MiddlewareFunc {
	return OAPIRequestValidatorWithOptions(swagger, nil)
}

type Options struct {
	Options      openapi3filter.Options
	ParamDecoder openapi3filter.ContentParameterDecoder
	UserData     interface{}
}

func OAPIRequestValidatorWithOptions(swagger *openapi3.Swagger, options *Options) echo.MiddlewareFunc {
	router := openapi3filter.NewRouter().WithSwagger(swagger)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := ValidateRequestFromContext(c, router, options)
			if err != nil {
				return err
			}

			return next(c)
		}
	}
}

func ValidateRequestFromContext(ctx echo.Context, router *openapi3filter.Router, options *Options) error {
	r := ctx.Request()

	route, pathParams, err := router.FindRoute(r.Method, r.URL)
	if err != nil {
		return err
	}

	validationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
	}

	requestContext := context.WithValue(context.Background(), EchoContextKey, ctx)
	if options != nil {
		validationInput.Options = &options.Options
		validationInput.ParamDecoder = options.ParamDecoder
		requestContext = context.WithValue(requestContext, UserDataKey, options.UserData)
	} else {
		validationInput.Options = openapi3filter.DefaultOptions
	}

	// Validate request
	operationParameters := validationInput.Route.Operation.Parameters
	pathItemParameters := route.PathItem.Parameters

	for _, parameterRef := range pathItemParameters {
		parameter := parameterRef.Value
		if operationParameters != nil {
			if override := operationParameters.GetByInAndName(parameter.In, parameter.Name); override != nil {
				continue
			}
			if err := openapi3filter.ValidateParameter(requestContext, validationInput, parameter); err != nil {
				return err
			}
		}
	}

	for _, parameter := range operationParameters {
		if err := openapi3filter.ValidateParameter(requestContext, validationInput, parameter.Value); err != nil {
			return err
		}
	}

	// Validate request body
	requestBody := validationInput.Route.Operation.RequestBody
	if requestBody != nil && !validationInput.Options.ExcludeRequestBody {
		compositeError := NewCompositeError(CompositeErrorTypeValidation)

		bb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return &openapi3filter.RequestError{
				Input:       validationInput,
				RequestBody: validationInput.Route.Operation.RequestBody.Value,
				Reason:      "request body reading failed",
				Err:         err,
			}
		}
		_ = r.Body.Close()

		if len(bb) == 0 {
			if requestBody.Value.Required {
				return &openapi3filter.RequestError{
					Input:       validationInput,
					RequestBody: requestBody.Value,
					Err:         openapi3filter.ErrInvalidRequired,
				}
			}

			return nil
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(bb))

		var value map[string]interface{}
		if err = json.Unmarshal(bb, &value); err != nil {
			return &openapi3filter.RequestError{
				Input:       validationInput,
				RequestBody: validationInput.Route.Operation.RequestBody.Value,
				Reason:      "json unmarshaling failed",
				Err:         err,
			}
		}

		if r.Body != nil {
			_ = r.Body.Close()
		}

		inputMIME := r.Header.Get("Content-Type")
		contentType := validationInput.Route.Operation.RequestBody.Value.Content.Get(inputMIME)
		if contentType == nil {
			return &openapi3filter.RequestError{
				Input:       validationInput,
				RequestBody: validationInput.Route.Operation.RequestBody.Value,
				Reason:      fmt.Sprintf("header 'Content-Type' has unexpected value: %q", inputMIME),
			}
		}

		if contentType.Schema == nil {
			return nil
		}

		refs := make(map[string]*openapi3.SchemaRef)
		dereferenceSchema(contentType.Schema, &refs)

		for key, val := range value {
			if refs[key] == nil || refs[key].Value == nil {
				continue
			}

			visitSchema(key, refs[key], val, compositeError)
		}

		if len(compositeError.Errors()) > 0 {
			return compositeError
		}
	}

	// Security
	security := validationInput.Route.Operation.Security

	if security == nil {
		if route.Swagger == nil {
			return errRouteMissingSwagger
		} else {
			security = &route.Swagger.Security
		}
	}

	if security != nil {
		if err := openapi3filter.ValidateSecurityRequirements(requestContext, validationInput, *security); err != nil {
			return err
		}
	}

	return nil
}

func visitSchema(key string, ref *openapi3.SchemaRef, value interface{}, compositeError *CompositeError) {
	visit := func(schema *openapi3.Schema, val interface{}) bool {
		if err := schema.VisitJSON(val); err != nil {
			origin, ok := err.(*openapi3.SchemaError)
			if ok {
				compositeError.add(&SchemaError{
					Key:    key,
					Origin: origin,
				})

				return false
			}
		}

		return true
	}

	if len(ref.Value.Properties) > 0 {
		valMap, ok := value.(map[string]interface{})
		if !ok {
			compositeError.add(&SchemaError{
				Key: key,
				Origin: &openapi3.SchemaError{
					SchemaField: "type",
					Schema: &openapi3.Schema{
						Type: "object",
					},
				},
			})
		} else {
			for childKey, childProperty := range ref.Value.Properties {
				if _, ok := valMap[childKey]; !ok {
					compositeError.add(&SchemaError{
						Key: key + "." + childKey,
						Origin: &openapi3.SchemaError{
							SchemaField: "required",
						},
					})
				} else {
					visitSchema(key+"."+childKey, childProperty, valMap[childKey], compositeError)
				}
			}
		}
	} else {
		visit(ref.Value, value)
	}
}

func dereferenceSchema(ref *openapi3.SchemaRef, refs *map[string]*openapi3.SchemaRef) {
	for _, item := range ref.Value.OneOf {
		dereferenceSchema(item, refs)
	}

	for _, item := range ref.Value.AnyOf {
		dereferenceSchema(item, refs)
	}

	for _, item := range ref.Value.AllOf {
		dereferenceSchema(item, refs)
	}

	if ref.Value.Items != nil {
		dereferenceSchema(ref.Value.Items, refs)
	}

	if ref.Value.AdditionalProperties != nil {
		dereferenceSchema(ref.Value.AdditionalProperties, refs)
	}

	for name, property := range ref.Value.Properties {
		(*refs)[name] = property

		if property.Ref != "" {
			dereferenceSchema(property, refs)
		}
	}
}
