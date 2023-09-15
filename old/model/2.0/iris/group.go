package iris

import (
	"metis/old/model/2.0/iris/internal/constant"
	"strings"
)

type Group struct {
	cols []string
}

func (g *Group) Literal() string {
	if g == nil {
		return ""
	}
	return strings.Join(g.cols, constant.CommaSpace.Literal())
}

func (e *Evaluator[T]) Group(cols ...*Column[T]) *Group {
	var snips = make([]string, len(cols))
	for i, col := range cols {
		snips[i] = col.Literal()
	}
	return &Group{
		cols: snips,
	}
}
