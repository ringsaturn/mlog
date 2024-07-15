// Package mlog provides MongoDB's client logging support via zap logger.
package mlog

import (
	"fmt"

	"go.uber.org/zap"
)

type LogSink struct{ logger *zap.Logger }

func New(logger *zap.Logger) *LogSink { return &LogSink{logger: logger} }

func toZapFields(keysAndValues []interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 >= len(keysAndValues) {
			break
		}

		key := keysAndValues[i]
		value := keysAndValues[i+1]

		keyStr, ok := key.(string)
		if !ok {
			keyStr = fmt.Sprintf("%v", key)
		}
		fields = append(fields, zap.Any(keyStr, value))
	}
	return fields
}

type logFn func(string, ...zap.Field)

func (l *LogSink) Info(level int, message string, keysAndValues ...interface{}) {
	var fn logFn = l.logger.Debug

	switch level {
	case 1:
		fn = l.logger.Info
	default:
		fn = l.logger.Debug
	}

	fn(message, toZapFields(keysAndValues)...)
}

func (l *LogSink) Error(err error, message string, keysAndValues ...interface{}) {
	l.logger.Error(message, append(toZapFields(keysAndValues), zap.Error(err))...)
}
