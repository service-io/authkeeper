// Package runerror
// @author tabuyos
// @since 2023/8/3
// @description runerror
package runerror

import (
	"errors"
	"fmt"
)

type IRunError interface {
	WithCode(code int) IRunError
	WithMessage(msg string) IRunError
	WithError(err error) IRunError
	TopMessage() string
	Error() string
	Code() int
	String() string
}

type runError struct {
	error
	code int
}

func New() IRunError {
	return &runError{}
}

func NewError(err string) error {
	return errors.New(err)
}

func NewAll(code int, error error) IRunError {
	return &runError{
		code:  code,
		error: error,
	}
}

func NewPlain(code int, msg string) IRunError {
	return &runError{
		code:  code,
		error: NewError(msg),
	}
}

func NewMessage(msg string) IRunError {
	return NewAll(DftError, errors.New(msg))
}

func NewCode(code int) IRunError {
	return NewAll(code, errors.New("操作失败"))
}

func (rec *runError) WithCode(code int) IRunError {
	rec.code = code
	return rec
}

func (rec *runError) WithMessage(msg string) IRunError {
	return rec.WithError(errors.New(msg))
}

func (rec *runError) WithError(err error) IRunError {
	rec.error = errors.Join(rec.error, err)
	return rec
}

func (rec *runError) TopMessage() string {
	if rec.error == nil {
		return ""
	}
	switch x := rec.error.(type) {
	case interface{ Unwrap() error }:
		err := x.Unwrap()
		if err == nil {
			return ""
		}
		return err.Error()
	case interface{ Unwrap() []error }:
		errs := x.Unwrap()
		if errs == nil {
			return ""
		}
		if len(errs) == 0 {
			return ""
		}
		return errs[0].Error()
	default:
		return rec.error.Error()
	}
}

func (rec *runError) Code() int {
	return rec.code
}

func (rec *runError) String() string {
	return fmt.Sprintf("(code: %d, message: %s)", rec.code, rec.Error())
}
