package iris

import "metis/model/3.0/iris/internal/keyword"

type Operator func(bool) string

func (op Operator) String() string {
	return op.Self()
}

func (op Operator) Literal() string {
	return op(true)
}

func (op Operator) Self() string {
	return op(false)
}

func AndOp(bool) string {
	return keyword.And.Literal()
}

func OrOp(bool) string {
	return keyword.Or.Literal()
}

func EqOp(has bool) string {
	if has {
		return "= ?"
	}
	return "="
}

func NqOp(has bool) string {
	if has {
		return "<> ?"
	}
	return "<>"
}
