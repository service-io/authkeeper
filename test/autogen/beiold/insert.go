// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"metis/test/autogen/beiold/keyword"
	"strings"
)

func (t *BaseEntity[T]) Insert(fds ...*FD[T]) *BaseEntity[T] {
	t.fds = append(t.fds, fds...)
	return t
}

func (t *BaseEntity[T]) Into(ref *RefTable) *BaseEntity[T] {
	t.ref = ref
	return t
}

func (t *BaseEntity[T]) InsertSQL() (sql string, values []any) {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column.")
	}

	if len(t.values)%len(t.fds) != 0 {
		panic("parameter not match.")
	}

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

	return
}

func (t *BaseEntity[T]) getIntoSQL() string {
	return t.ref.SQL()
}

func (t *BaseEntity[T]) getInsertPlaceholder() string {
	var snips = make([]string, len(t.fds))
	for i := 0; i < len(t.fds); i++ {
		snips[i] = "?"
	}
	return strings.Join(snips, ", ")
}

func (t *BaseEntity[T]) getTimesPlaceholder(times int, ph string) string {
	var snips = make([]string, times)
	for i := 0; i < times; i++ {
		snips[i] = LeftParentheses + ph + RightParentheses
	}
	return strings.Join(snips, ", ")
}
