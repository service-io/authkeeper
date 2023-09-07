// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"fmt"
	"metis/test/autogen/beiold/keyword"
	"strconv"
	"strings"
	"time"
)

type FD[T any] struct {
	fd     string
	as     string
	fn     func(*T) any
	render func(string) string
}

func (t *FD[T]) Fd() string {
	if t.render == nil {
		return t.fd
	}
	return t.render(t.fd)
}

func (t *FD[T]) As(as string) *FD[T] {
	t.as = as
	return t
}

func (t *FD[T]) Literal() string {
	if len(t.as) > 0 {
		return t.Fd() + " " + keyword.As.Literal() + " " + t.as
	}
	return t.Fd()
}

func (t *FD[T]) Func(fn func(string) string) *FD[T] {
	if fn == nil {
		return t
	}
	t.fd = fn(t.fd)
	return t
}

func (t *FD[T]) Inject(rt *RefTable) *FD[T] {
	t.render = rt.RefKey
	return t
}

func OfFD[T any](name string, fn func(*T) any) *FD[T] {
	return &FD[T]{
		fd: name,
		fn: fn,
	}
}

type BaseEntity[T any] struct {
	fds               []*FD[T]
	ref               *RefTable
	where             *Predicate
	having            *Predicate
	hints             []keyword.Keyword
	orders            []*Order
	groups            []*Group
	values            []any
	buf               strings.Builder
	limit             int64
	offset            int64
	logicDeleted      bool
	deletedKey        string
	deletedVal        string
	undeletedVal      string
	currentDeletedVal string
}

func (t *BaseEntity[T]) WrapFn(rt *RefTable) func(*FD[T]) *FD[T] {
	return func(f *FD[T]) *FD[T] {
		return f.Inject(rt)
	}
}

func (t *BaseEntity[T]) Select(fds ...*FD[T]) *BaseEntity[T] {
	t.fds = append(t.fds, fds...)
	return t
}

func (t *BaseEntity[T]) Hint(hints ...keyword.Keyword) *BaseEntity[T] {
	t.hints = append(t.hints, hints...)
	return t
}

func (t *BaseEntity[T]) From(rt *RefTable) *BaseEntity[T] {
	t.ref = rt
	return t
}

func (t *BaseEntity[T]) Where(pd *Predicate) *BaseEntity[T] {
	t.where = pd
	return t
}

func (t *BaseEntity[T]) GroupBy(gbs ...*Group) *BaseEntity[T] {
	t.groups = gbs
	return t
}

func (t *BaseEntity[T]) Having(pd *Predicate) *BaseEntity[T] {
	t.having = pd
	return t
}

func (t *BaseEntity[T]) OrderBy(obs ...*Order) *BaseEntity[T] {
	t.orders = obs
	return t
}

func (t *BaseEntity[T]) Limit(lmt int64) *BaseEntity[T] {
	t.limit = lmt
	return t
}

func (t *BaseEntity[T]) Offset(ost int64) *BaseEntity[T] {
	t.offset = ost
	return t
}

func (t *BaseEntity[T]) WithLogicDeleted(cdv ...string) *BaseEntity[T] {
	t.logicDeleted = true
	if len(cdv) == 1 {
		t.currentDeletedVal = cdv[0]
	}
	return t
}

const (
	Space            = " "
	Equal            = "="
	PrettyEqual      = " = "
	LeftParentheses  = "("
	RightParentheses = ")"
)

func (t *BaseEntity[T]) String() string {
	sql := t.buf.String()
	values := t.values
	var index = 0
	var buf strings.Builder
	for _, r := range sql {
		if r == '?' {
			v := values[index]
			switch vt := v.(type) {
			case string:
				buf.WriteString(fmt.Sprintf("'%v'", vt))
			case time.Time:
				buf.WriteString(fmt.Sprintf("'%v'", vt.Format(time.RFC3339Nano)))
			default:
				buf.WriteString(fmt.Sprintf("%v", vt))
			}
			index++
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func (t *BaseEntity[T]) QuerySQL() (execSQL, countSQL string, values []any, mappers []func(*T) any) {
	t.buf.Reset()
	if len(t.fds) == 0 {
		panic("not found any column")
	}
	var fieldSQL string

	hintSQL := t.getHintSQL()
	fieldSQL, mappers = t.getFieldSQL()
	fromSQL := t.getFromSQL()
	whereSQL, whereValues := t.getWhereSQL()
	logicDeletedSQL := t.getLogicDeletedSQL()
	groupBySQL := t.getGroupBySQL()
	havingSQL, havingValues := t.getHavingSQL()
	orderBySQL := t.getOrderBySQL()
	pageSQL := t.getPageSQL()

	t.buf.WriteString(keyword.Select.Literal())
	t.write(hintSQL)
	t.write(fieldSQL)
	t.write(fromSQL, keyword.From.Literal())
	if len(logicDeletedSQL) > 0 {
		t.write(LeftParentheses+whereSQL+RightParentheses, keyword.Where.Literal())
		t.write(keyword.And.Literal())
		t.write(logicDeletedSQL)
	} else {
		t.write(whereSQL, keyword.Where.Literal())
	}
	t.write(groupBySQL, keyword.Group.Literal(), keyword.By.Literal())
	t.write(havingSQL, keyword.Having.Literal())
	if len(pageSQL) > 0 {
		countSQL = t.buf.String() + ";"
	}
	t.write(orderBySQL, keyword.Order.Literal(), keyword.By.Literal())
	t.write(pageSQL)
	t.buf.WriteString(";")
	values = append(values, whereValues...)
	values = append(values, havingValues...)
	t.values = values
	execSQL = t.buf.String()
	return
}

func (t *BaseEntity[T]) write(sql string, kws ...string) {
	sql = strings.TrimSpace(sql)
	if len(sql) > 0 {
		if len(kws) > 0 {
			t.buf.WriteString(Space)
			t.buf.WriteString(strings.Join(kws, Space))
		}
		t.buf.WriteString(Space)
		t.buf.WriteString(sql)
	}
}

func (t *BaseEntity[T]) getHintSQL() string {
	var snips = make([]string, len(t.hints))
	for i, hint := range t.hints {
		snips[i] = hint.Literal()
	}
	return strings.Join(snips, Space)
}

func (t *BaseEntity[T]) getFieldSQL() (sql string, mappers []func(*T) any) {
	fields := make([]string, len(t.fds))
	mappers = make([]func(*T) any, len(t.fds))
	for i, fd := range t.fds {
		fields[i] = fd.Literal()
		mappers[i] = fd.fn
	}
	sql = strings.Join(fields, ", ")
	return
}

func (t *BaseEntity[T]) getFromSQL() string {
	return t.ref.SQL()
}

func (t *BaseEntity[T]) getWhereSQL() (string, []any) {
	return t.where.SQL()
}

func (t *BaseEntity[T]) getGroupBySQL() string {
	var snips = make([]string, len(t.groups))
	for i, group := range t.groups {
		snips[i] = group.SQL()
	}
	return strings.Join(snips, ", ")
}

func (t *BaseEntity[T]) getHavingSQL() (string, []any) {
	return t.having.SQL()
}

func (t *BaseEntity[T]) getOrderBySQL() string {
	var snips = make([]string, len(t.orders))
	for i, order := range t.orders {
		snips[i] = order.SQL()
	}
	return strings.Join(snips, ", ")
}

func (t *BaseEntity[T]) getPageSQL() string {
	if t.limit > 0 {
		if t.offset > 0 {
			return keyword.Limit.Literal() + Space + strconv.FormatInt(t.limit, 10) + keyword.Offset.Literal() + Space + strconv.FormatInt(t.offset, 10)
		} else {
			return keyword.Limit.Literal() + Space + strconv.FormatInt(t.limit, 10)
		}
	} else {
		return ""
	}
}

func (t *BaseEntity[T]) getLogicDeletedSQL() string {
	if t.logicDeleted {
		tables := t.ref.FlatAll()
		var snips = make([]string, len(tables))
		for i, table := range tables {
			key := table.RefKey(t.deletedKey)
			val := t.undeletedVal
			if len(t.currentDeletedVal) > 0 {
				val = t.currentDeletedVal
			}
			snips[i] = key + PrettyEqual + val
		}
		return strings.Join(snips, Space+keyword.And.Literal()+Space)
	}
	return ""
}
