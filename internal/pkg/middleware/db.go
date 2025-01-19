package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	"katydid_base_api/tools"
)

func ConnPgSql() *gorm.DB {
	config := configs.GetClient()
	host := fmt.Sprintf("host=%s", config.PgSql.Host)
	port := fmt.Sprintf("port=%s", config.PgSql.Port)
	database := fmt.Sprintf("database=%s", config.PgSql.Database)
	user := fmt.Sprintf("user=%s", config.PgSql.User)
	pwd := ""
	if len(config.PgSql.Pwd) > 0 {
		pwd = fmt.Sprintf("password=%s", config.PgSql.Pwd)
	}
	timeOut := ""
	if config.PgSql.Timeout > 0 {
		timeOut = fmt.Sprintf("connect_timeout=%d", config.PgSql.Timeout)
	}
	sslMode := ""
	if len(config.PgSql.SSLMode) > 0 {
		sslMode = fmt.Sprintf("sslmode=%s", config.PgSql.SSLMode)
	}
	timeZone := ""
	if len(config.PgSql.TimeZone) > 0 {
		timeZone = fmt.Sprintf("TimeZone=%s", config.PgSql.TimeZone)
	}

	dsn := fmt.Sprintf("%s %s %s %s %s %s %s %s", host, port, database, user, pwd, timeOut, sslMode, timeZone)
	tools.Debug("PgSql 尝试连接", zap.String("dsn", dsn))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn, PreferSimpleProtocol: false, // enable implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		tools.Panic("PgSql 连接失败", zap.String("dsn", dsn), zap.Error(err))
	}

	tools.Info("PgSql 连接成功", zap.String("dsn", dsn))
	return db
}
