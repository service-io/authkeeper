// Package page
// @author tabuyos
// @since 2023/7/18
// @description model
package page

import (
	"bytes"
	"strconv"
)

type Query struct {
	// 当前页
	Page int64 `json:"page"`
	// 页大小
	Size int64 `json:"size"`
}

type Result struct {
	// 分页数据
	Data any `json:"data"`
	// 总量
	Total int64 `json:"total"`
}

func NewQuery() *Query {
	return &Query{}
}

func NewResult(data any, total int64) *Result {
	return &Result{
		Data:  data,
		Total: total,
	}
}

func (receiver Query) String() string {
	var b bytes.Buffer
	b.WriteString("(")
	b.WriteString("page -> ")
	b.WriteString(strconv.FormatUint(uint64(receiver.Page), 10))
	b.WriteString(", ")
	b.WriteString("size -> ")
	b.WriteString(strconv.FormatUint(uint64(receiver.Size), 10))
	b.WriteString(")")
	return b.String()
}
