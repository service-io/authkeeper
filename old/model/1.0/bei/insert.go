// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/old/model/1.0/bei/keyword"
)

type InsertBuilder[T any] struct {
	*Evaluator[T]
}

func (t *Evaluator[T]) Insert(fds ...*FD[T]) *InsertBuilder[T] {
	t.fds = fds
	return &InsertBuilder[T]{t}
}

func (t *InsertBuilder[T]) Into(ref *RefTable) *InsertBuilder[T] {
	t.ref = ref
	return t
}

func (t *InsertBuilder[T]) Values(values ...any) *InsertBuilder[T] {
	t.values = append(t.values, values...)
	return t
}

func (t *InsertBuilder[T]) Value(values ...any) *InsertBuilder[T] {
	t.values = values
	return t
}

func (t *InsertBuilder[T]) WithSQLKey(key string) *InsertBuilder[T] {
	t.sqlKey = key
	return t
}

func (t *InsertBuilder[T]) WithLogicDeleted(cdv ...string) *InsertBuilder[T] {
	if !t.enableLogical() {
		t.logical.Enable()
	}
	// using last value
	for _, v := range cdv {
		t.logical.cdval = v
	}
	return t
}

func (t *InsertBuilder[T]) Eval(pss ...PersistService[T]) EvalInfoService[T] {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column.")
	}

	if len(t.values)%len(t.fds) != 0 {
		panic("parameter not match.")
	}

	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		lookup := ps.Lookup(t.sqlKey)
		if lookup != nil {
			t.ei = OfEvalInfo[T](lookup.SQL(), "", t.values, nil)
			return t.ei
		}
	}

	var (
		sql    string
		values []any
	)

	quotient := len(t.values) / len(t.fds)

	placeholder := t.getInsertPlaceholder()
	fieldSQL, _ := t.getFieldSQL()
	intoSQL := t.getIntoSQL()
	timesPlaceholder := t.getTimesPlaceholder(quotient, placeholder)

	if len(fieldSQL) == 0 {
		panic("not found into column")
	}

	t.buf.WriteString(keyword.Insert.Literal())
	t.write(intoSQL, keyword.Into.Literal())
	t.buf.WriteString(LeftParentheses)
	t.buf.WriteString(fieldSQL)
	t.buf.WriteString(RightParentheses)
	t.write(timesPlaceholder, keyword.Values.Literal())
	t.buf.WriteString(";")

	values = t.values
	sql = t.buf.String()

	t.ei = OfEvalInfo[T](sql, "", values, nil)
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		ps.Persistence(t.sqlKey, OfEvalInfo[T](sql, "", nil, nil))
	}
	return t.ei
}
