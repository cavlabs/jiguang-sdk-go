/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package adapter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is an adapter for the logrus.Logger.
type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	projectRoot, _ := filepath.Abs("..") // The parent directory of the submodule directory `example`.

	logger := logrus.New()

	logger.Level = logrus.DebugLevel
	logger.Out = os.Stderr
	logger.ExitFunc = os.Exit
	logger.ReportCaller = true
	logger.Formatter = &SDKTextFormatter{
		TextFormatter: logrus.TextFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
			PadLevelText:    true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				// Skip the first 11 layers of the stack to avoid the logrus adapter and library themselves being reported.
				_, file, line, ok := runtime.Caller(11)
				if !ok {
					file, line = frame.File, frame.Line
				}
				if rel, err := filepath.Rel(projectRoot, file); err == nil {
					file = rel
				}
				return "", fmt.Sprintf(" %s:%d", file, line)
			},
		},
		Prefix: jiguang.LogPrefix + " ",
	}

	return &LogrusLogger{logger}
}

func (l *LogrusLogger) Debug(_ context.Context, msg string) {
	l.logger.Debug(msg)
}

func (l *LogrusLogger) Info(_ context.Context, msg string) {
	l.logger.Info(msg)
}

func (l *LogrusLogger) Warn(_ context.Context, msg string) {
	l.logger.Warn(msg)
}

func (l *LogrusLogger) Error(_ context.Context, msg string) {
	l.logger.Error(msg)
}

func (l *LogrusLogger) Debugf(_ context.Context, format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *LogrusLogger) Infof(_ context.Context, format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *LogrusLogger) Warnf(_ context.Context, format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *LogrusLogger) Errorf(_ context.Context, format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

type SDKTextFormatter struct {
	logrus.TextFormatter
	Prefix string
}

func (f *SDKTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logText, err := f.TextFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	return append([]byte(f.Prefix), logText...), nil
}
