package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
	"log"
	"io"
)

type (
	coreStatus string
	coreResource string

	authStatus string
	authEventType string
)

const (
	CORECONNECTED coreStatus = "CONNECTED"
	COREFAILED coreStatus = "FAILED"
	CORESTARTED coreStatus = "STARTED"

	COREDB coreResource = "DB"
	COREINDEX coreResource = "index"

	AUTHSUCCESS authStatus = "SUCCESS"
	AUTHFAIL authStatus = "FAIL"

	AUTHUSERLOGIN authEventType = "USER_LOGIN"
	AUTHUSERCREATE authEventType = "USER_CREATE"
	AUTHUSERUPDATE authEventType = "USER_UPDATE"
	AUTHUSERUPDATESETTINGS authEventType = "USER_UPDATE_SETTINGS"
	AUTHUSERCREATECONTACT authEventType = "USER_CREATE_CONTACT"
	AUTHUSERUPDATECONTACT authEventType = "USER_UPDATE_CONTACT"
)

var logger Logger

func GetLogger() Logger {
	if logger == nil {
		log.Fatalln("Logger is not found")
	}
	return logger
}

// Logger Init

type Logger interface {
	Core() CoreEntrier
	Auth() AuthEntrier
}

type profileLog struct {
	lg *logrus.Logger
}

func NewLogger (out io.Writer, levStr string, format string) Logger {
	lg := logrus.New()
	level, err := logrus.ParseLevel(levStr)
	if err != nil {
		lg.WithError(err).Warnln("Can't set logging level")
	}

	lg.SetLevel(level)
	lg.SetOutput(out)

	if format == "JSON" {
		lg.SetFormatter(&logrus.JSONFormatter{})
	}

	logger := &profileLog{lg: lg}
	return logger
}

// Core Entrier

type CoreEntrier interface {
	Info(resource coreResource, host, port, version string, status coreStatus)
	Fatal(resource coreResource, host, port string, status coreStatus)
	Debug(...interface{})
	Error(resource coreResource, indexHost, path string, status coreStatus)
}

func (l *profileLog) Core() CoreEntrier {
	entry := l.lg.WithField("context", "CORE")

	return &coreEntry{
		entry,
	}
}

type coreEntry struct {
	entry *logrus.Entry
}

func (e *coreEntry) Info(resource coreResource, host, port, version string, status coreStatus) {
	if resource != "" {
		resource = "resource=" + resource
	}

	addr := host + ":" + port
	if addr != ":" {
		addr = " addr=" + strings.ToUpper(addr)
	} else {
		addr=""
	}

	if version != "" {
		version = " version=" + strings.ToUpper(version)
	}

	e.entry.Infof("%s%s%s status=%s", resource, addr, version, status)
}

func (e *coreEntry) Fatal(resource coreResource, host, port string, status coreStatus) {
	addr := host + ":" + port

	e.entry.Fatalf("resource=%s addr=%s status=%s", resource, addr, status)
}

func (e *coreEntry) Debug(val ...interface{}) {
	e.entry.Debugln(val...)
}

func (e *coreEntry) Error(resource coreResource, indexHost, path string, status coreStatus) {
	addr := indexHost + "/" + path

	e.entry.Errorf("resource=%s addr=%s status=%s", resource, addr, status)
}

// Auth Entrier

type AuthEntrier interface {
	Info(eventType authEventType, entityLogin string, status authStatus)
}

func (l *profileLog) Auth() AuthEntrier {
	entry := l.lg.WithField("context", "AUTH")

	return &authEntry{
		entry,
	}
}

type authEntry struct {
	entry *logrus.Entry
}

func (e *authEntry) Info(eventType authEventType, entityLogin string, status authStatus) {
	e.entry.Infof("eventType=%s login=%s status=%s", eventType, entityLogin, status)
}