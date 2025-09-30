package logger

import (
	"log"
	"os"
)

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

type DefaultLogger struct {
	logger *log.Logger
}

func NewLogger() Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *DefaultLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Printf("[DEBUG] %s %v", msg, fields)
}

func (l *DefaultLogger) Info(msg string, fields ...interface{}) {
	l.logger.Printf("[INFO] %s %v", msg, fields)
}

func (l *DefaultLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Printf("[WARN] %s %v", msg, fields)
}

func (l *DefaultLogger) Error(msg string, fields ...interface{}) {
	l.logger.Printf("[ERROR] %s %v", msg, fields)
}

func (l *DefaultLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatalf("[FATAL] %s %v", msg, fields)
}
