// Package encodingx
// @author tabuyos
// @since 2023/8/7
// @description encodingx
package encodingx

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API

func InitEncodingX() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func RegisterInt64ToString() {
	jsoniter.RegisterTypeEncoder("int64", &int64AsStringCodec{})
	// jsoniter.RegisterTypeDecoder("int64", &int64AsStringCodec{})
}

func FetchJSON() jsoniter.API {
	return json
}

func ToJSON(info any) string {
	rs, err := FetchJSON().MarshalToString(info)
	if err != nil {
		panic(errors.Join(err, errors.New("JSON 转换失败")))
	}
	return rs
}
