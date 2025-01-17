package tools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"log/slog"
)

func InitConfig(configs [][]string) {
	setUpConfigs(configs)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO:GG 配置会被更新，这里要做一些相关的re_init操作
		slog.Debug("config file changed", "name", e.Name)
	})
}

func setUpConfigs(configs [][]string) {
	for _, config := range configs {
		setUpConfig(config[0], config[1], config[2])
	}
}

func setUpConfig(path, name, suffix string) {
	slog.Debug("加载配置", "path", path, "name", name, "suffix", suffix)
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(suffix)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err) // TODO:GG 输出到logs目录
	}
}

func ConfigEnvKey(cloud, prod bool, key1, key2 string) string {
	if len(key1) <= 0 {
		panic("tag is empty") // TODO:GG 输出到logs目录
	}
	suffix1 := "local"
	if cloud {
		suffix1 = "cloud"
	}
	suffix2 := "dev"
	if prod {
		suffix2 = "prod"
	}

	suffixs := [4][2]string{
		{fmt.Sprintf(".%s", suffix1), fmt.Sprintf(".%s", suffix2)},
		{fmt.Sprintf(".%s", suffix1), ""},
		{"", fmt.Sprintf(".%s", suffix2)},
		{"", ""},
	}
	if len(key2) > 0 {
		key2 = fmt.Sprintf(".%s", key2)
	}

	for i := 0; i < len(suffixs); i++ {
		tag := fmt.Sprintf("%s%s%s", key1, suffixs[i][0], suffixs[i][1])
		key := fmt.Sprintf("%s%s", tag, key2)
		if viper.IsSet(key) {
			return key
		}
	}
	panic(fmt.Sprintf("not found key: %s %s", key1, key2)) // TODO:GG 输出到logs目录
}
