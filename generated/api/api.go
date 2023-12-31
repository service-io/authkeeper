package api

import (
	"deepsea/generated/helper"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"time"
)

type autogen struct {
	option *helper.Option
}

func New(option *helper.Option) helper.AutoGenService {
	return &autogen{option: option}
}

func (ag *autogen) RenderAuto() {
	file := jen.NewFile("api")
	helper.ImportPkg(ag.option.Module, file)

	file.HeaderComment("Code generated by tabuyos. DO NOT EDIT.")
	file.PackageComment("Package api")
	file.PackageComment("@author tabuyos")
	file.PackageComment("@since " + time.Now().Format("2006/01/02"))
	file.PackageComment("@description " + ag.option.Table)

	file.Add(ag.GenStructHandler())
	file.Add(ag.GenFuncNewHandler())
	file.Add(ag.GenFuncCreate())
	file.Add(ag.GenFuncRemove())
	file.Add(ag.GenFuncModify())
	file.Add(ag.GenFuncDetail())
	file.Add(ag.GenFuncListPage())

	helper.WriteToFile(file, fmt.Sprintf("%s/%s_autogen.go", ag.option.Package.Api, ag.option.Table), false)
}

func (ag *autogen) RenderSelf() {
	file := jen.NewFile("api")
	helper.ImportPkg(ag.option.Module, file)

	file.PackageComment("Package api")
	file.PackageComment("@author tabuyos")
	file.PackageComment("@since " + time.Now().Format("2006/01/02"))
	file.PackageComment("@description " + ag.option.Table)

	// file.Add(ag.GenFuncWhoami())

	helper.WriteToFile(file, fmt.Sprintf("%s/%s.go", ag.option.Package.Api, ag.option.Table), true)
}

func (ag *autogen) RenderBoth() {
	ag.RenderAuto()
	ag.RenderSelf()
}

func (ag *autogen) GenStructHandler() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("%sHandler API 处理器", ag.option.Entity)).Line().Type().Id(fmt.Sprintf("%sHandler", ag.option.Entity)).Struct()
}

func (ag *autogen) GenFuncNewHandler() jen.Code {
	return jen.Line().Comment(fmt.Sprintf("New%sHandler 创建 API 处理器", ag.option.Entity)).Line().Func().Id(fmt.Sprintf("New%sHandler", ag.option.Entity)).Params().
		Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).
		Block(jen.Return().Op("&").Id(fmt.Sprintf("%sHandler", ag.option.Entity)).Values())
}

func (ag *autogen) GenFuncCreate() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("Add 新增数据")).Line().
		Comment(fmt.Sprintf("@Summary      新增数据")).Line().
		Comment(fmt.Sprintf("@Description  新增数据")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		Comment(fmt.Sprintf("@Param        req_info    body     dto.%s  true  \"待新增的数据对象\"", ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply  \"操作结果\"")).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/add [put]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("Add").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorder()).Call(jen.Id("ctx")),
				jen.Id(fmt.Sprintf("new%s", ag.option.Entity)).Op(":=").Add(helper.UseDto(fmt.Sprintf("New%s", ag.option.Entity))).Call(),
				jen.Id("err").Op(":=").Id("ctx").Dot("ShouldBindJSON").Call(jen.Id(fmt.Sprintf("new%s", ag.option.Entity))),
				jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
				jen.List(jen.Id("svc"),
					jen.Id("release")).Op(":=").Add(helper.UseService(ag.option.Module, fmt.Sprintf("New%sSvc", ag.option.Entity))).Call(jen.Id("ctx")),
				jen.Defer().Id("release").Call(),
				jen.Id("id").Op(":=").Id("svc").Dot("Add").Call(jen.Id(fmt.Sprintf("new%s", ag.option.Entity))),
				jen.If(jen.Id("id").Op("!=").Lit(0)).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("OkPayload")).Call(jen.Id("id")))).Else().
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("FailedMessage")).Call(jen.Lit("新增失败"))))))
}

func (ag *autogen) GenFuncRemove() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("Remove 根据 ID 删除数据")).Line().
		Comment(fmt.Sprintf("@Summary      删除数据")).Line().
		Comment(fmt.Sprintf("@Description  根据 ID 删除数据")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		Comment(fmt.Sprintf("@Param        id    query     integer  true  \"待删除 ID\"")).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply  \"操作结果\"")).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/remove [delete]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("Remove").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorder()).Call(jen.Id("ctx")),
				jen.List(jen.Id("id"),
					jen.Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(jen.Id("ctx").Dot("Query").Call(jen.Lit("id")),
					jen.Lit(10),
					jen.Lit(64)),
				jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
				jen.If(jen.Id("id").Op("<=").Lit(0)).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("FailedMessage")).Call(jen.Lit("请传递正确的 ID"))),
						jen.Return()),
				jen.List(jen.Id("svc"),
					jen.Id("release")).Op(":=").Add(helper.UseService(ag.option.Module, fmt.Sprintf("New%sSvc", ag.option.Entity))).Call(jen.Id("ctx")),
				jen.Defer().Id("release").Call(),
				jen.Id("op").Op(":=").Id("svc").Dot("Remove").Call(jen.Id("id")),
				jen.If(jen.Id("op")).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("Ok")).Call().Dot("WithState").Call(jen.Id("op")).Dot("WithMessage").Call(jen.Lit("操作成功"))),
						jen.Return()),
				jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
					jen.Add(helper.UseReply("Failed")).Call().Dot("WithState").Call(jen.Id("op")).Dot("WithMessage").Call(jen.Lit("操作失败")))))
}

func (ag *autogen) GenFuncModify() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("Modify 根据 ID 修改数据")).Line().
		Comment(fmt.Sprintf("@Summary      修改数据")).Line().
		Comment(fmt.Sprintf("@Description  根据 ID 修改数据")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		Comment(fmt.Sprintf("@Param        req_info    body     dto.%s  true  \"待修改的数据对象\"", ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply  \"操作结果\"")).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/base-modify [post]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("Modify").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorder()).Call(jen.Id("ctx")),
				jen.Id("new"+ag.option.Entity).Op(":=").Add(helper.UseDto("New"+ag.option.Entity)).Call(),
				jen.Id("err").Op(":=").Id("ctx").Dot("ShouldBindJSON").Call(jen.Id("new"+ag.option.Entity)),
				jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
				jen.If(jen.Id("new"+ag.option.Entity).Dot("ID").Op("==").Id("nil")).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("FailedMessage")).Call(jen.Lit("请传递正确的 ID, 以及需要修改的信息"))),
						jen.Return()),
				jen.List(jen.Id("svc"),
					jen.Id("release")).Op(":=").Add(helper.UseService(ag.option.Module, fmt.Sprintf("New%sSvc", ag.option.Entity))).Call(jen.Id("ctx")),
				jen.Defer().Id("release").Call(),
				jen.Id("op").Op(":=").Id("svc").Dot("Modify").Call(jen.Id("new"+ag.option.Entity)),
				jen.If(jen.Id("op")).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("OkMessage")).Call(jen.Lit("操作成功"))),
						jen.Return()),
				jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
					jen.Add(helper.UseReply("FailedMessage")).Call(jen.Lit("操作失败")))))
}

func (ag *autogen) GenFuncDetail() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("Detail 根据 ID 获取详情")).Line().
		Comment(fmt.Sprintf("@Summary      获取详情")).Line().
		Comment(fmt.Sprintf("@Description  根据 ID 获取详情")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		Comment(fmt.Sprintf("@Param        id    query     integer  true  \"查询 ID\"")).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply{payload=dto.%s}  \"查询详情\"", ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/detail [get]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("Detail").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorder()).Call(jen.Id("ctx")),
				jen.List(jen.Id("id"),
					jen.Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(jen.Id("ctx").Dot("Query").Call(jen.Lit("id")),
					jen.Lit(10),
					jen.Lit(64)),
				jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
				jen.If(jen.Id("id").Op("<=").Lit(0)).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("FailedMessage")).Call(jen.Lit("请传递正确的 ID"))),
						jen.Return()),
				jen.List(jen.Id("svc"),
					jen.Id("release")).Op(":=").Add(helper.UseService(ag.option.Module, fmt.Sprintf("New%sSvc", ag.option.Entity))).Call(jen.Id("ctx")),
				jen.Defer().Id("release").Call(),
				jen.Id(ag.option.LowerCamel).Op(":=").Id("svc").Dot("Find").Call(jen.Id("id")),
				jen.If(jen.Id(ag.option.LowerCamel).Op("!=").Id("nil")).
					Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
						jen.Add(helper.UseReply("OkPayload")).Call(jen.Id(ag.option.LowerCamel))),
						jen.Return()),
				jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
					jen.Add(helper.UseReply("Failed")).Call().Dot("WithMessage").Call(jen.Lit("无对应数据")))))
}

func (ag *autogen) GenFuncListPage() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("ListPage 分页列表")).Line().
		Comment(fmt.Sprintf("@Summary      分页列表")).Line().
		Comment(fmt.Sprintf("@Description  获取分页列表")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		Comment(fmt.Sprintf("@Param        req_info    body     page.Query  false  \"分页信息\"")).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply{payload=[]page.Result}  \"分页列表\"")).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/list-page [post]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("ListPage").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("recorder").Op(":=").Add(helper.UseFetchRecorder()).Call(jen.Id("ctx")),
				jen.Id("query").Op(":=").Add(helper.UsePage("NewQuery")).Call(),
				jen.Id("query").Dot("Page").Op("=").Lit(1),
				jen.Id("query").Dot("Size").Op("=").Lit(20),
				jen.Id("err").Op(":=").Id("ctx").Dot("ShouldBindJSON").Call(jen.Id("query")),
				jen.Id("recorder").Dot("MaybePanic").Call(jen.Id("err")),
				jen.List(jen.Id("svc"),
					jen.Id("release")).Op(":=").Add(helper.UseService(ag.option.Module, fmt.Sprintf("New%sSvc", ag.option.Entity))).Call(jen.Id("ctx")),
				jen.Defer().Id("release").Call(),
				jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")),
					jen.Add(helper.UseReply("OkPayload")).Call(jen.Id("svc").Dot("FindWithPage").Call(jen.Op("*").Id("query"))))))
}

func (ag *autogen) GenFuncWhoami() jen.Code {
	return jen.Line().
		Comment(fmt.Sprintf("Whoami 示例")).Line().
		Comment(fmt.Sprintf("@Summary      示例")).Line().
		Comment(fmt.Sprintf("@Description  示例")).Line().
		// Comment(fmt.Sprintf("@Tags         %s,%s", ag.option.Module, ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Tags         %s", ag.option.RenderTag())).Line().
		Comment(fmt.Sprintf("@Accept       json")).Line().
		Comment(fmt.Sprintf("@Produce      json")).Line().
		// Comment(fmt.Sprintf("@Param        req_info    body     dto.%s  true  \"示例对象\"", ag.option.Entity)).Line().
		Comment(fmt.Sprintf("@Success      200  {object}  reply.Reply  \"操作结果\"")).Line().
		Comment(fmt.Sprintf("@Security     ApiKeyAuth")).Line().
		Comment(fmt.Sprintf("@Router       /%s/%s/whoami [get]", strcase.ToKebab(ag.option.Module), strcase.ToKebab(ag.option.GetRouterPrefix()))).Line().
		Func().Params(jen.Op("*").Id(fmt.Sprintf("%sHandler", ag.option.Entity))).Id("Whoami").Params().Params(helper.UseGinHandlerFunc()).
		Block(jen.Return().Func().Params(jen.Id("ctx").Op("*").Add(helper.UseGinCtx())).
			Block(jen.Id("ctx").Dot("JSON").Call(jen.Add(helper.UseHttp("StatusOK")), jen.Add(helper.UseReply("OkPayload")).Call(jen.Lit("你好")))))
}
