// Package entity
// @author tabuyos
// @since 2023/8/29
// @description entity
package entity

import (
	"deepsea/generated/helper"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"strings"
)

type autogen struct {
	option *helper.Option
}

func New(option *helper.Option) helper.AutoGenService {
	return &autogen{option: option}
}

func (ag *autogen) RenderAuto() {
	file := jen.NewFile("entity")
	helper.ImportPkg(ag.option.Module, file)

	file.Add(ag.GenStructEntity())
	file.Add(ag.GenFuncNew())
	file.Add(ag.GenFuncXCol()...)
	file.Add(ag.GenFuncConfigure())
	file.Add(ag.GenFuncColumnAndValue())
	file.Add(ag.GenFuncAsterisk())
	file.Add(ag.GenFuncPKey())
	file.Add(ag.GenFuncLogicDelKey())
	file.Add(ag.GenFuncEvaluator())
	file.Add(ag.GenFuncTable())
	file.Add(ag.GenFuncEnableDecorate())
	file.Add(ag.GenFuncDisableDecorate())
	file.Add(ag.GenFuncSelf())

	helper.WriteToFile(file, fmt.Sprintf("%s/%s_autogen.go", ag.option.Package.Ety, ag.option.Table), false)
}

func (ag *autogen) RenderSelf() {
}

func (ag *autogen) RenderBoth() {
	ag.RenderAuto()
	ag.RenderSelf()
}

func (ag *autogen) GenStructEntity() jen.Code {
	var codes []jen.Code //nolint:prealloc
	for _, column := range ag.option.Cols {
		cn := strcase.ToCamel(column.ColumnName)
		tags := make(map[string]string)
		tags["json"] = strcase.ToLowerCamel(strings.ToUpper(column.ColumnName))
		comment := column.ColumnComment
		if len(comment) == 0 {
			comment = column.ColumnName
		}
		codes = append(codes, jen.Comment(comment))
		ct := helper.TypeMappingMysqlToGo[column.Type]
		if ct == "time.Time" {
			codes = append(codes, jen.Id(cn).Op("*").Add(helper.UseTime("Time")).Tag(tags))
			continue
		}
		codes = append(codes, jen.Id(cn).Op("*").Add(jen.Id(ct)).Tag(tags))
	}

	codes = append(codes, jen.Line().Id("evaluator").Op("*").Add(helper.UseCellarEvaluator(ag.option.Entity)), jen.Id("decorate").Bool())
	return jen.Type().Id(ag.option.Entity).Struct(codes...)
}

func (ag *autogen) GenFuncNew() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("New%s 初始化", ag.option.Entity)).Line().Func().Id("New" + ag.option.Entity).Params().Op("*").Id(ag.option.Entity).Block(
		jen.Return(jen.Op("&").Id(ag.option.Entity).Values()),
	)
}

func (ag *autogen) GenFuncXCol() []jen.Code {
	var codes = make([]jen.Code, len(ag.option.Cols))
	for i, col := range ag.option.Cols {
		cn := strcase.ToCamel(col.ColumnName)
		code := jen.Line().Comment(fmt.Sprintf("%sCol %s 列", cn, cn)).Line().Func().
			Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id(fmt.Sprintf("%sCol", cn)).
			Params().Op("*").Add(helper.UseCellarColumn(ag.option.Entity)).Block(
			jen.Return(jen.Add(helper.UseCellar("WithColumn")).
				Call(
					jen.Lit(fmt.Sprintf("`%s`", col.ColumnName)),
					jen.Func().Params(jen.Id("rec").Op("*").
						Id(ag.option.Entity)).Params(jen.Any()).
						Block(jen.Return(jen.Add(helper.RenderAndField("rec", cn)))),
					jen.Func().Params(jen.Id("key").String()).String().
						Block(jen.If(jen.Op("!").Id("e").Dot("decorate")).Block(jen.Return(jen.Id("key"))),
							jen.Return(jen.Lit(fmt.Sprintf("`%s`.", ag.option.Table)).Op("+").Id("key"))),
				)),
		)

		codes[i] = code
	}
	return codes
}

func (ag *autogen) GenFuncConfigure() jen.Code {
	return jen.Line().Comment("Configure evaluator 配置").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("Configure").Params(jen.Id("fn").Func().Params(jen.Op("*").Add(helper.UseCellarEvaluator(ag.option.Entity)))).Block(
		jen.If(jen.Id("e").Dot("evaluator").Op("==").Nil()).Block(
			jen.Id("e").Dot("evaluator").Op("=").Add(helper.UseCellar("WithLogicalEvaluator")).Types(jen.Id(ag.option.Entity)).Call(),
		),
		jen.Id("fn").Call(jen.Id("e").Dot("evaluator")),
	)
}

func (ag *autogen) GenFuncColumnAndValue() jen.Code {
	var codes = make([]jen.Code, len(ag.option.Cols))

	for i, col := range ag.option.Cols {
		cn := strcase.ToCamel(col.ColumnName)
		code := jen.Line().If(jen.Id("fn").Call(jen.Id("e").Dot(fmt.Sprintf("%sCol", cn)).Call(), jen.Id("e").Dot(cn))).
			Block(
				jen.Id("selfishs").Op("=").Append(jen.Id("selfishs"), jen.Id("e").Dot(fmt.Sprintf("%sCol", cn)).Call()),
				jen.Id("values").Op("=").Append(jen.Id("values"), jen.Op("*").Id("e").Dot(cn)),
			)
		codes[i] = code
	}

	return jen.Line().Comment("ColumnAndValue 列值计算").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("ColumnAndValue").
		Params(jen.Id("fns").Op("...").Func().Params(jen.Op("*").Add(helper.UseCellarColumn(ag.option.Entity)), jen.Any()).Bool()).
		Params(jen.Id("selfishs").Index().Add(helper.UseCellar("Selfish")), jen.Id("values").Index().Any()).
		Block(
			jen.Id("fn").Op(":=").Func().Params(jen.Op("*").Add(helper.UseCellarColumn(ag.option.Entity)), jen.Any()).Params(jen.Bool()).Block(jen.Return(jen.True())),
			jen.If(jen.Len(jen.Id("fns")).Op(">").Lit(0).Block(jen.Id("fn").Op("=").Id("fns").Index(jen.Lit(0)))),
			jen.Add(codes...),
			jen.Return(),
		)
}

func (ag *autogen) GenFuncAsterisk() jen.Code {
	var codes = make([]jen.Code, len(ag.option.Cols))

	for i, col := range ag.option.Cols {
		cn := strcase.ToCamel(col.ColumnName)
		code := jen.Line().Id("e").Dot(fmt.Sprintf("%sCol", cn)).Call().Dot("Decorate").Call(jen.Id("fn"))
		codes[i] = code
	}
	codes = append(codes, jen.Line())
	return jen.Line().Comment("Asterisk 所有列").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("Asterisk").
		Params(jen.Id("fns").Op("...").Func().
			Params(jen.String()).Params(jen.String())).
		Params(jen.Index().Op("*").Add(helper.UseCellarColumn(ag.option.Entity))).
		Block(
			jen.Var().Id("fn").Func().Params(jen.String()).Params(jen.String()),
			jen.If(jen.Len(jen.Id("fns")).Op(">").Lit(0)).Block(jen.Id("fn").Op("=").Id("fns").Index(jen.Lit(0))),
			jen.Return(jen.Index().Op("*").Add(helper.UseCellarColumn(ag.option.Entity)).Values(codes...)),
		)
}

func (ag *autogen) GenFuncPKey() jen.Code {
	return jen.Line().Comment("PKey 主键").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("PKey").Params().Params(jen.Op("*").Add(helper.UseCellarColumn(ag.option.Entity))).Block(jen.Return(jen.Id("e").Dot("IDCol").Call()))
}

func (ag *autogen) GenFuncLogicDelKey() jen.Code {
	return jen.Line().Comment("LogicDelKey 逻辑删除").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("LogicDelKey").Params().Params(jen.Op("*").Add(helper.UseCellarColumn(ag.option.Entity))).Block(jen.Return(jen.Id("e").Dot("DeletedCol").Call()))
}

func (ag *autogen) GenFuncEvaluator() jen.Code {
	return jen.Line().Comment("Evaluator 计算器").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("Evaluator").Params().Params(jen.Op("*").Add(helper.UseCellar("Evaluator")).Types(jen.Id(ag.option.Entity))).
		Block(
			jen.If(jen.Id("e").Op("==").Nil()).Block(jen.Return(jen.Nil())),
			jen.Return(jen.Id("e").Dot("evaluator")),
		)
}

func (ag *autogen) GenFuncTable() jen.Code {
	return jen.Line().Comment("Table 表").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("Table").Params().Params(jen.Op("*").Add(helper.UseCellarRefTable())).
		Block(jen.Return(jen.Add(helper.UseCellar("WithTable")).Call(jen.Lit(ag.option.Table))))
}

func (ag *autogen) GenFuncEnableDecorate() jen.Code {
	return jen.Line().Comment("EnableDecorate 启用修饰符").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("EnableDecorate").Params().Params(jen.Op("*").Id(ag.option.Entity)).
		Block(jen.Id("e").Dot("decorate").Op("=").True(), jen.Return(jen.Id("e")))
}

func (ag *autogen) GenFuncDisableDecorate() jen.Code {
	return jen.Line().Comment("DisableDecorate 启用修饰符").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("DisableDecorate").Params().Params(jen.Op("*").Id(ag.option.Entity)).
		Block(jen.Id("e").Dot("decorate").Op("=").False(), jen.Return(jen.Id("e")))
}

func (ag *autogen) GenFuncSelf() jen.Code {
	return jen.Line().Comment("Self 原始信息").Line().
		Func().Params(jen.Id("e").Op("*").Id(ag.option.Entity)).Id("Self").Params().Params(jen.Op("*").Id(ag.option.Entity)).Block(jen.Return(jen.Id("e")))
}
