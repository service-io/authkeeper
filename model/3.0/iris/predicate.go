package iris

import (
	"metis/model/3.0/iris/internal/constant"
	"strings"
)

type Predicate struct {
	mod   Mode
	ars   []any
	tasks []Task
}

func WithPredicate(f string, s Operator, ars ...any) *Predicate {
	pred := &Predicate{
		mod: DftMode,
		ars: ars,
		tasks: []Task{
			func(buffers ...*strings.Builder) []any {
				for _, buffer := range buffers {
					buffer.WriteString(f)
					buffer.WriteString(constant.Space.Literal())
					buffer.WriteString(s.Literal())
				}
				return nil
			},
		},
	}
	return pred
}

func (p *Predicate) renderSQL(currentMode, inverseMode Mode, sql string) string {
	if currentMode == inverseMode {
		return strings.Join([]string{constant.LeftParentheses.Literal(), constant.RightParentheses.Literal()}, sql)
	}
	return sql
}

func (p *Predicate) predFunc(buf *strings.Builder, currentMode, inverseMode Mode, sym Operator, preds ...*Predicate) {
	var snips = make([]string, len(preds)+1)
	snips[0] = p.renderSQL(p.mod, inverseMode, buf.String())
	buf.Reset()
	for i, pred := range preds {
		mod := pred.mod
		sql, val := pred.Literal()
		p.ars = append(p.ars, val...)
		snips[i+1] = p.renderSQL(mod, inverseMode, sql)
	}
	delimiter := strings.Join([]string{constant.Space.Literal(), constant.Space.Literal()}, sym.Self())
	buf.WriteString(strings.Join(snips, delimiter))
	p.mod = currentMode
}

func (p *Predicate) And(preds ...*Predicate) *Predicate {
	p.tasks = append(
		p.tasks, func(buffers ...*strings.Builder) []any {
			for _, buffer := range buffers {
				p.predFunc(buffer, AndMode, OrMode, AndOp, preds...)
			}
			return nil
		},
	)
	return p
}

func (p *Predicate) Or(preds ...*Predicate) *Predicate {
	p.tasks = append(
		p.tasks, func(buffers ...*strings.Builder) []any {
			for _, buffer := range buffers {
				p.predFunc(buffer, OrMode, AndMode, OrOp, preds...)
			}
			return nil
		},
	)
	return p
}

func (p *Predicate) Literal() (sql string, values []any) {
	buf := &strings.Builder{}
	p.Render(buf)
	return buf.String(), p.ars
}

func (p *Predicate) Render(buffers ...*strings.Builder) []any {
	tempBuffers := make([]*strings.Builder, len(buffers))
	for i := range buffers {
		tempBuffers[i] = &strings.Builder{}
	}
	for _, task := range p.tasks {
		task.Idle(tempBuffers...)
	}
	for i, buffer := range buffers {
		buf := tempBuffers[i]
		buffer.WriteString(buf.String())
	}

	return p.ars
}
