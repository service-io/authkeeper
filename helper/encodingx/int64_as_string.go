// Package encodingx
// @author tabuyos
// @since 2023/8/8
// @description encodingx
package encodingx

import (
	enjson "encoding/json"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"unsafe"
)

type int64AsStringCodec struct{}

func (rec *int64AsStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NumberValue:
		var number enjson.Number
		iter.ReadVal(&number)
		i, err := strconv.ParseInt(string(number), 10, 64)
		if err == nil {
			*(*interface{})(ptr) = i
			return
		}
		f, err := strconv.ParseFloat(string(number), 64)
		if err == nil {
			*(*interface{})(ptr) = f
			return
		}
		// Not much we can do here.
	default:
		*(*interface{})(ptr) = iter.Read()
	}
}

func (rec *int64AsStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	m := (*int64)(ptr)
	if m == nil {
		return true
	}
	v := *m
	return v == 0
}

func (rec *int64AsStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := *((*int64)(ptr))
	stream.WriteString(strconv.FormatInt(v, 10))
}
