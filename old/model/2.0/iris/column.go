package iris

import (
	"metis/old/model/2.0/iris/internal/keyword"
)

type Column[T any] struct {
	name       string
	aliasFunc  func() string
	mapperFunc func(*T) any
	renderFunc func(string) string
}

func WithColumn[T any](name string) *Column[T] {
	return &Column[T]{name: name}
}

func (c *Column[T]) Name() string {
	if c.renderFunc == nil {
		return c.name
	}
	return c.renderFunc(c.name)
}

func (c *Column[T]) Mapper() func(*T) any {
	return c.mapperFunc
}

func (c *Column[T]) As(as string) *Column[T] {
	c.aliasFunc = func() string {
		return c.Name() + keyword.As.Pretty() + as
	}
	return c
}

func (c *Column[T]) RenderKey(fn func(string) string) *Column[T] {
	c.renderFunc = fn
	return c
}

func (c *Column[T]) Literal() string {
	if c.aliasFunc == nil {
		return c.Name()
	}
	return c.aliasFunc()
}

func (c *Column[T]) Apply(fn func(string) string) *Column[T] {
	c.name = fn(c.name)
	return c
}
