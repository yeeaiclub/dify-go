// Copyright 2025 yeeaiclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func Global() Logger {
	return globalLogger
}

func SetGlobalLogger(logger Logger) {
	globalMu.Lock()
	defer globalMu.Unlock()
	globalLogger = logger
}

func Debug(args ...any)                 { globalLogger.Debug(args...) }
func Info(args ...any)                  { globalLogger.Info(args...) }
func Warn(args ...any)                  { globalLogger.Warn(args...) }
func Error(args ...any)                 { globalLogger.Error(args...) }
func Fatal(args ...any)                 { globalLogger.Fatal(args...) }
func Debugf(format string, args ...any) { globalLogger.Debugf(format, args...) }
func Infof(format string, args ...any)  { globalLogger.Infof(format, args...) }
func Warnf(format string, args ...any)  { globalLogger.Warnf(format, args...) }
func Errorf(format string, args ...any) { globalLogger.Errorf(format, args...) }
func Fatalf(format string, args ...any) { globalLogger.Fatalf(format, args...) }

func SetLevel(level Level) {
	globalLogger.SetLevel(level)
}
