// Package logging implements a Zap-based logger, configured to be GCE/Stackdriver-compatible.
package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a Zap-based logger configured to be GCE/Stackdriver-compatible.
type Logger zap.Logger

// NewLog returns a new Logger instance.
// If 'productionLogging' is true, JSON logger is constructed which logs at the Info level (by default); if
// 'productionLogging' is false, a developent-oriented console logger is constructed which logs at the Debug level.
// 'source' is the optional source (e.g. application or service name) of the events; all logs under this logger and
// its domains will be tagged with the source, if one is provided.
// 'fields' is an optional set of fields to install in the logger instance.
func NewLog(productionLogging bool, source string, fields ...zapcore.Field) (*Logger, error) {

	var log *zap.Logger
	var err error

	options := []zap.Option{zap.AddCallerSkip(1)}

	if productionLogging {
		if log, err = newProductionConfig().Build(options...); err != nil {
			return nil, fmt.Errorf("failed to create production logger: %v", err)
		}
	} else {
		if log, err = newDevelopmentConfig().Build(options...); err != nil {
			return nil, fmt.Errorf("failed to create development logger: %v", err)
		}
	}

	if source != "" {
		fields = append(fields, zap.String(SourceKey, source))
	}

	if hostname, err := os.Hostname(); err == nil {
		fields = append(fields, zap.String(HostKey, hostname))
	}

	return (*Logger)(log.With(fields...)), nil
}

// Debug logs development-related information that is not typically necessary for production. When choosing a logging
// level, you should assume that debug logging will be disabled in production environments.
func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	(*zap.Logger)(l).Debug(msg, fields...)
}

// Info logs at a level suitable for production; the log should be *actionable information* that will be read by a
// human, or by a machine.
func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	(*zap.Logger)(l).Info(msg, fields...)
}

// ErrorWithTrace logs an error and captures the current stack trace (at the cost of a small performance hit).
// ErrorWithTrace is appropriate for unexpected errors, where the stack trace will be useful in diagnosing the
// root cause. For expected / non-critical errors, you may wish to prefer Info with a zap.Error() field.
func (l *Logger) ErrorWithTrace(err error, msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.Error(err))
	(*zap.Logger)(l).Error(msg, fields...)
}

// NewDomainLogger returns a new Logger instance, based on the current instance, with the provided domain.
// The domain is set via With() using a well-defined key.
func (l *Logger) NewDomainLogger(domain string) *Logger {
	newLogger := (*zap.Logger)(l).Named(domain)
	return (*Logger)(newLogger)
}

// Flush tries to commit any buffered log messages.
func (l *Logger) Flush() {
	if err := (*zap.Logger)(l).Sync(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to flush logger: %v", err)
	}
}
