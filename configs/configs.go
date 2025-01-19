package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"katydid_base_api/tools"
)

const (
	NameKey   = "name"
	EnableKey = "enable"
	CloudKey  = "cloud"
	ProdKey   = "prod"

	ClientKey = "client"
)

var files = map[string][][2]string{
	"./configs/": {
		{"app", "toml"},
		{"module", "toml"},
	},
	"./configs/client/": {
		{"common", "toml"},
		{"secret", "toml"},
	},
}

var filesLocal = map[string][][2]string{
	"./configs/client/": {
		{"local", "toml"},
	},
}

var filesCloud = map[string][][2]string{
	"./configs/client/": {
		{"cloud", "toml"},
	},
}

func FilesGet() map[string][][2]string {
	return files
}

func FilesGetByCloud(cloud bool) map[string][][2]string {
	if cloud {
		return filesCloud
	} else {
		return filesLocal
	}
}

var (
	client *ClientConfig
)

type ModuleConfig struct {
	name   string `mapstructure:"name"`
	enable bool   `mapstructure:"enable"`
	cloud  bool   `mapstructure:"cloud"`
	prod   bool   `mapstructure:"prod"`
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

func GetClient(force bool) *ClientConfig {
	if (client != nil) && !force {
		return client
	}
	module := &ModuleConfig{}
	if viper.IsSet(fmt.Sprintf("%s.%s", ClientKey, NameKey)) {
		module.name = viper.GetString(fmt.Sprintf("%s.%s", ClientKey, NameKey))
	} else {
		module.name = viper.GetString(NameKey)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", ClientKey, EnableKey)) {
		module.enable = viper.GetBool(fmt.Sprintf("%s.%s", ClientKey, EnableKey))
	} else {
		module.enable = viper.GetBool(EnableKey)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", ClientKey, CloudKey)) {
		module.cloud = viper.GetBool(fmt.Sprintf("%s.%s", ClientKey, CloudKey))
	} else {
		module.cloud = viper.GetBool(CloudKey)
	}
	if viper.IsSet(fmt.Sprintf("%s.%s", ClientKey, ProdKey)) {
		module.prod = viper.GetBool(fmt.Sprintf("%s.%s", ClientKey, ProdKey))
	} else {
		module.prod = viper.GetBool(ProdKey)
	}

	pgSql := newPgSql(module.prod, ClientKey)

	client = &ClientConfig{
		ModuleConfig: module,
		PgSql:        pgSql,
	}
	return client
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
