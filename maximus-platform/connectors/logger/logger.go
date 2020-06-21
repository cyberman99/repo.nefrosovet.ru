package logger

import (
	"os"
	"repo.nefrosovet.ru/maximus-platform/connectors/connector"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	logCoreStatus  string
	logEventStatus string
)

const (
	CORESTARTED   logCoreStatus  = "STARTED"
	CORECONNECTED logCoreStatus  = "CONNECTED"
	COREFAILED    logCoreStatus  = "FAILED"
	EVENTSUCCESS  logEventStatus = "SUCCESS"
	EVENTFAIL     logEventStatus = "FAILED"

	resource = "MQ"
)

type Logger interface {
	Core() CoreEntrier
	Event() EventEntrier
}

type connectorLog struct {
	lg *logrus.Logger
}

func NewLogger(out, levStr, format string) Logger {
	lg := logrus.New()
	lg.SetLevel(logrus.DebugLevel)
	lg.SetOutput(os.Stdout)

	if levStr != "" {
		level, err := logrus.ParseLevel(levStr)
		if err != nil {
			lg.WithError(err).Warnln("Can't set logging level")
		}
		lg.SetLevel(level)
	}
	if out != "STDOUT" && out != "" {
		logFile, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer func() {
			err = logFile.Close()
			if err != nil {
				lg.Fatal(err)
			}
		}()
		if err != nil {
			lg.Fatal(err)
		}
		lg.SetOutput(logFile)
	}

	if format == "JSON" {
		lg.SetFormatter(&logrus.JSONFormatter{})
	}
	return &connectorLog{lg: lg}
}

type CoreEntrier interface {
	MQConnected(host string, port string)
	MQFailed(host string, port string)
	AppStarted(version string)
	Debug(...interface{})
	Level() string
}

func (l *connectorLog) Core() CoreEntrier {
	entry := l.lg.WithFields(logrus.Fields{
		"context": "CORE",
	})

	return &coreEntry{
		entry,
	}
}

type coreEntry struct {
	entry *logrus.Entry
}

func (e *coreEntry) AppStarted(version string) {
	if version != "" {
		version = " version=" + strings.ToUpper(version)
	}

	status := " status=" + CORESTARTED


	e.entry.Infof("%s%s", version, status)
}

func (e *coreEntry) MQConnected(host string, port string) {
	res := "resource=" + strings.ToUpper(resource)

	addr := host + ":" + port
	if addr != ":" {
		addr = " addr=" + strings.ToUpper(addr)
	} else {
		addr = ""
	}

	status := " status=" + CORECONNECTED


	e.entry.Infof("%s%s%s", res, addr, status)
}

func (e *coreEntry) MQFailed(host string, port string) {
	res := "resource=" + strings.ToUpper(resource)

	addr := host + ":" + port
	if addr != ":" {
		addr = " addr=" + strings.ToUpper(addr)
	} else {
		addr = ""
	}
	status := " status=" + COREFAILED

	e.entry.Fatalf("%s%s%s", res, addr, status)
}


func (e *coreEntry) Debug(val ...interface{}) {
	e.entry.Debugln(val...)
}

func (e *coreEntry) Level() string {
	return e.entry.Level.String()
}

type EventEntrier interface {
	EventSuccess(
		txID string, connType connector.ConnectorType, connectorID string,
	)
	EventFail(
		txID string, connType connector.ConnectorType, connectorID string, cause string,
	)
	Debug(val ...interface{})
}

func (l *connectorLog) Event() EventEntrier {
	entry := l.lg.WithFields(logrus.Fields{
		"context": "EVENT",
	})
	return &eventEntry{
		entry,
	}
}

type eventEntry struct {
	entry *logrus.Entry
}

func (e *eventEntry) EventSuccess(
	txID string, connType connector.ConnectorType, connectorID string,
) {
	if txID != "" {
		txID = "transactionID=" + strings.ToUpper(txID)
	}
	if connType != "" {
		connType = connector.ConnectorType(" connectorType=" + strings.ToUpper(string(connType)))
	}
	if connectorID != "" {
		connectorID = " connectorID=" + strings.ToUpper(connectorID)
	}
	status := " status=" + EVENTSUCCESS

	e.entry.Infof("%s%s%s%s", txID, connType, connectorID, status)
}


func (e *eventEntry) EventFail(
	txID string, connType connector.ConnectorType, connectorID string, cause string,
) {
	if txID != "" {
		txID = "transactionID=" + strings.ToUpper(txID)
	}
	if connType != "" {
		connType = connector.ConnectorType(" connectorType=" + strings.ToUpper(string(connType)))
	}
	if connectorID != "" {
		connectorID = " connectorID=" + strings.ToUpper(connectorID)
	}
	status := " status=" + EVENTFAIL
	if cause != "" {
		cause = " cause=" + strings.ToUpper(cause)
	}

	e.entry.Infof("%s%s%s%s%s", txID, connType, connectorID, status, cause)
}


func (e *eventEntry) Debug(val ...interface{}) {
	e.entry.Debugln(val...)
}
