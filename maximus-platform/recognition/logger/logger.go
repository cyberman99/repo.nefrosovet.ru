package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

type (
	coreStatus string
	apiStatus  string
)

const (
	CORESTARTED   coreStatus = "STARTED"
	CORECONNECTED coreStatus = "CONNECTED"
	COREFAILED    coreStatus = "FAILED"
	APIUPLOADED   apiStatus  = "UPLOADED"
	APIDELETED    apiStatus  = "DELETED"
	APIREKOGNIZED apiStatus  = "REKOGNIZED"
)

type Logger interface {
	Core() CoreEntrier
	Api(sim float64) APIEntrier
}

type rekognLog struct {
	lg *logrus.Logger
}


func NewLogger(out io.Writer, levStr string, format string) Logger {
	lg := logrus.New()
	if levStr == "" {
		lg.SetLevel(logrus.DebugLevel)
	} else {
		level, err := logrus.ParseLevel(levStr)
		if err != nil {
			lg.WithError(err).Warnln("Can't set logging level")
		}
		lg.SetLevel(level)
	}

	lg.SetOutput(out)
	if format == "JSON" {
		lg.SetFormatter(&logrus.JSONFormatter{})
	}
	return &rekognLog{lg: lg}
}

type CoreEntrier interface {
	Info(resourse string, host string, port string, version string, status coreStatus)
	Fatal(resourse string, host string, port string, status coreStatus)
	Debug(...interface{})
	Level() string
}

func (l *rekognLog) Core() CoreEntrier {
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
	Photo() APIPhoto
	Rekognition() APIRekognition
	Infof(format string, args ...interface{})
}

func (l *rekognLog) Api(sim float64) APIEntrier {
	entry := l.lg.WithFields(logrus.Fields{
		"context": "API",
	})
	return &apiEntry{
		entry,
		sim,
	}
}

type apiEntry struct {
	entry      *logrus.Entry
	similarity float64
}

type APIPhoto interface {
	Info(imgID, personID string, status apiStatus)
	Debug(...interface{})
}

type APIRekognition interface {
	RekInfo(imgID string, status apiStatus)
	Debug(...interface{})
}

type endpointEntry struct {
	entry *apiEntry
}

func (e *apiEntry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

func (e *apiEntry) Photo() APIPhoto {
	return &endpointEntry{e}
}

func (e *apiEntry) Rekognition() APIRekognition {
	return &endpointEntry{e}
}

func (e *endpointEntry) Info(imgID, personID string, status apiStatus) {
	e.entry.entry.Infof("imageID=%s personID=%s status=%v",
		strings.ToUpper(imgID),
		strings.ToUpper(personID),
		status,
	)
}

func (e *endpointEntry) RekInfo(imgID string, status apiStatus) {
	e.entry.entry.Infof("imageID=%s similarity=%0.2f status=%v",
		strings.ToUpper(imgID),
		e.entry.similarity,
		status,
	)
}

func (e *endpointEntry) Debug(val ...interface{}) {
	e.entry.entry.Debugln(val...)
}
