package db

import (
	"gorm.io/gorm"
	"katydid_base_api/internal/client/model"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
	"strings"
)

func Table() (tx *gorm.DB) {
	return DB().Table(pgsqlTableName(tClient))
}

func InsertClient(client *model.Client) *tools.CodeError {
	if client == nil {
		return utils.MatchErrorByCode(utils.ErrorCodeDBInsNil)
	}
	err := Table().Create(client).Error
	return utils.MatchErrorByErr(err)
}

func SelectClient(id uint64) (*model.Client, *tools.CodeError) {
	var client model.Client
	err := Table().First(&client, id).Error
	if (err != nil) && strings.Contains(err.Error(), "record not found") {
		return nil, nil
	}
	return &client, utils.MatchErrorByErr(err)
}

func UpdateClient(client *model.Client) error {
	return Table().Save(client).Error
}

func DeleteClient(id int64) error {
	return Table().Delete(&model.Client{}, id).Error
}
