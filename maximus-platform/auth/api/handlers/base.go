package handlers

import (
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PayloadSuccessMessage used in 200 answers
var PayloadSuccessMessage = "SUCCESS"

// PayloadFailMessage - coomon fail message
var PayloadFailMessage = "FAIL"

// PayloadSuccessMessage used in 500 answers
var InternalServerErrorMessage = "Internal server error"

// PayloadAuthFailureMessage used for authentication failure responses
var PayloadAuthFailureMessage = "Authentication failure"

// PayloadValidationErrorMessage - used on composite errors
var PayloadValidationErrorMessage = "Validation error"

// NotFoundMessage used in 404 answers
var NotFoundMessage = "Entity not found"

// AccessDeniedMessage used in 401 answers
var AccessDeniedMessage = "Access denied"

// Version of service
var Version string

func getSourceIP(req *http.Request) string {
	var err error
	var addr string

	switch {
	case req.Header.Get("X-REAL-IP") != "":
		addr = req.Header.Get("X-REAL-IP")
	case req.Header.Get("X-FORWARDED-FOR") != "":
		addr = req.Header.Get("X-FORWARDED-FOR")
	default:
		addr, _, err = net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			logrus.WithField("address", req.RemoteAddr).Error("address is not IP:port")
		}
	}

	return addr
}
