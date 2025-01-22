package utils

import (
	"errors"
	"go.uber.org/zap"
	"katydid_base_api/tools"
	"strings"
)

const (
	ErrorCodeDBInsNil     = 1002
	ErrorCodeDBSelNil     = 1003
	ErrorCodeDBUpdNil     = 1004
	ErrorCodeDBDelNil     = 1005
	ErrorCodeDBFieldNil   = 1006
	ErrorCodeDBFieldLarge = 1007
	ErrorCodeDBFieldShort = 1008
	ErrorCodeDBDupPk      = 1001
)

// TODO:GG 国际化
var errorMessages = map[int]string{
	ErrorCodeDBInsNil:     "插入对象为空",
	ErrorCodeDBSelNil:     "查询对象为空",
	ErrorCodeDBUpdNil:     "更新对象为空",
	ErrorCodeDBDelNil:     "删除对象为空",
	ErrorCodeDBDupPk:      "数据库唯一约束冲突",
	ErrorCodeDBFieldNil:   "数据库字段为空",
	ErrorCodeDBFieldLarge: "数据库字段过长",
	ErrorCodeDBFieldShort: "数据库字段过短",
}

var errorCodes = map[string]int{
	"duplicate key value violates unique constraint": ErrorCodeDBDupPk,
}

func MatchErrorByCode(code int) *tools.CodeError {
	if message, ok := errorMessages[code]; ok {
		return &tools.CodeError{
			Code: code,
			Err:  errors.New(message),
		}
	}
	tools.Warn("MatchErrorByCode 没有匹配的错误Code:", zap.Int("code", code))
	return nil
}

func MatchErrorByErr(err error) *tools.CodeError {
	if err == nil {
		return nil
	}
	for msg, code := range errorCodes {
		if strings.Contains(err.Error(), msg) {
			if errorCode := MatchErrorByCode(code); errorCode != nil {
				return errorCode
			}
			tools.Warn("MatchErrorByErr 没有匹配的错误Msg:", zap.Error(err))
			return &tools.CodeError{
				Code: code,
				Err:  err,
			}
		}
	}
	return tools.NewCodeError(err)
}
