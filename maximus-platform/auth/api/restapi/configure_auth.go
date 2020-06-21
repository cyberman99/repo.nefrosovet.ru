// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/auth/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/admin"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/auth"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/backend"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/client"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/event"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/j_w_k"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/role"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/token"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
)

//go:generate swagger generate server --target ../../api --name auth --spec ../../docs/swagger.yaml --exclude-main

// Version of service
var Version string

func configureFlags(api *operations.AuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AuthAPI) http.Handler {
	handlers.Version = Version

	api.ServeError = handlers.ServeError
	api.Logger = log.Printf
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization: Bearer" header or the "access_token" query is set
	api.BearerAuth = func(token string) (interface{}, error) {
		// The header: Authorization: Bearer {base64 string} (or ?access_token={base 64 string} param) has already
		// been decoded by the runtime as a token
		// api.Logger("HasRoleAuth handler called")

		for _, prefix := range []string{"bearer ", "Bearer "} {
			if strings.HasPrefix(token, prefix) {
				token = strings.TrimPrefix(token, prefix)

				break
			}
		}

		_, claims, err := jwt.ParseToken(token)

		if err != nil {
			return nil, err
		}

		// Only access tokens is valid
		if claims["type"] != "access" {
			err = errors.New(401, "Can't use refresh token for access")
		}

		return claims, nil
	}

	// Roles
	api.RoleGetRolesHandler = role.GetRolesHandlerFunc(handlers.GetRoles)
	api.RolePostRolesHandler = role.PostRolesHandlerFunc(handlers.PostRoles)
	api.RoleGetRolesRoleIDHandler = role.GetRolesRoleIDHandlerFunc(handlers.GetRolesRoleID)
	api.RolePutRolesRoleIDHandler = role.PutRolesRoleIDHandlerFunc(handlers.PutRolesRoleID)
	api.RoleDeleteRolesRoleIDHandler = role.DeleteRolesRoleIDHandlerFunc(handlers.DeleteRolesRoleID)
	api.RoleGetRolesRoleIDUsersHandler = role.GetRolesRoleIDUsersHandlerFunc(handlers.GetRolesRoleIDUsers)
	api.RolePostRolesRoleIDUsersUserIDHandler = role.PostRolesRoleIDUsersUserIDHandlerFunc(handlers.PostRolesRoleIDUsersUserID)
	api.RoleDeleteRolesRoleIDUsersUserIDHandler = role.DeleteRolesRoleIDUsersUserIDHandlerFunc(handlers.DeleteRolesRoleIDUsersUserID)

	// Token
	api.TokenGetWhoamiHandler = token.GetWhoamiHandlerFunc(handlers.GetWhoami)

	// JWK
	api.JWKGetJwkHandler = j_w_k.GetJwkHandlerFunc(handlers.GetJwk)

	// Admin
	api.AdminPostAdminHandler = admin.PostAdminHandlerFunc(handlers.PostAdmin)

	// Auth
	api.AuthPostClientHandler = auth.PostClientHandlerFunc(handlers.PostAuthClient)
	api.AuthPostUserHandler = auth.PostUserHandlerFunc(handlers.PostUser)
	api.AuthPostRefreshHandler = auth.PostRefreshHandlerFunc(handlers.PostRefresh)
	api.AuthPostIdentifyHandler = auth.PostIdentifyHandlerFunc(handlers.PostIdentify)
	api.AuthGetOauth2BackendIDHandler = auth.GetOauth2BackendIDHandlerFunc(handlers.AuthGetOAuth2BackendID)
	api.AuthPostOauth2BackendIDConsentHandler = auth.PostOauth2BackendIDConsentHandlerFunc(handlers.AuthPostOAuth2Consent)

	// Event
	api.EventGetEventsHandler = event.GetEventsHandlerFunc(handlers.GetEvents)
	api.EventGetEventsEventIDHandler = event.GetEventsEventIDHandlerFunc(handlers.GetEventsEventID)

	// Client
	api.ClientGetClientsHandler = client.GetClientsHandlerFunc(handlers.GetClients)
	api.ClientGetClientsClientIDHandler = client.GetClientsClientIDHandlerFunc(handlers.GetClientsClientID)
	api.ClientPostClientsHandler = client.PostClientsHandlerFunc(handlers.PostClients)
	api.ClientPutClientsClientIDHandler = client.PutClientsClientIDHandlerFunc(handlers.PutClientsClientID)
	api.ClientPatchClientsClientIDHandler = client.PatchClientsClientIDHandlerFunc(handlers.PatchClientsClientID)
	api.ClientDeleteClientsClientIDHandler = client.DeleteClientsClientIDHandlerFunc(handlers.DeleteClientsClientID)

	// Backend
	api.BackendGetBackendsHandler = backend.GetBackendsHandlerFunc(handlers.GetBackends)
	api.BackendGetBackendsBackendIDHandler = backend.GetBackendsBackendIDHandlerFunc(handlers.GetBackendsBackendID)
	api.BackendPostBackendsBackendIDGroupsHandler = backend.PostBackendsBackendIDGroupsHandlerFunc(handlers.PostBackendsBackendIDGroups)
	api.BackendDeleteBackendsBackendIDGroupsHandler = backend.DeleteBackendsBackendIDGroupsHandlerFunc(handlers.DeleteBackendsBackendIDGroups)
	api.BackendGetBackendsBackendIDGroupsHandler = backend.GetBackendsBackendIDGroupsHandlerFunc(handlers.GetBackendsBackendIDGroups)
	api.BackendPostBackendsLdapHandler = backend.PostBackendsLdapHandlerFunc(handlers.PostBackendsLDAP)
	api.BackendPutBackendsLdapBackendIDHandler = backend.PutBackendsLdapBackendIDHandlerFunc(handlers.PutBackendsLdapBackendID)
	api.BackendPatchBackendsLdapBackendIDHandler = backend.PatchBackendsLdapBackendIDHandlerFunc(handlers.PatchBackendsLdapBackendID)
	api.BackendPostBackendsOauth2Handler = backend.PostBackendsOauth2HandlerFunc(handlers.PostBackendsOAuth2)
	api.BackendPutBackendsOauth2BackendIDHandler = backend.PutBackendsOauth2BackendIDHandlerFunc(handlers.PutBackendsOAuth2BackendID)
	api.BackendPatchBackendsOauth2BackendIDHandler = backend.PatchBackendsOauth2BackendIDHandlerFunc(handlers.PatchBackendsOAuth2BackendID)
	api.BackendDeleteBackendsBackendIDHandler = backend.DeleteBackendsBackendIDHandlerFunc(handlers.DeleteBackendsBackendID)
	api.BackendGetBackendsBackendIDTestHandler = backend.GetBackendsBackendIDTestHandlerFunc(handlers.GetBackendsBackendIDTest)
	api.BackendGetFlowHandler = backend.GetFlowHandlerFunc(handlers.GetFlow)
	api.BackendPostFlowHandler = backend.PostFlowHandlerFunc(handlers.PostFlow)

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
	if viper.GetString("sentryDSN") != "" {
		return handlers.NewRavenHandler(handler)
	}

	return handler
}
