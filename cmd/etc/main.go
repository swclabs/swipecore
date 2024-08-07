// Package main test
package main

import (
	"fmt"
	"swclabs/swix/pkg/lib/logger"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SyslogTimeEncoder is a custom time encoder
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateOnly))
}

// CustomLevelEncoder is a custom level encoder
func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Blue.Add(level.CapitalString())))
		return
	case zapcore.DebugLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Magenta.Add(level.CapitalString())))
		return
	case zapcore.WarnLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Yellow.Add(level.CapitalString())))
		return
	}
	enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Red.Add(level.CapitalString())))
}

func main() {
	// Tạo encoder config cho console
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CustomLevelEncoder,
		EncodeTime:     SyslogTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger.Info("Test")
	logger.Debug("Test")
	logger.Warn("Test")
}
