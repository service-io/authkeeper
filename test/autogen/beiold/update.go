// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"metis/test/autogen/beiold/keyword"
	"strings"
)

func (t *BaseEntity[T]) Update(ref *RefTable, fds ...*FD[T]) *BaseEntity[T] {
	t.fds = append(t.fds, fds...)
	t.ref = ref
	return t
}

func (t *BaseEntity[T]) WithValues(values ...any) {
	t.values = append(t.values, values...)
}

func (t *BaseEntity[T]) UpdateSQL(times ...int) (sql string, values []any) {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column.")
	}

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

	return
}

func (t *BaseEntity[T]) getSetSQL() string {
	var snips = make([]string, len(t.fds))
	for i, update := range t.fds {
		snips[i] = update.Fd() + " = ?"
	}
	return strings.Join(snips, ", ")
}
