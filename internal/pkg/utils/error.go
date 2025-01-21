package utils

import (
	"errors"
	"go.uber.org/zap"
	"katydid_base_api/tools"
	"strings"
)

const (
	ErrorCodeDuplicateKey = 1001
	ErrorCodeInsertNil    = 1002
	ErrorCodeSelectNil    = 1003
	ErrorCodeUpdateNil    = 1004
	ErrorCodeDeleteNil    = 1005
)

var errorCodes = map[string]int{
	"duplicate key value violates unique constraint": ErrorCodeDuplicateKey,
}

// TODO:GG 国际化
var errorMessages = map[int]string{
	ErrorCodeInsertNil:    "插入对象为空",
	ErrorCodeSelectNil:    "查询对象为空",
	ErrorCodeUpdateNil:    "更新对象为空",
	ErrorCodeDeleteNil:    "删除对象为空",
	ErrorCodeDuplicateKey: "数据库唯一约束冲突",
}

func MatchErrorCode(code int) *tools.CodeError {
	if message, ok := errorMessages[code]; ok {
		return &tools.CodeError{
			Code: code,
			Err:  errors.New(message),
		}
	}
	tools.Warn("MatchErrorCode 没有匹配的错误Code:", zap.Int("code", code))
	return nil
}

func MatchCodeError(err error) *tools.CodeError {
	if err == nil {
		return nil
	}
	for msg, code := range errorCodes {
		if strings.Contains(err.Error(), msg) {
			if errorCode := MatchErrorCode(code); errorCode != nil {
				return errorCode
			}
			tools.Warn("MatchCodeError 没有匹配的错误Msg:", zap.Error(err))
			return &tools.CodeError{
				Code: code,
				Err:  err,
			}
		}
	}
	return tools.NewCodeError(err)
}
