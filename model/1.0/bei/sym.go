// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

import (
	"metis/model/1.0/bei/keyword"
	"strings"
)

type SymService interface {
	I(bool) string
}

type Sym func(bool) string

func (s Sym) String() string {
	return s.Self()
}

func (s Sym) Ph() string {
	return s(true)
}

func (s Sym) Self() string {
	return s(false)
}

func andSym(bool) string {
	return keyword.And.Literal()
}

func orSym(bool) string {
	return keyword.Or.Literal()
}

func eqSym(has bool) string {
	if has {
		return "= ?"
	}
	return "="
}

func nqSym(has bool) string {
	if has {
		return "<> ?"
	}
	return "<>"
}

func ltSym(has bool) string {
	if has {
		return "< ?"
	}
	return "<"
}

func gtSym(has bool) string {
	if has {
		return "> ?"
	}
	return ">"
}

func geSym(has bool) string {
	if has {
		return ">= ?"
	}
	return ">="
}

func leSym(has bool) string {
	if has {
		return "<= ?"
	}
	return "<="
}

func multiSym(bool) string {
	return "*"
}

func addSym(bool) string {
	return "+"
}

func minusSym(bool) string {
	return "-"
}

func likeSym(has bool) string {
	if has {
		return strings.Join([]string{keyword.Like.Literal(), "?"}, " ")
	}
	return keyword.Like.Literal()
}

func isNullSym(bool) string {
	return strings.Join([]string{keyword.Is.Literal(), keyword.Null.Literal()}, " ")
}

func isNotNullSym(bool) string {
	return strings.Join([]string{keyword.Is.Literal(), keyword.Not.Literal(), keyword.Null.Literal()}, " ")
}
