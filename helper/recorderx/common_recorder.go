// Package recorderx
// @author tabuyos
// @since 2023/8/22
// @description recorderx
package recorderx

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"time"
)

type commonRecorder struct {
	ctx     context.Context
	logger  *slog.Logger
	options []*Option
	rf      Recorder
}

func (rec *commonRecorder) With(attrs ...any) Recorder {
	return &commonRecorder{
		ctx:     rec.ctx,
		logger:  rec.logger.With(attrs...),
		options: rec.options,
		rf:      rec.rf,
	}
}

func (rec *commonRecorder) WithGroup(name string) Recorder {
	return &commonRecorder{
		ctx:     rec.ctx,
		logger:  rec.logger.WithGroup(name),
		options: rec.options,
		rf:      rec.rf,
	}
}

func (rec *commonRecorder) WithOptions(options ...*Option) Recorder {
	return &commonRecorder{
		ctx:     rec.ctx,
		logger:  rec.logger,
		options: options,
		rf:      rec.rf,
	}
}

func (rec *commonRecorder) Tracef(format string, args ...any) {
	rec.rf.Trace(fmt.Sprintf(format, args...))
}

func (rec *commonRecorder) Debugf(format string, args ...any) {
	rec.rf.Debug(fmt.Sprintf(format, args...))
}

// Infof uses fmt.Sprintf to log a templated message.
func (rec *commonRecorder) Infof(format string, args ...any) {
	rec.rf.Info(fmt.Sprintf(format, args...))
}

func (rec *commonRecorder) Warnf(format string, args ...any) {
	rec.rf.Warn(fmt.Sprintf(format, args...))
}

func (rec *commonRecorder) Errorf(format string, args ...any) {
	rec.rf.Error(fmt.Sprintf(format, args...))
}

func (rec *commonRecorder) Fatalf(format string, args ...any) {
	rec.rf.Fatal(fmt.Sprintf(format, args...))
}

func (rec *commonRecorder) Trace(msg string) {
	rec.log(rec.ctx, LevelTrace, msg)
}

func (rec *commonRecorder) Debug(msg string) {
	rec.log(rec.ctx, LevelDebug, msg)
}

func (rec *commonRecorder) Info(msg string) {
	rec.log(rec.ctx, LevelInfo, msg)
}

func (rec *commonRecorder) Warn(msg string) {
	rec.log(rec.ctx, LevelWarn, msg)
}

func (rec *commonRecorder) Error(msg string) {
	rec.log(rec.ctx, LevelError, msg)
}

func (rec *commonRecorder) Fatal(msg string) {
	rec.log(rec.ctx, LevelFatal, msg)
}

func (rec *commonRecorder) MaybePanic(errs ...error) {
	for _, err := range errs {
		if err != nil {
			rec.log(rec.ctx, LevelError, err.Error())
			panic(err)
		}
	}
}

func (rec *commonRecorder) Panic(msg string) {
	if len(msg) != 0 {
		rec.log(rec.ctx, LevelError, msg)
		panic(msg)
	}
}

func (rec *commonRecorder) Context() context.Context {
	return rec.ctx
}

func (rec *commonRecorder) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !rec.logger.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	//goland:noinspection GoBoolExpressions
	if !IgnorePC {
		skip := 3
		if len(rec.options) > 0 {
			option := rec.options[0]
			skip += option.AddCallerSkip
		}
		var pcs [1]uintptr
		// skip [runtime.Callers, this function, this function's caller]
		runtime.Callers(skip, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = rec.logger.Handler().Handle(ctx, r)
}
