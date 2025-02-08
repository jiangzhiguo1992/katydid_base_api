package base

import (
	"gorm.io/gorm"
	"katydid_base_api/tools"
)

type (
	DBModel struct {
		//gorm.Model
		//IDBModel

		Id       uint64 `json:"id" gorm:"primarykey"`
		CreateAt int64  `json:"createAt" gorm:"autoCreateTime:milli"`
		UpdateAt int64  `json:"updateAt" gorm:"autoUpdateTime:milli"`

		// TODO:GG 所有的查询都带上index `gorm:"index"`
		DeleteBy int64  `json:"deleteBy"` // <-1=管理员 -1=系统 0=未删除 1=自己 >1=用户
		DeleteAt *int64 `json:"deleteAt"`

		FieldsCheck func() []*tools.CodeError `json:"-" gorm:"-:all"`
	}

	//IDBModel interface {
	//	CheckFields() []*tools.CodeError
	//}
)

//func NewDBModel(
//	id uint64, createAt int64, updateAt int64,
//	fieldsCheck func() []*tools.CodeError,
//) *DBModel {
//	return &DBModel{
//		Id: id, CreateAt: createAt, UpdateAt: updateAt, DeleteAt: nil,
//		FieldsCheck: fieldsCheck,
//	}
//}

func NewDBModelEmpty() *DBModel {
	return &DBModel{
		//Id: nil,
		//CreateAt: time.Now().UnixMilli(),
		//UpdateAt: time.Now().UnixMilli(),
		DeleteBy: 0,
		DeleteAt: nil,
	}
}

//func (b *DBModel) CheckFields() []*tools.CodeError {
//	panic("implement me")
//}

func (b *DBModel) BeforeSave(tx *gorm.DB) (err error) {
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
