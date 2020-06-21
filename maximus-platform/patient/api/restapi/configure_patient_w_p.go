// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/access"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/appointment"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/employee"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/profile"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/program"
)

//go:generate swagger generate server --target ../../api --name PatientWP --spec ../../docs/swagger.yaml --exclude-main

func configureFlags(api *operations.PatientWPAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PatientWPAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AppointmentAppointmentCollectionHandler == nil {
		api.AppointmentAppointmentCollectionHandler = appointment.AppointmentCollectionHandlerFunc(func(params appointment.AppointmentCollectionParams) middleware.Responder {
			return middleware.NotImplemented("operation appointment.AppointmentCollection has not yet been implemented")
		})
	}
	if api.AppointmentAppointmentParamsCollectionHandler == nil {
		api.AppointmentAppointmentParamsCollectionHandler = appointment.AppointmentParamsCollectionHandlerFunc(func(params appointment.AppointmentParamsCollectionParams) middleware.Responder {
			return middleware.NotImplemented("operation appointment.AppointmentParamsCollection has not yet been implemented")
		})
	}
	if api.ProgramAppointmentProgramCollectionHandler == nil {
		api.ProgramAppointmentProgramCollectionHandler = program.AppointmentProgramCollectionHandlerFunc(func(params program.AppointmentProgramCollectionParams) middleware.Responder {
			return middleware.NotImplemented("operation program.AppointmentProgramCollection has not yet been implemented")
		})
	}
	if api.ProgramAppointmentProgramViewHandler == nil {
		api.ProgramAppointmentProgramViewHandler = program.AppointmentProgramViewHandlerFunc(func(params program.AppointmentProgramViewParams) middleware.Responder {
			return middleware.NotImplemented("operation program.AppointmentProgramView has not yet been implemented")
		})
	}
	if api.AppointmentAppointmentViewHandler == nil {
		api.AppointmentAppointmentViewHandler = appointment.AppointmentViewHandlerFunc(func(params appointment.AppointmentViewParams) middleware.Responder {
			return middleware.NotImplemented("operation appointment.AppointmentView has not yet been implemented")
		})
	}
	if api.AccessCodeConfirmationHandler == nil {
		api.AccessCodeConfirmationHandler = access.CodeConfirmationHandlerFunc(func(params access.CodeConfirmationParams) middleware.Responder {
			return middleware.NotImplemented("operation access.CodeConfirmation has not yet been implemented")
		})
	}
	if api.EmployeeEmployeeViewHandler == nil {
		api.EmployeeEmployeeViewHandler = employee.EmployeeViewHandlerFunc(func(params employee.EmployeeViewParams) middleware.Responder {
			return middleware.NotImplemented("operation employee.EmployeeView has not yet been implemented")
		})
	}
	if api.AccessPasswordRecoveryHandler == nil {
		api.AccessPasswordRecoveryHandler = access.PasswordRecoveryHandlerFunc(func(params access.PasswordRecoveryParams) middleware.Responder {
			return middleware.NotImplemented("operation access.PasswordRecovery has not yet been implemented")
		})
	}
	if api.ProfilePatientContactUpdateHandler == nil {
		api.ProfilePatientContactUpdateHandler = profile.PatientContactUpdateHandlerFunc(func(params profile.PatientContactUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.PatientContactUpdate has not yet been implemented")
		})
	}
	if api.ProfilePatientProfileUpdateHandler == nil {
		api.ProfilePatientProfileUpdateHandler = profile.PatientProfileUpdateHandlerFunc(func(params profile.PatientProfileUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.PatientProfileUpdate has not yet been implemented")
		})
	}
	if api.ProfilePatientProfileViewHandler == nil {
		api.ProfilePatientProfileViewHandler = profile.PatientProfileViewHandlerFunc(func(params profile.PatientProfileViewParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.PatientProfileView has not yet been implemented")
		})
	}
	if api.AccessPatientRegistrationHandler == nil {
		api.AccessPatientRegistrationHandler = access.PatientRegistrationHandlerFunc(func(params access.PatientRegistrationParams) middleware.Responder {
			return middleware.NotImplemented("operation access.PatientRegistration has not yet been implemented")
		})
	}
	if api.ProfilePatientcontactViewHandler == nil {
		api.ProfilePatientcontactViewHandler = profile.PatientcontactViewHandlerFunc(func(params profile.PatientcontactViewParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.PatientcontactView has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
