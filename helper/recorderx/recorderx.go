// Package recorderx
// @author tabuyos
// @since 2023/8/22
// @description recorderx
package recorderx

import (
	"context"
	"deepsea/config/constant"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"path"
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

var accessFile = &lumberjack.Logger{
	Filename:   path.Join(constant.LogDir, "access.log"),
	MaxSize:    20,
	MaxBackups: 0,
	MaxAge:     30,
	Compress:   false,
}
var IgnorePC = false
var defaultHandlerSupplier = func() slog.Handler {
	return NewConsoleHandler(io.MultiWriter(accessFile, os.Stdout), &slog.HandlerOptions{
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

func FetchRecorder(ctx *gin.Context) Recorder {
	if ctx == nil {
		return DefaultRecorder()
	}
	recorder, ok := ctx.Get(constant.RecorderGinKey)
	if ok {
		return recorder.(Recorder)
	}
	return DefaultRecorder()
}

func FetchVisitor(ctx *gin.Context) Recorder {
	if ctx == nil {
		return DefaultRecorder()
	}
	recorder, ok := ctx.Get(constant.RecorderVisitorGinKey)
	if ok {
		return recorder.(Recorder)
	}
	return FetchRecorder(ctx)
}

func FetchOperate(ctx *gin.Context) Recorder {
	if ctx == nil {
		return DefaultRecorder()
	}
	recorder, ok := ctx.Get(constant.RecorderOperateGinKey)
	if ok {
		return recorder.(Recorder)
	}
	return DefaultRecorder()
}

func WithGinContext(ctx *gin.Context, signSupplier SignService) {
	WithIntoGinContext(ctx, signSupplier)
	WithVisitorIntoGinContext(ctx, signSupplier)
	WithOperateIntoGinContext(ctx, signSupplier)
}

func WithVisitorIntoGinContext(ctx *gin.Context, signSupplier SignService) {
	withGinContext(ctx, constant.RecorderVisitorGinKey, VisitorMode, signSupplier)
}

func WithOperateIntoGinContext(ctx *gin.Context, signSupplier SignService) {
	withGinContext(ctx, constant.RecorderOperateGinKey, OperateMode, signSupplier)
}

func WithIntoGinContext(ctx *gin.Context, signSupplier SignService) {
	withGinContext(ctx, constant.RecorderGinKey, CommonMode, signSupplier)
}

func jsonChanHandler(ctx context.Context, mode Mode) slog.Handler {
	return slog.NewJSONHandler(NewChanWriter(ctx, mode), &slog.HandlerOptions{AddSource: true, ReplaceAttr: replaceAttr()})
}

func withGinContext(ctx *gin.Context, key string, mode Mode, signSupplier SignService) {
	var jsonHandler slog.Handler = nil
	if mode != CommonMode {
		deliver := &deliver{
			ginCtx:       ctx,
			signSupplier: signSupplier,
		}
		writerCtx := context.WithValue(context.Background(), constant.RecorderDeliverKey, deliver)
		jsonHandler = jsonChanHandler(writerCtx, mode)
	}
	traceID := ctx.GetString(constant.TraceIdKey)
	handler := NewConsoleHandler(io.MultiWriter(accessFile, os.Stdout), &slog.HandlerOptions{Level: LevelInfo, AddSource: true}, jsonHandler)

	ginRecorder := WithHandler(handler, slog.String(constant.TraceIdKey, traceID))
	ctx.Set(key, ginRecorder)
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
