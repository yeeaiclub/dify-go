// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"os"
	"sync"
)

var (
	globalLogger Logger
	globalMu     sync.Mutex
)

func init() {
	globalLogger = New(os.Stderr, DebugLevel)
}

// Global returns the global logger instance.
// Returns the current global Logger implementation.
func Global() Logger {
	return globalLogger
}

// SetGlobalLogger sets the global logger instance.
// logger: The logger implementation to use globally.
func SetGlobalLogger(logger Logger) {
	globalMu.Lock()
	defer globalMu.Unlock()
	globalLogger = logger
}

// Debug logs a message at Debug level.
// args: Arguments to be logged.
func Debug(args ...any) { globalLogger.Debug(args...) }

// Info logs a message at Info level.
// args: Arguments to be logged.
func Info(args ...any) { globalLogger.Info(args...) }

// Warn logs a message at Warn level.
// args: Arguments to be logged.
func Warn(args ...any) { globalLogger.Warn(args...) }

// Error logs a message at Error level.
// args: Arguments to be logged.
func Error(args ...any) { globalLogger.Error(args...) }

// Fatal logs a message at Fatal level and exits the program.
// args: Arguments to be logged.
func Fatal(args ...any) { globalLogger.Fatal(args...) }

// Debugf logs a formatted message at Debug level.
// format: Format string for the log message.
// args: Arguments to be formatted into the string.
func Debugf(format string, args ...any) { globalLogger.Debugf(format, args...) }

// Infof logs a formatted message at Info level.
// format: Format string for the log message.
// args: Arguments to be formatted into the string.
func Infof(format string, args ...any) { globalLogger.Infof(format, args...) }

// Warnf logs a formatted message at Warn level.
// format: Format string for the log message.
// args: Arguments to be formatted into the string.
func Warnf(format string, args ...any) { globalLogger.Warnf(format, args...) }

// Errorf logs a formatted message at Error level.
// format: Format string for the log message.
// args: Arguments to be formatted into the string.
func Errorf(format string, args ...any) { globalLogger.Errorf(format, args...) }

// Fatalf logs a formatted message at Fatal level and exits the program.
// format: Format string for the log message.
// args: Arguments to be formatted into the string.
func Fatalf(format string, args ...any) { globalLogger.Fatalf(format, args...) }

// SetLevel sets the log level for the global logger.
// level: The log level to use.
func SetLevel(level Level) {
	globalLogger.SetLevel(level)
}
