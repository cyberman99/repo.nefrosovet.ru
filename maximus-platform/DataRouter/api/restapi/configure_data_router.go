// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/events"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"

	"github.com/go-openapi/runtime"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/clients"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/permissions"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/replies"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/routes"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations/status"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
)

//go:generate swagger generate server --target ../../api --name DataRouter --spec ../../docs/swagger.yaml --exclude-main

var (
	Version string
	Lg      logger.Logger
)

func configureFlags(api *operations.DataRouterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DataRouterAPI) http.Handler {
	// configure the api here
	api.ServeError = handlers.ServeDataRouterError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:

	api.Logger = Lg.Printf
	handlers.Version = Version

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	var storage mongod.Storer
	storage = mongod.GetStorage()

	api.StatusStatusViewHandler = status.StatusViewHandlerFunc(handlers.NewStatusView(storage).Get)

	var cliRepo repos.ClientRepository
	cliRepo = repos.NewClientRepo(storage, nil)

	var routeRepo repos.RouteRepository
	routeRepo = repos.NewRouteRepo(storage)

	var replyRepo repos.ReplyRepository
	replyRepo = repos.NewReplyRepo(storage)

	var (
		eventsRepo influx.EventRepository
		inflx      influx.Influxer
	)
	inflx = influx.GetInfluxer()
	eventsRepo = influx.NewEventRepo(inflx.DBName(), inflx)

	var cliCtrl *handlers.ClientController
	cliCtrl = handlers.NewClientController(cliRepo, Lg)

	api.ClientsClientCollectionHandler = clients.ClientCollectionHandlerFunc(cliCtrl.List)
	api.ClientsClientCreateHandler = clients.ClientCreateHandlerFunc(cliCtrl.Post)
	api.ClientsClientDeleteHandler = clients.ClientDeleteHandlerFunc(cliCtrl.Delete)
	api.ClientsClientPatchHandler = clients.ClientPatchHandlerFunc(cliCtrl.Patch)
	api.ClientsClientViewHandler = clients.ClientViewHandlerFunc(cliCtrl.Get)

	var permCtrl *handlers.ClientPermissionsController
	permCtrl = handlers.NewClientPermissionsController(cliRepo, Lg)

	api.PermissionsClientPermissionViewHandler =
		permissions.ClientPermissionViewHandlerFunc(permCtrl.Get)
	api.PermissionsClientPermissionCreateHandler =
		permissions.ClientPermissionCreateHandlerFunc(permCtrl.Post)

	var routeCtrl *handlers.RouteController
	routeCtrl = handlers.NewRouteController(routeRepo, replyRepo, Lg)

	api.RoutesRouteCollectionHandler = routes.RouteCollectionHandlerFunc(routeCtrl.List)
	api.RoutesRouteCreateHandler = routes.RouteCreateHandlerFunc(routeCtrl.Post)
	api.RoutesRouteDeleteHandler = routes.RouteDeleteHandlerFunc(routeCtrl.Delete)
	api.RoutesRoutePutHandler = routes.RoutePutHandlerFunc(routeCtrl.Put)
	api.RoutesRouteViewHandler = routes.RouteViewHandlerFunc(routeCtrl.Get)

	var replyCtrl *handlers.ReplyController
	replyCtrl = handlers.NewReplyController(replyRepo, routeRepo, Lg)

	api.RepliesReplyCollectionHandler = replies.ReplyCollectionHandlerFunc(replyCtrl.List)
	api.RepliesReplyCreateHandler = replies.ReplyCreateHandlerFunc(replyCtrl.Post)
	api.RepliesReplyDeleteHandler = replies.ReplyDeleteHandlerFunc(replyCtrl.Delete)
	api.RepliesReplyPutHandler = replies.ReplyPutHandlerFunc(replyCtrl.Put)
	api.RepliesReplyViewHandler = replies.ReplyViewHandlerFunc(replyCtrl.Get)

	var eventsCtrl *handlers.EventsController
	eventsCtrl = handlers.NewEventsController(eventsRepo, Lg)

	api.EventsEventCollectionHandler = events.EventCollectionHandlerFunc(eventsCtrl.List)
	api.EventsEventViewHandler = events.EventViewHandlerFunc(eventsCtrl.Get)

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
