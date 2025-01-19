package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"katydid_base_api/tools"
)

const (
	ModuleTag = "module"
	ClientTag = "client"

	EnableKey = "enable"
	CloudKey  = "cloud"
	ProdKey   = "prod"
)

var (
	client *ClientConfig
)

type ModuleConfig struct {
	enable bool `mapstructure:"enable"`
	cloud  bool `mapstructure:"cloud"`
	prod   bool `mapstructure:"prod"`
}

type PgSqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Pwd      string `mapstructure:"pwd"`
	Timeout  int    `mapstructure:"timeout"`
	TimeZone string `mapstructure:"time_zone"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type RedisConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Pwd  string `mapstructure:"pwd"`
}

type ClientConfig struct {
	*ModuleConfig
	// TODO:GG 额外加 (属于module的扩展之类的)
	PgSql *PgSqlConfig `mapstructure:"pgsql"`
	// TODO:GG 额外加 (属于非module的扩展之类的)
}

func GetClient() *ClientConfig {
	if client != nil {
		return client
	}
	return RefreshClient()
}

func RefreshClient() *ClientConfig {
	module := newModule(ClientTag)
	pgSql := newPgSql(module.prod, ClientTag)

	client = &ClientConfig{
		ModuleConfig: module,
		PgSql:        pgSql,
	}
	return client
}

func newModule(subTag string) *ModuleConfig {
	module := &ModuleConfig{}
	tag := fmt.Sprintf("%s.%s", ModuleTag, subTag)
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, EnableKey)) {
		module.enable = viper.GetBool(fmt.Sprintf("%s.%s", tag, EnableKey))
	} else {
		module.enable = viper.GetBool(tools.ModuleEnableKey)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, CloudKey)) {
		module.cloud = viper.GetBool(fmt.Sprintf("%s.%s", tag, CloudKey))
	} else {
		module.cloud = viper.GetBool(tools.ModuleCloudKey)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, ProdKey)) {
		module.prod = viper.GetBool(fmt.Sprintf("%s.%s", tag, ProdKey))
	} else {
		module.prod = viper.GetBool(tools.ModuleProdKey)
	}
	return module
}

func newPgSql(prod bool, module string) *PgSqlConfig {
	if len(module) <= 0 {
		panic(fmt.Sprintf("configs pgsql module is empty: %v, %s", prod, module))
	}
	pgsql := &PgSqlConfig{}
	tag := fmt.Sprintf("%s.pgsql", module)
	host := tools.ConfigEnvKey(prod, tag, "host")
	pgsql.Host = viper.GetString(host)
	port := tools.ConfigEnvKey(prod, tag, "port")
	pgsql.Port = viper.GetString(port)
	database := tools.ConfigEnvKey(prod, tag, "database")
	pgsql.Database = viper.GetString(database)
	user := tools.ConfigEnvKey(prod, tag, "user")
	pgsql.User = viper.GetString(user)
	pwd := tools.ConfigEnvKey(prod, tag, "pwd")
	pgsql.Pwd = viper.GetString(pwd)
	timeout := tools.ConfigEnvKey(prod, tag, "timeout")
	pgsql.Timeout = viper.GetInt(timeout)
	timeZone := tools.ConfigEnvKey(prod, tag, "timezone")
	pgsql.TimeZone = viper.GetString(timeZone)
	sslMode := tools.ConfigEnvKey(prod, tag, "sslmode")
	pgsql.SSLMode = viper.GetString(sslMode)
	return pgsql
}
