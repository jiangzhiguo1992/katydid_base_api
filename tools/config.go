package tools

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
)

const (
	configsDir = "./configs" // 目录

	keyInit   = "init"   // 初始化配置 (最先加载的)
	keyCommon = "common" // 公共配置 (始终加载，不被ignore)
	keyLocal  = "local"  // 本地配置 (本地加载，不被ignore)
	keyCloud  = "cloud"  // 云端配置 (云端加载，被ignore)
	keySecret = "secret" // 秘密配置 (始终加载，被ignore)

	ConfigKeyAppEnable = "app.enable" // 按照规则写
	ConfigKeyAppCloud  = "app.cloud"  // 按照规则写
	ConfigKeyAppProd   = "app.prod"   // 按照规则写
)

// InitConfigs 初始化配置
func InitConfigs() (bool, bool, bool) {
	mapFiles := getConfigsFiles()
	var settings map[string]any

	for _, f := range mapFiles[keyInit] {
		settings = setUpConfig(false, f[0], f[1], f[2])
	}
	enable, cloud, prod := viper.GetBool(ConfigKeyAppEnable), viper.GetBool(ConfigKeyAppCloud), viper.GetBool(ConfigKeyAppProd)
	viper.Reset()

	others := append(mapFiles[keyCommon], mapFiles[keySecret]...)
	if cloud {
		others = append(others, mapFiles[keyCloud]...)
	} else {
		others = append(others, mapFiles[keyLocal]...)
	}

	for _, f := range others {
		setUpConfig(true, f[0], f[1], f[2])
		sets := viper.AllSettings()
		for k, v := range sets {
			settings[k] = v
		}
		viper.Reset()
	}

	for k, v := range settings {
		viper.Set(k, v)
	}
	//viper.AutomaticEnv()

	//tests := viper.AllSettings()
	//fmt.Printf("------>%v\n", tests)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO:GG 配置会被更新，这里要做一些相关的re_init操作
		fmt.Printf("config file changed name:%s\n", e.Name)
	})
	return enable, cloud, prod
}

// getConfigsFiles 获取配置文件
func getConfigsFiles() map[string][][3]string {
	// creates
	dirs := []string{
		path.Join(configsDir, keyInit),
		path.Join(configsDir, keyCommon),
		path.Join(configsDir, keyLocal),
		path.Join(configsDir, keyCloud),
		path.Join(configsDir, keySecret),
	}
	for _, v := range dirs {
		if err := os.MkdirAll(v, os.ModePerm); err != nil {
			panic(errors.New(fmt.Sprintf("failed to create configs_dir %s: %s", v, err)))
		}
	}

	// init
	initFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, keyInit))
	if err != nil {
		panic(err)
	}
	initParams := splitFiles(initFiles)

	// common
	commonFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, keyCommon))
	if err != nil {
		panic(err)
	}
	commonParams := splitFiles(commonFiles)

	// local
	localFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, keyLocal))
	if err != nil {
		panic(err)
	}
	localParams := splitFiles(localFiles)

	// cloud (ignore)
	cloudFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, keyCloud))
	if err != nil {
		panic(err)
	}
	cloudParams := splitFiles(cloudFiles)

	// secret (ignore)
	secretFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, keySecret))
	if err != nil {
		panic(err)
	}
	secretParams := splitFiles(secretFiles)

	return map[string][][3]string{
		keyInit:   initParams,
		keyCommon: commonParams,
		keyLocal:  localParams,
		keyCloud:  cloudParams,
		keySecret: secretParams,
	}
}

// splitFiles 分割文件
func splitFiles(files []string) [][3]string {
	var params [][3]string
	for _, file := range files {
		dir := filepath.Dir(file)
		name := filepath.Base(file)
		ext := filepath.Ext(file)
		name = name[:len(name)-len(ext)]
		params = append(params, [3]string{dir, name, ext[1:]})
	}
	return params
}

// setUpConfig 设置配置
func setUpConfig(merge bool, path, name, suffix string) map[string]any {
	fmt.Printf("加载配置, path:%s/%s.%s\n", path, name, suffix)
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(suffix)
	if merge {
		if err := viper.MergeInConfig(); err != nil {
			log.Fatalf("merge config failed: %v", err)
		}
	} else {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("read config failed: %v", err)
		}
	}
	return viper.AllSettings()
}

// InitConfigsRemotes 初始化远程配置
func InitConfigsRemotes() {
	// TODO:GG 远程配置
}

// ConfigEnvKey 获取配置环境变量
func ConfigEnvKey(prod bool, tag, key string) string {
	if len(tag) <= 0 {
		panic(fmt.Sprintf("config tag is empty: %s", tag))
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
	panic(fmt.Sprintf("config not found tag.key: %s%s", tag, key))
}
