// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/model/bei/keyword"
)

type InsertBuilder[T any] struct {
	*BaseEntity[T]
}

func (t *BaseEntity[T]) Insert(fds ...*FD[T]) *InsertBuilder[T] {
	t.fds = append(t.fds, fds...)
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

func (t *InsertBuilder[T]) Persist() *Persist[T] {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column.")
	}

	if len(t.values)%len(t.fds) != 0 {
		panic("parameter not match.")
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

	return OfPersist[T](sql, "", values, nil)
}
