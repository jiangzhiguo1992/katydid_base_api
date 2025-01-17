package configs

import (
	"github.com/spf13/viper"
	"katydid_base_api/tools"
)

type AppConfig struct {
	cloud bool `mapstructure:"cloud"`
	prod  bool `mapstructure:"prod"`
}

type PgSqlConfig struct {
	Name   string `mapstructure:"name"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pwd    string `mapstructure:"pwd"`
	DBName string `mapstructure:"db_name"`
}

var (
	app   *AppConfig
	pgsql *PgSqlConfig
)

func GetPgSqlConfig(cloud, prod bool) *PgSqlConfig {
	if pgsql != nil {
		return pgsql
	}
	pgsql = &PgSqlConfig{}
	name := tools.ConfigEnvKey(cloud, prod, "pgsql", "name")
	pgsql.Name = viper.GetString(name)
	host := tools.ConfigEnvKey(cloud, prod, "pgsql", "host")
	pgsql.Host = viper.GetString(host)
	port := tools.ConfigEnvKey(cloud, prod, "pgsql", "port")
	pgsql.Port = viper.GetString(port)
	user := tools.ConfigEnvKey(cloud, prod, "pgsql", "user")
	pgsql.User = viper.GetString(user)
	pwd := tools.ConfigEnvKey(cloud, prod, "pgsql", "pwd")
	pgsql.Pwd = viper.GetString(pwd)
	dbName := tools.ConfigEnvKey(cloud, prod, "pgsql", "db_name")
	pgsql.DBName = viper.GetString(dbName)
	return pgsql
}
