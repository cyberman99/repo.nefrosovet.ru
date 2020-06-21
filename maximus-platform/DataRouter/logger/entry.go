package logger

import (
	"github.com/sirupsen/logrus"
)

type Entry interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})

	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type asyncEntry struct {
	*logrus.Entry

	Logger *asyncLogger
}

func newEntry(logger *asyncLogger) Entry {
	e := &asyncEntry{
		Logger: logger,
		Entry:  logrus.NewEntry(logger.Logger),
	}

	return e
}

func (e *asyncEntry) Debug(args ...interface{}) {
	e.Logger.queue(func() {
		e.Entry.Debug(args...)
	})
}

func (e *asyncEntry) Debugf(format string, args ...interface{}) {
	e.Logger.queue(func() {
		e.Entry.Debugf(format, args...)
	})
}

func (e *asyncEntry) Printf(format string, args ...interface{}) {
	e.Logger.queue(func() {
		e.Entry.Printf(format, args...)
	})
}

func (e *asyncEntry) Fatal(args ...interface{}) {
	// not async
	e.Entry.Fatal(args...)
}

func (e *asyncEntry) Debugln(args ...interface{}) {
	e.Logger.queue(func() {
		e.Entry.Debugln(args...)
	})
}

func (e *asyncEntry) Println(args ...interface{}) {
	e.Logger.queue(func() {
		e.Entry.Println(args...)
	})
}

func (e *asyncEntry) Fatalln(args ...interface{}) {
	// not async
	e.Entry.Fatalln(args...)
}
