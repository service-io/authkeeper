// Package entity
// @author tabuyos
// @since 2023/8/29
// @description entity
package entity

import (
	"github.com/dave/jennifer/jen"
	"metis/generated/helper"
)

type autogen struct {
	option *helper.Option
}

func New(option *helper.Option) helper.AutoGenService {
	return &autogen{option: option}
}

func (ag *autogen) RenderAuto() {
}

func (ag *autogen) RenderSelf() {
}

func (ag *autogen) GenStructEntity() jen.Code {
	return nil
}

func (ag *autogen) GenFuncNew() jen.Code {
	return nil
}

func (ag *autogen) GenFuncXCol() jen.Code {
	return nil
}

func (ag *autogen) GenFuncConfigure() jen.Code {
	return nil
}

func (ag *autogen) GenFuncColumnAndValue() jen.Code {
	return nil
}

func (ag *autogen) GenFuncAsterisk() jen.Code {
	return nil
}

func (ag *autogen) GenFuncPKey() jen.Code {
	return nil
}

func (ag *autogen) GenFuncLogicDelKey() jen.Code {
	return nil
}

func (ag *autogen) GenFuncEvaluator() jen.Code {
	return nil
}

func (ag *autogen) GenFuncTable() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelf() jen.Code {
	return nil
}
