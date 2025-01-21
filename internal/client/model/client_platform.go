package model

import (
	"katydid_base_api/internal/pkg/database"
	"time"
)

// ClientPlatform 客户端平台
type ClientPlatform struct {
	*database.BaseModel
	Cid      uint64 `json:"cid"`      // 客户端id TODO:GG idx
	Platform uint16 `json:"platform"` // 平台 TODO:GG idx_1
	Area     uint16 `json:"area"`     // 区域编号 TODO:GG idx_1

	Enable    bool  `json:"enable"`    // 是否可用 (一般不用，下架之类的，没有reason) TODO:GG idx
	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页，提示bulletins) TODO:GG idx
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制下线+升级/等待/...) TODO:GG idx

	AppId   string `json:"appId"`   // 各平台应用唯一标识 (pkg/bundle，海外和大陆可以同时安装!) TODO:GG idx
	AppName string `json:"appName"` // app名称 (不同Area的区别) TODO:GG idx

	Extra map[string]interface{} `json:"extra" gorm:"serializer:json"` // 额外信息

	LatestVersion map[uint16]*ClientVersion `json:"latestVersion" gorm:"-:all"` // [market]最新publish版本号
}

func NewClientPlatformDefault(
	Cid uint64, platform uint16, area uint16,
	enable bool,
	appId string, appName string,
) *ClientPlatform {
	if !isPlatformTypeOk(platform) || !isAreaTypeOk(area) {
		return nil
	} else if (len(appId) <= 0) || (len(appName) <= 0) {
		return nil
	}
	return &ClientPlatform{
		BaseModel: database.NewBaseModelEmpty(),
		Cid:       Cid, Platform: platform, Area: area,
		Enable: enable, OnlineAt: -1, OfflineAt: -1,
		AppId: appId, AppName: appName,
		Extra:         map[string]interface{}{},
		LatestVersion: make(map[uint16]*ClientVersion),
	}
}

// IsOnline 是否上线
func (c *ClientPlatform) IsOnline() bool {
	currentTime := time.Now().Unix()
	return (c.OnlineAt > 0 && (c.OnlineAt <= currentTime)) && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

// IsOffline 是否下线
func (c *ClientPlatform) IsOffline() bool {
	currentTime := time.Now().Unix()
	return (c.OfflineAt > 0 && (c.OfflineAt <= currentTime)) && (c.OnlineAt == -1 || c.OnlineAt > currentTime)
}

// IsComingOnline 是否即将上线
func (c *ClientPlatform) IsComingOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

// IsComingOffline 是否即将下线
func (c *ClientPlatform) IsComingOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt > currentTime && (c.OnlineAt == -1 || c.OnlineAt < currentTime)
}

// SetMarkets 应用市场页面 (方便控制台跳转)
func (c *ClientPlatform) SetMarkets(appMarkets *map[uint16]string) int {
	var count int
	if (appMarkets != nil) && (len(*appMarkets) > 0) {
		for k := range *appMarkets {
			ok := c.SetMarket(k, (*appMarkets)[k])
			if ok {
				count++
			}
		}
	} else {
		delete(c.Extra, "appMarkets")
	}
	return count
}

func (c *ClientPlatform) SetMarket(tp uint16, market string) bool {
	if !isMarketTypeOk(c.Platform, tp) {
		return false
	} else if len(market) <= 0 {
		if c.Extra["appMarkets"] != nil {
			delete((c.Extra["appMarkets"]).(map[uint16]string), tp)
		}
		return true
	}
	if c.Extra["appMarkets"] == nil {
		c.Extra["appMarkets"] = map[uint16]string{}
	}
	(c.Extra["appMarkets"]).(map[uint16]string)[tp] = market
	return true
}

func (c *ClientPlatform) GetMarkets() map[uint16]string {
	if c.Extra["appMarkets"] == nil {
		return map[uint16]string{}
	}
	return (c.Extra["appMarkets"]).(map[uint16]string)
}

func (c *ClientPlatform) GetMarket(tp uint16) string {
	if v, ok := c.GetMarkets()[tp]; ok {
		return v
	}
	return ""
}

// SetSocials 社交链接 (方便控制台跳转)
func (c *ClientPlatform) SetSocials(socials *map[uint16]string) int {
	var count int
	if (socials != nil) && (len(*socials) > 0) {
		for k := range *socials {
			ok := c.SetSocial(k, (*socials)[k])
			if ok {
				count++
			}
		}
	} else {
		delete(c.Extra, "socials")
	}
	return count
}

func (c *ClientPlatform) SetSocial(tp uint16, social string) bool {
	if !isSocialTypeOk(tp) {
		return false
	} else if len(social) <= 0 {
		if c.Extra["socials"] != nil {
			delete((c.Extra["socials"]).(map[uint16]string), tp)
		}
		return true
	}
	if c.Extra["socials"] == nil {
		c.Extra["socials"] = map[uint16]string{}
	}
	(c.Extra["socials"]).(map[uint16]string)[tp] = social
	return true
}

func (c *ClientPlatform) GetSocials() map[uint16]string {
	if c.Extra["socials"] == nil {
		return map[uint16]string{}
	}
	return (c.Extra["socials"]).(map[uint16]string)
}

func (c *ClientPlatform) GetSocial(tp uint16) string {
	if v, ok := c.GetSocials()[tp]; ok {
		return v
	}
	return ""
}

// SetIosId apple应用市场id
func (c *ClientPlatform) SetIosId(iosId *string) {
	if (iosId != nil) && (len(*iosId) > 0) {
		c.Extra["iosId"] = *iosId
	} else {
		delete(c.Extra, "iosId")
	}
}

func (c *ClientPlatform) GetIosId() string {
	if c.Extra["iosId"] == nil {
		return ""
	}
	return c.Extra["iosId"].(string)
}

func (c *ClientPlatform) GetLatestVersion(market uint16) *ClientVersion {
	if v, ok := c.LatestVersion[market]; ok {
		return v
	}
	return nil
}

const (
	PlatformTypeLinux   uint16 = 1
	PlatformTypeWindows uint16 = 2
	PlatformTypeMacOS   uint16 = 3
	PlatformTypeWeb     uint16 = 4
	PlatformTypeAndroid uint16 = 5
	PlatformTypeIOS     uint16 = 6
	PlatformTypeApplet  uint16 = 7
	//PlatformTypeTvAnd   uint16 = 8
	//PlatformTypeTvIOS   uint16 = 9
)

const (
	AreaTypeWord      uint16 = 1 // 全球 (默认英文，泛指海外)
	AreaTypeChinaLand uint16 = 2 // 中国大陆 (简体中文)
	AreaTypeChinaHMT  uint16 = 3 // 中国港澳台 (繁体中文)
	AreaTypeEurope    uint16 = 4 // 欧洲 (默认英文，GDPR)
	// 后面可能会划分更多的区域
)

const (
	SocialTypeEmail     uint16 = 1
	SocialTypePhone     uint16 = 2
	SocialTypeQQ        uint16 = 3
	SocialTypeWeChat    uint16 = 4
	SocialTypeWeibo     uint16 = 5
	SocialTypeFacebook  uint16 = 6
	SocialTypeTwitter   uint16 = 7
	SocialTypeTelegram  uint16 = 8
	SocialTypeDiscord   uint16 = 9
	SocialTypeInstagram uint16 = 10
	SocialTypeYouTube   uint16 = 11
	SocialTypeTikTok    uint16 = 12
)

func isPlatformTypeOk(platformType uint16) bool {
	switch platformType {
	case PlatformTypeLinux,
		PlatformTypeWindows,
		PlatformTypeMacOS,
		PlatformTypeWeb,
		PlatformTypeAndroid,
		PlatformTypeIOS,
		PlatformTypeApplet:
		//PlatformTypeTvAnd,
		//PlatformTypeTvIOS:
		return true
	}
	return false
}

func isAreaTypeOk(areaType uint16) bool {
	switch areaType {
	case AreaTypeWord,
		AreaTypeChinaLand,
		AreaTypeChinaHMT,
		AreaTypeEurope:
		return true
	}
	return false
}

func isSocialTypeOk(socialType uint16) bool {
	switch socialType {
	case SocialTypeEmail,
		SocialTypePhone,
		SocialTypeQQ,
		SocialTypeWeChat,
		SocialTypeWeibo,
		SocialTypeFacebook,
		SocialTypeTwitter,
		SocialTypeTelegram,
		SocialTypeDiscord,
		SocialTypeInstagram,
		SocialTypeYouTube,
		SocialTypeTikTok:
		return true
	}
	return false
}

var platformInfos = map[uint16]string{
	PlatformTypeLinux:   "Linux",
	PlatformTypeWindows: "Windows",
	PlatformTypeMacOS:   "MacOS",
	PlatformTypeWeb:     "Web",
	PlatformTypeAndroid: "Android",
	PlatformTypeIOS:     "IOS",
	PlatformTypeApplet:  "Applet",
	//PlatformTypeTvAnd: "TvAnd",
	//PlatformTypeTvIOS: "TvIOS",
}

var areaInfos = map[uint16]string{
	AreaTypeWord:      "Word",
	AreaTypeChinaLand: "ChinaLand",
	AreaTypeChinaHMT:  "ChinaHMT",
	AreaTypeEurope:    "Europe",
}

var socialInfos = map[uint16]string{
	SocialTypeEmail:     "Email",
	SocialTypePhone:     "Phone",
	SocialTypeQQ:        "QQ",
	SocialTypeWeChat:    "WeChat",
	SocialTypeWeibo:     "Weibo",
	SocialTypeFacebook:  "Facebook",
	SocialTypeTwitter:   "Twitter",
	SocialTypeTelegram:  "Telegram",
	SocialTypeDiscord:   "Discord",
	SocialTypeInstagram: "Instagram",
	SocialTypeYouTube:   "YouTube",
	SocialTypeTikTok:    "TikTok",
}

func platformName(platformType uint16) string {
	if v, ok := platformInfos[platformType]; ok {
		return v
	}
	return ""
}

func areaName(areaType uint16) string {
	if v, ok := areaInfos[areaType]; ok {
		return v
	}
	return ""
}

func socialName(socialType uint16) string {
	if v, ok := socialInfos[socialType]; ok {
		return v
	}
	return ""
}
