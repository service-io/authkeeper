// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/model/bei/keyword"
)

type UpdateBuilder[T any] struct {
	*BaseEntity[T]
}

func (t *BaseEntity[T]) UpdateRef(ref *RefTable, fds ...*FD[T]) *UpdateBuilder[T] {
	t.fds = append(t.fds, fds...)
	t.ref = ref
	return &UpdateBuilder[T]{t}
}

func (t *BaseEntity[T]) Update(ref *RefTable) *UpdateBuilder[T] {
	t.ref = ref
	return &UpdateBuilder[T]{t}
}

func (t *UpdateBuilder[T]) Set(fd *FD[T], v any) *UpdateBuilder[T] {
	t.fds = append(t.fds, fd)
	t.values = append(t.values, v)
	return t
}

func (t *UpdateBuilder[T]) SetValues(values ...any) *UpdateBuilder[T] {
	t.values = append(t.values, values...)
	return t
}

func (t *UpdateBuilder[T]) Where(pd *Predicate) *UpdateBuilder[T] {
	t.where = pd
	return t
}

func (t *UpdateBuilder[T]) Persist() *Persist[T] {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column.")
	}

	var (
		sql    string
		values []any
	)

	setSQL := t.getSetSQL()
	fromSQL := t.getFromSQL()
	whereSQL, whereValues := t.getWhereSQL()
	logicDeletedSQL := t.getLogicDeletedSQL()

	t.buf.WriteString(keyword.Update.Literal())
	t.write(fromSQL)
	t.write(setSQL, keyword.Set.Literal())
	if len(logicDeletedSQL) > 0 {
		t.write(LeftParentheses+whereSQL+RightParentheses, keyword.Where.Literal())
		t.write(keyword.And.Literal())
		t.write(logicDeletedSQL)
	} else {
		t.write(whereSQL, keyword.Where.Literal())
	}
	t.buf.WriteString(";")

	sql = t.buf.String()
	values = append(t.values, whereValues...)
	t.values = values

	return OfPersist[T](sql, "", values, nil)
}
