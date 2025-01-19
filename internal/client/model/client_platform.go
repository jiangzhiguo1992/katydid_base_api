package model

import "time"

const (
	PlatformTypeWeb     = 1
	PlatformTypeAndroid = 2
	PlatformTypeIOS     = 3
	PlatformTypePC      = 4
	PlatformTypeMACOS   = 5
	PlatformTypeWeChat  = 6
	PlatformTypeDouYin  = 7
	PlatformTypeTaoBao  = 8
)

const (
	LinkTypeEmail     = 1
	LinkTypePhone     = 2
	LinkTypeQQ        = 3
	LinkTypeWeChat    = 4
	LinkTypeWeibo     = 5
	LinkTypeFacebook  = 6
	LinkTypeTwitter   = 7
	LinkTypeInstagram = 8
	LinkTypeYouTube   = 9
	LinkTypeTikTok    = 10
)

// ClientPlatform 客户端平台
type ClientPlatform struct {
	*Base
	Cid      int64 `json:"cid"`      // 客户端id
	Area     int   `json:"area"`     // 区域编号 (从conf读取，国家/洲际/...)
	Platform int   `json:"platform"` // 平台

	AppId     string         `json:"appId"`     // 标识 (各平台pkg/bundle)
	AppMarket map[int]string `json:"appMarket"` // [market]app市场页面 (方便控制台跳转)
	Links     map[int]string `json:"links"`     // [link]链接

	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页)
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制升级/等待/etc)
	Enable    bool  `json:"enable"`    // 是否可用 (没有reason)

	Extra map[string]interface{} `json:"extra"` // 额外信息

	LatestVersion map[int]*ClientVersion `json:"latestVersion"` // [market]最新publish版本号
}

func NewClientPlatform(
	base *Base,
	Cid int64, area int, platform int,
	appId string, appMarket map[int]string, links map[int]string,
	onlineAt int64, offlineAt int64, enable bool,
	extra map[string]interface{},
) *ClientPlatform {
	return &ClientPlatform{
		Base: base,
		Cid:  Cid, Area: area, Platform: platform,
		AppId: appId, AppMarket: appMarket, Links: links,
		OnlineAt: onlineAt, OfflineAt: offlineAt, Enable: enable,
		Extra:         extra,
		LatestVersion: map[int]*ClientVersion{},
	}
}

func (c *ClientPlatform) GetAppMarket(marketType int) string {
	if v, ok := c.AppMarket[marketType]; ok {
		return v
	}
	return ""
}

func (c *ClientPlatform) GetLink(linkType int) string {
	if v, ok := c.Links[linkType]; ok {
		return v
	}
	return ""
}

func (c *ClientPlatform) IsOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt <= currentTime && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

func (c *ClientPlatform) IsOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt <= currentTime && (c.OnlineAt == -1 || c.OnlineAt > currentTime)
}

func (c *ClientPlatform) IsComingOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

func (c *ClientPlatform) IsComingOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt > currentTime && (c.OnlineAt == -1 || c.OnlineAt < currentTime)
}

func (c *ClientPlatform) GetLatestVersion(market int) *ClientVersion {
	if v, ok := c.LatestVersion[market]; ok {
		return v
	}
	return nil
}
