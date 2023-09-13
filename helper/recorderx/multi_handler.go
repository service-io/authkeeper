// Package recorderx
// @author tabuyos
// @since 2023/8/23
// @description recorderx
package recorderx

import (
	"context"
	"errors"
	"log/slog"
)

type multiHandler struct {
	handlers []slog.Handler
	opts     slog.HandlerOptions
}

func NewMultiHandler(opts *slog.HandlerOptions, handlers ...slog.Handler) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &multiHandler{handlers: handlers, opts: *opts}
}

func (m *multiHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := LevelInfo
	if m.opts.Level != nil {
		minLevel = m.opts.Level.Level()
	}
	return level >= minLevel
}

func (m *multiHandler) Handle(ctx context.Context, record slog.Record) error {
	var err error
	for _, handler := range m.handlers {
		err = errors.Join(err, handler.Handle(ctx, record))
	}
	return err
}

func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var hs = make([]slog.Handler, len(m.handlers))
	for i, handler := range m.handlers {
		hs[i] = handler.WithAttrs(attrs)
	}

	return NewMultiHandler(&m.opts, hs...)
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	var hs = make([]slog.Handler, len(m.handlers))
	for i, handler := range m.handlers {
		hs[i] = handler.WithGroup(name)
	}

	return NewMultiHandler(&m.opts, hs...)
}
