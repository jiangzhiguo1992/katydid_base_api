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

func InitConfigsRemotes() {
	// TODO:GG 远程配置
}

// TODO:GG 修改
func setUpConfig(path, name, suffix string) {
	fmt.Printf("加载配置, path:%s%s.%s\n", path, name, suffix)
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(suffix)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

func ConfigEnvKey(prod bool, module, tag, key string) string {
	if len(tag) <= 0 {
		panic(fmt.Sprintf("key1 is empty: %s", tag))
	}

	// suffix_1
	suffix1 := ""
	if len(module) > 0 {
		suffix1 = fmt.Sprintf(".%s", module)
	}
	var suffix1s []string
	if len(suffix1) > 0 {
		suffix1s = []string{suffix1, ""}
	} else {
		suffix1s = []string{""}
	}

	// suffix_2
	suffix2 := "dev"
	if prod {
		suffix2 = "prod"
	}
	suffix2s := [2]string{fmt.Sprintf(".%s", suffix2), ""}

	// key
	if len(key) > 0 {
		key = fmt.Sprintf(".%s", key)
	}

	for n := 0; n < len(suffix1s); n++ {
		for i := 0; i < len(suffix2s); i++ {
			t := fmt.Sprintf("%s%s%s", tag, suffix1s[n], suffix2s[i])
			name := fmt.Sprintf("%s%s", t, key)
			if viper.IsSet(name) {
				return name
			}
		}
	}
	panic(fmt.Sprintf("not found key: %s.%s", tag, key))
}
