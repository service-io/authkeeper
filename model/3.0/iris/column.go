package iris

import (
	"metis/model/3.0/iris/internal/keyword"
)

type Column[T any] struct {
	name       string
	aliasFunc  func(string) string
	mapperFunc func(*T) any
	renderFunc func(string) string
}

func WithColumn[T any](name string, mapper func(*T) any) *Column[T] {
	return &Column[T]{name: name, mapperFunc: mapper}
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
	c.aliasFunc = func(name string) string {
		return name + keyword.As.Pretty() + as
	}
	return c
}

func (c *Column[T]) RenderKey(fn func(string) string) *Column[T] {
	c.renderFunc = fn
	return c
}

func (c *Column[T]) Apply(fn func(string) string) *Column[T] {
	c.name = fn(c.name)
	return c
}

func (c *Column[T]) Literal() string {
	if c.aliasFunc == nil {
		return c.Name()
	}
	return c.aliasFunc(c.Name())
}

func render[T any](asc bool, col *Column[T]) *Order {
	return &Order{
		col: col.Name(),
		asc: asc,
	}
}

func (c *Column[T]) Desc() *Order {
	return render(false, c)
}

func (c *Column[T]) Asc() *Order {
	return render(true, c)
}
