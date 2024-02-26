package logger

import (
	"context"
	"weather-api/models"

	"github.com/sirupsen/logrus"
)

// Logger struct
type Logger struct {
	*logrus.Logger
}

// NewLogger creates new instance of Logger
func NewLogger() *Logger {
	var log = &Logger{logrus.New()}
	log.Formatter = &logrus.JSONFormatter{}
	return log
}

// InfoC logs info with context request id
func (l *Logger) InfoC(c context.Context, args ...interface{}) {
	l.WithFields(logrus.Fields{"Request-id": c.Value(models.RequestID)}).Info(args...)
}

// WarnC logs warn with context request id
func (l *Logger) WarnC(c context.Context, args ...interface{}) {
	l.WithFields(logrus.Fields{"Request-id": c.Value(models.RequestID)}).Warn(args...)
}

// ErrorC logs error with context request id
func (l *Logger) ErrorC(c context.Context, args ...interface{}) {
	l.WithFields(logrus.Fields{"Request-id": c.Value(models.RequestID)}).Error(args...)
}

// Log will be accessed across different packages
var Log = NewLogger()
