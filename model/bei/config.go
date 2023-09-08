// Package bei
// @author tabuyos
// @since 2023/9/8
// @description bei
package bei

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
	Asterisk() []*FD[T]
	PKey() *FD[T]
	LogicDelKey() *FD[T]
	Table() *RefTable
	Self() *T
}
