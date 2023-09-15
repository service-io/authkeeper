// Package validator
// @author tabuyos
// @since 2023/8/17
// @description validator
package validator

import (
	"deepsea/model/page"
)

type IValidator[T any] interface {
	ValidateAdd(info T)
	ValidateRemove(id int64)
	ValidateModify(info T)
	ValidateFind(id int64)
	ValidateFindWithPage(query page.Query)
}

type AutoGenValidator[T any] struct{}

func (rec *AutoGenValidator[T]) ValidateAdd(info T) {
}

func (rec *AutoGenValidator[T]) ValidateRemove(id int64) {
}

func (rec *AutoGenValidator[T]) ValidateModify(info T) {
}

func (rec *AutoGenValidator[T]) ValidateFind(id int64) {
}

func (rec *AutoGenValidator[T]) ValidateFindWithPage(query page.Query) {
}
