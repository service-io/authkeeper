package iris

import (
	"metis/model/3.0/iris/internal/constant"
	"metis/model/3.0/iris/internal/keyword"
	"strconv"
	"strings"
)

type SelectEvaluator[T any] struct {
	*Evaluator[T]
	pageable bool
	pageTask Task
	tasks    []Task
	hintTask Task
}

func (e *Evaluator[T]) Select(cols ...*Column[T]) *SelectEvaluator[T] {
	se := &SelectEvaluator[T]{Evaluator: e}
	task := func(buffers ...*strings.Builder) []any {
		execFn := func() {
			buf := buffers[0]
			buf.WriteString(keyword.Select.Literal())
			se.hintTask.Idle(buf)
			colLit := make([]string, len(cols))
			for i, col := range cols {
				colLit[i] = col.Literal()
			}
			buf.WriteString(constant.Space.Literal())
			buf.WriteString(strings.Join(colLit, constant.CommaSpace.Literal()))
		}
		switch len(buffers) {
		case 1:
			execFn()
		case 2:
			execFn()
			if se.pageable {
				buf := buffers[1]
				buf.WriteString(keyword.Select.Literal())
				buf.WriteString(constant.Space.Literal())
				buf.WriteString(keyword.Count.Literal())
				buf.WriteString(constant.LeftParentheses.Literal())
				buf.WriteString("1")
				buf.WriteString(constant.RightParentheses.Literal())
			}
		}
		return nil
	}
	se.tasks = append(se.tasks, task)
	return se
}

func (e *SelectEvaluator[T]) Hint(ks ...keyword.Keyword) *SelectEvaluator[T] {
	e.hintTask = func(buffers ...*strings.Builder) []any {
		buf := buffers[0]
		for _, k := range ks {
			buf.WriteString(constant.Space.Literal())
			buf.WriteString(k.Literal())
		}
		return nil
	}
	return e
}

func (e *SelectEvaluator[T]) From(rt *RefTable) *SelectEvaluator[T] {
	task := func(buffers ...*strings.Builder) []any {
		for _, buffer := range buffers {
			buffer.WriteString(constant.Space.Literal())
			buffer.WriteString(keyword.From.Literal())
		}
		values := rt.Render(buffers...)
		return values
	}
	e.tasks = append(e.tasks, task)
	return e
}

func (e *SelectEvaluator[T]) Where(pred *Predicate) *SelectEvaluator[T] {
	task := func(buffers ...*strings.Builder) []any {
		buf := &strings.Builder{}
		values := pred.Render(buf)
		for _, buffer := range buffers {
			buffer.WriteString(keyword.Where.Pretty())
			buffer.WriteString(buf.String())
		}
		return values
	}
	e.tasks = append(e.tasks, task)
	return e
}

func (e *SelectEvaluator[T]) GroupBy(cols ...*Column[T]) *SelectEvaluator[T] {
	task := func(buffers ...*strings.Builder) []any {
		snips := make([]string, len(cols))
		for i, col := range cols {
			snips[i] = col.Literal()
		}
		for _, buffer := range buffers {
			buffer.WriteString(keyword.GroupBy.Pretty())
			buffer.WriteString(strings.Join(snips, constant.CommaSpace.Literal()))
		}
		return nil
	}
	e.tasks = append(e.tasks, task)
	return e
}

func (e *SelectEvaluator[T]) Having(pred *Predicate) *SelectEvaluator[T] {
	task := func(buffers ...*strings.Builder) []any {
		for _, buffer := range buffers {
			buffer.WriteString(keyword.Having.Pretty())
		}
		values := pred.Render(buffers...)
		return values
	}
	e.tasks = append(e.tasks, task)
	return e
}

func (e *SelectEvaluator[T]) OrderBy(orders ...*Order) *SelectEvaluator[T] {
	task := func(buffers ...*strings.Builder) []any {
		for _, buffer := range buffers {
			buffer.WriteString(keyword.OrderBy.Pretty())
			var snips = make([]string, len(orders))
			for i, order := range orders {
				snips[i] = order.Literal()
			}
			buffer.WriteString(strings.Join(snips, constant.CommaSpace.Literal()))
		}
		return nil
	}
	e.tasks = append(e.tasks, task)
	return e
}

func (e *SelectEvaluator[T]) Limit(limit int64) *SelectEvaluator[T] {
	e.pageTask = func(buffers ...*strings.Builder) []any {
		buf := buffers[0]
		buf.WriteString(keyword.Limit.Pretty())
		buf.WriteString(strconv.FormatInt(limit, 10))
		return nil
	}
	return e
}

func (e *SelectEvaluator[T]) Offset(offset int64) *SelectEvaluator[T] {
	e.pageable = true
	preTask := e.pageTask
	e.pageTask = func(buffers ...*strings.Builder) []any {
		preTask.Idle(buffers...)
		buf := buffers[0]
		buf.WriteString(keyword.Offset.Pretty())
		buf.WriteString(strconv.FormatInt(offset, 10))
		return nil
	}
	return e
}

func (e *SelectEvaluator[T]) Eval(pss ...PersistService[T]) EvalInfoService[T] {
	//var buffers []*strings.Builder
	//var values []any
	//execSQLBuf := &strings.Builder{}
	//buffers = append(buffers, execSQLBuf)
	//if e.pageable {
	//	buffers = append(buffers, &strings.Builder{})
	//}
	//for _, task := range e.tasks {
	//	values = append(values, task.Idle(buffers...)...)
	//}
	//e.pageTask.Idle(execSQLBuf)
	//fmt.Println(execSQLBuf.String())
	//if e.pageable {
	//	fmt.Println(buffers[1].String())
	//}
	//fmt.Println(values)
	return nil
}
