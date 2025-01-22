package db

import (
	"fmt"
	"gorm.io/gorm"
	"katydid_base_api/internal/client/model"
	"katydid_base_api/internal/pkg/setup"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
	"strings"
)

// TODO:GG idx

// TODO:GG 放哪里
const (
	scheme = "clients"

	tClient         = "client"
	tClientPlatform = "client_platform"
	tClientVersion  = "client_version"
)

func TX() (tx *gorm.DB) {
	return setup.ClientDB().Table(fmt.Sprintf("%s.%s", scheme, tClient))
}

func AddClient(client *model.Client) *tools.CodeError {
	if client == nil {
		return utils.MatchErrorByCode(utils.ErrorCodeDBInsNil)
	}
	err := TX().Create(client).Error
	return utils.MatchErrorByErr(err)
}

func QueClient(id uint64) (*model.Client, *tools.CodeError) {
	var client model.Client
	err := TX().First(&client, id).Error
	if (err != nil) && strings.Contains(err.Error(), "record not found") {
		return nil, nil
	}
	return &client, utils.MatchErrorByErr(err)
}

func QueClientByIpPart(ip, part uint) (*model.Client, *tools.CodeError) {
	if ip <= 0 {
		return nil, utils.MatchErrorByCode(utils.ErrorCodeDBQueryParams).WithSuffix("ip")
	} else if part <= 0 {
		return nil, utils.MatchErrorByCode(utils.ErrorCodeDBQueryParams).WithPrefix("part")
	}
	var client model.Client
	tx := TX()
	tx = tx.Where("ip = ?", ip).Where("part = ?", part)
	err := tx.First(&client).Error
	return &client, utils.MatchErrorByErr(err)
}

func QueClientsByIp(ip uint) ([]model.Client, *tools.CodeError) {
	if ip <= 0 {
		return nil, utils.MatchErrorByCode(utils.ErrorCodeDBQueryParams).WithSuffix("ip")
	}
	var clients []model.Client
	err := TX().Where("ip = ?", ip).Find(&clients).Error
	return clients, utils.MatchErrorByErr(err)
}

func UpdClient(client *model.Client) error {
	return TX().Save(client).Error
}

func DelClient(id int64) error {
	return TX().Delete(&model.Client{}, id).Error
}
