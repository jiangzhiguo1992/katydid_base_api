package db

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	_ "katydid_base_api/init"
	"katydid_base_api/internal/pkg/dababase"
	"katydid_base_api/tools"
)

const (
	scheme = "clients"

	tClient = "client"
)

var (
	conn *gorm.DB
	ctx  = context.Background()
)

func DB() *gorm.DB {
	if conn != nil {
		return conn
	}
	RefreshDB(ctx)
	return conn
}

func RefreshDB(ctx context.Context) {
	if conn != nil {
		err := dababase.DisConnPgSql(conn)
		if err != nil {
			tools.Panic("RefreshDB 断开连接 failed", zap.Error(err))
		}
	}
	config := configs.GetClient()
	conn = dababase.ConnPgSql(ctx, config.PgSql)
}

func tableName(table string) string {
	return fmt.Sprintf("%s.%s", scheme, table)
}
