// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

import (
	"metis/model/1.0/bei/keyword"
	"strings"
)

type QueryBuilder[T any] struct {
	*Evaluator[T]
}

func (t *Evaluator[T]) Select(fds ...*FD[T]) *QueryBuilder[T] {
	t.fds = fds
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

func (t *QueryBuilder[T]) WithSQLKey(key string) *QueryBuilder[T] {
	t.sqlKey = key
	return t
}

func (t *QueryBuilder[T]) WithLogicDeleted(cdv ...string) *QueryBuilder[T] {
	if !t.enableLogical() {
		t.logical = t.logical.Enable()
	}
	// using last value
	for _, v := range cdv {
		t.logical.cdval = v
	}
	return t
}

func (t *QueryBuilder[T]) Eval(pss ...PersistService[T]) EvalInfoService[T] {
	t.buf.Reset()
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		lookup := ps.Lookup(t.sqlKey)
		if lookup != nil {
			_, whereValues := t.getWhereSQL()
			_, havingValues := t.getHavingSQL()
			values := append([]any{}, whereValues...)
			values = append(values, havingValues...)
			t.ei = OfEvalInfo(lookup.SQL(), lookup.TotalSQL(), values, lookup.Mappers())
			return t.ei
		}
	}
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

	ts := &strings.Builder{}
	if enablePage {
		ts.WriteString(keyword.Select.Literal())
		ts.WriteString(Space)
		ts.WriteString(keyword.Count.Literal())
		ts.WriteString("(")
		t.writeToBuf(ts, hintSQL)
		ts.WriteString("1")
		ts.WriteString(")")
	}

	t.buf.WriteString(keyword.Select.Literal())
	t.write(hintSQL)
	t.write(fieldSQL)
	t.writeAppend(ts, fromSQL, keyword.From.Literal())
	if len(logicDeletedSQL) > 0 {
		if strings.Contains(whereSQL, keyword.And.Literal()) || strings.Contains(whereSQL, keyword.Or.Literal()) {
			t.writeAppend(ts, LeftParentheses+whereSQL+RightParentheses, keyword.Where.Literal())
		} else {
			t.writeAppend(ts, whereSQL, keyword.Where.Literal())
		}
		t.writeAppend(ts, keyword.And.Literal())
		t.writeAppend(ts, logicDeletedSQL)
	} else {
		t.writeAppend(ts, whereSQL, keyword.Where.Literal())
	}
	t.writeAppend(ts, groupBySQL, keyword.Group.Literal(), keyword.By.Literal())
	t.writeAppend(ts, havingSQL, keyword.Having.Literal())
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
	t.ei = OfEvalInfo(execSQL, totalSQL, values, mappers)
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		ps.Persistence(t.sqlKey, OfEvalInfo(t.ei.SQL(), t.ei.TotalSQL(), nil, t.ei.Mappers()))
	}
	return t.ei
}
