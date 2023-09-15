// Package bei
// @author tabuyos
// @since 2023/9/6
// @description bei
package bei

import (
	"metis/old/model/1.0/bei/keyword"
)

type DeleteBuilder[T any] struct {
	*Evaluator[T]
}

func (t *Evaluator[T]) Delete() *DeleteBuilder[T] {
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

func (t *DeleteBuilder[T]) WithSQLKey(key string) *DeleteBuilder[T] {
	t.sqlKey = key
	return t
}

func (t *DeleteBuilder[T]) WithLogicDeleted(cdv ...string) *DeleteBuilder[T] {
	if !t.enableLogical() {
		t.logical.Enable()
	}
	// using last value
	for _, v := range cdv {
		t.logical.cdval = v
	}
	return t
}

func (t *DeleteBuilder[T]) Eval(pss ...PersistService[T]) EvalInfoService[T] {
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

	var (
		sql    string
		values []any
	)

	fromSQL := t.getFromSQL()
	whereSQL, whereValues := t.getWhereSQL()
	logicDeletedSQL := t.getLogicDeletedSQL()

	if t.enableLogical() {
		setSQL := t.logical.key + PrettyEqual + t.logical.ddval

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
	t.ei = OfEvalInfo[T](sql, "", values, nil)
	for _, ps := range pss {
		if len(t.sqlKey) == 0 {
			break
		}
		ps.Persistence(t.sqlKey, OfEvalInfo[T](sql, "", nil, nil))
	}
	return t.ei
}
