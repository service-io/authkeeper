// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import "metis/test/autogen/beiold/keyword"

func (t *BaseEntity[T]) Delete(fds ...*FD[T]) *BaseEntity[T] {
	t.fds = append(t.fds, fds...)
	return t
}

func (t *BaseEntity[T]) DeleteSQL() (sql string, values []any) {
	t.buf.Reset()

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
	return
}
