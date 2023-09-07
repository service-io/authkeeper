// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/model/bei/keyword"
)

type DeleteBuilder[T any] struct {
	*BaseEntity[T]
}

func (t *BaseEntity[T]) Delete() *DeleteBuilder[T] {
	return &DeleteBuilder[T]{t}
}

func (t *DeleteBuilder[T]) From(rt *RefTable) *DeleteBuilder[T] {
	t.ref = rt
	return t
}

func (t *DeleteBuilder[T]) Where(pd *Predicate) *DeleteBuilder[T] {
	t.where = pd
	return t
}

func (t *DeleteBuilder[T]) Persist() *Persist[T] {
	t.buf.Reset()

	var (
		sql    string
		values []any
	)

	if t.logicDeleted {
		setSQL := t.deletedKey + " = " + t.deletedVal
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
		values = whereValues
	} else {
		fromSQL := t.getFromSQL()
		whereSQL, whereValues := t.getWhereSQL()
		logicDeletedSQL := t.getLogicDeletedSQL()

		t.buf.WriteString(keyword.Delete.Literal())
		t.write(fromSQL, keyword.From.Literal())
		if len(logicDeletedSQL) > 0 {
			t.write(LeftParentheses+whereSQL+RightParentheses, keyword.Where.Literal())
			t.write(keyword.And.Literal())
			t.write(logicDeletedSQL)
		} else {
			t.write(whereSQL, keyword.Where.Literal())
		}
		t.buf.WriteString(";")

		sql = t.buf.String()
		values = whereValues
	}
	t.values = values
	return OfPersist[T](sql, "", values, nil)
}
