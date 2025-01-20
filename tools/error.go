package tools

import (
	"fmt"
)

// MultiError 是一个可以包裹多层错误的结构体
type MultiError struct {
	Code   int
	Errors []*CodeError
}

func NewMultiError(err error) *MultiError {
	return &MultiError{Code: 0, Errors: []*CodeError{NewCodeError(err)}}
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
func (m *MultiError) WrapError(err error) *MultiError {
	return m.WrapCodeError(NewCodeError(err))
}

func (m *MultiError) WrapCodeError(err *CodeError) *MultiError {
	m.Errors = append(m.Errors, err)
	return m
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

func NewCodeError(err error) *CodeError {
	return &CodeError{Code: 0, Err: err}
}

func (c *CodeError) WithCode(code int) *CodeError {
	c.Code = code
	return c
}

func (c *CodeError) Error() string {
	return c.Err.Error()
}
