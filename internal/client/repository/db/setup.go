package db

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	_ "katydid_base_api/init"
	"katydid_base_api/internal/pkg/database"
	"katydid_base_api/tools"
	"sync"
)

const (
	scheme = "clients"

	tClient         = "client"
	tClientPlatform = "client_platform"
	tClientVersion  = "client_version"
)

var (
	lock sync.Mutex
	ctx  = context.Background()
	conn *gorm.DB
)

func DB() *gorm.DB {
	if conn != nil {
		return conn
	}
	lock.Lock()
	defer lock.Unlock()
	if conn != nil {
		return conn
	}
	RefreshDB(ctx)
	return conn
}

func RefreshDB(ctx context.Context) {
	if conn != nil {
		err := database.DisConnPgSql(conn)
		if err != nil {
			tools.Panic("RefreshDB 断开连接 failed", zap.Error(err))
		}
	}
	config := configs.GetClient()
	conn = database.ConnPgSql(ctx, config.PgSql)
}

func tableName(table string) string {
	return fmt.Sprintf("%s.%s", scheme, table)
}
