package goswagger_errors

type Response struct {
	Version *string `json:"version"`

	Data interface{} `json:"data"`
}

// Internal server error (500)

type InternalServerErrorResponse struct {
	Response

	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
}

// Validation error (400)

type ValidationError struct {
	Core string `json:"core,omitempty"`
	JSON string `json:"json,omitempty"`

	Validation map[string]interface{} `json:"validation"`
}

type ValidationErrorResponse struct {
	Response

	Errors  *ValidationError `json:"errors,omitempty"`
	Message string           `json:"message,omitempty"`
}

// Method not allowed (405)

type MethodNotAllowedResponse struct {
	Response

	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Not found (404)

type NotFoundResponse struct {
	Response

	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}
