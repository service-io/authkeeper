// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/model/bei/keyword"
)

type UpdateBuilder[T any] struct {
	*Evaluator[T]
}

func (t *Evaluator[T]) UpdateRef(ref *RefTable, fds ...*FD[T]) *UpdateBuilder[T] {
	t.fds = fds
	t.ref = ref
	return &UpdateBuilder[T]{t}
}

func (t *Evaluator[T]) Update(ref *RefTable) *UpdateBuilder[T] {
	t.ref = ref
	return &UpdateBuilder[T]{Evaluator: t}
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

func (t *UpdateBuilder[T]) WithSQLKey(key string) *UpdateBuilder[T] {
	t.sqlKey = key
	return t
}

func (t *UpdateBuilder[T]) WithLogicDeleted(cdv ...string) *UpdateBuilder[T] {
	if !t.enableLogical() {
		t.logical.Enable()
	}
	// using last value
	for _, v := range cdv {
		t.logical.cdval = v
	}
	return t
}

func (t *UpdateBuilder[T]) Eval(pss ...PersistService[T]) EvalInfoService[T] {
	t.buf.Reset()
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		lookup := ps.Lookup(t.sqlKey)
		if lookup != nil {
			_, whereValues := t.getWhereSQL()
			values := append([]any{}, whereValues...)
			t.ei = OfEvalInfo[T](lookup.SQL(), "", values, nil)
			return t.ei
		}
	}
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

	t.ei = OfEvalInfo[T](sql, "", values, nil)
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		ps.Persistence(t.sqlKey, OfEvalInfo[T](sql, "", nil, nil))
	}
	return t.ei
}
