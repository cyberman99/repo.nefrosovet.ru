package goswagger_errors

import (
	"fmt"
	"net/http"
)

type RecoverMiddleware struct {
	handler     *Handler
	httpHandler http.Handler
}

func newRecoverMiddleware(handler *Handler, httpHandler http.Handler) *RecoverMiddleware {
	return &RecoverMiddleware{
		handler:     handler,
		httpHandler: httpHandler,
	}
}

func (m *RecoverMiddleware) SetHTTPHandler(httpHandler http.Handler) {
	m.httpHandler = httpHandler
}

func (m *RecoverMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			err, ok := rec.(error)
			if !ok {
				err = fmt.Errorf("%s", rec)
			}

			// Return internal server error
			m.handler.internalServerError(w, r, err)
		}
	}()

	if m.httpHandler != nil {
		m.httpHandler.ServeHTTP(w, r)
	}
}
