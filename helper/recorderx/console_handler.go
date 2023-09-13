// Package recorderx
// @author tabuyos
// @since 2023/8/22
// @description recorderx
package recorderx

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

type consoleHandler struct {
	opts        slog.HandlerOptions
	writer      io.Writer
	attrs       []slog.Attr
	group       *string
	jsonHandler slog.Handler
	mu          *sync.Mutex
}

const (
	TimeFmt = "2006-01-02T15:04:05.000"
)

func NewConsoleHandler(writer io.Writer, opts *slog.HandlerOptions, jsonHandler slog.Handler) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	return &consoleHandler{
		opts:        *opts,
		writer:      writer,
		attrs:       nil,
		group:       nil,
		jsonHandler: jsonHandler,
		mu:          &sync.Mutex{},
	}
}

type handleContent struct {
	buf bytes.Buffer
	sep string
}

func (h *handleContent) appendTime(ts time.Time) {
	h.buf.WriteString("")
}

func (h *handleContent) appendAttr(attr slog.Attr) {
	h.buf.WriteString(attr.Value.String())
}

func (h *handleContent) appendError(err error) {
	h.buf.WriteString(err.Error())
}

func (h *handleContent) appendString(str string) {
	h.buf.WriteString(str)
}

func (h *handleContent) appendValue(val slog.Value) {
	h.buf.WriteString(val.String())
}

func (h *consoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

var contentPool = &sync.Pool{
	New: func() any {
		return &handleContent{
			sep: " ",
			buf: bytes.Buffer{},
		}
	},
}

func getContent() (*handleContent, func()) {
	content := contentPool.Get().(*handleContent)
	rel := func() {
		content.buf.Reset()
	}
	return content, rel
}

func (h *consoleHandler) Handle(ctx context.Context, r slog.Record) error {
	content, release := getContent()
	defer release()
	// time
	if !r.Time.IsZero() {
		val := r.Time.Round(0)
		content.appendTime(val)
		content.buf.WriteString(content.sep)
	}

	// level
	val := r.Level
	content.appendString(renderLevel(val))
	content.buf.WriteString(content.sep)

	// source
	if h.opts.AddSource {
		content.appendAttr(slog.Any(slog.SourceKey, simpleSource(source(r))))
		// content.buf.WriteString(":")
		content.buf.WriteString(content.sep)
	}

	// message
	msg := r.Message
	content.appendString(msg)

	for _, attr := range h.attrs {
		content.buf.WriteString(content.sep)
		content.buf.WriteString(attr.Value.String())
	}
	content.buf.WriteString("\n")

	_, err := h.writer.Write(content.buf.Bytes())

	go func() {
		if h.jsonHandler != nil {
			r.AddAttrs(h.attrs...)
			_ = h.jsonHandler.Handle(ctx, r)
		}
	}()

	return err
}

func (h *consoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := h.clone()
	if h2.group != nil {
		name := *h2.group
		h2.attrs = append(h2.attrs, slog.Group(name, toAny(attrs)...))
	} else {
		h2.attrs = append(h2.attrs, attrs...)
	}
	return h2
}

func (h *consoleHandler) WithGroup(name string) slog.Handler {
	h2 := h.clone()
	h2.group = &name
	return h2
}

func (h *consoleHandler) clone() *consoleHandler {
	return &consoleHandler{
		opts:        h.opts,
		writer:      h.writer,
		attrs:       h.attrs,
		group:       h.group,
		jsonHandler: h.jsonHandler,
		mu:          h.mu,
	}
}

func toAny(items []slog.Attr) []any {
	var rs = make([]any, len(items))
	for i, item := range items {
		rs[i] = item
	}
	return rs
}

func source(r slog.Record) *slog.Source {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	return &slog.Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}

const sourceLen = 40

var sourcePlh = func() string {
	ss := make([]rune, sourceLen)
	for i := 0; i < sourceLen; i++ {
		ss[i] = ' '
	}
	return string(ss)
}()

func simpleSource(s *slog.Source) string {
	return ""
}

var ()

func renderLevel(l slog.Level) (level string) {
	return
}
