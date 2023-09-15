// Package helper
// @author tabuyos
// @since 2023/7/20
// @description helper
package helper

import (
	"crypto/rand"
	"database/sql"
	"deepsea/helper/recorderx"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"strings"
	"unsafe"
)

type eface struct {
	v   int64
	ptr unsafe.Pointer
}

func DeferClose(closer io.Closer, errHandler ...func(...error)) {
	err := closer.Close()
	if err != nil {
		if len(errHandler) > 0 {
			for _, eh := range errHandler {
				eh(err)
			}
		}
		panic(err)
	}
}

func HandleRollback(err error, tx *sql.Tx, eh func(...error)) {
	if err != nil {
		err := tx.Rollback()
		eh(err)
	}
}

func ErrToLog(logger recorderx.Recorder) func(...error) {
	return func(errs ...error) {
		for _, err := range errs {
			logger.WithOptions(recorderx.AddCallerSkip(1)).Error(err.Error())
		}
	}
}

func ErrToLogAndPanic(logger recorderx.Recorder) func(...error) {
	return func(errs ...error) {
		for _, err := range errs {
			logger.WithOptions(recorderx.AddCallerSkip(1)).Error(err.Error())
			panic(err)
		}
	}
}

func Shuffle[S ~[]E, E any](slice S) {
	for len(slice) > 0 {
		n := len(slice)
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
		if err != nil {
			panic(err)
		}
		ri := randIndex.Int64()
		slice[n-1], slice[ri] = slice[ri], slice[n-1]
		slice = slice[:n-1]
	}
}

func HandleErr(err error, fn func()) {
	if err != nil {
		fn()
	}
}

func LogErr(logger recorderx.Recorder, err error) {
	if err != nil {
		logger.WithOptions(recorderx.AddCallerSkip(1)).Error(err.Error())
	}
}

func PanicErr(logger recorderx.Recorder, err error) {
	if err != nil {
		logger.WithOptions(recorderx.AddCallerSkip(1)).Error(err.Error())
		panic(err)
	}
}

func HandleTx(tx *sql.Tx, eh func(err ...error)) {
	err := recover()
	if err != nil {
		switch v := err.(type) {
		case error:
			eh(v)
		case string:
			eh(errors.New(v))
		}
		rbErr := tx.Rollback()
		eh(rbErr)
	} else {
		cmErr := tx.Commit()
		eh(cmErr)
	}
}

func Rows[T any](rows *sql.Rows, supplier func() (*T, []any)) []*T {
	// rs := make([]T, 0)
	var rs []*T
	for rows.Next() {
		r, cs := supplier()
		if err := rows.Scan(cs...); err != nil {
			panic(err)
		}
		rs = append(rs, r)
	}
	return rs
}

func Row[T any](row *sql.Row, supplier func() (*T, []any)) *T {
	r, cs := supplier()
	if err := row.Scan(cs...); err != nil {
		panic(err)
	}
	return r
}

func Reduce[T ~[]E, R, E any](ls T, rs []R, fn func(E) R) {
	for i, e := range ls {
		rs[i] = fn(e)
	}
}

func GenPlaceholder(ids []int64) string {
	if len(ids) == 0 {
		panic("无 ID 信息")
	}
	ph := make([]string, len(ids))
	for i := range ids {
		ph[i] = "?"
	}
	return strings.Join(ph, ", ")
}

func ToAnyItems[T any](ps []T) []any {
	nps := make([]any, len(ps))
	for i, p := range ps {
		nps[i] = p
	}
	return nps
}

func ToAny[T any](p T) any {
	var np any = p
	return np
}

func ToString(e any) string {
	switch t := e.(type) {
	case string:
		return t
	default:
		if IsNil(e) {
			return ""
		}
		marshal, err := json.Marshal(t)
		if err != nil {
			return ""
		}
		return string(marshal)
	}
}

func ToPtr[T any](p T) *T {
	return &p
}

// SplitFunc 使用函数进行分割, 注意: 并不会移除符合谓词的字符,
// 具体实现参考 strings.FieldsFunc 进行修改的,
// strings.FieldsFunc 会移除符合谓词的字符
func SplitFunc(s string, f func(rune) bool) []string {
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	start := -1
	for end, char := range s {
		if f(char) {
			if start == -1 {
				start = 0
			} else if start >= 0 {
				spans = append(spans, span{start, end})
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	} else {
		spans = append(spans, span{0, len(s)})
	}

	a := make([]string, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}

func RecorderSign(aid, ip string) (rs string) {
	s := make([]string, 0)
	if len(aid) != 0 {
		s = append(s, "id:"+aid)
	}
	if len(ip) != 0 {
		s = append(s, "ip:"+ip)
	}

	return strings.Join(s, ",")
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func IsNil(b any) bool {
	if b == nil {
		return true
	}
	efptr := (*eface)(unsafe.Pointer(&b))
	if efptr == nil {
		return true
	}
	return uintptr(efptr.ptr) == 0x0
}

func IsNonNil(b any) bool {
	return !IsNil(b)
}

// func CalcFieldValue[T any](a *T) T {
//   if IsNil(a) {
//     return nil
//   }
//   return *a
// }
