package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	"katydid_base_api/tools"
	"time"
)

func ConnPgSql(ctx context.Context) *gorm.DB {
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

	var db *gorm.DB
	var err error

	maxRetries := config.PgSql.MaxRetries
	retryInterval := time.Duration(config.PgSql.RetryDelay) * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dsn, PreferSimpleProtocol: false, // enable implicit prepared statement usage
		}), &gorm.Config{})
		if (db != nil) && (err == nil) {
			break
		}

		tools.Warn("PgSql 连接失败，重试中...", zap.Int("times", i), zap.String("dsn", dsn), zap.Error(err))
		//time.Sleep(retryInterval)
		select {
		case <-time.After(retryInterval):
		case <-ctx.Done():
			break
		}
	}

	if err != nil {
		tools.Panic("PgSql 连接失败", zap.String("dsn", dsn), zap.Error(err))
	}

	tools.Info("PgSql 连接成功", zap.String("dsn", dsn))
	return db
}
