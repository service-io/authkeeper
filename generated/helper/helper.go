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
)

type AutoGenService interface {
	RenderAuto()
	RenderSelf()
}

type Option struct {
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
	return jen.Op("&").Id(sn).Dot(strcase.ToCamel(field))
}

func RenderStarField(sn, field string) jen.Code {
	return jen.Op("*").Id(sn).Dot(strcase.ToCamel(field))
}

func RenderField(sn, field string) jen.Code {
	return jen.Id(sn).Dot(strcase.ToCamel(field))
}
