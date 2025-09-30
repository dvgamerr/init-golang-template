package logger

import (
	"log"
	"os"
)

// Logger interface for logging
type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

// DefaultLogger is a simple logger implementation
type DefaultLogger struct {
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Debug logs debug level message
func (l *DefaultLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Printf("[DEBUG] %s %v", msg, fields)
}

// Info logs info level message
func (l *DefaultLogger) Info(msg string, fields ...interface{}) {
	l.logger.Printf("[INFO] %s %v", msg, fields)
}

// Warn logs warning level message
func (l *DefaultLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Printf("[WARN] %s %v", msg, fields)
}

// Error logs error level message
func (l *DefaultLogger) Error(msg string, fields ...interface{}) {
	l.logger.Printf("[ERROR] %s %v", msg, fields)
}

// Fatal logs fatal level message and exits
func (l *DefaultLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatalf("[FATAL] %s %v", msg, fields)
}
