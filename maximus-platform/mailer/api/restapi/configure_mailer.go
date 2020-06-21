// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
    "crypto/tls"
    "fmt"
    "net/http"

    log "github.com/Sirupsen/logrus"
    "github.com/getsentry/raven-go"
    "github.com/go-openapi/runtime"
    "github.com/spf13/viper"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/telegram"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/viber"

    "repo.nefrosovet.ru/maximus-platform/mailer/api/handlers"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/channels"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/manage"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/messages"
    "repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/service"
)

//go:generate swagger generate server --target ../../api --name mailer --spec ../../docs/swagger.yaml

func configureFlags(_ *operations.MailerAPI) {
    // api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MailerAPI) http.Handler {
    // configure the api here
    // api.ServeError = errors.ServeError
    api.ServeError = handlers.ServeError

    // Set your custom logger if needed. Default one is log.Printf
    // Expected interface func(string, ...interface{})
    api.Logger = log.Printf

    // Start telegram pool
    sender.TgPool = telegram.NewPool()
    if err := sender.TgPool.AddFromStorage(handlers.GetStorage().ChannelStorage); err != nil {
        log.WithFields(log.Fields{
            "context": "TGPOOL",
            "action":  "AddFromStorage",
        }).Fatal(err)
    }
    sender.TgPool.Start()

    // Start viber pool
    sender.ViberPool = viber.NewPool()
    if err := sender.ViberPool.AddFromStorage(handlers.GetStorage().ChannelStorage); err != nil {
        log.WithFields(log.Fields{
            "context": "ViberPool",
            "action":  "AddFromStorage",
        }).Fatal(err)
    }
    sender.ViberPool.Start()

    api.JSONConsumer = runtime.JSONConsumer()

    api.JSONProducer = runtime.JSONProducer()

    /*
    	Channels routes
    */

    api.ChannelsDeleteChannelsChannelIDHandler = channels.DeleteChannelsChannelIDHandlerFunc(handlers.DeleteChannelsChannelIDHandler)

    api.ChannelsGetChannelsHandler = channels.GetChannelsHandlerFunc(handlers.GetChannelsHandler)
    api.ChannelsGetChannelsChannelIDHandler = channels.GetChannelsChannelIDHandlerFunc(handlers.GetChannelsChannelIDHandler)

    api.ChannelsPostChannelsEmailHandler = channels.PostChannelsEmailHandlerFunc(handlers.PostChannelsEmailHandler)
    api.ChannelsPostChannelsLocalSmsHandler = channels.PostChannelsLocalSmsHandlerFunc(handlers.PostChannelsLocalSmsHandler)
    api.ChannelsPostChannelsMtsSmsHandler = channels.PostChannelsMtsSmsHandlerFunc(handlers.PostChannelsMtsSmsHandler)
    api.ChannelsPostChannelsSLACKHandler = channels.PostChannelsSLACKHandlerFunc(handlers.PostChannelsSLACKHandler)
    api.ChannelsPostChannelsTelegramHandler = channels.PostChannelsTelegramHandlerFunc(handlers.PostChannelsTelegramHandler)
    api.ChannelsPostChannelsViberHandler = channels.PostChannelsViberHandlerFunc(handlers.PostChannelsViberHandler)

    api.ChannelsPutChannelsEmailChannelIDHandler = channels.PutChannelsEmailChannelIDHandlerFunc(handlers.PutChannelsEmailChannelIDHandler)
    api.ChannelsPutChannelsLocalSmsChannelIDHandler = channels.PutChannelsLocalSmsChannelIDHandlerFunc(handlers.PutChannelsLocalSmsChannelIDHandler)
    api.ChannelsPutChannelsMtsSmsChannelIDHandler = channels.PutChannelsMtsSmsChannelIDHandlerFunc(handlers.PutChannelsMtsSmsChannelIDHandler)
    api.ChannelsPutChannelsSLACKChannelIDHandler = channels.PutChannelsSLACKChannelIDHandlerFunc(handlers.PutChannelsSLACKChannelIDHandler)
    api.ChannelsPutChannelsTelegramChannelIDHandler = channels.PutChannelsTelegramChannelIDHandlerFunc(handlers.PutChannelsTelegramChannelIDHandler)
    api.ChannelsPutChannelsViberChannelIDHandler = channels.PutChannelsViberChannelIDHandlerFunc(handlers.PutChannelsViberChannelIDHandler)

    /*
    	Messages routes
    */

    api.MessagesPostSendHandler = messages.PostSendHandlerFunc(handlers.PostSendHandler)
    api.MessagesGetMessagesHandler = messages.GetMessagesHandlerFunc(handlers.GetMessagesHandler)
    api.MessagesGetMessagesMessageIDHandler = messages.GetMessagesMessageIDHandlerFunc(handlers.GetMessagesMessageIDHandler)

    /*
    	Tokens routes
    */

    api.ManageGetTokensHandler = manage.GetTokensHandlerFunc(handlers.GetTokensHandler)
    api.ManagePostTokensHandler = manage.PostTokensHandlerFunc(handlers.PostTokensHandler)
    api.ManageGetTokensTokenIDHandler = manage.GetTokensTokenIDHandlerFunc(handlers.GetTokensTokenIDHandler)
    api.ManagePutTokensTokenIDHandler = manage.PutTokensTokenIDHandlerFunc(handlers.PutTokensTokenIDHandler)
    api.ManageDeleteTokensTokenIDHandler = manage.DeleteTokensTokenIDHandlerFunc(handlers.DeleteTokensTokenIDHandler)

    /*
    	Service routes
    */

    api.ServiceGetStatusHandler = service.GetStatusHandlerFunc(handlers.GetStatusHandler)

    /*
    	Server shutdown tasks
    */
    api.ServerShutdown = func() {
        // Stop telegram pool
        sender.TgPool.Stop()
        // Stop viber pool
        sender.ViberPool.Stop()
    }

    return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(_ *tls.Config) {
    // Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(_ *http.Server, _, _ string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
    return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
    if viper.GetString("sentryDSN") != "" {
        return NewRavenHandler(handler)
    }
    return handler
}

// RavenHandler is http.Handler with raven control
type RavenHandler struct {
    o http.Handler
}

// NewRavenHandler returns RavenHandler
func NewRavenHandler(h http.Handler) *RavenHandler {
    return &RavenHandler{
        o: h,
    }
}

// ServeHTTP realizes http.Handler interface
func (rh *RavenHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    defer func() {
        if rval := recover(); rval != nil {
            rvalStr := fmt.Sprint(rval)
            packet := raven.NewPacket(rvalStr, raven.NewException(fmt.Errorf("%s", rvalStr), raven.GetOrNewStacktrace(rval.(error), 2, 3, nil)), raven.NewHttp(r))
            raven.Capture(packet, nil)
        }
    }()

    rh.o.ServeHTTP(rw, r)
}
