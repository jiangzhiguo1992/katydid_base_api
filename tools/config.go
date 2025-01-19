package tools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

const (
	configsDir = "./configs"

	InitKey   = "init"
	CommonKey = "common"
	LocalKey  = "local"
	CloudKey  = "cloud"
	SecretKey = "secret"

	ModuleEnableKey = "module.enable"
	ModuleCloudKey  = "module.cloud"
	ModuleProdKey   = "module.prod"
)

func InitConfigs() (bool, bool) {
	mapFiles := getConfigsFiles()
	var settings map[string]any

	for _, f := range mapFiles[InitKey] {
		settings = setUpConfig(false, f[0], f[1], f[2])
	}
	cloud, prod := viper.GetBool(ModuleCloudKey), viper.GetBool(ModuleProdKey)
	viper.Reset()

	others := append(mapFiles[CommonKey], mapFiles[SecretKey]...)
	if cloud {
		others = append(others, mapFiles[CloudKey]...)
	} else {
		others = append(others, mapFiles[LocalKey]...)
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
	return cloud, prod
}

func getConfigsFiles() map[string][][3]string {
	// init
	initFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, InitKey))
	if err != nil {
		panic(err)
	}
	initParams := splitFiles(initFiles)

	// common
	commonFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, CommonKey))
	if err != nil {
		panic(err)
	}
	commonParams := splitFiles(commonFiles)

	// local
	localFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, LocalKey))
	if err != nil {
		panic(err)
	}
	localParams := splitFiles(localFiles)

	// cloud (ignore)
	cloudFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, CloudKey))
	if err != nil {
		panic(err)
	}
	cloudParams := splitFiles(cloudFiles)

	// secret (ignore)
	secretFiles, err := filepath.Glob(fmt.Sprintf("%s/%s/*.toml", configsDir, SecretKey))
	if err != nil {
		panic(err)
	}
	secretParams := splitFiles(secretFiles)

	return map[string][][3]string{
		InitKey:   initParams,
		CommonKey: commonParams,
		LocalKey:  localParams,
		CloudKey:  cloudParams,
		SecretKey: secretParams,
	}
}

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

func InitConfigsRemotes() {
	// TODO:GG 远程配置
}

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
