// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"metis/test/autogen/beiold/keyword"
	"strings"
)

type Order struct {
	cols []string
	asc  bool
}

func (t *BaseEntity[T]) Desc(cols ...string) *Order {
	return &Order{
		cols: cols,
		asc:  false,
	}
}

func (t *BaseEntity[T]) Asc(cols ...string) *Order {
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
