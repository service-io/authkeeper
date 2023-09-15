// Package helper
// @author tabuyos
// @since 2023/8/29
// @description helper
package helper

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"strings"
)

type AutoGenService interface {
	RenderAuto()
	RenderSelf()
	RenderBoth()
}

type Option struct {
	Table        string
	Module       string
	Entity       string
	LowerCamel   string
	Variable     string
	Variables    string
	routerPrefix string
	EnableShip   bool
	shipKey      string
	Left         *Option
	Right        *Option
	Tags         []string
	Cols         []Column
	Package      *Package
}

type Package struct {
	Dto string
	Ety string
	Rty string
	Svc string
	Api string
}

type Column struct {
	ColumnName      string
	Type            string
	Nullable        string
	TableName       string
	ColumnComment   string
	MaxLength       int64
	NumberPrecision int64
	ColumnType      string
	ColumnKey       string
	Default         interface{}
}

func BaseOption(module string, table string, tags ...string) *Option {
	return &Option{
		Module: module, Table: table, Entity: strcase.ToCamel(table), LowerCamel: strcase.ToLowerCamel(table), Tags: tags, Variable: "eto", Variables: "ets",
		Package: &Package{
			Dto: "model/dto",
			Ety: "model/entity",
			Rty: "module/" + module + "/repository",
			Svc: "module/" + module + "/service",
			Api: "module/" + module + "/api",
		},
	}
}

func (o *Option) RouterPrefix(prefix string) *Option {
	o.routerPrefix = prefix
	return o
}

func (o *Option) ShipKey(key string) *Option {
	o.shipKey = key
	return o
}

func (o *Option) GetShipKey() string {
	return strcase.ToCamel(o.shipKey)
}

func (o *Option) GetRouterPrefix() string {
	if len(o.routerPrefix) == 0 {
		return o.Table
	}
	return o.routerPrefix
}

func (o *Option) JoinShip(left, right *Option) *Option {
	o.EnableShip = true
	o.Left = left
	o.Right = right
	return o
}

func (o *Option) AddTag(tag string) *Option {
	o.Tags = append(o.Tags, tag)
	return o
}

func (o *Option) RenderTag(tags ...string) string {
	o.Tags = append(o.Tags, tags...)
	if len(o.Tags) == 0 {
		o.Tags = append(o.Tags, o.Module, o.Entity)
	}
	return strings.Join(o.Tags, ",")
}

func WriteToFile(f *jen.File, path string, skipExist bool) {
	_, ok := IsExists(path)
	if ok && skipExist {
		return
	}
	fmt.Printf("写入: %#v\n", path)
	if err := os.MkdirAll(filepath.Dir(path), 0766); err != nil {
		panic(err)
	}
	wr, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	_ = f.Render(wr)
}

// IsExists 文件是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

const (
	PKey            = "id"
	RKey            = "right"
	LKey            = "left"
	LlKey           = "level"
	TnKey           = "tree_no"
	DKey            = "deleted"
	NKey            = "name"
	TtKey           = "tenant_id"
	PwdKey          = "pwd"
	AccountTableKey = "plat_account"
	TtCondKey       = "`tenant_id` = ?"
	UdCondKey       = "`deleted` = 0"
	DdCondKey       = "`deleted` = 1"
	CbKey           = "create_by"
	CaKey           = "create_at"
	MbKey           = "modify_by"
	MaKey           = "modify_at"
	QuotedKey       = "`"
)

var TypeMappingMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int8",
	"smallint unsigned":  "int16",
	"mediumint unsigned": "int32",
	"bigint unsigned":    "int64",
	"bit":                "int8",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"json":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // string
	"datetime":           "time.Time", // string
	"timestamp":          "time.Time", // string
	"time":               "time.Time", // string
	"float":              "float32",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "[]byte",
	"varbinary":          "[]byte",
}

const (
	constantPkg    = "deepsea/config/constant"
	databasePkg    = "deepsea/helper/database"
	securityPkg    = "deepsea/helper/security"
	sensitivexPkg  = "deepsea/helper/sensitivex"
	snowflakeidPkg = "deepsea/helper/snowflakeid"
	validatorPkg   = "deepsea/helper/validator"
	recorderxPkg   = "deepsea/helper/recorderx"
	helperPkg      = "deepsea/helper"
	pagePkg        = "deepsea/model/page"
	replyPkg       = "deepsea/model/reply"
	irisPkg        = "deepsea/model/iris"
	dtoPkg         = "deepsea/model/dto"
	entityPkg      = "deepsea/model/entity"
	repositoryPkg  = "deepsea/module/%s/repository"
	servicePkg     = "deepsea/module/%s/service"
)

func ImportPkg(module string, f *jen.File) {
	f.ImportName("context", "context")
	f.ImportName("errors", "errors")
	f.ImportName("fmt", "fmt")
	f.ImportName("time", "time")
	f.ImportName("sync", "sync")
	f.ImportName("strings", "strings")
	f.ImportName("strconv", "strconv")
	f.ImportName("net/http", "http")
	f.ImportName("database/sql", "sql")

	f.ImportName("github.com/jinzhu/copier", "copier")
	f.ImportName("github.com/gin-gonic/gin", "gin")

	f.ImportName(constantPkg, "constant")
	f.ImportName(databasePkg, "database")
	f.ImportName(securityPkg, "security")
	f.ImportName(sensitivexPkg, "sensitivex")
	f.ImportName(snowflakeidPkg, "snowflakeid")
	f.ImportName(validatorPkg, "validator")
	f.ImportName(recorderxPkg, "recorderx")
	f.ImportName(dtoPkg, "dto")
	f.ImportName(entityPkg, "entity")
	f.ImportName(pagePkg, "page")
	f.ImportName(replyPkg, "reply")
	f.ImportName(irisPkg, "iris")

	f.ImportName(fmt.Sprintf(repositoryPkg, module), "repository")
	f.ImportName(fmt.Sprintf(servicePkg, module), "service")

	f.ImportName("deepsea/helper", "helper")
}

func DecoratorField(f string) string {
	return QuotedKey + f + QuotedKey
}

func InferCode(en bool, code jen.Code) jen.Code {
	if en {
		return code
	}
	return jen.Null()
}

func RenderAndField(sn, field string) jen.Code {
	return jen.Op("&").Id(sn).Dot(field)
}

func RenderStarField(sn, field string) jen.Code {
	return jen.Op("*").Id(sn).Dot(field)
}

func RenderField(sn, field string) jen.Code {
	return jen.Id(sn).Dot(field)
}

func use(path, name string) jen.Code {
	return jen.Qual(path, name)
}

func UseIris(name string) jen.Code {
	return use(irisPkg, name)
}

func UseIrisColumn(name string) jen.Code {
	return jen.Add(UseIris("Column")).Types(jen.Id(name))
}

func UseIrisColumnCode(types ...jen.Code) jen.Code {
	return jen.Add(UseIris("Column")).Types(types...)
}

func UseIrisRefTable() jen.Code {
	return jen.Add(UseIris("RefTable"))
}

func UseIrisEvaluator(name string) jen.Code {
	return jen.Add(UseIris("Evaluator")).Types(jen.Id(name))
}

func UseIrisEvaluatorCode(types ...jen.Code) jen.Code {
	return jen.Add(UseIris("Evaluator")).Types(types...)
}

func UseIrisConfigService(name string) jen.Code {
	return jen.Add(UseIris("ConfigService")).Types(jen.Id(name))
}

func UseIrisConfigServiceCode(types ...jen.Code) jen.Code {
	return jen.Add(UseIris("ConfigService")).Types(types...)
}

func UseIrisNamedConfigServiceCode(types ...jen.Code) jen.Code {
	return jen.Id("config").Add(UseIris("ConfigService")).Types(types...)
}

func UseSQL(name string) jen.Code {
	return use("database/sql", name)
}

func UseSQLTx() jen.Code {
	return UseSQL("Tx")
}

func UseHelper(name string) jen.Code {
	return use(helperPkg, name)
}

func UseDeferClose() jen.Code {
	return UseHelper("DeferClose")
}

func UseRow() jen.Code {
	return UseHelper("Row")
}

func UseRows() jen.Code {
	return UseHelper("Rows")
}

func UseHandleTx() jen.Code {
	return UseHelper("HandleTx")
}

func UseIsNonNil() jen.Code {
	return UseHelper("IsNonNil")
}

func UseRecorderX(name string) jen.Code {
	return use(recorderxPkg, name)
}
func UseErrors(name string) jen.Code {
	return use("errors", name)
}

func UseContext(name string) jen.Code {
	return use("context", name)
}

func UseContextContext() jen.Code {
	return UseContext("Context")
}

func UseStrings(name string) jen.Code {
	return use("strings", name)
}

func UseFmt(name string) jen.Code {
	return use("fmt", name)
}

func UseGin(name string) jen.Code {
	return use("github.com/gin-gonic/gin", name)
}

func UseGinCtx() jen.Code {
	return UseGin("Context")
}

func UseGinHandlerFunc() jen.Code {
	return UseGin("HandlerFunc")
}

func UseFetchRecorder() jen.Code {
	return UseRecorderX("FetchRecorder")
}

func UseFetchRecorderByCtx() jen.Code {
	return jen.Add(UseRecorderX("FetchRecorder")).Call(jen.Id("ag").Dot("ctx"))
}

func UseDefaultRecorder() jen.Code {
	return UseRecorderX("DefaultRecorder")
}

func UseDatabase(name string) jen.Code {
	return use(databasePkg, name)
}

func UseFetchDB() jen.Code {
	return jen.Add(UseDatabase("FetchDB")).Call()
}

func GetPwdCode(lowerCamel string) jen.Code {
	return jen.Line().Comment("加密密码信息").Line().If(jen.Add(RenderField(lowerCamel, PwdKey)).Op("!=").Nil()).Block(
		jen.Id("password").Op(":=").Add(UseSecurity("GeneratePassword")).Call(RenderStarField(lowerCamel, PwdKey)),
		jen.Id(lowerCamel).Dot(strcase.ToCamel(PwdKey)).Op("=").Op("&").Id("password"),
	).Line()
}

func IsAccountTable(table string, code jen.Code) jen.Code {
	if table == AccountTableKey {
		return jen.Add(UseSensitivexEraseAccountsSensitiveCall(code))
	}
	return jen.Null()
}

func UseSecurity(name string) jen.Code {
	return use(securityPkg, name)
}

func UseSensitivex(name string) jen.Code {
	return use(sensitivexPkg, name)
}

func UseSensitivexEraseAccountsSensitive() jen.Code {
	return UseSensitivex("EraseAccountsSensitive")
}

func UseSensitivexEraseAccountsSensitiveCall(code jen.Code) jen.Code {
	return jen.Add(UseSensitivexEraseAccountsSensitive()).Call(code)
}

func UseReply(name string) jen.Code {
	return use(replyPkg, name)
}

func UseValidator(name string) jen.Code {
	return use(validatorPkg, name)
}

func UseValidatorIValidator(name string) jen.Code {
	return jen.Add(UseValidator("IValidator")).Types(jen.Op("*").Id("dto").Op(".").Id(name))
}

func UseValidatorAutoGenValidator(name string) jen.Code {
	return jen.Add(UseValidator("AutoGenValidator")).Types(jen.Op("*").Id("dto").Op(".").Id(name))
}

func UseSvcValidateAdd(name string) jen.Code {
	return jen.Add(RenderField("svc", "ValidateAdd")).Call(jen.Id(name))
}

func UseSvcValidateRemove() jen.Code {
	return jen.Add(RenderField("svc", "ValidateRemove")).Call(jen.Id("id"))
}

func UseSvcValidateModify(name string) jen.Code {
	return jen.Add(RenderField("svc", "ValidateModify")).Call(jen.Id(name))
}

func UseSvcValidateFind() jen.Code {
	return jen.Add(RenderField("svc", "ValidateFind")).Call(jen.Id("id"))
}

func UseSvcValidateFindWithPage() jen.Code {
	return jen.Add(RenderField("svc", "ValidateFindWithPage")).Call(jen.Id("query"))
}

func UseStrconv(name string) jen.Code {
	return use("strconv", name)
}

func UseHttp(name string) jen.Code {
	return use("net/http", name)
}

func UseDto(name string) jen.Code {
	return use(dtoPkg, name)
}

func UsePage(name string) jen.Code {
	return use(pagePkg, name)
}

func UseEntity(name string) jen.Code {
	return use(entityPkg, name)
}

func UseService(module, name string) jen.Code {
	return use(fmt.Sprintf(servicePkg, module), name)
}

func UseRepository(module, name string) jen.Code {
	return use(fmt.Sprintf(repositoryPkg, module), name)
}

func UseTime(name string) jen.Code {
	return use("time", name)
}

func UseSync(name string) jen.Code {
	return use("sync", name)
}

func UseSnowflakeidToGenPtr() jen.Code {
	return UseSnowflakeid("GeneratePtr")
}

func UseSnowflakeid(name string) jen.Code {
	return use(snowflakeidPkg, name)
}

func UseConstant(name string) jen.Code {
	return use(constantPkg, name)
}

func HasColumn(col string, columns []Column) bool {
	for _, column := range columns {
		if column.ColumnName == col {
			return true
		}
	}
	return false
}
