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
