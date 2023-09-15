// Package helper
// @author tabuyos
// @since 2023/8/29
// @description helper
package helper

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"path/filepath"
)

type AutoGenService interface {
	RenderAuto()
	RenderSelf()
}

type Option struct {
	Table      string
	Entity     string
	Variable   string
	Variables  string
	EnableShip bool
	Left       string
	Right      string
	Tags       []string
	Cols       []Column
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

	f.ImportName("metis/config/constant", "constant")
	f.ImportName("metis/helper/database", "database")
	f.ImportName("metis/helper/security", "security")
	f.ImportName("metis/helper/sensitivex", "sensitivex")
	f.ImportName("metis/helper/snowflakeid", "snowflakeid")
	f.ImportName("metis/helper/validator", "validator")
	f.ImportName("metis/helper/recorderx", "recorderx")

	f.ImportName("metis/model/dto", "dto")
	f.ImportName("metis/test/autogen/3.0/entity", "entity")
	f.ImportName("metis/model/page", "page")
	f.ImportName("metis/model/reply", "reply")

	f.ImportName("metis/model/3.0/iris", "iris")

	f.ImportName(fmt.Sprintf("metis/module/%s/repository", module), "repository")
	f.ImportName(fmt.Sprintf("metis/module/%s/service", module), "service")

	f.ImportName("metis/helper", "helper")
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

func UseTime(name string) jen.Code {
	return use("time", name)
}

func UseIris(name string) jen.Code {
	return use("metis/model/3.0/iris", name)
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

func UseIrisConfigService(name string) jen.Code {
	return jen.Add(UseIris("ConfigService")).Types(jen.Id(name))
}

func UseIrisConfigServiceCode(types ...jen.Code) jen.Code {
	return jen.Add(UseIris("ConfigService")).Types(types...)
}

func UseIrisNamedConfigServiceCode(types ...jen.Code) jen.Code {
	return jen.Id("config").Add(UseIris("ConfigService")).Types(types...)
}

func UseEntity(name string) jen.Code {
	return use("metis/test/autogen/3.0/entity", name)
}

func UseSQL(name string) jen.Code {
	return use("database/sql", name)
}

func UseSQLTx() jen.Code {
	return UseSQL("Tx")
}

func UseHelper(name string) jen.Code {
	return use("metis/helper", name)
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
	return use("metis/helper/recorderx", name)
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
	return use("metis/helper/database", name)
}

func UseFetchDB() jen.Code {
	return jen.Add(UseDatabase("FetchDB")).Call()
}
