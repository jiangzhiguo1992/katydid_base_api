package db

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"katydid_base_api/configs"
	"katydid_base_api/tools"
	"strings"
	"time"
)

func ConnPgSql(ctx context.Context, config *configs.PgSqlConfig) *gorm.DB {
	if config == nil {
		tools.Panic("PgSql 配置为空", zap.Error(errors.New("PgSql 配置为空")))
	}
	host := fmt.Sprintf("host=%s", config.Host)
	port := fmt.Sprintf("port=%s", config.Port)
	database := fmt.Sprintf("database=%s", config.Database)
	user := fmt.Sprintf("user=%s", config.User)
	pwd := ""
	if len(config.Pwd) > 0 {
		pwd = fmt.Sprintf("password=%s", config.Pwd)
	}
	timeOut := ""
	if config.Timeout > 0 {
		timeOut = fmt.Sprintf("connect_timeout=%d", config.Timeout)
	}
	sslMode := ""
	if len(config.SSLMode) > 0 {
		sslMode = fmt.Sprintf("sslmode=%s", config.SSLMode)
	}
	timeZone := ""
	if len(config.TimeZone) > 0 {
		timeZone = fmt.Sprintf("TimeZone=%s", config.TimeZone)
	}

	dsn := fmt.Sprintf("%s %s %s %s %s %s %s %s", host, port, database, user, pwd, timeOut, sslMode, timeZone)
	tools.Debug("PgSql 尝试连接", zap.String("dsn", dsn))

	var db *gorm.DB
	var err error

	maxRetries := config.MaxRetries
	retryInterval := time.Duration(config.RetryDelay) * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dsn, PreferSimpleProtocol: false, // enable implicit prepared statement usage
			// TODO:GG 连接池(conf里读?)，看看gorm的文档
		}), &gorm.Config{
			// TODO:GG 配置 连接池，超时时间等，看看gorm的文档
		})
		if (db != nil) && (err == nil) {
			break
		}

		log := strings.Replace(dsn, pwd, "", 1)
		tools.Warn("PgSql 连接失败，重试中...", zap.Int("times", i), zap.String("dsn", log), zap.Error(err))
		//time.Sleep(retryInterval)
		select {
		case <-time.After(retryInterval):
		case <-ctx.Done():
			break
		}
	}

	if err != nil {
		log := strings.Replace(dsn, pwd, "", 1)
		tools.Panic("PgSql 连接失败", zap.String("dsn", log), zap.Error(err))
	}

	log := strings.Replace(dsn, pwd, "", 1)
	tools.Info("PgSql 连接成功", zap.String("dsn", log))
	return db
}

func DisConnPgSql(db *gorm.DB) error {
	if db == nil {
		return errors.New("PgSql 断开连接失败，连接为空")
	}
	sqlDB, err := db.DB()
	if err != nil {
		tools.Error("PgSql 断开连接失败, 获取连接失败", zap.Error(err))
		return tools.NewMultiCodeError(err).WrapError(errors.New("PgSql 断开连接失败, 获取连接失败"))
	}
	err = sqlDB.Close()
	if err != nil {
		tools.Error("PgSql 断开连接失败, 关闭连接失败", zap.Error(err))
		return tools.NewMultiCodeError(err).WrapError(errors.New("PgSql 断开连接失败, 关闭连接失败"))
	}
	tools.Info("PgSql 断开连接成功")
	return nil
}
