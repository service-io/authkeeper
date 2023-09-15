package iris

import (
	"deepsea/model/iris/internal/token"
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
	return o.col + token.Space.Join(token.Desc).Literal()
}
