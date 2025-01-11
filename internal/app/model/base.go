package model

type (
	// IBase 基类接口
	IBase interface {
		Id() int64       // id
		CreateAt() int64 // 首次创建时间
		UpdateAt() int64 // 最后更新时间
	}
	// Base 基类
	Base struct {
		Id       int64 `json:"id"`        // id
		CreateAt int64 `json:"createAt"`  // 首次创建时间
		UpdateAt int64 `json:"update_at"` // 最后更新时间
	}
)

func NewBase(
	id int64,
	createAt int64,
	updateAt int64,
) *Base {
	return &Base{
		Id:       id,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}
