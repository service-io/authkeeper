// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

import (
	"fmt"
	"metis/old/model/1.0/bei/keyword"
	"strconv"
	"strings"
	"time"
)

type Evaluator[T any] struct {
	fds     []*FD[T]
	ref     *RefTable
	where   *Predicate
	having  *Predicate
	hints   []keyword.Keyword
	orders  []*Order
	groups  []*Group
	values  []any
	buf     strings.Builder
	limit   int64
	offset  int64
	logical *Logical
	sqlKey  string
	ei      *EvalInfo[T]
}

func Default[T any]() *Evaluator[T] {
	return &Evaluator[T]{}
}

func WithLogical[T any]() *Evaluator[T] {
	return &Evaluator[T]{logical: OfLogical()}
}

func (t *Evaluator[T]) WrapFn(rt *RefTable) func(*FD[T]) *FD[T] {
	return func(f *FD[T]) *FD[T] {
		return f.Inject(rt)
	}
}

func OfFD[T any](name string, fn func(*T) any) *FD[T] {
	return &FD[T]{
		fd: name,
		fn: fn,
	}
}

func (t *Evaluator[T]) OfFD(name string, fn func(*T) any) *FD[T] {
	return &FD[T]{
		fd: name,
		fn: fn,
	}
}

func (t *Evaluator[T]) WithValues(values ...any) {
	t.values = append(t.values, values...)
}

func (t *Evaluator[T]) String() string {
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

func (t *Evaluator[T]) writeToBuf(buf *strings.Builder, sql string, kws ...string) {
	sql = strings.TrimSpace(sql)
	if len(sql) > 0 {
		if len(kws) > 0 {
			buf.WriteString(Space)
			buf.WriteString(strings.Join(kws, Space))
		}
		buf.WriteString(Space)
		buf.WriteString(sql)
	}
}

func (t *Evaluator[T]) write(sql string, kws ...string) {
	t.writeToBuf(&t.buf, sql, kws...)
}

func (t *Evaluator[T]) writeAppend(appendBuf *strings.Builder, sql string, kws ...string) {
	t.writeToBuf(&t.buf, sql, kws...)
	t.writeToBuf(appendBuf, sql, kws...)
}

func (t *Evaluator[T]) writeAppendPred(en bool, appendBuf *strings.Builder, sql string, kws ...string) {
	t.writeToBuf(&t.buf, sql, kws...)
	if en {
		t.writeToBuf(appendBuf, sql, kws...)
	}
}

func (t *Evaluator[T]) getHintSQL() string {
	var snips = make([]string, len(t.hints))
	for i, hint := range t.hints {
		snips[i] = hint.Literal()
	}
	return strings.Join(snips, Space)
}

func (t *Evaluator[T]) getFieldSQL() (sql string, mappers []func(*T) any) {
	fields := make([]string, len(t.fds))
	mappers = make([]func(*T) any, len(t.fds))
	for i, fd := range t.fds {
		fields[i] = fd.Literal()
		mappers[i] = fd.fn
	}
	sql = strings.Join(fields, ", ")
	return
}

func (t *Evaluator[T]) getFromSQL() string {
	return t.ref.SQL()
}

func (t *Evaluator[T]) getWhereSQL() (string, []any) {
	return t.where.SQL()
}

func (t *Evaluator[T]) getGroupBySQL() string {
	var snips = make([]string, len(t.groups))
	for i, group := range t.groups {
		snips[i] = group.SQL()
	}
	return strings.Join(snips, ", ")
}

func (t *Evaluator[T]) getHavingSQL() (string, []any) {
	return t.having.SQL()
}

func (t *Evaluator[T]) getOrderBySQL() string {
	var snips = make([]string, len(t.orders))
	for i, order := range t.orders {
		snips[i] = order.SQL()
	}
	return strings.Join(snips, ", ")
}

func (t *Evaluator[T]) getPageSQL() string {
	if t.limit > 0 {
		if t.offset > 0 {
			return keyword.Limit.Literal() + Space + strconv.FormatInt(
				t.limit, 10,
			) + keyword.Offset.Literal() + Space + strconv.FormatInt(t.offset, 10)
		}
		return keyword.Limit.Literal() + Space + strconv.FormatInt(t.limit, 10)
	} else {
		return ""
	}
}

func (t *Evaluator[T]) enableLogical() bool {
	if t.logical == nil {
		return false
	}
	return t.logical.enable
}

func (t *Evaluator[T]) getLogicDeletedSQL() string {
	if t.enableLogical() {
		tables := t.ref.FlatAll()
		var snips = make([]string, len(tables))
		for i, table := range tables {
			key := table.RefKey(t.logical.key)
			val := t.logical.udval
			if len(t.logical.cdval) > 0 {
				val = t.logical.cdval
			}
			snips[i] = key + PrettyEqual + val
		}
		return strings.Join(snips, Space+keyword.And.Literal()+Space)
	}
	return ""
}

func (t *Evaluator[T]) getIntoSQL() string {
	return t.ref.SQL()
}

func (t *Evaluator[T]) getInsertPlaceholder() string {
	var snips = make([]string, len(t.fds))
	for i := 0; i < len(t.fds); i++ {
		snips[i] = "?"
	}
	return strings.Join(snips, ", ")
}

func (t *Evaluator[T]) getTimesPlaceholder(times int, ph string) string {
	var snips = make([]string, times)
	for i := 0; i < times; i++ {
		snips[i] = LeftParentheses + ph + RightParentheses
	}
	return strings.Join(snips, ", ")
}

func (t *Evaluator[T]) getSetSQL() string {
	var snips = make([]string, len(t.fds))
	for i, update := range t.fds {
		snips[i] = update.Fd() + " = ?"
	}
	return strings.Join(snips, ", ")
}

func (t *Evaluator[T]) EvalInfo() *EvalInfo[T] {
	if t == nil {
		return nil
	}
	return t.ei
}

func (t *Evaluator[T]) Replace(ei *EvalInfo[T]) {
	if t == nil {
		return
	}
	t.ei = ei
}
