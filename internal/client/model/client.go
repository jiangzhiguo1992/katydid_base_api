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

	Website    string `json:"website"`    // 官网
	Company    string `json:"company"`    // 公司
	Copyright  string `json:"copyright"`  // 版权
	SupportUrl string `json:"supportUrl"` // 服务条款URL
	PrivacyUrl string `json:"privacyUrl"` // 隐私政策URL

	UserAccountMax int `json:"userAccountMax"` // 用户最多账户数 (身份证/护照/...)
	UserTokenMax   int `json:"userTokenMax"`   // 用户最多令牌数 (同时登录最大数，防止工作室?)

	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页)
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制升级/等待/etc)
	Enable    bool  `json:"enable"`    // 是否可用 (一般不用，下架之类的，从conf读取)

	Extra map[string]interface{} `json:"extra"` // 额外信息

	Platforms   map[int]map[int]*ClientPlatform        `json:"platforms"`   // [area][platform]平台列表
	LatestCodes map[int]map[int]map[int]*ClientVersion `json:"latestCodes"` // [area][platform][market]最新publish版本号
}

func NewClient(
	base *Base,
	IP int, part int,
	website string, company string, copyright string, supportUrl string, privacyUrl string,
	userAccountMax int, userTokenMax int,
	onlineAt int64, offlineAt int64, enable bool,
	extra map[string]interface{},
) *Client {
	return &Client{
		Base: base,
		IP:   IP, Part: part,
		Website: website, Company: company, Copyright: copyright, SupportUrl: supportUrl, PrivacyUrl: privacyUrl,
		UserAccountMax: userAccountMax, UserTokenMax: userTokenMax,
		OnlineAt: onlineAt, OfflineAt: offlineAt, Enable: enable,
		Extra:       extra,
		Platforms:   map[int]map[int]*ClientPlatform{},
		LatestCodes: map[int]map[int]map[int]*ClientVersion{},
	}
}

func (c *Client) IsOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt <= currentTime && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

func (c *Client) IsOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt <= currentTime && (c.OnlineAt == -1 || c.OnlineAt > currentTime)
}

func (c *Client) IsComingOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

func (c *Client) IsComingOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt > currentTime && (c.OnlineAt == -1 || c.OnlineAt < currentTime)
}

func (c *Client) GetPlatform(area, platform int) *ClientPlatform {
	if _, ok := c.Platforms[area]; !ok {
		return nil
	}
	return c.Platforms[area][platform]
}

func (c *Client) GetLatestCode(area, platform, market int) *ClientVersion {
	if _, ok := c.LatestCodes[area]; !ok {
		return nil
	}
	if _, ok := c.LatestCodes[area][platform]; !ok {
		return nil
	}
	return c.LatestCodes[area][platform][market]
}
