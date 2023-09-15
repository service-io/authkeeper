// Package iris
// @author tabuyos
// @since 2023/9/11
// @description iris
package iris

import (
	"strings"
	"sync"
)

var builderPool sync.Pool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

func getBuilder() (*strings.Builder, func()) {
	builder := builderPool.Get().(*strings.Builder)
	release := func() {
		builder.Reset()
		builderPool.Put(builder)
	}
	return builder, release
}
