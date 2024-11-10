package main

import "go.uber.org/zap"

type loggerFacade struct {
	logger zap.Logger
}

func NewLoggerFacade() Log {
	return &loggerFacade{
		logger: *zap.NewExample(),
	}
}

func (l *loggerFacade) Info(message string) {
	l.logger.Info(message)
}

func (l *loggerFacade) Error(message string) {
	l.logger.Error(message)
}
