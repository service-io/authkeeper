// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/model/bei/keyword"
	"strings"
)

type Order struct {
	cols []string
	asc  bool
}

func (t *Evaluator[T]) Desc(cols ...string) *Order {
	return &Order{
		cols: cols,
		asc:  false,
	}
}

func (t *Evaluator[T]) Asc(cols ...string) *Order {
	return &Order{
		cols: cols,
		asc:  true,
	}
}

func (o *Order) SQL() string {
	if o == nil {
		return ""
	}
	var snips = make([]string, len(o.cols))
	for i, col := range o.cols {
		if !o.asc {
			col += " " + keyword.Desc.Literal()
		}
		snips[i] = col
	}
	return strings.Join(snips, ", ")
}
