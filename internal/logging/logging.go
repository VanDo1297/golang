package logging

import (
	"github.com/go-errors/errors"
	"go.uber.org/zap"
)

type DefaultLogger struct {
	zapLogger *zap.SugaredLogger
}

func NewDefaultLogger(development bool) (Logger, error) {
	var logger *zap.Logger
	var e error
	if development {
		logger, e = zap.NewDevelopment()
	} else {
		logger, e = zap.NewProduction()
	}
	if e != nil {
		return nil, errors.Wrap(e, 0)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return &DefaultLogger{zapLogger: sugar}, nil
}

func (logger DefaultLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Infow(msg, keysAndValues...)
}
func (logger DefaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Warnw(msg, keysAndValues...)
}
func (logger DefaultLogger) Error(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Errorw(msg, keysAndValues...)
}
func (logger DefaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Debugw(msg, keysAndValues...)
}
func (logger DefaultLogger) Panic(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Panicw(msg, keysAndValues...)
}

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Panic(msg string, keysAndValues ...interface{})
}
