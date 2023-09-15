package iris

type EvalService[T any] interface {
	Eval() EvalInfoService[T]
}

type EvalInfoService[T any] interface {
	EvalInfo() *EvalInfo[T]
	Replace(ei *EvalInfo[T])
}

type ConfigService[T any] interface {
	ColumnAndValue(fns ...func(*Column[T], any) bool) (selfishs []Selfish, values []any)
	Configure(func(*Evaluator[T]))
	Evaluator() EvalInfoService[T]
	Asterisk(fns ...func(string) string) []*Column[T]
	PKey() *Column[T]
	LogicDelKey() *Column[T]
	Table() *RefTable
	Self() *T
}
