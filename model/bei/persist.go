// Package bei
// @author tabuyos
// @since 2023/9/7
// @description bei
package bei

type PersistService[T any] interface {
	Persist() *Persist[T]
}

type Persist[T any] struct {
	execSQL  string
	countSQL string
	values   []any
	mappers  []func(*T) any
}

func OfPersist[T any](execSQL, countSQL string, values []any, mappers []func(*T) any) *Persist[T] {
	return &Persist[T]{
		execSQL:  execSQL,
		countSQL: countSQL,
		values:   values,
		mappers:  mappers,
	}
}

func (p *Persist[T]) SQL() string {
	return p.execSQL
}

func (p *Persist[T]) TotalSQL() string {
	return p.countSQL
}

func (p *Persist[T]) Values() []any {
	return p.values
}

func (p *Persist[T]) Mappers() []func(*T) any {
	return p.mappers
}
