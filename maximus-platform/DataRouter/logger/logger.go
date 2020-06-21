package logger

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type (
	coreStatus   string
	coreResource string

	apiStatus string

	eventStatus string
	eventType   string
)

const (
	CORECONNECTED coreStatus = "CONNECTED"
	COREFAILED    coreStatus = "FAILED"
	CORESTARTED   coreStatus = "STARTED"

	CORECONFIGDB coreResource = "configDB"
	COREEVENTDB  coreResource = "eventDB"
	COREMQ       coreResource = "mq"

	APICREATED apiStatus = "CREATED"
	APIEDITED  apiStatus = "EDITED"
	APIDELETED apiStatus = "DELETED"

	EVENTPASS   eventStatus = "PASS"
	EVENTFAILED eventStatus = "FAILED"

	EVENTREPLY   eventType = "REPLY"
	EVENTREQUEST eventType = "REQUEST"
)

type Logger interface {
	Core() CoreLogger
	API() APILogger
	Event() EventLogger

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})

	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type asyncLogger struct {
	*logrus.Logger

	mu      sync.Mutex
	queueCh chan func()
	exitCh  chan interface{}

	closed bool
	once   sync.Once
}

type Config struct {
	Workers int
	Queue   int
	Level   logrus.Level
}

func New(workers, queue int, level, output, format string) Logger {
	l := &asyncLogger{
		Logger: logrus.New(),

		queueCh: make(chan func(), queue),
		exitCh:  make(chan interface{}, 1),
	}

	switch strings.ToLower(level) {
	case "info":
		l.Logger.SetLevel(logrus.InfoLevel)
	case "debug":
		l.Logger.SetLevel(logrus.DebugLevel)
	case "warn":
		l.Logger.SetLevel(logrus.WarnLevel)
	case "trace":
		l.Logger.SetLevel(logrus.TraceLevel)
	case "fatal":
		l.Logger.SetLevel(logrus.FatalLevel)
	default:
		l.Logger.SetLevel(logrus.InfoLevel)
	}

	l.Logger.SetOutput(setOutput(output))
	if format == "JSON" {
		l.Logger.SetFormatter(&logrus.JSONFormatter{})
	}
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < workers; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				for {
					select {
					case <-l.exitCh:
						return
					case fn := <-l.queueCh:
						fn()
					}
				}
			}()
		}

		<-l.exitCh
		wg.Wait()

		l.mu.Lock()
		close(l.queueCh)

		l.closed = true
		l.mu.Unlock()
	}()

	return l
}

func setOutput(opt string) *os.File {
	if opt != "STDOUT" && opt != "" {
		logFile, err := os.OpenFile(opt, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer func() {
			err = logFile.Close()
			if err != nil {
				log.Fatalf("error closing file: %v", err)
			}
		}()
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		return logFile

	}
	return os.Stdout
}

// CoreLogger -> Info(), Fatal(), Debug()
type CoreLogger interface {
	Info(resource coreResource, host, port, version string, status coreStatus)
	Fatal(resource coreResource, host, port, message string, status coreStatus)
	Debug(...interface{})
}

func (l *asyncLogger) Core() CoreLogger {
	entry := l.Logger.WithField("context", "CORE")
	return &coreEntry{
		entry:  entry,
		Logger: l,
	}
}

type coreEntry struct {
	entry  *logrus.Entry
	Logger *asyncLogger
}

func (e *coreEntry) Info(resource coreResource, host, port, version string, status coreStatus) {
	if resource != "" {
		resource = "resource=" + resource
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

	e.Logger.queue(func() {
		e.entry.Infof("%s%s%s status=%s", resource, addr, version, status)
	})
}

func (e *coreEntry) Fatal(resource coreResource, host, port, message string, status coreStatus) {
	addr := host + ":" + port

	if message != "" {
		message = " message=" + message
	}

	// not async
	e.entry.Fatalf("resource=%s addr=%s%s status=%s",
		resource,
		strings.ToUpper(addr),
		message,
		status)
}

func (e *coreEntry) Debug(val ...interface{}) {
	e.Logger.queue(func() {
		e.entry.Debugln(val...)
	})
}

// APILogger[clientID || routeID || replyID] -> APIEntrier -> Info(), Debug()
type APILogger interface {
	Client() APIEntrier
	Route() APIEntrier
	Reply() APIEntrier
}

func (l *asyncLogger) API() APILogger {
	entry := l.Logger.WithField("context", "API")
	return &apiEntry{
		entry:  entry,
		Logger: l,
	}
}

type apiEntry struct {
	entry  *logrus.Entry
	Logger *asyncLogger
}

type APIEntrier interface {
	Info(guid, resource string, status apiStatus)
	Debug(...interface{})
}

type endpointEntry struct {
	entry     *apiEntry
	whichGUID string
}

func (e *apiEntry) Client() APIEntrier {
	return &endpointEntry{
		entry:     e,
		whichGUID: "clientID=",
	}
}

func (e *apiEntry) Route() APIEntrier {
	return &endpointEntry{
		entry:     e,
		whichGUID: "routeID=",
	}
}

func (e *apiEntry) Reply() APIEntrier {
	return &endpointEntry{
		entry:     e,
		whichGUID: "replyID=",
	}
}

func (e *endpointEntry) Info(guid, resource string, status apiStatus) {
	if resource != "" {
		resource = " resource=" + resource
	}

	e.entry.Logger.queue(func() {
		e.entry.entry.Infof("%s%s%s status=%s",
			e.whichGUID,
			strings.ToUpper(guid),
			resource,
			status)
	})
}

func (e *endpointEntry) Debug(val ...interface{}) {
	e.entry.Logger.queue(func() {
		e.entry.entry.Debugln(val...)
	})
}

// EventLogger -> Info(), Error(), Debug()
type EventLogger interface {
	Info(eventID, transactionID, routeID, replyID, cause string, eventType eventType, status eventStatus)
	Error(eventID, transactionID, cause string, eventType eventType, status eventStatus)
	Debug(srcTopicName interface{}, dstTopicName, body string)
}

func (l *asyncLogger) Event() EventLogger {
	entry := l.Logger.WithField("context", "EVENT")
	return &eventEntry{
		entry:  entry,
		Logger: l,
	}
}

type eventEntry struct {
	entry  *logrus.Entry
	Logger *asyncLogger
}

func (e *eventEntry) Info(eventID, transactionID, routeID, replyID,
	cause string, eventType eventType, status eventStatus) {
	if transactionID != "" {
		transactionID = " transactionID=" + strings.ToUpper(transactionID)
	}

	if replyID != "" {
		replyID = " replyID=" + strings.ToUpper(replyID)
	}

	if cause != "" {
		cause = " cause=" + strings.ToUpper(cause)
	}

	e.Logger.queue(func() {
		e.entry.Logger.Infof("eventID=%s%s status=%s routeID=%s%s type=%s%s",
			strings.ToUpper(eventID),
			transactionID,
			status,
			strings.ToUpper(routeID),
			replyID,
			eventType,
			cause)
	})
}

func (e *eventEntry) Error(eventID, transactionID, cause string, eventType eventType, status eventStatus) {
	if transactionID != "" {
		transactionID = " transactionID=" + strings.ToUpper(transactionID)
	}

	if cause != "" {
		cause = " cause=" + strings.ToUpper(cause)
	}

	e.Logger.queue(func() {
		e.entry.Logger.Errorf("eventID=%s%s status=%s type=%s%s",
			strings.ToUpper(eventID),
			transactionID,
			status,
			eventType,
			cause)
	})
}

func (e *eventEntry) Debug(srcTopicName interface{}, dstTopicName, body string) {
	e.Logger.queue(func() {
		e.entry.Logger.Debugf("src.topic.name=%s dst.topic.name=%s body=%s",
			srcTopicName,
			strings.ToUpper(dstTopicName),
			strings.ToUpper(body))
	})
}

func (l *asyncLogger) Debug(args ...interface{}) {
	l.queue(func() {
		newEntry(l).Debug(args...)
	})
}

func (l *asyncLogger) Debugf(format string, args ...interface{}) {
	l.queue(func() {
		newEntry(l).Debugf(format, args...)
	})
}

func (l *asyncLogger) Debugln(args ...interface{}) {
	l.queue(func() {
		newEntry(l).Debugln(args...)
	})
}

func (l *asyncLogger) Fatal(args ...interface{}) {
	l.queue(func() {
		newEntry(l).Fatal(args...)
	})
}

func (l *asyncLogger) Fatalln(args ...interface{}) {
	l.queue(func() {
		newEntry(l).Fatalln(args...)
	})
}

func (l *asyncLogger) Printf(format string, args ...interface{}) {
	l.queue(func() {
		newEntry(l).Printf(format, args...)
	})
}

func (l *asyncLogger) Println(args ...interface{}) {
	l.queue(func() {
		newEntry(l).Println(args...)
	})
}

func (l *asyncLogger) queue(fn func()) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return
	}

	l.queueCh <- fn
}
