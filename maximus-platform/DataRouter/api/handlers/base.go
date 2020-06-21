package handlers

// PayloadSuccessMessage used in 200 answers
var success = "SUCCESS"
var PayloadSuccessMessage = &success

// PayloadFailMessage - common fail message
var fail = "FAIL"
var PayloadFailMessage = &fail

// PayloadSuccessMessage used in 500 answers
var serverError = "Internal server error"
var InternalServerErrorMessage = &serverError

// PayloadValidationErrorMessage - used on composite errors
var validError = "Validation error"
var PayloadValidationErrorMessage = &validError

// NotFoundMessage used in 404 answers
var notFound = "Entity not found"
var NotFoundMessage = &notFound

// Version of service
var Version string
