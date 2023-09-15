// Package runerror
// @author tabuyos
// @since 2023/8/5
// @description runerror
package runerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap(t *testing.T) {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	fmt.Println(err2)
	err3 := errors.Join(err1, errors.New("a"))
	fmt.Println(errors.Unwrap(err2))
	fmt.Println("---------")
	fmt.Println(err3.Error())
	fmt.Println(err3)
	fmt.Println("---------")
	fmt.Println(errors.Unwrap(err3))

	u, ok := err3.(interface {
		Unwrap() []error
	})
	if ok {
		fmt.Println("ok")
		fmt.Println(u.Unwrap())
	}
	fmt.Println(u.Unwrap())
}
