// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

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

func (t *Evaluator[T]) Group(cols ...string) *Group {
	return &Group{
		cols: cols,
	}
}
