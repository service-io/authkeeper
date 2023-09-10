package iris

import (
	"metis/model/3.0/iris/internal/constant"
	"metis/model/3.0/iris/internal/keyword"
	"strings"
)

type RefTable struct {
	nameTask   Task
	aliasTask  Task
	onTask     Task
	refTask    Task
	typeTask   Task
	renderTask Task
}

func WithRefTable(name string) *RefTable {
	return &RefTable{
		nameTask: func(buffers ...*strings.Builder) []any {
			buf := buffers[0]
			buf.WriteString(constant.Space.Literal())
			buf.WriteString(name)
			return nil
		},
	}
}

func (t *RefTable) As(as string) *RefTable {
	t.aliasTask = func(buffers ...*strings.Builder) []any {
		switch len(buffers) {
		case 2:
			buf := buffers[1]
			buf.WriteString(as)
		default:
			buf := buffers[0]
			snip := keyword.As.Literal() + constant.Space.Literal() + as
			buf.WriteString(constant.Space.Literal())
			buf.WriteString(snip)
		}
		return nil
	}
	return t
}

func (t *RefTable) On(at string, sm Operator, bt string) *RefTable {
	t.onTask = func(buffers ...*strings.Builder) []any {
		buf := buffers[0]
		snip := strings.Join([]string{keyword.On.Literal(), at, sm.Self(), bt}, constant.Space.Literal())
		buf.WriteString(constant.Space.Literal())
		buf.WriteString(snip)
		return nil
	}
	return t
}

func (t *RefTable) OnEQ(at string, bt string) *RefTable {
	return t.On(at, EqOp, bt)
}

func (t *RefTable) Ref(refs ...*RefTable) *RefTable {
	t.refTask = func(buffers ...*strings.Builder) []any {
		buf := buffers[0]
		for _, ref := range refs {
			buf.WriteString(constant.Space.Literal())
			ref.Render(buf)
		}
		return nil
	}
	return t
}

func (t *RefTable) RenderKey(key string) string {
	if t.aliasTask == nil {
		return key
	}
	buf := &strings.Builder{}
	t.aliasTask(nil, buf)
	return buf.String() + constant.Dot.Literal() + key
}

func (t *RefTable) JoinType(jt JointType) *RefTable {
	t.typeTask = func(buffers ...*strings.Builder) []any {
		snip := jt.Literal()
		for _, buffer := range buffers {
			buffer.WriteString(snip)
		}
		return nil
	}
	return t
}

func (t *RefTable) Literal() string {
	buf := &strings.Builder{}
	t.Render(buf)
	return buf.String()
}

func (t *RefTable) Render(buffers ...*strings.Builder) []any {
	for _, buffer := range buffers {
		t.typeTask.Idle(buffer)
		t.nameTask.Idle(buffer)
		t.aliasTask.Idle(buffer)
		t.refTask.Idle(buffer)
		t.onTask.Idle(buffer)
	}
	return nil
}
