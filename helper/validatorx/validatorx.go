// Package validatorx
// @author tabuyos
// @since 2023/8/24
// @description validatorx
package validatorx

import (
	"deepsea/helper/runerror"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Required struct {
	Key string
	Val any
}

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator.New()}
}

func (v *Validator) WithRequiredAndPanic(val any, key ...string) {
	err := v.WithRequired(val, key...)
	if err != nil {
		code := runerror.GetUsrErp(runerror.ModNil, runerror.ValidFailedError)
		panic(runerror.NewAll(code, err))
	}
}

func (v *Validator) WithRequired(val any, key ...string) error {
	err := v.Var(val, "required")
	if err != nil || val == "" {
		// _, ok := err.(validator.ValidationErrors)
		// if ok {
		// 	return errors.New(strings.Join(key, ",") + " 未传递!")
		// }
		return errors.New(strings.Join(key, ",") + "未传递!")
	}
	return err
}

func ValidateRequiredField(args ...any) (errs []error) {
	validatorx := NewValidator()
	for _, arg := range args {
		errs = append(errs, validatorx.Var(arg, "required"))
	}
	return
}
