package model

import "time"

// TODO:GG PGSQL <- Clients = 100 * Client
// TODO:GG PGSQL <- Versions = Clients * Version
// TODO:GG Mongo <- Stats = Versions * (24*365*10) * 4, 懒惰add没有就不add, 数据来源于应用商场?某些渠道没有数据,启动可以自己做？
// TODO:GG Fetch <- Comments = 需要和Market同步，不存DB，api拉取

// Client 客户端 (UIndex = id/ IP+Part)
type Client struct {
	*Base
	IP         int    `json:"IP"`         // 系列 (eg:大富翁IP)
	Part       int    `json:"part"`       // 类型 (eg:单机版)
	BundleId   string `json:"bundleId"`   // 包名
	Name       string `json:"name"`       // 名称 (eg:大富翁联单机版/大富翁1)
	Website    string `json:"website"`    // 官网
	SupportUrl string `json:"supportUrl"` // 服务条款URL
	PrivacyUrl string `json:"privacyUrl"` // 隐私政策URL

	SSO    bool `json:"sso"`    // 是否单点登录 (Single Sign-On)
	SBO    bool `json:"sbo"`    // 身份唯一 (身份证/护照/...) (Single Bio)
	Enable bool `json:"enable"` // 是否可用 (一般不用，下架之类的，从conf读取)

	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页)
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制升级/等待/etc)
	Watched   int   `json:"watched"`   // 总查看数量 (整点更新)
	Download  int   `json:"download"`  // 总下载数量 (整点更新)
	Opener    int   `json:"opener"`    // 总启动数量 (整点更新)
	Score     int   `json:"score"`     // 当前总评分 (整点更新)
	Comments  int   `json:"comments"`  // 当前总评数 (整点更新)

	LatestVersionNames map[int]map[int]map[int]string `json:"latestVersionNames"` // [area][platform][market]最新版本名 (发版的时候自身被动更新)
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
	SSO bool,
	SBO bool,
	enable bool,
) *Client {
	return &Client{
		Base:               base,
		IP:                 IP,
		Part:               part,
		BundleId:           bundleId,
		Name:               name,
		Website:            website,
		SupportUrl:         supportUrl,
		PrivacyUrl:         privacyUrl,
		SSO:                SSO,
		SBO:                SBO,
		Enable:             enable,
		OnlineAt:           -1,
		OfflineAt:          -1,
		Watched:            -1,
		Download:           -1,
		Opener:             -1,
		Score:              -1,
		Comments:           -1,
		LatestVersionNames: map[int]map[int]map[int]string{},
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
