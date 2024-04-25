package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(message string) {
	l.write(Trace, message)
}

func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template+"\n", vals...))
}
func (l *DefaultLogger) Debug(message string) {
	l.write(Debug, message)
}

func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template+"\n", vals...))
}

func (l *DefaultLogger) Info(message string) {
	l.write(Information, message)
}

func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(template+"\n", vals...))
}

func (l *DefaultLogger) Warn(message string) {
	l.write(Warning, message)
}

func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(template+"\n", vals...))
}

func (l *DefaultLogger) Panic(message string) {
	l.write(Fatal, message)
	if l.triggerPanic {
		panic(message)
	}
}

func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	l.write(Fatal, fmt.Sprintf(template+"\n", vals...))
}
