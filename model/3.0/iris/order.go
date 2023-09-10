package iris

import (
	"metis/model/3.0/iris/internal/keyword"
)

type Order struct {
	col string
	asc bool
}

func (o *Order) Literal() string {
	if o == nil {
		return ""
	}
	if o.asc {
		return o.col
	}
	return o.col + " " + keyword.Desc.Literal()
}
