package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"katydid_base_api/internal/client/model"
	"katydid_base_api/internal/pkg/base"
	"katydid_base_api/internal/pkg/setup"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
	"strings"
	"time"
)

// TODO:GG idx?

// TODO:GG 放哪里?
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
	result := TX().Create(client) // .Omit("DeleteAt")
	return utils.MatchErrorByErr(result.Error)
}

func DropClient(id uint64) (bool, *tools.CodeError) {
	tx := TX().Unscoped()
	tx = tx.Delete(model.NewClientJustId(id))
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func DelClient(id uint64, by int64) (bool, *tools.CodeError) {
	now := time.Now().UnixMilli()
	tx := TX().Model(model.NewClientJustId(id))
	tx = tx.Updates(model.Client{DBModel: &base.DBModel{DeleteBy: by, DeleteAt: &now}})
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func UpdClientEnable(client *model.Client, enable bool) (bool, *tools.CodeError) {
	tx := TX().Model(&client).Clauses(clause.Returning{})
	tx = tx.Update("enable", enable)
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func UpdClientLineAt(client *model.Client, offlineAt, onlineAt int64) (bool, *tools.CodeError) {
	tx := TX().Model(&client).Clauses(clause.Returning{})
	tx = tx.Updates(model.Client{OfflineAt: offlineAt, OnlineAt: onlineAt})
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func UpdClientIPName(client *model.Client, IPName string) (bool, *tools.CodeError) {
	tx := TX().Model(&client).Clauses(clause.Returning{})
	tx = tx.Update("ip_name", IPName)
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func UpdClientOrganization(client *model.Client, organization string) (bool, *tools.CodeError) {
	tx := TX().Model(&client).Clauses(clause.Returning{})
	tx = tx.Update("organization", organization)
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
}

func UpdClientExtra(client *model.Client, extra map[string]interface{}) (bool, *tools.CodeError) {
	tx := TX().Model(&client).Clauses(clause.Returning{})
	tx = tx.Update("extra", extra)
	return tx.RowsAffected > 0, utils.MatchErrorByErr(tx.Error)
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
