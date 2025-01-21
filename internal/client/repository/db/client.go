package db

import (
	"gorm.io/gorm"
	"katydid_base_api/internal/client/model"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
)

func Table() (tx *gorm.DB) {
	return DB().Table(tableName(tClient))
}

func InsertClient(client *model.Client) *tools.CodeError {
	if client == nil {
		return utils.MatchErrorCode(utils.ErrorCodeInsertNil)
	}
	err := Table().Create(client).Error
	return utils.MatchCodeError(err)
}

func SelectClient(id int64) (*model.Client, error) {
	var client model.Client
	err := Table().First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func UpdateClient(client *model.Client) error {
	return Table().Save(client).Error
}

func DeleteClient(id int64) error {
	return Table().Delete(&model.Client{}, id).Error
}
