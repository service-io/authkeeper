package iris

import (
	"metis/old/model/2.0/iris/internal/keyword"
	"strings"
)

type Order struct {
	cols []string
	asc  bool
}

func render[T any](asc bool, cols ...*Column[T]) *Order {
	var snips = make([]string, len(cols))
	for i, col := range cols {
		snips[i] = col.Literal()
	}
	return &Order{
		cols: snips,
		asc:  asc,
	}
}

func (e *Evaluator[T]) Desc(cols ...*Column[T]) *Order {
	return render(false, cols...)
}

func (e *Evaluator[T]) Asc(cols ...*Column[T]) *Order {
	return render(true, cols...)
}

func (o *Order) Literal() string {
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
