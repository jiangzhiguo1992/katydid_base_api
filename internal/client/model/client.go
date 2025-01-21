package model

import (
	"katydid_base_api/internal/pkg/dababase"
	"time"
)

// Client 客户端
type Client struct {
	*dababase.BaseModel
	IP   uint `json:"IP"`   // 系列 (eg:大富翁IP)
	Part uint `json:"part"` // 类型 (eg:单机版)

	Enable    bool  `json:"enable"`    // 是否可用 (一般不用，下架之类的)
	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页)
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制下线+升级/等待/...)

	Organization string `json:"organization"` // 组织

	Extra map[string]interface{} `json:"extra" gorm:"serializer:json"` // 额外信息

	Platforms   map[int]map[int]*ClientPlatform        `json:"platforms" gorm:"-:all"`   // [area][platform]平台列表
	LatestCodes map[int]map[int]map[int]*ClientVersion `json:"latestCodes" gorm:"-:all"` // [area][platform][market]最新publish版本号
}

func NewClientDefault(
	IP uint, part uint,
	enable bool,
	organization string,
) *Client {
	return &Client{
		BaseModel: dababase.NewBaseModelEmpty(),
		IP:        IP, Part: part,
		Enable: enable, OnlineAt: -1, OfflineAt: -1,
		Organization: organization,
		Extra:        map[string]interface{}{},
		Platforms:    make(map[int]map[int]*ClientPlatform),
		LatestCodes:  make(map[int]map[int]map[int]*ClientVersion),
	}
}

// IsOnline 是否上线
func (c *Client) IsOnline() bool {
	currentTime := time.Now().Unix()
	return (c.OnlineAt > 0 && (c.OnlineAt <= currentTime)) && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

// IsOffline 是否下线
func (c *Client) IsOffline() bool {
	currentTime := time.Now().Unix()
	return (c.OfflineAt > 0 && (c.OfflineAt <= currentTime)) && (c.OnlineAt == -1 || c.OnlineAt > currentTime)
}

// IsComingOnline 是否即将上线
func (c *Client) IsComingOnline() bool {
	currentTime := time.Now().Unix()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

// IsComingOffline 是否即将下线
func (c *Client) IsComingOffline() bool {
	currentTime := time.Now().Unix()
	return c.OfflineAt > currentTime && (c.OnlineAt == -1 || c.OnlineAt < currentTime)
}

// SetWebsite 官网
func (c *Client) SetWebsite(website *string) {
	if (website != nil) && (len(*website) > 0) {
		c.Extra["website"] = *website
	} else {
		delete(c.Extra, "website")
	}
}

func (c *Client) GetWebsite() string {
	if c.Extra == nil || c.Extra["website"] == nil {
		return ""
	}
	return c.Extra["website"].(string)
}

// SetCopyrights 版权
func (c *Client) SetCopyrights(copyrights *[]string) {
	if (copyrights != nil) && (len(*copyrights) > 0) {
		c.Extra["copyrights"] = *copyrights
	} else {
		delete(c.Extra, "copyrights")
	}
}

func (c *Client) GetCopyrights() []string {
	if c.Extra == nil || c.Extra["copyrights"] == nil {
		return []string{}
	}
	return c.Extra["copyrights"].([]string)
}

// SetSupportUrl 服务条款URL
func (c *Client) SetSupportUrl(supportUrl *string) {
	if (supportUrl != nil) && (len(*supportUrl) > 0) {
		c.Extra["supportUrl"] = *supportUrl
	} else {
		delete(c.Extra, "supportUrl")
	}
}

func (c *Client) GetSupportUrl() string {
	if c.Extra == nil || c.Extra["supportUrl"] == nil {
		return ""
	}
	return c.Extra["supportUrl"].(string)
}

// SetPrivacyUrl 隐私政策URL
func (c *Client) SetPrivacyUrl(privacyUrl *string) {
	if (privacyUrl != nil) && (len(*privacyUrl) > 0) {
		c.Extra["privacyUrl"] = *privacyUrl
	} else {
		delete(c.Extra, "privacyUrl")
	}
}

func (c *Client) GetPrivacyUrl() string {
	if c.Extra == nil || c.Extra["privacyUrl"] == nil {
		return ""
	}
	return c.Extra["privacyUrl"].(string)
}

// SetBulletins 维护公告 (不同于version_log，不需要升级)
func (c *Client) SetBulletins(bulletins *[]string) {
	if (bulletins != nil) && (len(*bulletins) > 0) {
		c.Extra["bulletins"] = *bulletins
	} else {
		delete(c.Extra, "bulletins")
	}
}

func (c *Client) GetBulletins() []string {
	if c.Extra == nil || c.Extra["bulletins"] == nil {
		return []string{}
	}
	return c.Extra["bulletins"].([]string)
}

func (c *Client) GetBulletinLatest() string {
	bulletins := c.GetBulletins()
	if len(bulletins) <= 0 {
		return ""
	}
	return bulletins[len(bulletins)-1]
}

// SetUserMaxAccount 用户最多账户数 (身份证/护照/...)
func (c *Client) SetUserMaxAccount(userMaxAccount *int) {
	if (userMaxAccount != nil) && (*userMaxAccount > 0) {
		c.Extra["userMaxAccount"] = *userMaxAccount
	} else {
		delete(c.Extra, "userMaxAccount")
	}
}

func (c *Client) GetUserMaxAccount() int {
	if c.Extra == nil || c.Extra["userMaxAccount"] == nil {
		return -1
	}
	return c.Extra["userMaxAccount"].(int)
}

func (c *Client) OverUserMaxAccount(count int) bool {
	maxCount := c.GetUserMaxAccount()
	if maxCount < 0 {
		return false
	}
	return count > maxCount
}

// SetUserMaxToken 用户最多令牌数 (同时登录最大数，防止工作室?)
func (c *Client) SetUserMaxToken(userMaxToken *int) {
	if (userMaxToken != nil) && (*userMaxToken > 0) {
		c.Extra["userMaxToken"] = *userMaxToken
	} else {
		delete(c.Extra, "userMaxToken")
	}
}

func (c *Client) GetUserMaxToken() int {
	if c.Extra == nil || c.Extra["userMaxToken"] == nil {
		return -1
	}
	return c.Extra["userMaxToken"].(int)
}

func (c *Client) OverUserMaxToken(count int) bool {
	maxCount := c.GetUserMaxToken()
	if maxCount < 0 {
		return false
	}
	return count > maxCount
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
