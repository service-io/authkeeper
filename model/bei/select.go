// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

import (
	"metis/model/bei/keyword"
	"strings"
)

type QueryBuilder[T any] struct {
	*BaseEntity[T]
}

func (t *BaseEntity[T]) Select(fds ...*FD[T]) *QueryBuilder[T] {
	t.fds = append(t.fds, fds...)
	return &QueryBuilder[T]{t}
}

func (t *QueryBuilder[T]) Hint(hints ...keyword.Keyword) *QueryBuilder[T] {
	t.hints = append(t.hints, hints...)
	return t
}

func (t *QueryBuilder[T]) From(rt *RefTable) *QueryBuilder[T] {
	t.ref = rt
	return t
}

func (t *QueryBuilder[T]) Where(pd *Predicate) *QueryBuilder[T] {
	t.where = pd
	return t
}

func (t *QueryBuilder[T]) GroupBy(gbs ...*Group) *QueryBuilder[T] {
	t.groups = gbs
	return t
}

func (t *QueryBuilder[T]) Having(pd *Predicate) *QueryBuilder[T] {
	t.having = pd
	return t
}

func (t *QueryBuilder[T]) OrderBy(obs ...*Order) *QueryBuilder[T] {
	t.orders = obs
	return t
}

func (t *QueryBuilder[T]) Limit(lmt int64) *QueryBuilder[T] {
	t.limit = lmt
	return t
}

func (t *QueryBuilder[T]) Offset(ost int64) *QueryBuilder[T] {
	t.offset = ost
	return t
}

func (t *QueryBuilder[T]) WithLogicDeleted(cdv ...string) *QueryBuilder[T] {
	t.logicDeleted = true
	if len(cdv) == 1 {
		t.currentDeletedVal = cdv[0]
	}
	return t
}

func (t *QueryBuilder[T]) Persist() *Persist[T] {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column")
	}
	var fieldSQL string

	var (
		execSQL  string
		totalSQL string
		values   []any
		mappers  []func(*T) any
	)

	hintSQL := t.getHintSQL()
	fieldSQL, mappers = t.getFieldSQL()
	fromSQL := t.getFromSQL()
	whereSQL, whereValues := t.getWhereSQL()
	logicDeletedSQL := t.getLogicDeletedSQL()
	groupBySQL := t.getGroupBySQL()
	havingSQL, havingValues := t.getHavingSQL()
	orderBySQL := t.getOrderBySQL()
	pageSQL := t.getPageSQL()

	enablePage := len(pageSQL) > 0

	var ts strings.Builder
	if enablePage {
		ts.WriteString(keyword.Select.Literal())
		ts.WriteString(Space)
		ts.WriteString(keyword.Count.Literal())
		ts.WriteString("(")
		t.writeToBuf(&ts, hintSQL)
		ts.WriteString("1")
		ts.WriteString(")")
	}

	t.buf.WriteString(keyword.Select.Literal())
	t.write(hintSQL)
	t.write(fieldSQL)
	t.writeAppend(&ts, fromSQL, keyword.From.Literal())
	if len(logicDeletedSQL) > 0 {
		t.writeAppend(&ts, LeftParentheses+whereSQL+RightParentheses, keyword.Where.Literal())
		t.writeAppend(&ts, keyword.And.Literal())
		t.writeAppend(&ts, logicDeletedSQL)
	} else {
		t.writeAppend(&ts, whereSQL, keyword.Where.Literal())
	}
	t.writeAppend(&ts, groupBySQL, keyword.Group.Literal(), keyword.By.Literal())
	t.writeAppend(&ts, havingSQL, keyword.Having.Literal())
	if len(pageSQL) > 0 {
		ts.WriteString(";")
		totalSQL = ts.String()
	}
	t.write(orderBySQL, keyword.Order.Literal(), keyword.By.Literal())
	t.write(pageSQL)
	t.buf.WriteString(";")
	values = append(values, whereValues...)
	values = append(values, havingValues...)
	t.values = values
	execSQL = t.buf.String()
	return OfPersist(execSQL, totalSQL, values, mappers)
}
