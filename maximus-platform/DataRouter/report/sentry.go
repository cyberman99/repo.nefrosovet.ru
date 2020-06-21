package report

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"net/http"
)

type Reporter interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type RavenHandler struct {
	o http.Handler
}

func NewRavenHandler(h http.Handler) Reporter {
	return &RavenHandler{
		o: h,
	}
}

func (rh *RavenHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			Capture(r, req)
		}
	}()

	rh.o.ServeHTTP(rw, req)
}

func Capture(r interface{}, req *http.Request) {
	str := fmt.Sprint(r)
	err, ok := r.(error)
	if !ok {
		err = fmt.Errorf("error: %v", r)
	}
	packet := raven.NewPacket(
		str,
		raven.NewException(errors.New(str),
			raven.GetOrNewStacktrace(
				err,
				2,
				3,
				nil,
			),
		),
		raven.NewHttp(req),
	)
	raven.Capture(packet, nil)
}
