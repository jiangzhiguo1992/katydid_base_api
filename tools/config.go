package tools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitConfigStarts(configs map[string][][2]string, cloudKey, prodKey string) (bool, bool) {
	for dir, files := range configs {
		for _, f := range files {
			setUpConfig(dir, f[0], f[1])
		}
	}
	cloud, prod := viper.GetBool(cloudKey), viper.GetBool(prodKey)
	return cloud, prod
}

func InitConfigEnds(configs map[string][][2]string) {
	for dir, files := range configs {
		for _, f := range files {
			setUpConfig(dir, f[0], f[1])
		}
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
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

func InitConfigsRemotes() {
	// TODO:GG 远程配置
}

func ConfigEnvKey(prod bool, tag, key string) string {
	if len(tag) <= 0 {
		panic(fmt.Sprintf("donfig tag is empty: %s", tag))
	}

	// env
	suffix := "prod"
	if !prod {
		suffix = "dev"
	}
	suffixs := [2]string{fmt.Sprintf(".%s", suffix), ""}

	// key
	if len(key) > 0 {
		key = fmt.Sprintf(".%s", key)
	}

	for i := 0; i < len(suffixs); i++ {
		t := fmt.Sprintf("%s%s", tag, suffixs[i])
		name := fmt.Sprintf("%s%s", t, key)
		if viper.IsSet(name) {
			return name
		}
	}
	panic(fmt.Sprintf("donfig not found tag.key: %s%s", tag, key))
}
