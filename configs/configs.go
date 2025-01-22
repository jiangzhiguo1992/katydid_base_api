package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"katydid_base_api/tools"
)

const (
	tagModule = "module"
	tagClient = "client"

	keyEnable = "enable"
	keyCloud  = "cloud"
	keyProd   = "prod"
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
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	User       string `mapstructure:"user"`
	Pwd        string `mapstructure:"pwd"`
	Timeout    int    `mapstructure:"timeout"`
	TimeZone   string `mapstructure:"timezone"`
	SSLMode    string `mapstructure:"sslMode"`
	MaxRetries int    `mapstructure:"maxRetries"`
	RetryDelay int    `mapstructure:"retryDelay"`
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
	module := newModule(tagClient)
	pgSql := newPgSql(module.prod, tagClient)

	client = &ClientConfig{
		ModuleConfig: module,
		PgSql:        pgSql,
	}
	return client
}

func newModule(subTag string) *ModuleConfig {
	module := &ModuleConfig{}
	tag := fmt.Sprintf("%s.%s", tagModule, subTag)
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, keyEnable)) {
		module.enable = viper.GetBool(fmt.Sprintf("%s.%s", tag, keyEnable))
	} else {
		module.enable = viper.GetBool(tools.ConfigKeyAppEnable)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, keyCloud)) {
		module.cloud = viper.GetBool(fmt.Sprintf("%s.%s", tag, keyCloud))
	} else {
		module.cloud = viper.GetBool(tools.ConfigKeyAppCloud)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", tag, keyProd)) {
		module.prod = viper.GetBool(fmt.Sprintf("%s.%s", tag, keyProd))
	} else {
		module.prod = viper.GetBool(tools.ConfigKeyAppProd)
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
	sslMode := tools.ConfigEnvKey(prod, tag, "sslMode")
	pgsql.SSLMode = viper.GetString(sslMode)
	maxRetries := tools.ConfigEnvKey(prod, tag, "maxRetries")
	pgsql.MaxRetries = viper.GetInt(maxRetries)
	if pgsql.MaxRetries <= 0 {
		pgsql.MaxRetries = 1
	}
	retryDelay := tools.ConfigEnvKey(prod, tag, "retryDelay")
	pgsql.RetryDelay = viper.GetInt(retryDelay)
	return pgsql
}
