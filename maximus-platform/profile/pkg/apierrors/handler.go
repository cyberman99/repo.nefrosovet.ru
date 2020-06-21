package apierrors

import (
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/profile/pkg/middleware"
)

type Handler struct {
	version string

	sentryHub *sentry.Hub
}

func NewHandler(version string) *Handler {
	return &Handler{
		version: version,
	}
}

func (h *Handler) WithSentry(hub *sentry.Hub) *Handler {
	return &Handler{
		version: h.version,

		sentryHub: hub,
	}
}

func (h *Handler) Handle(err error, ctx echo.Context) {
	code := 500

	if requestError, ok := err.(*openapi3filter.RequestError); ok {
		code = requestError.HTTPStatus()

		// Get original error
		err = requestError.Err
	}

	switch err := err.(type) {
	case *echo.HTTPError:
		code = err.Code
	case *openapi3filter.RouteError:
		response := NotFoundResponse{
			Response: Response{
				Version: h.version,
				Message: UnknownPathMessage,
			},
			Errors: []interface{}{
				err.Error(),
			},
		}

		_ = ctx.JSON(http.StatusNotFound, response)
		return
	case *middleware.CompositeError:
		if err.Type != middleware.CompositeErrorTypeValidation {
			break
		}

		var core []string
		validation := make(map[string]string)
		for _, err := range err.Errors() {
			switch err := err.(type) {
			case *middleware.SchemaError:
				core = append(core, err.Origin.Reason)

				switch err.Origin.SchemaField {
				case "minLength":
					validation[err.Key] = "min"
				case "maxLength":
					validation[err.Key] = "max"
				case "enum":
					validation[err.Key] = "enum"
				case "type":
					validation[err.Key] = err.Origin.Schema.Type
				case "format":
					validation[err.Key] = "format"
				}
			}
		}

		response := ValidationErrorResponse{
			Response: Response{
				Version: h.version,
				Message: InternalServerErrorMessage,
			},
			Errors: &ValidationError{
				Core:       strings.Join(core, "\n"),
				Validation: validation,
			},
		}

		_ = ctx.JSON(http.StatusBadRequest, response)
		return
	}

	switch code {
	case http.StatusBadRequest:
		response := ValidationErrorResponse{
			Response: Response{
				Version: h.version,
				Message: InternalServerErrorMessage,
			},
			Errors: &ValidationError{},
		}

		if err.Error() == "EOF" {
			response.Errors.JSON = "EOF"
		} else {
			response.Errors.Core = err.Error()
		}

		_ = ctx.JSON(http.StatusBadRequest, response)
	case http.StatusInternalServerError:
		h.capture(ctx.Request(), err)

		response := InternalServerErrorResponse{
			Response: Response{
				Version: h.version,
				Message: InternalServerErrorMessage,
			},
		}

		if err != nil {
			response.Errors = []interface{}{
				err.Error(),
			}
		}

		_ = ctx.JSON(http.StatusInternalServerError, response)
	}

	// fmt.Printf("Error: (%T) %v \n", err, err)
}

func (h *Handler) capture(r *http.Request, err error) {
	if h.sentryHub == nil {
		return
	}

	stacktrace := sentry.NewStacktrace()

	if len(stacktrace.Frames) > 4 {
		stacktrace.Frames = stacktrace.Frames[:len(stacktrace.Frames)-4]
	}

	event := sentry.Event{
		User:    sentry.User{},
		Request: sentryRequestFromHTTP(r),
		Exception: []sentry.Exception{{
			Type:       fmt.Sprintf("%T", err),
			Value:      err.Error(),
			Stacktrace: stacktrace,
		}},
	}

	eventID := h.sentryHub.CaptureEvent(&event)

	log.WithError(err).
		WithFields(log.Fields{
			"sentryEventID": eventID,
		}).
		Error("Panic captured!")

	if log.StandardLogger().GetLevel() == log.DebugLevel {
		debug.PrintStack()
	}
}

func sentryRequestFromHTTP(r *http.Request) sentry.Request {
	proto := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		proto = "https"
	}

	sentryRequest := sentry.Request{
		URL:         proto + "://" + r.Host + r.URL.Path,
		Method:      r.Method,
		QueryString: r.RequestURI,
		Cookies:     r.Header.Get("Cookie"),
		Headers:     map[string]string{},
	}

	for k, v := range r.Header {
		sentryRequest.Headers[k] = strings.Join(v, ",")
	}

	if addr, port, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		sentryRequest.Env = map[string]string{
			"REMOTE_ADDR": addr,
			"REMOTE_PORT": port,
		}
	}

	return sentryRequest
}
