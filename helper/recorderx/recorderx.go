// Package recorderx
// @author tabuyos
// @since 2023/8/22
// @description recorderx
package recorderx

import (
	"context"
	"io"
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8)
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = slog.Level(12)
)

type Recorder interface {
	With(attrs ...any) Recorder
	WithGroup(name string) Recorder
	WithOptions(options ...*Option) Recorder

	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	// Infof uses fmt.Sprintf to log a templated message.
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	Trace(msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)

	Panic(msg string)
	MaybePanic(errs ...error)

	Context() context.Context
}

type Option struct {
	AddCallerSkip int
}

var IgnorePC = false
var defaultHandlerSupplier = func() slog.Handler {
	return NewConsoleHandler(io.MultiWriter(os.Stdout), &slog.HandlerOptions{
		Level:     LevelInfo,
		AddSource: true,
	}, nil)
}

func InitRecorder() {

}

func AddCallerSkip(skip int) *Option {
	return &Option{
		AddCallerSkip: skip,
	}
}

func WithDefault(attrs ...any) Recorder {
	return WithContext(context.Background(), nil, attrs...)
}

func DefaultRecorder() Recorder {
	return WithDefault()
}

func WithHandler(handler slog.Handler, attrs ...any) Recorder {
	return WithContext(context.Background(), func() slog.Handler {
		return handler
	}, attrs...)
}

func WithContext(ctx context.Context, handlerSupplier func() slog.Handler, attrs ...any) Recorder {
	var handler slog.Handler

	if handlerSupplier != nil {
		handler = handlerSupplier()
	} else {
		handler = defaultHandlerSupplier()
	}

	self := &commonRecorder{
		ctx:    ctx,
		logger: slog.New(handler).With(attrs...),
	}

	self.rf = self.WithOptions(AddCallerSkip(1))

	return self
}

func replaceAttr() func(groups []string, a slog.Attr) slog.Attr {
	var LevelNames = map[slog.Leveler]string{
		LevelTrace: "TRACE",
		LevelFatal: "FATAL",
	}
	return func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.LevelKey {
			level := attr.Value.Any().(slog.Level)
			levelLabel, exists := LevelNames[level]
			if !exists {
				levelLabel = level.String()
			}
			attr.Value = slog.StringValue(levelLabel)
		} else if attr.Key == slog.TimeKey {
			at := attr.Value.Time()
			attr.Value = slog.StringValue(at.Format(TimeFmt))
		}
		return attr
	}
}
