package init

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	_ "katydid_base_api/init"
	"katydid_base_api/internal/pkg/middleware"
	"katydid_base_api/tools"
	"sync"
)

var (
	lock  sync.Mutex
	ctx   = context.Background()
	conns = make(map[string]*gorm.DB)
)

//func init() {
//	RefreshDB(ctx, "client")
//}

func ClientDB() *gorm.DB {
	return conn("client")
}

// conn 获取+创建数据库连接
func conn(key string) *gorm.DB {
	if conns[key] != nil {
		return conns[key]
	}
	lock.Lock()
	defer lock.Unlock()
	if conns[key] != nil {
		return conns[key]
	}
	restartConn(ctx, key)
	return conns[key]
}

// restartConn 重启数据库连接 TODO:GG 什么时候调用，怎么保持断了重连?
func restartConn(ctx context.Context, key string) {
	if conns[key] != nil {
		err := middleware.DisConnPgSql(conns[key])
		if err != nil {
			tools.Panic("restartDB 断开连接 failed", zap.Error(err))
		}
	}
	config := configs.GetClient()
	conns[key] = middleware.ConnPgSql(ctx, config.PgSql)
}
