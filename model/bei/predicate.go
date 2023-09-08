// Package bei
// @author tabuyos
// @since 2023/9/5
// @description bei
package bei

import (
	"fmt"
	"strings"
)

type Mode int

const (
	DftMode Mode = iota + 1
	AndMode
	OrMode
)

type Predicate struct {
	mod Mode
	fid string
	lft *Predicate
	sym Sym
	rht *Predicate
	ars []any
	buf strings.Builder
}

func Once(f string, s Sym, ars ...any) *Predicate {
	return &Predicate{
		mod: DftMode,
		fid: f,
		sym: s,
		ars: ars,
	}
}

func (t *Evaluator[T]) Once(f string, s Sym, ars ...any) *Predicate {
	return &Predicate{
		mod: DftMode,
		fid: f,
		sym: s,
		ars: ars,
	}
}

func (p *Predicate) And(preds ...*Predicate) *Predicate {
	var fp = p
	for _, pred := range preds {
		fp = &Predicate{
			mod: AndMode,
			lft: fp,
			sym: andSym,
			rht: pred,
		}
	}
	return fp
}

func (p *Predicate) Or(preds ...*Predicate) *Predicate {
	var fp = p
	for _, pred := range preds {
		fp = &Predicate{
			mod: OrMode,
			lft: fp,
			sym: orSym,
			rht: pred,
		}
	}
	return fp
}

func (p *Predicate) String() string {
	sql, values := p.SQL()
	var index = 0
	var buf strings.Builder
	for _, r := range sql {
		if r == '?' {
			buf.WriteString(fmt.Sprintf("%v", values[index]))
			index++
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func (p *Predicate) SQL() (sql string, values []any) {
	if p == nil {
		return "", nil
	}

	if p.lft == nil {
		p.buf.Reset()
		p.buf.WriteString(p.fid)
		p.buf.WriteString(Space)
		p.buf.WriteString(p.sym.Ph())
		sql = p.buf.String()
		values = append(values, p.ars...)
		return
	}
	lftSQL, lftValues := p.lft.SQL()
	rhtSQL, rhtValues := p.rht.SQL()
	if p.lft.mod == DftMode {
		p.lft.mod = p.mod
	}
	if p.rht.mod == DftMode {
		p.rht.mod = p.mod
	}
	p.buf.Reset()
	if p.mod == p.lft.mod {
		if p.mod == p.rht.mod {
			p.buf.WriteString(lftSQL)
			p.buf.WriteString(Space)
			p.buf.WriteString(p.sym.Ph())
			p.buf.WriteString(Space)
			p.buf.WriteString(rhtSQL)
		} else {
			p.buf.WriteString(lftSQL)
			p.buf.WriteString(Space)
			p.buf.WriteString(p.sym.Ph())
			p.buf.WriteString(Space)
			p.buf.WriteString(LeftParentheses)
			p.buf.WriteString(rhtSQL)
			p.buf.WriteString(RightParentheses)
		}
	} else {
		if p.mod == p.rht.mod {
			p.buf.WriteString(LeftParentheses)
			p.buf.WriteString(lftSQL)
			p.buf.WriteString(RightParentheses)
			p.buf.WriteString(Space)
			p.buf.WriteString(p.sym.Ph())
			p.buf.WriteString(Space)
			p.buf.WriteString(rhtSQL)
		} else {
			p.buf.WriteString(LeftParentheses)
			p.buf.WriteString(lftSQL)
			p.buf.WriteString(RightParentheses)
			p.buf.WriteString(Space)
			p.buf.WriteString(p.sym.Ph())
			p.buf.WriteString(Space)
			p.buf.WriteString(LeftParentheses)
			p.buf.WriteString(rhtSQL)
			p.buf.WriteString(RightParentheses)
		}
	}
	sql = p.buf.String()
	values = append(values, lftValues...)
	values = append(values, rhtValues...)
	return
}
