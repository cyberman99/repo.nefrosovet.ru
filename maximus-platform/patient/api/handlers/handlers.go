package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	goswaggererrors "repo.nefrosovet.ru/libs/goswagger-errors"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/access"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/appointment"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/employee"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/profile"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations/program"
	profileClient "repo.nefrosovet.ru/maximus-platform/patient/client/profile"
	"repo.nefrosovet.ru/maximus-platform/patient/db"
	"repo.nefrosovet.ru/maximus-platform/patient/db/sqlc"
	"repo.nefrosovet.ru/maximus-platform/patient/logger"
)

var (
	EntityNotFoundMessage = "Entity Not Found"
	SuccessMessage        = "SUCCESS"
	InternalServerError = "Internal server error"
)

func ConfigureAPI(
	api *operations.PatientWPAPI,
	l logger.APIEntrier,
	db db.Storer,
	profileClient *profileClient.ClientWithResponses,
	version string,
	queries *sqlc.Queries,
) http.Handler {
	errHandler := goswaggererrors.New(version)

	// TODO: error handler sentry support (https://repo.nefrosovet.ru/libs/goswagger-errors/issues/1)
	// errHandler = errHandler.WithSentry(sentryClient)

	api.ServeError = errHandler.Serve
	api.Logger = l.Infof

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AppointmentAppointmentCollectionHandler = appointment.AppointmentCollectionHandlerFunc(func(params appointment.AppointmentCollectionParams) middleware.Responder {
		// look at service/service.go

		// colls, err := service.Appointment().GetAll(params.Limit, param.Offset)
		// if err != nil {return handleErr(err)}
		// return handleResponse(colls)
		return middleware.NotImplemented("operation appointment.AppointmentCollection has not yet been implemented")
	})
	api.AppointmentAppointmentParamsCollectionHandler = appointment.AppointmentParamsCollectionHandlerFunc(func(params appointment.AppointmentParamsCollectionParams) middleware.Responder {
		return middleware.NotImplemented("operation appointment.AppointmentParamsCollection has not yet been implemented")
	})
	api.ProgramAppointmentProgramCollectionHandler = program.AppointmentProgramCollectionHandlerFunc(func(params program.AppointmentProgramCollectionParams) middleware.Responder {
		return middleware.NotImplemented("operation program.AppointmentProgramCollection has not yet been implemented")
	})
	api.ProgramAppointmentProgramViewHandler = program.AppointmentProgramViewHandlerFunc(func(params program.AppointmentProgramViewParams) middleware.Responder {
		return middleware.NotImplemented("operation program.AppointmentProgramView has not yet been implemented")
	})
	api.AppointmentAppointmentViewHandler = appointment.AppointmentViewHandlerFunc(func(params appointment.AppointmentViewParams) middleware.Responder {
		return middleware.NotImplemented("operation appointment.AppointmentView has not yet been implemented")
	})
	api.AccessCodeConfirmationHandler = access.CodeConfirmationHandlerFunc(func(params access.CodeConfirmationParams) middleware.Responder {
		return middleware.NotImplemented("operation access.CodeConfirmation has not yet been implemented")
	})

	employeeViewHandler := NewEmployeeView(version, queries, l)
	api.EmployeeEmployeeViewHandler = employee.EmployeeViewHandlerFunc(employeeViewHandler.Get)
	api.AccessPasswordRecoveryHandler = access.PasswordRecoveryHandlerFunc(func(params access.PasswordRecoveryParams) middleware.Responder {
		return middleware.NotImplemented("operation access.PasswordRecovery has not yet been implemented")
	})
	api.ProfilePatientContactUpdateHandler = profile.PatientContactUpdateHandlerFunc(func(params profile.PatientContactUpdateParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.PatientContactUpdate has not yet been implemented")
	})
	api.ProfilePatientProfileUpdateHandler = profile.PatientProfileUpdateHandlerFunc(func(params profile.PatientProfileUpdateParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.PatientProfileUpdate has not yet been implemented")
	})
	api.ProfilePatientProfileViewHandler = profile.PatientProfileViewHandlerFunc(func(params profile.PatientProfileViewParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.PatientProfileView has not yet been implemented")
	})
	api.AccessPatientRegistrationHandler = access.PatientRegistrationHandlerFunc(func(params access.PatientRegistrationParams) middleware.Responder {
		return middleware.NotImplemented("operation access.PatientRegistration has not yet been implemented")
	})
	api.ProfilePatientcontactViewHandler = profile.PatientcontactViewHandlerFunc(func(params profile.PatientcontactViewParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.PatientcontactView has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(nil), errHandler)
}

func setupGlobalMiddleware(handler http.Handler, errHandler *goswaggererrors.Handler) http.Handler {
	if errHandler != nil {
		return errHandler.NewRecoverMiddleware(handler)
	}

	return handler
}
