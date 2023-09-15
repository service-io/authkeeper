package iris

type LockupService[T any] interface {
	Put(key string, val *T)
	Get(key string) *T
}

type PersistService[T any] interface {
	Persistence(sqlKey string, info *EvalInfo[T])
	Lookup(sqlKey string) *EvalInfo[T]
}

type EvalInfo[T any] struct {
	execSQL  string
	totalSQL string
	values   []any
	mappers  []func(*T) any
}

func OfEvalInfo[T any](execSQL, countSQL string, values []any, mappers []func(*T) any) *EvalInfo[T] {
	return &EvalInfo[T]{
		execSQL:  execSQL,
		totalSQL: countSQL,
		values:   values,
		mappers:  mappers,
	}
}

func (p *EvalInfo[T]) SQL() string {
	return p.execSQL
}

func (p *EvalInfo[T]) TotalSQL() string {
	return p.totalSQL
}

func (p *EvalInfo[T]) Values() []any {
	return p.values
}

func (p *EvalInfo[T]) Mappers() []func(*T) any {
	return p.mappers
}

func (p *EvalInfo[T]) EvalInfo() *EvalInfo[T] {
	return p
}

func (p *EvalInfo[T]) Replace(ei *EvalInfo[T]) {
	p.execSQL = ei.execSQL
	p.totalSQL = ei.totalSQL
	p.values = ei.values
	p.mappers = ei.mappers
}
