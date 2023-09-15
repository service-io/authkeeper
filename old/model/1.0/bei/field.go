// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

import (
	"fmt"
	"metis/old/model/1.0/bei/keyword"
	"strings"
)

type FD[T any] struct {
	fd     string
	as     string
	fn     func(*T) any
	render func(string) string
}

// Fd func
func (t *FD[T]) Fd() string {
	if t.render == nil {
		return t.fd
	}
	return t.render(t.fd)
}

// As func
func (t *FD[T]) As(as string) *FD[T] {
	t.as = as
	return t
}

// Literal func
func (t *FD[T]) Literal() string {
	if len(t.as) > 0 {
		return strings.Join([]string{t.Fd(), keyword.As.Literal(), t.as}, " ")
	}
	return t.Fd()
}

// Inject func
func (t *FD[T]) Inject(rt *RefTable) *FD[T] {
	t.render = rt.RefKey
	return t
}

// Func func
func (t *FD[T]) Func(fn func(string) string) *FD[T] {
	if fn == nil {
		return t
	}
	t.fd = fn(t.fd)
	return t
}

// Apply func
func (t *FD[T]) Apply(sym Sym, val ...any) *Predicate {
	return Once(t.Fd(), sym, val...)
}

// Nq func
func (t *FD[T]) Nq(val ...any) *Predicate {
	return Once(t.Fd(), nqSym, val...)
}

// Eq func
func (t *FD[T]) Eq(val ...any) *Predicate {
	return Once(t.Fd(), eqSym, val...)
}

// Lt func
func (t *FD[T]) Lt(val ...any) *Predicate {
	return Once(t.Fd(), ltSym, val...)
}

// Gt func
func (t *FD[T]) Gt(val ...any) *Predicate {
	return Once(t.Fd(), gtSym, val...)
}

// Le func
func (t *FD[T]) Le(val ...any) *Predicate {
	return Once(t.Fd(), leSym, val...)
}

// Ge func
func (t *FD[T]) Ge(val ...any) *Predicate {
	return Once(t.Fd(), geSym, val...)
}

// Like func
func (t *FD[T]) Like(val ...any) *Predicate {
	values := make([]any, len(val))
	for i, a := range val {
		values[i] = fmt.Sprintf("%%%v%%", a)
	}
	return Once(t.Fd(), likeSym, values...)
}

// LikeLeft func
func (t *FD[T]) LikeLeft(val ...any) *Predicate {
	values := make([]any, len(val))
	for i, a := range val {
		values[i] = fmt.Sprintf("%%%v", a)
	}
	return Once(t.Fd(), likeSym, values...)
}

// LikeRight func
func (t *FD[T]) LikeRight(val ...any) *Predicate {
	values := make([]any, len(val))
	for i, a := range val {
		values[i] = fmt.Sprintf("%v%%", a)
	}
	return Once(t.Fd(), likeSym, values...)
}
