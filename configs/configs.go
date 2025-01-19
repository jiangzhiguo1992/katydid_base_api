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
	Name   string `mapstructure:"name"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pwd    string `mapstructure:"pwd"`
	DBName string `mapstructure:"db_name"`
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
	name := tools.ConfigEnvKey(prod, tag, "name")
	pgsql.Name = viper.GetString(name)
	host := tools.ConfigEnvKey(prod, tag, "host")
	pgsql.Host = viper.GetString(host)
	port := tools.ConfigEnvKey(prod, tag, "port")
	pgsql.Port = viper.GetString(port)
	user := tools.ConfigEnvKey(prod, tag, "user")
	pgsql.User = viper.GetString(user)
	pwd := tools.ConfigEnvKey(prod, tag, "pwd")
	pgsql.Pwd = viper.GetString(pwd)
	dbName := tools.ConfigEnvKey(prod, tag, "db_name")
	pgsql.DBName = viper.GetString(dbName)
	return pgsql
}
