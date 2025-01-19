package tools

import (
	"fmt"
)

// MultiError 是一个可以包裹多层错误的结构体
type MultiError struct {
	Code   int
	Errors []*CodeError
}

func NewMultiError(errs ...*CodeError) *MultiError {
	return &MultiError{Code: 0, Errors: errs}
}

func (m *MultiError) WithCode(code int) *MultiError {
	m.Code = code
	return m
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	result := "Multiple:"
	for _, err := range m.Errors {
		result += fmt.Sprintf("\n\t-%s", err.Error())
	}
	return result
}

// WrapError 包裹一个新的错误
func (m *MultiError) WrapError(err error) {
	m.WrapCodeError(NewCodeError(0, err))
}

func (m *MultiError) WrapCodeError(err *CodeError) {
	m.Errors = append(m.Errors, err)
}

// Unwrap 返回第一个错误
func (m *MultiError) Unwrap() error {
	if len(m.Errors) > 0 {
		return m.Errors[0]
	}
	return nil
}

// CodeError 是一个可以包裹错误码的结构体
type CodeError struct {
	Code int
	Err  error
}

func NewCodeError(code int, err error) *CodeError {
	return &CodeError{Err: err, Code: code}
}

func (c *CodeError) Error() string {
	return c.Err.Error()
}
