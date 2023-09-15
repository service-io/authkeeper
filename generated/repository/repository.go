// Package repository
// @author tabuyos
// @since 2023/9/13
// @description repository
package repository

import (
	"deepsea/generated/helper"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"strings"
	"time"
)

type autogen struct {
	option *helper.Option
}

func New(option *helper.Option) helper.AutoGenService {
	return &autogen{option: option}
}

func (ag *autogen) RenderAuto() {
	file := jen.NewFile("repository")
	helper.ImportPkg(ag.option.Module, file)

	file.Add(ag.GenInterfaceRepository())
	file.Add(ag.GenStructEntity())
	file.Add(ag.GenFuncSelectOneByConfig())
	file.Add(ag.GenFuncSelectManyByConfig())
	file.Add(ag.GenFuncSelectPageByConfig())
	file.Add(ag.GenFuncInsertByConfig())
	file.Add(ag.GenFuncUpdateByConfig())
	file.Add(ag.GenFuncDeleteByConfig())
	file.Add(ag.GenFuncSelectByID())
	file.Add(ag.GenFuncSelectByIDs())
	file.Add(ag.GenFuncBatchSelectByID())
	file.Add(ag.GenFuncSelectByXXX()...)
	file.Add(ag.GenFuncSelectAllWithPage())
	file.Add(ag.GenFuncInsert())
	file.Add(ag.GenFuncInsertNonNil())
	file.Add(ag.GenFuncInsertWithFunc())
	file.Add(ag.GenFuncBatchInsert())
	file.Add(ag.GenFuncBatchInsertNonNil())
	file.Add(ag.GenFuncBatchInsertWithFunc())
	file.Add(ag.GenFuncDeleteByID())
	file.Add(ag.GenFuncDeleteByIDs())
	file.Add(ag.GenFuncBatchDeleteByID())
	file.Add(ag.GenFuncUpdateByID())
	file.Add(ag.GenFuncUpdateNonNilByID())
	file.Add(ag.GenFuncUpdateByIDWithFunc())
	file.Add(ag.GenFuncBatchUpdateByIDWithFunc())
	file.Add(ag.GenFuncGetDBCtx())

	helper.WriteToFile(file, fmt.Sprintf("%s/%s_autogen.go", ag.option.Package.Rty, ag.option.Table), false)
}

func (ag *autogen) RenderSelf() {
	file := jen.NewFile("repository")
	helper.ImportPkg(ag.option.Module, file)

	file.PackageComment("Package repository")
	file.PackageComment("@author tabuyos")
	file.PackageComment("@since " + time.Now().Format("2006/01/02"))
	file.PackageComment("@description " + ag.option.Table)

	file.Add(ag.GenVarPool())
	file.Add(ag.GenInterfaceRty())
	file.Add(ag.GenStructRty())
	file.Add(ag.GenFuncNewRty())

	helper.WriteToFile(file, fmt.Sprintf("%s/%s.go", ag.option.Package.Rty, ag.option.Table), true)
}

func (ag *autogen) RenderBoth() {
	ag.RenderAuto()
	ag.RenderSelf()
}

func (ag *autogen) GenVarPool() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("%sRtyPool 持久池", ag.option.LowerCamel)).Line().Var().Id(fmt.Sprintf("%sRtyPool", ag.option.LowerCamel)).Op("=").Op("&").Add(helper.UseSync("Pool")).
		Values(jen.Id("New").Op(":").
			Func().
			Params().
			Params(jen.Interface()).
			Block(jen.Return().Id("new").Call(jen.Id(fmt.Sprintf("%sRty", ag.option.LowerCamel)))))
}

func (ag *autogen) GenInterfaceRty() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("%sRty 持久层接口", ag.option.Entity)).Line().Type().Id(fmt.Sprintf("%sRty", ag.option.Entity)).Interface(jen.Id(fmt.Sprintf("i%sAutoGen", ag.option.Entity)))
}

func (ag *autogen) GenStructRty() jen.Code {
	return jen.Type().Id(fmt.Sprintf("%sRty", ag.option.LowerCamel)).Struct(jen.Id(fmt.Sprintf("%sAutoGen", ag.option.LowerCamel)))
}

func (ag *autogen) GenFuncNewRty() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("New%sRty 从池中创建", ag.option.Entity)).Line().
		Func().Id(fmt.Sprintf("New%sRty", ag.option.Entity)).Params(jen.Id("ctx").Op("*").Add(helper.UseGin("Context"))).
		Params(jen.Id(fmt.Sprintf("%sRty", ag.option.Entity)), jen.Func().Params()).
		Block(jen.Id("rty").Op(":=").Id(fmt.Sprintf("%sRtyPool", ag.option.LowerCamel)).Dot("Get").Call().
			Assert(jen.Op("*").Id(fmt.Sprintf("%sRty", ag.option.LowerCamel))),
			jen.Id("rty").Dot("ctx").Op("=").Id("ctx"),
			jen.Id("rel").Op(":=").Func().Params().Block(jen.Id("rty").Dot("ctx").Op("=").Id("nil"),
				jen.Id(fmt.Sprintf("%sRtyPool", ag.option.LowerCamel)).Dot("Put").Call(jen.Id("rty"))),
			jen.Return().List(jen.Id("rty"),
				jen.Id("rel")))
}

func (ag *autogen) GenInterfaceRepository() jen.Code {
	var idCodes []jen.Code
	var strCodes []jen.Code

	for _, col := range ag.option.Cols {
		if strings.HasSuffix(col.ColumnName, "_id") {
			cn := strcase.ToCamel(strings.TrimSuffix(col.ColumnName, "_id"))
			code := jen.Line().Line().Comment(fmt.Sprintf("SelectBy%sID Query by %s ID", cn, cn)).Line().Id(fmt.Sprintf("SelectBy%sID", cn)).Params(jen.Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)))
			idCodes = append(idCodes, code)
		}
		if !strings.HasSuffix(col.ColumnName, "_id") && helper.TypeMappingMysqlToGo[col.Type] == "string" {
			cn := strcase.ToCamel(col.ColumnName)
			code := jen.Line().Line().Comment(fmt.Sprintf("SelectBy%s Query by %s", cn, cn)).Line().Id(fmt.Sprintf("SelectBy%s", cn)).Params(jen.Id("string")).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)))
			strCodes = append(strCodes, code)
		}
	}

	return jen.Line().Comment(fmt.Sprintf("i%sAutoGen 该接口自动生成, 请勿修改", ag.option.Entity)).Line().
		Type().Id(fmt.Sprintf("i%sAutoGen", ag.option.Entity)).
		Interface(
			jen.Comment("SelectOneByConfig Use config service to execute(select statement)").Line().Id("SelectOneByConfig").Params(helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity))),
			jen.Line().Comment("SelectManyByConfig Use config service to execute(select statement)").Line().Id("SelectManyByConfig").Params(helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))),
			jen.Line().Comment("SelectPageByConfig Use config service to execute(select statement)").Line().Id("SelectPageByConfig").Params(helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Int64()),
			jen.Line().Comment("InsertByConfig Use config service to execute(insert statement)").Line().Id("InsertByConfig").Params(jen.Op("*").Add(helper.UseSQLTx()), helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Bool(),
			jen.Line().Comment("UpdateByConfig Use config service to execute(update statement)").Line().Id("UpdateByConfig").Params(jen.Op("*").Add(helper.UseSQLTx()), helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Bool(),
			jen.Line().Comment("DeleteByConfig Use config service to execute(delete statement)").Line().Id("DeleteByConfig").Params(jen.Op("*").Add(helper.UseSQLTx()), helper.UseIrisConfigServiceCode(helper.UseEntity(ag.option.Entity))).Bool(),

			jen.Line().Comment("SelectByID Query and return the related entries of id(p key)").Line().Id("SelectByID").Params(jen.Int64()).Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity))),
			jen.Line().Comment("SelectByIDs Query and return the related entries of id(p key list).").Line().Comment("Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID").Line().Id("SelectByIDs").Params(jen.Op("...").Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))),
			jen.Line().Comment("BatchSelectByID Query and return the related entries of id(p key list).").Line().Comment("Call BatchSelectByID if the total number of id is more than one").Line().Id("BatchSelectByID").Params(jen.Index().Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))),

			jen.Add(idCodes...),
			jen.Add(strCodes...),

			jen.Line().Comment("SelectAllWithPage Query all entries by page").Line().Id("SelectAllWithPage").Params(jen.Int64(), jen.Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Int64()),
			jen.Line().Comment("Insert Matches all columns and insert into the table").Line().Id("Insert").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).Int64(),
			jen.Line().Comment("InsertNonNil Matches non-nil columns and insert into the table").Line().Id("InsertNonNil").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).Int64(),
			jen.Line().Comment("InsertWithFunc Matches all columns matching this function and insert into the table").Line().Id("InsertWithFunc").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Func().Params(jen.Op("*").Add(helper.UseIrisColumnCode(helper.UseEntity(ag.option.Entity))), jen.Id("any")).Bool()).Int64(),
			jen.Line().Comment("BatchInsert Batch insert all columns of table into table").Line().Id("BatchInsert").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Int64()),
			jen.Line().Comment("BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table").Line().Id("BatchInsertNonNil").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Int64()),
			jen.Line().Comment("BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table").Line().Id("BatchInsertWithFunc").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Func().Params(jen.Op("*").Add(helper.UseIrisColumnCode(helper.UseEntity(ag.option.Entity))), jen.Id("any")).Bool()).Params(jen.Index().Int64()),
			jen.Line().Comment("DeleteByID Delete the related entries of id(p key)").Line().Id("DeleteByID").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Int64()).Bool(),
			jen.Line().Comment("DeleteByIDs Delete the related entries of id(p key list)").Line().Id("DeleteByIDs").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("...").Int64()).Bool(),
			jen.Line().Comment("BatchDeleteByID Batch delete the related entries of id(p key list)").Line().Id("BatchDeleteByID").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Index().Int64()).Bool(),
			jen.Line().Comment("UpdateByID Update all columns of the related entries of id(p key)").Line().Id("UpdateByID").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).Bool(),
			jen.Line().Comment("UpdateNonNilByID Update non-nil columns of the related entries of id(p key)").Line().Id("UpdateNonNilByID").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).Bool(),
			jen.Line().Comment("UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)").Line().Id("UpdateByIDWithFunc").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Func().Params(jen.Op("*").Add(helper.UseIrisColumnCode(helper.UseEntity(ag.option.Entity))), jen.Id("any")).Bool()).Bool(),
			jen.Line().Comment("BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)").Line().Id("BatchUpdateByIDWithFunc").Params(jen.Op("*").Add(helper.UseSQLTx()), jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)), jen.Func().Params(jen.Op("*").Add(helper.UseIrisColumnCode(helper.UseEntity(ag.option.Entity))), jen.Id("any")).Bool()).Bool())
}

func (ag *autogen) GenStructEntity() jen.Code {
	cn := strcase.ToLowerCamel(ag.option.Entity)
	return jen.Line().Comment(fmt.Sprintf("%sAutoGen 该结构体自动生成, 请勿修改", cn)).Line().Type().Id(fmt.Sprintf("%sAutoGen", cn)).Struct(jen.Id("ctx").Op("*").Add(helper.UseGinCtx()))
}

func (ag *autogen) funcPrefix(codes ...jen.Code) *jen.Statement {
	return jen.Add(codes...).Line().Func().Params(jen.Id("ag").Op("*").Id(fmt.Sprintf("%sAutoGen", strcase.ToLowerCamel(ag.option.Entity))))
}

func (ag *autogen) GenFuncSelectOneByConfig() jen.Code {
	return ag.funcPrefix().Id("SelectOneByConfig").Params(helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).
		Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().Nil()),
			jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
			jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().Nil()),
			jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
			jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
			jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
			jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
			jen.Id("db").Op(":=").Add(helper.UseFetchDB()),
			jen.List(jen.Id("stmt"),
				jen.Id("err")).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("execSQL")),
			jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
				jen.Id("recorder").Dot("MaybePanic")),
			jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
			jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
				jen.Id("values").Op("...")),
			jen.Id(ag.option.Variable).Op(":=").Add(helper.UseRow()).Call(jen.Id("row"),
				jen.Func().Params().Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity)),
					jen.Index().Id("any")).Block(jen.Id(ag.option.Variable).Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
					jen.Id("mappers").Op(":=").Id("evalInfo").Dot("MapperRows").Call(jen.Id(ag.option.Variable)),
					jen.Return().List(jen.Id(ag.option.Variable),
						jen.Id("mappers")))),
			jen.Return().Id(ag.option.Variable))
}

func (ag *autogen) GenFuncSelectManyByConfig() jen.Code {
	return ag.funcPrefix().Id("SelectManyByConfig").Params(helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).
		Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().Nil()),
		jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
		jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().Nil()),
		jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
		jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
		jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
		jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
		jen.Id("db").Op(":=").Add(helper.UseFetchDB()),
		jen.List(jen.Id("stmt"),
			jen.Id("err")).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("execSQL")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.List(jen.Id("rows"),
			jen.Id("err")).Op(":=").Id("stmt").Dot("QueryContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
			jen.Id("values").Op("...")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Id(ag.option.Variables).Op(":=").Add(helper.UseRows()).Call(jen.Id("rows"),
			jen.Func().Params().Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity)),
				jen.Index().Id("any")).Block(jen.Id(ag.option.Variable).Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
				jen.Id("mappers").Op(":=").Id("evalInfo").Dot("MapperRows").Call(jen.Id(ag.option.Variable)),
				jen.Return().List(jen.Id(ag.option.Variable),
					jen.Id("mappers")))),
		jen.Return().Id(ag.option.Variables))
}

func (ag *autogen) GenFuncSelectPageByConfig() jen.Code {
	return ag.funcPrefix().Id("SelectPageByConfig").Params(helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Int64()).Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().List(jen.Nil(),
		jen.Lit(0))),
		jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
		jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().List(jen.Nil(),
			jen.Lit(0))),
		jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
		jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
		jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
		jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
		jen.List(jen.Id("tx"),
			jen.Id("err")).Op(":=").Add(helper.UseFetchDB()).Dot("Begin").Call(),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Defer().Add(helper.UseHandleTx()).Call(jen.Id("tx"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.List(jen.Id("stmt"),
			jen.Id("err")).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("execSQL")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.List(jen.Id("rows"),
			jen.Id("err")).Op(":=").Id("stmt").Dot("QueryContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
			jen.Id("values").Op("...")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Id(ag.option.Variables).Op(":=").Add(helper.UseRows()).Call(jen.Id("rows"),
			jen.Func().Params().Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity)),
				jen.Index().Id("any")).Block(jen.Id(ag.option.Variable).Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
				jen.Id("mappers").Op(":=").Id("evalInfo").Dot("MapperRows").Call(jen.Id(ag.option.Variable)),
				jen.Return().List(jen.Id(ag.option.Variable),
					jen.Id("mappers")))),
		jen.If(jen.Id("evalInfo").Dot("Pageable").Call()).Block(jen.Id("totalSQL").Op(":=").Id("evalInfo").Dot("TotalSQL").Call(),
			jen.List(jen.Id("stmt"),
				jen.Id("err")).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("totalSQL")),
			jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
				jen.Id("recorder").Dot("MaybePanic")),
			jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
			jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
				jen.Id("values").Op("...")),
			jen.Id("total").Op(":=").Add(helper.UseRow()).Call(jen.Id("row"),
				jen.Func().Params().Params(jen.Op("*").Op("*").Int64(),
					jen.Index().Id("any")).Block(jen.Null().Var().Id("r").Op("*").Int64(),
					jen.Null().Var().Id("cs").Op("=").Index().Id("any").Values(jen.Op("&").Id("r")),
					jen.Return().List(jen.Op("&").Id("r"),
						jen.Id("cs")))),
			jen.If(jen.Op("*").Id("total").Op("==").Nil()).Block(jen.Return().List(jen.Id(ag.option.Variables),
				jen.Lit(0))),
			jen.Return().List(jen.Id(ag.option.Variables),
				jen.Op("*").Op("*").Id("total"))),
		jen.Return().List(jen.Id(ag.option.Variables),
			jen.Lit(0)))
}

func (ag *autogen) GenFuncInsertByConfig() jen.Code {
	return ag.funcPrefix().Id("InsertByConfig").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()), helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).
		Params(jen.Id("bool")).Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
		jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
		jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
		jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
		jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
		jen.Id("self").Op(":=").Id("config").Dot("Self").Call(),
		jen.List(jen.Id("stmt"),
			jen.Id("err")).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("execSQL")),
		jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.List(jen.Id("result"),
			jen.Id("err")).Op(":=").Id("stmt").Dot("ExecContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
			jen.Id("values").Op("...")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.If(jen.Id("self").Dot("ID").Op("!=").Nil()).Block(jen.Return().Id("true")),
		jen.List(jen.Id("id"),
			jen.Id("err")).Op(":=").Id("result").Dot("LastInsertId").Call(),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.If(jen.Id("id").Op("==").Lit(0)).Block(jen.Id("panic").Call(jen.Lit("插入失败"))),
		jen.Id("self").Dot("ID").Op("=").Op("&").Id("id"),
		jen.Return().Id("true"))
}

func (ag *autogen) GenFuncUpdateByConfig() jen.Code {
	return ag.funcPrefix().Id("UpdateByConfig").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()), helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).
		Params(jen.Id("bool")).Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
		jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
		jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
		jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
		jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
		jen.List(jen.Id("stmt"),
			jen.Id("err")).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("execSQL")),
		jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.List(jen.Id("result"),
			jen.Id("err")).Op(":=").Id("stmt").Dot("ExecContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
			jen.Id("values").Op("...")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.List(jen.Id("_"),
			jen.Id("err")).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Return().Id("true"))
}

func (ag *autogen) GenFuncDeleteByConfig() jen.Code {
	return ag.funcPrefix().Id("DeleteByConfig").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()), helper.UseIrisNamedConfigServiceCode(helper.UseEntity(ag.option.Entity))).
		Params(jen.Id("bool")).Block(jen.If(jen.Id("config").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("evaluator").Op(":=").Id("config").Dot("Evaluator").Call(),
		jen.If(jen.Id("evaluator").Op("!=").Nil()).Block(jen.Return().Id("false")),
		jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorderByCtx()),
		jen.Id("evalInfo").Op(":=").Id("evaluator").Dot("EvalInfo").Call(),
		jen.Id("execSQL").Op(":=").Id("evalInfo").Dot("SQL").Call(),
		jen.Id("values").Op(":=").Id("evalInfo").Dot("Values").Call(),
		jen.List(jen.Id("stmt"),
			jen.Id("err")).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("execSQL")),
		jen.Defer().Add(helper.UseDeferClose()).Call(jen.Id("stmt"),
			jen.Id("recorder").Dot("MaybePanic")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.List(jen.Id("result"),
			jen.Id("err")).Op(":=").Id("stmt").Dot("ExecContext").Call(jen.Id("ag").Dot("getDBCtx").Call(),
			jen.Id("values").Op("...")),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.List(jen.Id("_"),
			jen.Id("err")).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
		jen.Return().Id("true"))
}

func (ag *autogen) GenFuncSelectByID() jen.Code {
	return ag.funcPrefix().Id("SelectByID").Params(jen.Id("id").Int64()).Params(jen.Op("*").Add(helper.UseEntity(ag.option.Entity))).
		Block(jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
			jen.Id("config").Dot("Configure").Call(
				jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
					Block(jen.Id("eval").Dot("Select").Call(jen.Id("config").Dot("Asterisk").Call().Op("...")).
						Dot("From").Call(jen.Id("config").Dot("Table").Call()).
						Dot("Where").Call(jen.Id("config").Dot("IDCol").Call().Dot("EQ").Call(jen.Id("id"))).
						Dot("Eval").Call())),
			jen.Return().Id("ag").Dot("SelectOneByConfig").Call(jen.Id("config")))
}

func (ag *autogen) GenFuncSelectByIDs() jen.Code {
	return ag.funcPrefix().Id("SelectByIDs").Params(jen.Id("ids").Op("...").Int64()).
		Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).
		Block(jen.Return().Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")))
}

func (ag *autogen) GenFuncBatchSelectByID() jen.Code {
	return ag.funcPrefix().Id("BatchSelectByID").Params(jen.Id("ids").Index().Int64()).
		Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).
		Block(jen.Id("values").Op(":=").Id("make").Call(jen.Index().Id("any"),
			jen.Id("len").Call(jen.Id("ids"))),
			jen.For(jen.List(jen.Id("i"),
				jen.Id("id")).Op(":=").Range().Id("ids")).Block(jen.Id("values").Index(jen.Id("i")).Op("=").Id("id")),
			jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
			jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
				Block(jen.Id("eval").Dot("Select").Call(jen.Id("config").Dot("Asterisk").Call().Op("...")).
					Dot("From").Call(jen.Id("config").Dot("Table").Call()).
					Dot("Where").Call(jen.Id("config").Dot("IDCol").Call().Dot("IN").Call(jen.Id("values").Op("..."))).
					Dot("Eval").Call())),
			jen.Return().Id("ag").Dot("SelectManyByConfig").Call(jen.Id("config")))
}

func (ag *autogen) GenFuncSelectByXXX() []jen.Code {
	var idCodes []jen.Code
	var strCodes []jen.Code

	for _, col := range ag.option.Cols {
		if strings.HasSuffix(col.ColumnName, "_id") {
			cn := strcase.ToCamel(strings.TrimSuffix(col.ColumnName, "_id"))
			code := ag.funcPrefix(jen.Line(), jen.Line()).Id(fmt.Sprintf("SelectBy%sID", cn)).Params(jen.Id("id").Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).
				Block(jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
					jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
						Block(jen.Id("eval").Dot("Select").Call(jen.Id("config").Dot("Asterisk").Call().Op("...")).
							Dot("From").Call(jen.Id("config").Dot("Table").Call()).
							Dot("Where").Call(jen.Id("config").Dot(fmt.Sprintf("%sIdCol", cn)).Call().Dot("EQ").Call(jen.Id("id"))).Dot("Eval").Call())),
					jen.Return().Id("ag").Dot("SelectManyByConfig").Call(jen.Id("config")))

			idCodes = append(idCodes, code)
		}
		if !strings.HasSuffix(col.ColumnName, "_id") && helper.TypeMappingMysqlToGo[col.Type] == "string" {
			cn := strcase.ToCamel(col.ColumnName)
			code := ag.funcPrefix(jen.Line(), jen.Line()).Id(fmt.Sprintf("SelectBy%s", cn)).Params(jen.Id("val").Id("string")).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).
				Block(jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
					jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
						Block(jen.Id("eval").Dot("Select").Call(jen.Id("config").Dot("Asterisk").Call().Op("...")).
							Dot("From").Call(jen.Id("config").Dot("Table").Call()).
							Dot("Where").Call(jen.Id("config").Dot(fmt.Sprintf("%sCol", cn)).Call().
							Dot("Like").Call(jen.Id("val"))).Dot("Eval").Call())),
					jen.Return().Id("ag").Dot("SelectManyByConfig").Call(jen.Id("config")))
			strCodes = append(strCodes, code)
		}
	}

	return append(idCodes, strCodes...)
}

func (ag *autogen) GenFuncSelectAllWithPage() jen.Code {
	return ag.funcPrefix().Id("SelectAllWithPage").Params(jen.Id("page").Int64(),
		jen.Id("size").Int64()).Params(jen.Index().Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Int64()).Block(jen.If(jen.Id("page").Op("<").Lit(1)).Block(jen.Return().List(jen.Nil(),
		jen.Lit(0))),
		jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
		jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
			Block(jen.Id("eval").Dot("Select").Call(jen.Id("config").Dot("Asterisk").Call().Op("...")).
				Dot("From").Call(jen.Id("config").Dot("Table").Call()).Dot("Page").Call(jen.Id("size"),
				jen.Parens(jen.Id("page").Op("-").Lit(1)).Op("*").Id("size")).Dot("Eval").Call())),
		jen.Return().Id("ag").Dot("SelectPageByConfig").Call(jen.Id("config")))
}

func (ag *autogen) GenFuncInsert() jen.Code {
	return ag.funcPrefix().Id("Insert").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity))).Int64().Block(jen.Return().Id("ag").Dot("InsertWithFunc").Call(jen.Id("tx"),
		jen.Id(ag.option.Variable),
		jen.Func().Params(jen.Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("any")).Bool().Block(jen.Return().Id("true"))))
}

func (ag *autogen) GenFuncInsertNonNil() jen.Code {
	return ag.funcPrefix().Id("InsertNonNil").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity))).Int64().Block(jen.Return().Id("ag").Dot("InsertWithFunc").Call(jen.Id("tx"),
		jen.Id(ag.option.Variable),
		jen.Func().Params(jen.Id("c").Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("v").Id("any")).Bool().Block(jen.Return().Add(helper.UseIsNonNil()).Call(jen.Id("v")))))
}

func (ag *autogen) GenFuncInsertWithFunc() jen.Code {
	return ag.funcPrefix().Id("InsertWithFunc").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Id("fn").Func().Params(jen.Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("any")).Bool()).Int64().
		Block(
			jen.If(jen.Id(ag.option.Variable).Dot("Evaluator").Call().Op("!=").Nil()).
				Block(jen.Id("ag").Dot("InsertByConfig").Call(jen.Id("tx"),
					jen.Id(ag.option.Variable)),
					jen.Return().Op("*").Id(ag.option.Variable).Dot("ID")),
			jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
			jen.List(jen.Id("selfishs"),
				jen.Id("values")).Op(":=").Id(ag.option.Variable).Dot("ColumnAndValue").Call(jen.Id("fn")),
			jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
				Block(jen.Id("eval").Dot("Insert").Call(jen.Id("selfishs").Op("...")).
					Dot("Into").Call(jen.Id("config").Dot("Table").Call()).Dot("Values").Call(jen.Id("values").Op("...")).
					Dot("Eval").Call())),
			jen.Id("ag").Dot("InsertByConfig").Call(jen.Id("tx"), jen.Id("config")),
			jen.Return().Op("*").Id("config").Dot("ID"))
}

func (ag *autogen) GenFuncBatchInsert() jen.Code {
	return ag.funcPrefix().Id("BatchInsert").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variables).Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Int64()).
		Block(jen.Return().Id("ag").Dot("BatchInsertWithFunc").Call(jen.Id("tx"),
			jen.Id(ag.option.Variables),
			jen.Func().Params(jen.Id("c").Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
				jen.Id("v").Id("any")).Bool().Block(jen.Return().Id("true"))))
}

func (ag *autogen) GenFuncBatchInsertNonNil() jen.Code {
	return ag.funcPrefix().Id("BatchInsertNonNil").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variables).Index().Op("*").Add(helper.UseEntity(ag.option.Entity))).Params(jen.Index().Int64()).
		Block(jen.Return().Id("ag").Dot("BatchInsertWithFunc").Call(jen.Id("tx"),
			jen.Id(ag.option.Variables),
			jen.Func().Params(jen.Id("c").Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
				jen.Id("v").Id("any")).Bool().Block(jen.Return().Add(helper.UseIsNonNil()).Call(jen.Id("v")))))
}

func (ag *autogen) GenFuncBatchInsertWithFunc() jen.Code {
	return ag.funcPrefix().Id("BatchInsertWithFunc").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variables).Index().Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Id("fn").Func().Params(jen.Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("any")).Bool()).Params(jen.Index().Int64()).
		Block(jen.If(jen.Id("len").Call(jen.Id(ag.option.Variables)).Op("==").Lit(0)).Block(jen.Return().Nil()),
			jen.Id(ag.option.Variable).Op(":=").Id(ag.option.Variables).Index(jen.Lit(0)),
			jen.Id("ids").Op(":=").Id("make").Call(jen.Index().Int64(),
				jen.Id("len").Call(jen.Id(ag.option.Variables))),
			jen.For(jen.List(jen.Id("i"),
				jen.Id("e")).Op(":=").Range().Id(ag.option.Variables)).Block(jen.Id("ids").Index(jen.Id("i")).Op("=").Op("*").Id("e").Dot("ID")),
			jen.If(jen.Id(ag.option.Variable).Dot("Evaluator").Call().Op("!=").Nil()).Block(jen.Id("ag").Dot("InsertByConfig").Call(jen.Id("tx"),
				jen.Id(ag.option.Variable)),
				jen.Return().Id("ids")),
			jen.Id("values").Op(":=").Id("make").Call(jen.Index().Id("any"),
				jen.Lit(0)),
			jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
			jen.List(jen.Id("selfishs"),
				jen.Id("_")).Op(":=").Id(ag.option.Variable).Dot("ColumnAndValue").Call(jen.Id("fn")),
			jen.For(jen.List(jen.Id("_"),
				jen.Id("e")).Op(":=").Range().Id(ag.option.Variables)).Block(jen.List(jen.Id("_"),
				jen.Id("snipValues")).Op(":=").Id("e").Dot("ColumnAndValue").Call(jen.Id("fn")),
				jen.Id("values").Op("=").Id("append").Call(jen.Id("values"),
					jen.Id("snipValues").Op("..."))),
			jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
				Block(jen.Id("eval").Dot("Insert").Call(jen.Id("selfishs").Op("...")).
					Dot("Into").Call(jen.Id("config").Dot("Table").Call()).
					Dot("Values").Call(jen.Id("values").Op("...")).Dot("Eval").Call())),
			jen.Id("ag").Dot("InsertByConfig").Call(jen.Id("tx"),
				jen.Id("config")),
			jen.Return().Id("ids"))
}

func (ag *autogen) GenFuncDeleteByID() jen.Code {
	return ag.funcPrefix().Id("DeleteByID").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id("id").Int64()).Bool().Block(jen.Return().Id("ag").Dot("DeleteByIDs").Call(jen.Id("tx"),
		jen.Id("id")))
}

func (ag *autogen) GenFuncDeleteByIDs() jen.Code {
	return ag.funcPrefix().Id("DeleteByIDs").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id("ids").Op("...").Int64()).Bool().Block(jen.Return().Id("ag").Dot("BatchDeleteByID").Call(jen.Id("tx"),
		jen.Id("ids")))
}

func (ag *autogen) GenFuncBatchDeleteByID() jen.Code {
	return ag.funcPrefix().Id("BatchDeleteByID").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id("ids").Index().Int64()).Bool().Block(jen.Id("values").Op(":=").Id("make").Call(jen.Index().Id("any"),
		jen.Id("len").Call(jen.Id("ids"))),
		jen.For(jen.List(jen.Id("i"),
			jen.Id("id")).Op(":=").Range().Id("ids")).Block(jen.Id("values").Index(jen.Id("i")).Op("=").Id("id")),
		jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
		jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
			Block(jen.Id("eval").Dot("Delete").Call().Dot("From").Call(jen.Id("config").Dot("Table").Call()).
				Dot("Where").Call(jen.Id("config").Dot("IDCol").Call().Dot("IN").Call(jen.Id("values").Op("..."))).Dot("Eval").Call())),
		jen.Return().Id("ag").Dot("DeleteByConfig").Call(jen.Id("tx"),
			jen.Id("config")))
}

func (ag *autogen) GenFuncUpdateByID() jen.Code {
	return ag.funcPrefix().Id("UpdateByID").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity))).Bool().Block(jen.Return().Id("ag").Dot("UpdateByIDWithFunc").Call(jen.Id("tx"),
		jen.Id(ag.option.Variable),
		jen.Func().Params(jen.Id("c").Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("v").Id("any")).Bool().Block(jen.Return().Id("true"))))
}

func (ag *autogen) GenFuncUpdateNonNilByID() jen.Code {
	return ag.funcPrefix().Id("UpdateNonNilByID").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity))).Bool().Block(jen.Return().Id("ag").Dot("UpdateByIDWithFunc").Call(jen.Id("tx"),
		jen.Id(ag.option.Variable),
		jen.Func().Params(jen.Id("c").Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("v").Id("any")).Bool().Block(jen.Return().Add(helper.UseIsNonNil()).Call(jen.Id("v")))))
}

func (ag *autogen) GenFuncUpdateByIDWithFunc() jen.Code {
	return ag.funcPrefix().Id("UpdateByIDWithFunc").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variable).Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Id("fn").Func().Params(jen.Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("any")).Bool()).Bool().Block(jen.If(jen.Id(ag.option.Variable).Dot("Evaluator").Call().Op("!=").Nil()).
		Block(jen.Return().Id("ag").Dot("UpdateByConfig").Call(jen.Id("tx"),
			jen.Id(ag.option.Variable))),
		jen.Id("config").Op(":=").Add(helper.UseEntity(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
		jen.List(jen.Id("selfishs"),
			jen.Id("values")).Op(":=").Id(ag.option.Variable).Dot("ColumnAndValue").Call(jen.Id("fn")),
		jen.Id("config").Dot("Configure").Call(jen.Func().Params(jen.Id("eval").Op("*").Id("iris").Dot("Evaluator").Index(helper.UseEntity(ag.option.Entity))).
			Block(jen.Id("eval").Dot("UpdateRef").Call(jen.Id("config").Dot("Table").Call(),
				jen.Id("selfishs").Op("...")).Dot("SetValues").Call(jen.Id("values").Op("...")).Dot("Eval").Call())),
		jen.Return().Id("ag").Dot("UpdateByConfig").Call(jen.Id("tx"),
			jen.Id("config")))
}

func (ag *autogen) GenFuncBatchUpdateByIDWithFunc() jen.Code {
	return ag.funcPrefix().Id("BatchUpdateByIDWithFunc").Params(jen.Id("tx").Op("*").Add(helper.UseSQLTx()),
		jen.Id(ag.option.Variables).Index().Op("*").Add(helper.UseEntity(ag.option.Entity)),
		jen.Id("fn").Func().Params(jen.Op("*").Id("iris").Dot("Column").Index(helper.UseEntity(ag.option.Entity)),
			jen.Id("any")).Bool()).Bool().Block(jen.For(jen.List(jen.Id("_"),
		jen.Id(ag.option.Variable)).Op(":=").Range().Id(ag.option.Variables)).Block(jen.Id("ag").Dot("UpdateByIDWithFunc").Call(jen.Id("tx"),
		jen.Id(ag.option.Variable),
		jen.Id("fn"))),
		jen.Return().Id("true"))
}

func (ag *autogen) GenFuncGetDBCtx() jen.Code {
	return ag.funcPrefix().Id("getDBCtx").Params().Params(helper.UseContextContext()).Block(jen.Return().Add(helper.UseContext("Background")).Call())
}
