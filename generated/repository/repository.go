// Package repository
// @author tabuyos
// @since 2023/9/13
// @description repository
package repository

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

func (ag *autogen) GenInterfaceRepository() jen.Code {
	return nil
}

func (ag *autogen) GenStructEntity() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectOneByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectManyByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectAccountByRoleID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectRoleByAccountID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectPageByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsertByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncDeleteByConfig() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectByIDs() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchSelectByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectByXXXID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectByXXXStr() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectAllWithPage() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsert() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsertNonNil() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsertWithFunc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchInsert() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchInsertNonNil() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchInsertWithFunc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncDeleteByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncDeleteByIDs() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchDeleteByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateNonNilByID() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateByIDWithFunc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncBatchUpdateByIDWithFunc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncGetDBCtx() jen.Code {
	return nil
}
