package goswagger_errors

import (
	"net/http"
	"strings"

	"github.com/getsentry/raven-go"

	oapiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

type Handler struct {
	version string

	ravenClient *raven.Client
}

func New(version string) *Handler {
	return &Handler{
		version: version,
	}
}

func (s *Handler) WithRaven(ravenClient *raven.Client) *Handler {
	return &Handler{
		version: s.version,

		ravenClient: ravenClient,
	}
}

func (s *Handler) Serve(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch e := err.(type) {
	case *oapiErrors.CompositeError:
		response := ValidationErrorResponse{
			Response: Response{
				Version: &s.version,
			},
			Errors:  parseComposite(e),
			Message: ValidationErrorMessage,
		}

		w.WriteHeader(400)
		if err := runtime.JSONProducer().Produce(w, response); err != nil {
			s.internalServerError(w, r, err)
		}
	case *oapiErrors.MethodNotAllowedError:
		w.Header().Add("Allow", strings.Join(err.(*oapiErrors.MethodNotAllowedError).Allowed, ","))

		if r == nil || r.Method != "HEAD" {
			response := MethodNotAllowedResponse{
				Response: Response{
					Version: &s.version,
				},
			}

			switch r {
			case nil:
				response.Message = MethodNotAllowedMessage
			default:
				response.Message = WrapMethodNotAllowedMessage(r.Method)
			}

			w.WriteHeader(405)
			if err := runtime.JSONProducer().Produce(w, response); err != nil {
				s.internalServerError(w, r, err)
			}
		}
	case oapiErrors.Error:
		switch code := e.Code(); code {
		case 404:
			response := NotFoundResponse{
				Response: Response{
					Version: &s.version,
				},
				Message: UnknownPathMessage,
			}

			w.WriteHeader(404)
			if err := runtime.JSONProducer().Produce(w, response); err != nil {
				s.internalServerError(w, r, err)
			}
		default:
			s.internalServerError(w, r, err)
		}
	default:
		s.internalServerError(w, r, err)
	}
}

func (s *Handler) NewRecoverMiddleware(httpHandler http.Handler) http.Handler {
	return newRecoverMiddleware(s, httpHandler)
}

func (s *Handler) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Raven capture
	s.capture(r, err)

	w.Header().Set("Content-Type", "application/json")

	response := &InternalServerErrorResponse{
		Response: Response{
			Version: &s.version,
		},
		Errors: map[string]string{
			"core": err.Error(),
		},
		Message: InternalServerErrorMessage,
	}

	w.WriteHeader(500)
	_ = runtime.JSONProducer().Produce(w, &response)
}

func (s *Handler) capture(r *http.Request, err error) {
	if s.ravenClient == nil {
		return
	}

	packet := raven.NewPacket(err.Error(),
		raven.NewException(err,
			raven.GetOrNewStacktrace(err, 2, 3, nil),
		),
		raven.NewHttp(r),
	)

	s.ravenClient.Capture(packet, nil)
}
