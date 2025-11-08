// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Level represents the log level type, used to control the verbosity of log output
type Level int32

// Predefined log level constants
const (
	// DebugLevel is used for the most detailed debug information
	DebugLevel Level = iota
	// InfoLevel is used for general informational logs
	InfoLevel
	// WarnLevel is used for warning messages that may indicate potential issues but don't prevent program execution
	WarnLevel
	// ErrorLevel is used for error messages that indicate problems but allow the program to continue running
	ErrorLevel
	// FatalLevel is used for critical errors that cause the program to exit after logging
	FatalLevel
)

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

// Logger defines the interface for a logger, providing multiple log levels and formatting methods
type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	SetLevel(level Level)
	GetLevel() Level
}

// BaseLogger is the basic implementation of the Logger interface, providing thread-safe log level control
// and common logging logic
type BaseLogger struct {
	level Level
	mu    sync.Mutex
	impl  BaseImplementation
}

// BaseImplementation defines the interface for underlying log output, used to implement specific log output methods
type BaseImplementation interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

// New creates a new logger instance
// writer: the log output destination
// level: the log level threshold, logs below this level will be ignored
func New(writer io.Writer, level Level) Logger {
	return &BaseLogger{
		level: level,
		impl:  newStdImplementation(writer),
		mu:    sync.Mutex{},
	}
}

// SetLevel sets the level of the logger
// level: the new log level, panics if the level is invalid
func (l *BaseLogger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < DebugLevel || level > FatalLevel {
		panic("invalid log level")
	}
	l.level = level
}

// GetLevel gets the current level of the logger
func (l *BaseLogger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// canLogAt checks if logs at the specified level should be recorded
// v: the log level to check
// returns true if logs at this level should be recorded
func (l *BaseLogger) canLogAt(v Level) bool {
	return v >= l.GetLevel()
}

// Debug logs a message at debug level
func (l *BaseLogger) Debug(args ...any) {
	if !l.canLogAt(DebugLevel) {
		return
	}
	l.impl.Debug(fmt.Sprint(args...))
}

// Info logs a message at info level
func (l *BaseLogger) Info(args ...any) {
	if !l.canLogAt(InfoLevel) {
		return
	}
	l.impl.Info(fmt.Sprint(args...))
}

// Warn logs a message at warn level
func (l *BaseLogger) Warn(args ...any) {
	if !l.canLogAt(WarnLevel) {
		return
	}
	l.impl.Warn(fmt.Sprint(args...))
}

// Error logs a message at error level
func (l *BaseLogger) Error(args ...any) {
	if !l.canLogAt(ErrorLevel) {
		return
	}
	l.impl.Error(fmt.Sprint(args...))
}

// Fatal logs a message at fatal level and exits the program
func (l *BaseLogger) Fatal(args ...any) {
	l.impl.Fatal(fmt.Sprint(args...))
	os.Exit(1)
}

// Debugf logs a formatted message at debug level
func (l *BaseLogger) Debugf(format string, args ...any) {
	if !l.canLogAt(DebugLevel) {
		return
	}
	l.impl.Debug(fmt.Sprintf(format, args...))
}

// Infof logs a formatted message at info level
func (l *BaseLogger) Infof(format string, args ...any) {
	if !l.canLogAt(InfoLevel) {
		return
	}
	l.impl.Info(fmt.Sprintf(format, args...))
}

// Warnf logs a formatted message at warn level
func (l *BaseLogger) Warnf(format string, args ...any) {
	if !l.canLogAt(WarnLevel) {
		return
	}
	l.impl.Warn(fmt.Sprintf(format, args...))
}

// Errorf logs a formatted message at error level
func (l *BaseLogger) Errorf(format string, args ...any) {
	if !l.canLogAt(ErrorLevel) {
		return
	}
	l.impl.Error(fmt.Sprintf(format, args...))
}

// Fatalf logs a formatted message at fatal level and exits the program
func (l *BaseLogger) Fatalf(format string, args ...any) {
	l.impl.Fatal(fmt.Sprintf(format, args...))
	os.Exit(1)
}
