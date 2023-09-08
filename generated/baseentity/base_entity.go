// Package baseentity
// @author tabuyos
// @since 2023/8/29
// @description baseentity
package baseentity

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
	file := jen.NewFile("baseentity")

	file.Add(ag.GenInterfaceInfoService())
	file.Add(ag.GenInterfaceConfigService())
	file.Add(ag.GenInterfaceOperateService())
	file.Add(ag.GenStructBaseEntity())
	file.Add(ag.GenStructFS())
	file.Add(ag.GenFuncOfFS())

	helper.WriteToFile(file, "test/autogen/baseentity.go", false)
}

func (ag *autogen) RenderSelf() {
}

func (ag *autogen) GenMode() jen.Code {
	return nil
}

func (ag *autogen) GenFuncModeString() jen.Code {
	return nil
}

func (ag *autogen) GenConst() jen.Code {
	return nil
}

func (ag *autogen) GenStructCond() jen.Code {
	return nil
}

func (ag *autogen) GenStructPred() jen.Code {
	return nil
}

func (ag *autogen) GenStructOrder() jen.Code {
	return nil
}

func (ag *autogen) GenFuncDesc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncAsc() jen.Code {
	return nil
}

func (ag *autogen) GenFuncOfCond() jen.Code {
	return nil
}

func (ag *autogen) GenFuncOf() jen.Code {
	return nil
}

func (ag *autogen) GenInterfaceInfoService() jen.Code {
	return jen.Id(`
type InfoService[T any] interface {
	Pkey() *FS[T]

	LogicDeleted() (bool, *FS[T])

	TableName() string

	Alias() string
	AliasQualifier()
	SetAlias(alias string)

	AllFS() []*FS[T]

	CalcInsertFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any)
	CalcInsertFields(fns ...func(string, any) bool) (fields []string)
	CalcInsertValues(fns ...func(string, any) bool) (values []any)
	CalcUpdateFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any)
	CalcUpdateFields(fns ...func(string, any) bool) (fields []string)
	CalcUpdateValues(fns ...func(string, any) bool) (values []any)
}
`)
}

func (ag *autogen) GenInterfaceConfigService() jen.Code {
	return jen.Id(`
type ConfigService[T any] interface {
  Configure(configFn func(entity *Evaluator[T]))
  BEI() *Evaluator[T]
}
`)
}

func (ag *autogen) GenInterfaceOperateService() jen.Code {
	return jen.Id(`
type OperateService[T any] interface {
  Insert(fss ...*FS[T]) *Evaluator[T]
  InsertFields() (fields []string, placeholders []string)
  InsertSql(times ...int) (string, []any)

  Update(fss ...*FS[T]) *Evaluator[T]
  UpdateFields() (fields []string)
  UpdateSql() (string, []any)

  WithValues(values ...any)

  WithWhere(pred *Pred)
  Where() (plain string, values []any)

  WithPageable(limit, offset int64)
  Pageable() (bool, int64, int64)

  WithOrder(orders ...*Order)
  Order() (plain []string)

  Select(fss ...*FS[T])
  SelectFS() []*FS[T]
  SelectSql() (pageSql, countSql string, values []any, mappers []func(*T) any)

  Values() []any
}
`)
}

func (ag *autogen) GenStructBaseEntity() jen.Code {
	return jen.Id(`
type Evaluator[T any] struct {
  inserts      []*FS[T]
  updates      []*FS[T]
  selects      []*FS[T]
  pred         *Pred
  orders       []*Order
  values       []any
  table        string
  buf          strings.Builder
  alias        string
  pageable     bool
  limit        int64
  offset       int64
  qualifier    bool
  logicDeleted bool
  deletedKey   string
  deletedVal   string
}
`)
}

func (ag *autogen) GenStructFS() jen.Code {
	return jen.Id(`
type FS[T any] struct {
  Name string
  Fn  func(*T) any
}
`)
}

func (ag *autogen) GenFuncOfFS() jen.Code {
	return jen.Id(`
func OfFS[T any](name string, fn func(*T) any) *FS[T] {
  return &FS[T]{
    Name: name,
    Fn:  fn,
  }
}
`)
}

func (ag *autogen) GenFuncInsert() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsertFields() jen.Code {
	return nil
}

func (ag *autogen) GenFuncInsertSql() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdate() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateFields() jen.Code {
	return nil
}

func (ag *autogen) GenFuncUpdateSql() jen.Code {
	return nil
}

func (ag *autogen) GenFuncWithValues() jen.Code {
	return nil
}

func (ag *autogen) GenFuncWithWhere() jen.Code {
	return nil
}

func (ag *autogen) GenFuncWhere() jen.Code {
	return nil
}

func (ag *autogen) GenFuncWithPageable() jen.Code {
	return nil
}

func (ag *autogen) GenFuncPageable() jen.Code {
	return nil
}

func (ag *autogen) GenFuncWithOrder() jen.Code {
	return nil
}

func (ag *autogen) GenFuncOrder() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelect() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectFS() jen.Code {
	return nil
}

func (ag *autogen) GenFuncSelectSql() jen.Code {
	return nil
}

func (ag *autogen) GenFuncValues() jen.Code {
	return nil
}
