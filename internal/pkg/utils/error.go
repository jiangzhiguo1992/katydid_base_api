package utils

import (
	"errors"
	"katydid_base_api/tools"
	"strings"
)

const (
	ErrorCodeDuplicateKey = 1001
)

var errorCodes = map[string]int{
	"duplicate key value violates unique constraint": ErrorCodeDuplicateKey,
}

// TODO:GG 国际化
var errorMessages = map[int]string{
	ErrorCodeDuplicateKey: "数据库唯一约束冲突",
}

func CodeError(err error) *tools.CodeError {
	for msg, code := range errorCodes {
		if strings.Contains(err.Error(), msg) {
			if message, ok := errorMessages[code]; ok {
				return &tools.CodeError{
					Code: code,
					Err:  errors.New(message),
				}
			}
			return &tools.CodeError{
				Code: code,
				Err:  err,
			}
		}
	}
	return tools.NewCodeError(err)
}
