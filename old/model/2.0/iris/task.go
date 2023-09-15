package iris

import "strings"

type Task func(...*strings.Builder) []any

func (task Task) Idle(buffers ...*strings.Builder) []any {
	if task == nil {
		return nil
	}
	return task(buffers...)
}
