package iris

import (
	"metis/model/3.0/iris/internal/keyword"
	"strings"
)

type MRRunner func() (string, []any)
type JointType int
type Mode int

const (
	_ JointType = iota
	Join
	LeftJoin
	RightJoin
)

const (
	DftMode Mode = iota + 1
	AndMode
	OrMode
)

type Render interface {
	Render(...*strings.Builder) []any
}

type Literaler interface {
	Literal() string
}

func (j JointType) Literal() string {
	switch j {
	case Join:
		return keyword.Join.Literal()
	case LeftJoin:
		return keyword.LeftJoin.Literal()
	case RightJoin:
		return keyword.RightJoin.Literal()
	default:
		return ""
	}
}

func Once(f string, s Operator, ars ...any) *Predicate {
	pred := &Predicate{
		mod: DftMode,
		ars: ars,
	}
	//pred.buf.WriteString(f + constant.Space.Literal() + s.Literal())
	//pred.buf.WriteString(Space)
	//pred.buf.WriteString(s.Ph())
	return pred
}

type Evaluator[T any] struct {
	logical *Logical
	sqlKey  string
	ei      *EvalInfo[T]
}

func (e *Evaluator[T]) Cache(key string, pss ...PersistService[T]) *Evaluator[T] {
	e.sqlKey = key
	return e
}
