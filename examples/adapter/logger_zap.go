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
	"runtime"
	"strings"
	"time"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an adapter for the zap.Logger.
type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			// EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05"),
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(strings.TrimRight(jiguang.LogPrefix, " "))

				type appendTimeEncoder interface {
					AppendTimeLayout(time.Time, string)
				}

				layout := "2006/01/02 15:04:05"

				if enc, ok := enc.(appendTimeEncoder); ok {
					enc.AppendTimeLayout(t, layout)
					return
				}

				enc.AppendString(t.Format(layout))
			},
		},
	}

	baseLogger, err := config.Build()
	if err != nil {
		panic(err)
	}
	// Skip the first layer of the stack to avoid zap adapter itself being reported.
	logger := baseLogger.WithOptions(zap.AddCallerSkip(1))
	zapLogger := &ZapLogger{logger.Sugar()}

	// defer zapLogger.logger.Sync() // flushes buffer, if any
	runtime.SetFinalizer(zapLogger, func(z *ZapLogger) {
		_ = z.logger.Sync() // flushes buffer, if any
	})

	return zapLogger
}

func (z *ZapLogger) Debug(_ context.Context, msg string) {
	z.logger.Debug(msg)
}

func (z *ZapLogger) Info(_ context.Context, msg string) {
	z.logger.Info(msg)
}

func (z *ZapLogger) Warn(_ context.Context, msg string) {
	z.logger.Warn(msg)
}

func (z *ZapLogger) Error(_ context.Context, msg string) {
	z.logger.Error(msg)
}

func (z *ZapLogger) Debugf(_ context.Context, format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

func (z *ZapLogger) Infof(_ context.Context, format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z *ZapLogger) Warnf(_ context.Context, format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z *ZapLogger) Errorf(_ context.Context, format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}
