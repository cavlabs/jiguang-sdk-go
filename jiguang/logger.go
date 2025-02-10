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

package jiguang

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// Logger 是 SDK 使用的日志记录器的接口。
type Logger interface {
	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)

	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}

// =====================================================================================================================

// 用于在终端输出彩色日志的 ANSI 转义码。
const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34m" // DEBUG
	colorGreen  = "\033[32m" // INFO
	colorYellow = "\033[33m" // WARN
	colorRed    = "\033[31m" // ERROR
)

// StdLogger 使用标准库的 log 包实现了 Logger 接口。
type StdLogger struct {
	*log.Logger
}

func NewStdLogger(opts ...StdLoggerOption) *StdLogger {
	logger := &StdLogger{log.New(os.Stdout, "[JGSDK] ", log.LstdFlags)}
	for _, opt := range opts {
		_ = opt(logger)
	}
	return logger
}

// ---------------------------------------------------------------------------------------------------------------------

// 用于配置 StdLogger 的选项。
type StdLoggerOption func(*StdLogger) error

// 自定义设置日志输出，默认为 os.Stdout。
func WithLogOutput(logOutput io.Writer) StdLoggerOption {
	return func(logger *StdLogger) error {
		logger.SetOutput(logOutput)
		return nil
	}
}

// 自定义设置日志前缀，默认为 "[JGSDK] "。
func WithLogPrefix(logPrefix string) StdLoggerOption {
	return func(logger *StdLogger) error {
		logger.SetPrefix(logPrefix)
		return nil
	}
}

// 自定义设置日志标志，默认为 log.LstdFlags。
func WithLogFlags(logFlags int) StdLoggerOption {
	return func(logger *StdLogger) error {
		logger.SetFlags(logFlags)
		return nil
	}
}

// ---------------------------------------------------------------------------------------------------------------------

var ModPathRegex = regexp.MustCompile(`github\.com/[^@]+(@[^/]+)?/.+`) // 匹配模块路径

// 格式化带有日志级别、文件路径、行号和颜色的日志消息，然后打印它。
func (s *StdLogger) logMessage(_ context.Context, level, color, msg string) {
	_, file, line, _ := runtime.Caller(2)
	matches := ModPathRegex.FindStringSubmatch(file)
	if len(matches) > 0 {
		file = matches[0]
	} else {
		file = filepath.Base(file)
	}
	s.Printf("%s%-5s%s %s:%d %s", color, level, colorReset, file, line, msg)
}

func (s *StdLogger) Debug(ctx context.Context, msg string) {
	s.logMessage(ctx, "DEBUG", colorBlue, msg)
}

func (s *StdLogger) Info(ctx context.Context, msg string) {
	s.logMessage(ctx, "INFO", colorGreen, msg)
}

func (s *StdLogger) Warn(ctx context.Context, msg string) {
	s.logMessage(ctx, "WARN", colorYellow, msg)
}

func (s *StdLogger) Error(ctx context.Context, msg string) {
	s.logMessage(ctx, "ERROR", colorRed, msg)
}

func (s *StdLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	s.logMessage(ctx, "DEBUG", colorBlue, fmt.Sprintf(format, args...))
}

func (s *StdLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	s.logMessage(ctx, "INFO", colorGreen, fmt.Sprintf(format, args...))
}

func (s *StdLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	s.logMessage(ctx, "WARN", colorYellow, fmt.Sprintf(format, args...))
}

func (s *StdLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	s.logMessage(ctx, "ERROR", colorRed, fmt.Sprintf(format, args...))
}
