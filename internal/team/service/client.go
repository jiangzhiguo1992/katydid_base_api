package service

import (
	"katydid_base_api/internal/team/model"
	"katydid_base_api/internal/team/repository/db"
	"katydid_base_api/tools"
)

func GetClient(id uint64) (*model.Client, *tools.CodeError) {
	// TODO:GG permission
	// TODO:GG cache
	client, codeError := db.QueClient(id)
	return client, codeError
}
