package util

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"time"
)

func InitLogger() *zap.Logger {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    "func",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"}, // 输出到标准输出
		ErrorOutputPaths: []string{"stdout"}, // 错误信息输出到标准错误输出
	}

	// 构建日志记录器
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

type ZapGormLogger struct {
	Logger *zap.Logger
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Sugar().Infof(msg, data...)
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Sugar().Warnf(msg, data...)
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Sugar().Errorf(msg, data...)
}
