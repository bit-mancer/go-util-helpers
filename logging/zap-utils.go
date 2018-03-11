package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const gkeTimeKey = "timestamp"
const gkeLevelKey = "severity"
const gkeNameKey = DomainKey
const gkeCallerKey = "caller"
const gkeMessageKey = "message"
const gkeStacktraceKey = "trace"

// FieldSet is an alias for a slice of zapcore Fields. Functions that aggregate fields can return a FieldSet to allow
// further aggregation.
type FieldSet []zapcore.Field

// AppendFieldSet appends a FieldSet to the current FieldSet.
func (f FieldSet) AppendFieldSet(fields FieldSet) FieldSet {
	return append([]zapcore.Field(f), fields...)
}

// Append appends fields to the current FieldSet.
func (f FieldSet) Append(fields ...zapcore.Field) FieldSet {
	return append([]zapcore.Field(f), fields...)
}

func rfc3339UTCZuluTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Timestamp

	// Sample Stackdriver format: 2014-10-02T15:01:23.045123456Z

	// Golang reference layout:
	// Mon Jan 2 15:04:05 MST 2006
	// 01/02 03:04:05PM '06 -0700

	enc.AppendString(t.Format("2006-01-02T15:04:05.000000000Z"))
}

func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        gkeTimeKey,
		LevelKey:       gkeLevelKey,
		NameKey:        gkeNameKey,
		CallerKey:      gkeCallerKey,
		MessageKey:     gkeMessageKey,
		StacktraceKey:  gkeStacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     rfc3339UTCZuluTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func newDevelopmentConfig() zap.Config {
	return zap.NewDevelopmentConfig()
}
