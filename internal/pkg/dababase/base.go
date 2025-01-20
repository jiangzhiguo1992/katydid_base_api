package dababase

type (
	BaseModel struct {
		Id       uint64 `json:"id"`
		CreateAt int64  `json:"create_at" gorm:"autoCreateTime:milli"`
		UpdateAt int64  `json:"update_at" gorm:"autoUpdateTime:milli"`
		DeleteAt *int64
	}
)

func NewBaseModel(id uint64, createAt int64, updateAt int64) *BaseModel {
	return &BaseModel{Id: id, CreateAt: createAt, UpdateAt: updateAt, DeleteAt: nil}
}
