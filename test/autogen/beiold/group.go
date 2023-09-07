// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import "strings"

type Group struct {
	cols []string
}

func (g *Group) SQL() string {
	if g == nil {
		return ""
	}
	return strings.Join(g.cols, Space)
}

func (t *BaseEntity[T]) Group(cols ...string) *Group {
	return &Group{
		cols: cols,
	}
}
