package model

import "time"

// TODO:GG PGSQL <- Clients = 100 * Client
// TODO:GG PGSQL <- Versions = Clients * Version
// TODO:GG Mongo <- Stats = Versions * (24*365*10) * 4, 懒惰add没有就不add, 数据来源于应用商场?某些渠道没有数据,启动可以自己做？ (数据量过多可以合并旧数据，时->日->月->年)
// TODO:GG Fetch <- Comments = 需要和Market同步，不存DB，api拉取

// Client 客户端
type Client struct {
	*Base
	IP   int `json:"IP"`   // 系列 (eg:大富翁IP)
	Part int `json:"part"` // 类型 (eg:单机版)

	BundleId   string `json:"bundleId"`   // 包名
	Name       string `json:"name"`       // 名称 (eg:大富翁联单机版/大富翁1)
	Website    string `json:"website"`    // 官网
	SupportUrl string `json:"supportUrl"` // 服务条款URL
	PrivacyUrl string `json:"privacyUrl"` // 隐私政策URL

	OnlineAt       int64 `json:"onlineAt"`       // 上线时间 (时间没到时，只能停留在首页)
	OfflineAt      int64 `json:"offlineAt"`      // 下线时间 (时间到后，强制升级/等待/etc)
	UserAccountMax int   `json:"userAccountMax"` // 用户最多账户数 (身份证/护照/...)
	UserTokenMax   int   `json:"userTokenMax"`   // 用户最多令牌数 (同时登录最大数，防止工作室?)
	Enable         bool  `json:"enable"`         // 是否可用 (一般不用，下架之类的，从conf读取)

	LatestVCodes map[int]map[int]map[int]int `json:"latestVCodes"` // [area][platform][market]最新运行版本号 (发版的时候自身被动更新)
	StatsList    map[int]int                 `json:"statsList"`    // [stats_kind]统计数据 (整点更新)
}

func NewClient(
	base *Base,
	IP int,
	part int,
	bundleId string,
	name string,
	website string,
	supportUrl string,
	privacyUrl string,
	userAccountMax int,
	userTokenMax int,
	enable bool,
) *Client {
	return &Client{
		Base:           base,
		IP:             IP,
		Part:           part,
		BundleId:       bundleId,
		Name:           name,
		Website:        website,
		SupportUrl:     supportUrl,
		PrivacyUrl:     privacyUrl,
		OnlineAt:       -1,
		OfflineAt:      -1,
		UserAccountMax: userAccountMax,
		UserTokenMax:   userTokenMax,
		Enable:         enable,
		LatestVCodes:   map[int]map[int]map[int]int{},
		StatsList:      map[int]int{},
	}
}

func (c *Client) IsOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt <= currentTime && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

func (c *Client) IsComingOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

func (c *Client) IsComingOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt > currentTime && (c.OnlineAt == -1 || c.OnlineAt < currentTime)
}

func (c *Client) GetVersionCode(area, platform, market int) int {
	if _, ok := c.LatestVCodes[area]; !ok {
		return 0
	}
	if _, ok := c.LatestVCodes[area][platform]; !ok {
		return 0
	}
	return c.LatestVCodes[area][platform][market]
}
