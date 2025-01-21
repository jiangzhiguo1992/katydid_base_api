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

// SetMarketHomes 应用市场页面 (方便控制台跳转)
func (c *ClientPlatform) SetMarketHomes(marketHomes *map[uint16]string) int {
	var count int
	if (marketHomes != nil) && (len(*marketHomes) > 0) {
		for k := range *marketHomes {
			ok := c.SetMarketHome(k, (*marketHomes)[k])
			if ok {
				count++
			}
		}
	} else {
		delete(c.Extra, "marketHomes")
	}
	return count
}

func (c *ClientPlatform) SetMarketHome(tp uint16, market string) bool {
	if !isMarketTypeOk(c.Platform, tp) {
		return false
	} else if len(market) <= 0 {
		if c.Extra["marketHomes"] != nil {
			delete((c.Extra["marketHomes"]).(map[uint16]string), tp)
		}
		return true
	}
	if c.Extra["marketHomes"] == nil {
		c.Extra["marketHomes"] = map[uint16]string{}
	}
	(c.Extra["marketHomes"]).(map[uint16]string)[tp] = market
	return true
}

func (c *ClientPlatform) GetMarketHomes() map[uint16]string {
	if c.Extra["marketHomes"] == nil {
		return map[uint16]string{}
	}
	return (c.Extra["marketHomes"]).(map[uint16]string)
}

func (c *ClientPlatform) GetMarket(tp uint16) string {
	if v, ok := c.GetMarketHomes()[tp]; ok {
		return v
	}
	return ""
}

// SetSocialLinks 社交链接 (方便控制台跳转)
func (c *ClientPlatform) SetSocialLinks(socialLinks *map[uint16]string) int {
	var count int
	if (socialLinks != nil) && (len(*socialLinks) > 0) {
		for k := range *socialLinks {
			ok := c.SetSocialLink(k, (*socialLinks)[k])
			if ok {
				count++
			}
		}
	} else {
		delete(c.Extra, "socialLinks")
	}
	return count
}

func (c *ClientPlatform) SetSocialLink(tp uint16, socialLink string) bool {
	if !isSocialLinkTypeOk(tp) {
		return false
	} else if len(socialLink) <= 0 {
		if c.Extra["socialLinks"] != nil {
			delete((c.Extra["socialLinks"]).(map[uint16]string), tp)
		}
		return true
	}
	if c.Extra["socialLinks"] == nil {
		c.Extra["socialLinks"] = map[uint16]string{}
	}
	(c.Extra["socialLinks"]).(map[uint16]string)[tp] = socialLink
	return true
}

func (c *ClientPlatform) GetSocialLinks() map[uint16]string {
	if c.Extra["socialLinks"] == nil {
		return map[uint16]string{}
	}
	return (c.Extra["socialLinks"]).(map[uint16]string)
}

func (c *ClientPlatform) GetSocialLink(tp uint16) string {
	if v, ok := c.GetSocialLinks()[tp]; ok {
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
	SocialLinkTypeEmail     uint16 = 1
	SocialLinkTypePhone     uint16 = 2
	SocialLinkTypeQQ        uint16 = 3
	SocialLinkTypeWeChat    uint16 = 4
	SocialLinkTypeWeibo     uint16 = 5
	SocialLinkTypeFacebook  uint16 = 6
	SocialLinkTypeTwitter   uint16 = 7
	SocialLinkTypeTelegram  uint16 = 8
	SocialLinkTypeDiscord   uint16 = 9
	SocialLinkTypeInstagram uint16 = 10
	SocialLinkTypeYouTube   uint16 = 11
	SocialLinkTypeTikTok    uint16 = 12
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

func isSocialLinkTypeOk(socialLinkType uint16) bool {
	switch socialLinkType {
	case SocialLinkTypeEmail,
		SocialLinkTypePhone,
		SocialLinkTypeQQ,
		SocialLinkTypeWeChat,
		SocialLinkTypeWeibo,
		SocialLinkTypeFacebook,
		SocialLinkTypeTwitter,
		SocialLinkTypeTelegram,
		SocialLinkTypeDiscord,
		SocialLinkTypeInstagram,
		SocialLinkTypeYouTube,
		SocialLinkTypeTikTok:
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

var socialLinkInfos = map[uint16]string{
	SocialLinkTypeEmail:     "Email",
	SocialLinkTypePhone:     "Phone",
	SocialLinkTypeQQ:        "QQ",
	SocialLinkTypeWeChat:    "WeChat",
	SocialLinkTypeWeibo:     "Weibo",
	SocialLinkTypeFacebook:  "Facebook",
	SocialLinkTypeTwitter:   "Twitter",
	SocialLinkTypeTelegram:  "Telegram",
	SocialLinkTypeDiscord:   "Discord",
	SocialLinkTypeInstagram: "Instagram",
	SocialLinkTypeYouTube:   "YouTube",
	SocialLinkTypeTikTok:    "TikTok",
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

func socialLinkName(socialLinkType uint16) string {
	if v, ok := socialLinkInfos[socialLinkType]; ok {
		return v
	}
	return ""
}
