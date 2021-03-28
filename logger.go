package gpsd

import (
	"log"
	"os"
	"strconv"
)

// Logger gives ability to use different logger implementations.
type Logger interface {
	Errorf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

func newStdLogger() *envLogger {
	return &envLogger{
		d: debugFromEnv(),
	}
}

type envLogger struct {
	d bool
}

func (l *envLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[gpsd] E "+format, v...)
}

func (l *envLogger) Debugf(format string, v ...interface{}) {
	if l.d {
		log.Printf("[gpsd] D "+format, v...)
	}
}

func debugFromEnv() bool {
	ok, _ := strconv.ParseBool(os.Getenv("GPSD_DEBUG"))
	return ok
}
