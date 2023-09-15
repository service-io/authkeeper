// Package recorderx
// @author tabuyos
// @since 2023/8/23
// @description recorderx
package recorderx

import (
	"context"
	"deepsea/config/constant"
	"io"
	"slices"
	"strconv"
	"sync"
	"time"
)

type Mode int

func (m *Mode) String() string {
	if m == nil {
		return ""
	}
	return strconv.Itoa(int(*m))
}

const (
	CommonMode Mode = iota
	VisitorMode
	OperateMode
)

var (
	QueueSize = 10000
	queue     = make(chan Entry, QueueSize)
)

type Entry struct {
	content []byte
	mode    Mode
	sign    *string
}

func (e *Entry) BytesToString() string {
	if len(e.content) == 0 {
		return ""
	}
	return string(e.content)
}

func (e *Entry) Bytes() []byte {
	return e.content
}

func (e *Entry) Mode() Mode {
	return e.mode
}

func (e *Entry) Sign() string {
	if e.sign == nil {
		return ""
	}
	return *e.sign
}

func NewChanWriter(ctx context.Context, mode Mode) io.Writer {
	return &chanWriter{ctx: ctx, mode: mode, mu: &sync.Mutex{}}
}

func QueueChan() <-chan Entry {
	return queue
}

type chanWriter struct {
	ctx  context.Context
	sign *string
	mode Mode
	mu   *sync.Mutex
}

func (w *chanWriter) Write(bytes []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.sign == nil {
		value := w.ctx.Value(constant.RecorderDeliverKey)
		deliver, ok := value.(*deliver)
		if ok {
			sign := deliver.signSupplier(deliver.ginCtx)
			w.sign = &sign
		}
	}

	// 高性能, 但是无序
	// 需要有序处理, 可修改为非协程模式
	go func(bytes []byte) {
		entry := Entry{
			content: bytes,
			mode:    w.mode,
			sign:    w.sign,
		}
		timer := time.NewTimer(5 * time.Second)
		select {
		case queue <- entry:
		case <-timer.C:
			DefaultRecorder().Warn("信道超时!")
			break
		}
	}(slices.Clone(bytes))

	return 0, nil
}
