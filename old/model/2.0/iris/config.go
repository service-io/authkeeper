package iris

type EvalService[T any] interface {
	Eval(pss ...PersistService[T]) EvalInfoService[T]
}

type EvalInfoService[T any] interface {
	EvalInfo() *EvalInfo[T]
	Replace(ei *EvalInfo[T])
}

type ConfigService[T any] interface {
	Configure(func(*Evaluator[T]))
	Evaluator() EvalInfoService[T]
	Asterisk() []*Column[T]
	PKey() *Column[T]
	LogicDelKey() *Column[T]
	Table() *RefTable
	Self() *T
}
