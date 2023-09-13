// Package recorderx
// @author tabuyos
// @since 2023/8/23
// @description recorderx
package recorderx

import (
	"io"
)

func NewEmptyWriter() io.Writer {
	return &emptyWriter{}
}

type emptyWriter struct {
}

func (w *emptyWriter) Write(_ []byte) (int, error) {
	return 0, nil
}
