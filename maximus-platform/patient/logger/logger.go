package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	coreStatus string
	apiStatus  string
)

const (
	CORESTARTED   coreStatus = "STARTED"
	CORECONNECTED coreStatus = "CONNECTED"
	COREFAILED    coreStatus = "FAILED"
	APIPOSTED     apiStatus  = "POSTED"
	APIUPDATED    apiStatus  = "UPDATED"
	APIDELETED    apiStatus  = "DELETED"
)

type Logger interface {
	Core() CoreEntrier
	Api() APIEntrier
}

type patientLog struct {
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
	if out != "STDOUT" {
		logFile, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
	return &patientLog{lg: lg}
}

type CoreEntrier interface {
	Info(resourse string, host string, port string, version string, status coreStatus)
	Fatal(resourse string, host string, port string, status coreStatus)
	Debug(...interface{})
	Level() string
}

func (l *patientLog) Core() CoreEntrier {
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

func (e *coreEntry) Info(
	resourse string, host string, port string, version string, status coreStatus,
) {
	if resourse != "" {
		resourse = "resource=" + strings.ToUpper(resourse)
	}

	addr := host + ":" + port
	if addr != ":" {
		addr = " addr=" + strings.ToUpper(addr)
	} else {
		addr = ""
	}
	if version != "" {
		version = " version=" + strings.ToUpper(version)
	}
	if status != "" {
		status = " status=" + status
	}

	e.entry.Infof("%s%s%s%s", resourse, addr, version, status)
}

func (e *coreEntry) Fatal(resourse string, host string, port string, status coreStatus) {
	addr := host + ":" + port
	e.entry.Fatalf("resource=%s addr=%s status=%s",
		strings.ToUpper(resourse),
		strings.ToUpper(addr),
		status,
	)
}

func (e *coreEntry) Debug(val ...interface{}) {
	e.entry.Debugln(val...)
}

func (e *coreEntry) Level() string {
	return e.entry.Level.String()
}

type APIEntrier interface {
	Patient() APIPatient
	Appointment() APIAppointment
	Employee() APIEmployee
	Infof(format string, args ...interface{})
}

func (l *patientLog) Api() APIEntrier {
	entry := l.lg.WithFields(logrus.Fields{
		"context": "API",
	})
	return &apiEntry{
		entry,
	}
}

type apiEntry struct {
	entry *logrus.Entry
}

type APIPatient interface {
	Info(text string, status apiStatus)
	Debug(...interface{})
}

type APIAppointment interface {
	Info(text string, status apiStatus)
	Debug(...interface{})
}

type APIEmployee interface {
	Info(text string, status apiStatus)
	Debug(...interface{})
}

type endpointEntry struct {
	entry *apiEntry
}

func (e *apiEntry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

func (e *apiEntry) Patient() APIPatient {
	return &endpointEntry{e}
}

func (e *apiEntry) Appointment() APIAppointment {
	return &endpointEntry{e}
}

func (e *apiEntry) Employee() APIEmployee {
	return &endpointEntry{e}
}

func (e *endpointEntry) Info(text string, status apiStatus) {
	e.entry.entry.Infof("%s status=%v",
		text,
		status,
	)
}

func (e *endpointEntry) Debug(val ...interface{}) {
	e.entry.entry.Debugln(val...)
}
