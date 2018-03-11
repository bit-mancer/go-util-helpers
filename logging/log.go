package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is implemented by... loggers.
type Logger interface {

	// Debug logs development-related information that is not typically necessary for production. When choosing a
	// logging level, you should assume that debug logging will be disabled in production environments.
	Debug(msg string, fields ...zapcore.Field)

	// Info logs at a level suitable for production; the log should be *actionable information* that will be read by
	// a human, or by a machine.
	Info(msg string, fields ...zapcore.Field)

	// ErrorWithTrace logs an error and captures the current stack trace (at the cost of a small performance hit).
	// ErrorWithTrace is appropriate for unexpected errors, where the stack trace will be useful in diagnosing the
	// root cause. For expected / non-critical errors, you may wish to prefer Info with a zap.Error() field.
	ErrorWithTrace(err error, msg string, fields ...zapcore.Field)

	// NewDomainLogger returns a new Logger instance, based on the current instance, with the provided domain.
	// The domain is set via With() using a well-defined key.
	NewDomainLogger(domain string) Logger

	// Flush tries to commit any buffered log messages.
	Flush()
}

type logger zap.Logger

// NewLog returns a new Logger instance.
// If 'productionLogging' is true, JSON logger is constructed which logs at the Info level (by default); if
// 'productionLogging' is false, a developent-oriented console logger is constructed which logs at the Debug level.
// 'source' is the optional source (e.g. application or service name) of the events; all logs under this logger and
// its domains will be tagged with the source, if one is provided.
// 'fields' is an optional set of fields to install in the logger instance.
func NewLog(productionLogging bool, source string, fields ...zapcore.Field) (Logger, error) {

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

	return (*logger)(log.With(fields...)), nil
}

func (l *logger) Debug(msg string, fields ...zapcore.Field) {
	(*zap.Logger)(l).Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zapcore.Field) {
	(*zap.Logger)(l).Info(msg, fields...)
}

func (l *logger) ErrorWithTrace(err error, msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.Error(err))
	(*zap.Logger)(l).Error(msg, fields...)
}

func (l *logger) NewDomainLogger(domain string) Logger {
	newLogger := (*zap.Logger)(l).Named(domain)
	return (*logger)(newLogger)
}

func (l *logger) Flush() {
	if err := (*zap.Logger)(l).Sync(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to flush logger: %v", err)
	}
}
