// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"metis/test/autogen/beiold/keyword"
	"strings"
)

type RefTable struct {
	ref  []*RefTable
	name string
	as   string
	on   string
	jt   JointType
}

func OfRef(name string) *RefTable {
	return &RefTable{name: name}
}

type JointType int

const (
	Nil JointType = iota
	Join
	LeftJoin
)

func (j *JointType) String() string {
	if j == nil {
		return ""
	}
	switch *j {
	case Join:
		return keyword.Join.String()
	case LeftJoin:
		return keyword.Left.String() + " " + keyword.Join.String()
	default:
		return ""
	}
}

func (t *RefTable) As(as string) *RefTable {
	t.as = as
	return t
}

func (t *RefTable) RefKey(key string) string {
	if len(t.as) > 0 {
		return t.as + "." + key
	}
	return key
}

func (t *RefTable) JoinType(jt JointType) *RefTable {
	t.jt = jt
	return t
}
func (t *RefTable) On(a string, s Sym, b string) *RefTable {
	var buf strings.Builder
	buf.WriteString(keyword.On.String())
	buf.WriteString(" ")
	buf.WriteString(a)
	buf.WriteString(" ")
	buf.WriteString(s.Self())
	buf.WriteString(" ")
	buf.WriteString(b)
	t.on = buf.String()
	return t
}

func (t *RefTable) Ref(ref ...*RefTable) *RefTable {
	t.ref = ref
	return t
}

func (t *RefTable) selfSQL() (sql string) {
	var name string
	var bs []string
	if len(t.as) == 0 {
		name = t.name
	} else {
		name = t.name + " " + keyword.As.String() + " " + t.as
	}
	if len(t.jt.String()) > 0 {
		bs = append(bs, t.jt.String())
	}
	bs = append(bs, name)
	if len(t.on) > 0 {
		bs = append(bs, t.on)
	}
	return strings.Join(bs, " ")
}

func (t *RefTable) SQL() (sql string) {
	if t == nil {
		return ""
	}
	if len(t.ref) == 0 {
		return t.selfSQL()
	}
	var refs = make([]string, len(t.ref))
	for i, table := range t.ref {
		if i == 0 {
			refs[i] = t.selfSQL() + " " + table.SQL()
			continue
		}
		refs[i] = table.SQL()
	}
	return strings.Join(refs, " ")
}

func (t *RefTable) FlatAll() (tables []*RefTable) {
	if t == nil {
		return nil
	}
	tables = append(tables, t)
	for _, ref := range t.ref {
		flatAll := ref.FlatAll()
		tables = append(tables, flatAll...)
	}
	return
}
