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
	"strconv"
	"strings"
	"time"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
	"github.com/rs/zerolog"
)

// ZeroLogger is an adapter for the zerolog.Logger.
type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatLevel: func(i interface{}) string {
			var l string
			if ll, ok := i.(string); ok {
				level := strings.ToUpper(ll)
				switch ll {
				case zerolog.LevelDebugValue:
					l = colorize(level, 33) // yellow
				case zerolog.LevelInfoValue:
					l = colorize(level, 32) // green
				case zerolog.LevelWarnValue:
					l = colorize(level, 31) // red
				case zerolog.LevelErrorValue:
					l = colorize(colorize(level, 31), 1) // red, bold
				case zerolog.LevelFatalValue:
					l = colorize(colorize(level, 31), 1) // red, bold
				case zerolog.LevelPanicValue:
					l = colorize(colorize(level, 31), 1) // red, bold
				default:
					l = colorize(level, 1) // bold
				case zerolog.LevelTraceValue:
					l = colorize(level, 35) // magenta
				}
			} else {
				if i == nil {
					l = colorize("???", 1) // bold
				} else {
					l = strings.ToUpper(fmt.Sprintf("%s", i))
				}
			}
			return l
		},
		// TimeFormat: "2006/01/02 15:04:05",
		FormatTimestamp: func(i interface{}) string {
			t, tt := "<nil>", i.(string)
			if ts, err := time.ParseInLocation(zerolog.TimeFieldFormat, tt, time.Local); err != nil {
				t = tt
			} else {
				t = ts.Local().Format("2006/01/02 15:04:05")
			}
			return jiguang.LogPrefix + t
		},
	}

	projectRoot, _ := filepath.Abs("..") // The parent directory of the submodule directory `example`.
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		if rel, err := filepath.Rel(projectRoot, file); err == nil {
			file = rel
		}
		return file + ":" + strconv.Itoa(line)
	}

	logger := zerolog.New(consoleWriter).With().Timestamp().CallerWithSkipFrameCount(3).Logger()

	return &ZeroLogger{logger}
}

// colorize returns the string s wrapped in ANSI code c.
func colorize(s interface{}, c int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

func (z *ZeroLogger) Debug(_ context.Context, msg string) {
	z.logger.Debug().Msg(msg)
}

func (z *ZeroLogger) Info(_ context.Context, msg string) {
	z.logger.Info().Msg(msg)
}

func (z *ZeroLogger) Warn(_ context.Context, msg string) {
	z.logger.Warn().Msg(msg)
}

func (z *ZeroLogger) Error(_ context.Context, msg string) {
	z.logger.Error().Msg(msg)
}

func (z *ZeroLogger) Debugf(_ context.Context, format string, args ...interface{}) {
	z.logger.Debug().Msgf(format, args...)
}

func (z *ZeroLogger) Infof(_ context.Context, format string, args ...interface{}) {
	z.logger.Info().Msgf(format, args...)
}

func (z *ZeroLogger) Warnf(_ context.Context, format string, args ...interface{}) {
	z.logger.Warn().Msgf(format, args...)
}

func (z *ZeroLogger) Errorf(_ context.Context, format string, args ...interface{}) {
	z.logger.Error().Msgf(format, args...)
}
