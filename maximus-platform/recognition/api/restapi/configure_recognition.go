// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"repo.nefrosovet.ru/maximus-platform/recognition/services"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/photo"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/recognize"
)

//go:generate swagger generate server --target ../../api --name Recognition --spec ../../docs/swagger.yaml --exclude-main

func configureFlags(api *operations.RecognitionAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}
func configureAPI(api *operations.RecognitionAPI) http.Handler { return nil } // do nothing. Will be overrided

func ConfigureAPI(
	api *operations.RecognitionAPI,
	l logger.APIEntrier,
	servicer services.CloudServicer,
	version string,
) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = l.Infof
	api.ServeError = handlers.NewErrorHandler(version).ServeError

	var rekCtrl *handlers.RekognitionController
	rekCtrl = handlers.NewRekognitionController(servicer, l, version)

	var photoCtrl *handlers.PhotoController
	photoCtrl = handlers.NewPhotoController(servicer, l, version)

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.PhotoDeleteHandler = photo.DeleteHandlerFunc(photoCtrl.Delete)

	api.PhotoCollectionHandler = photo.CollectionHandlerFunc(photoCtrl.List)

	api.PhotoViewHandler = photo.ViewHandlerFunc(photoCtrl.Get)
	api.PhotoCreateHandler = photo.CreateHandlerFunc(photoCtrl.Post)

	api.RecognizeRecognizeHandler = recognize.RecognizeHandlerFunc(rekCtrl.Rekognize)

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
