package tools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func InitConfig(cloud bool, configs [][3]string) {
	skipSuffix := "_cloud"
	if cloud {
		skipSuffix = "_local"
	}
	for _, config := range configs {
		if strings.HasSuffix(config[1], skipSuffix) {
			fmt.Printf("跳过配置, path:%s%s.%s\n", config[0], config[1], config[2])
			continue
		}
		setUpConfig(config[0], config[1], config[2])
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO:GG 配置会被更新，这里要做一些相关的re_init操作
		fmt.Printf("config file changed name:%s\n", e.Name)
	})
}

func setUpConfig(path, name, suffix string) {
	fmt.Printf("加载配置, path:%s%s.%s\n", path, name, suffix)
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(suffix)
	if err := viper.ReadInConfig(); err != nil {
		if !strings.HasSuffix(name, "_cloud") {
			log.Fatalf("read config failed: %v", err)
		}
	}
}

func ConfigEnvKey(prod bool, key1, key2 string) string {
	if len(key1) <= 0 {
		panic(fmt.Sprintf("key1 is empty: %s", key1))
	}
	suffix := "dev"
	if prod {
		suffix = "prod"
	}
	suffixs := [2]string{fmt.Sprintf(".%s", suffix), ""}
	if len(key2) > 0 {
		key2 = fmt.Sprintf(".%s", key2)
	}

	for i := 0; i < len(suffixs); i++ {
		tag := fmt.Sprintf("%s%s", key1, suffixs[i])
		key := fmt.Sprintf("%s%s", tag, key2)
		if viper.IsSet(key) {
			return key
		}
	}
	panic(fmt.Sprintf("not found key: %s.%s", key1, key2))
}
