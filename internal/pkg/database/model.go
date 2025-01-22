package database

import (
	"gorm.io/gorm"
	"katydid_base_api/tools"
)

type (
	BaseModel struct {
		Id       uint64 `json:"id"`
		CreateAt int64  `json:"createAt" gorm:"autoCreateTime:milli"`
		UpdateAt int64  `json:"updateAt" gorm:"autoUpdateTime:milli"`
		DeleteAt *int64 // invisible TODO:GG 所有的查询都带上index

		FieldsCheck func() []*tools.CodeError `json:"-" gorm:"-:all"`
	}
)

func NewBaseModel(
	id uint64, createAt int64, updateAt int64,
	fieldsCheck func() []*tools.CodeError,
) *BaseModel {
	return &BaseModel{
		Id: id, CreateAt: createAt, UpdateAt: updateAt, DeleteAt: nil,
		FieldsCheck: fieldsCheck,
	}
}

func NewBaseModelEmpty() *BaseModel {
	return &BaseModel{}
}

func (b *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	if b.FieldsCheck == nil {
		return nil
	}
	errors := b.FieldsCheck()
	if (errors == nil) || (len(errors) <= 0) {
		return nil
	}
	multiError := tools.NewMultiCodeError(errors[0])
	for i := 0; i < len(errors); i++ {
		_ = multiError.WrapCodeError(errors[i])
	}
	return multiError
}
