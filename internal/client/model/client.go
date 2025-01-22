package model

import (
	"fmt"
	"katydid_base_api/internal/pkg/base"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
	"time"
)

// Client 客户端
type Client struct {
	*base.DBModel
	IP   uint `json:"IP"`   // 系列 (eg:大富翁IP)
	Part uint `json:"part"` // 类型 (eg:单机版)

	Enable    bool  `json:"enable"`    // 是否可用 (一般不用，下架之类的，没有reason)
	OnlineAt  int64 `json:"onlineAt"`  // 上线时间 (时间没到时，只能停留在首页，提示bulletins)
	OfflineAt int64 `json:"offlineAt"` // 下线时间 (时间到后，强制下线+升级/等待/...)

	IPName       string `json:"IPName"`       // ip名称
	Organization string `json:"organization"` // 组织

	Extra map[string]interface{} `json:"extra" gorm:"serializer:json"` // 额外信息

	Platforms   map[uint16]map[uint16]*ClientPlatform           `json:"platforms" gorm:"-:all"`   // [platform][area]平台列表
	LatestCodes map[uint16]map[uint16]map[uint16]*ClientVersion `json:"latestCodes" gorm:"-:all"` // [platform][area][market]最新publish版本号
}

func NewClientDefault(
	IP uint, part uint,
	enable bool,
	IPName string, organization string,
) *Client {
	client := &Client{
		DBModel: base.NewDBModelEmpty(),
		IP:      IP, Part: part,
		Enable: enable, OnlineAt: -1, OfflineAt: -1,
		IPName: IPName, Organization: organization,
		Extra:       map[string]interface{}{},
		Platforms:   make(map[uint16]map[uint16]*ClientPlatform),
		LatestCodes: make(map[uint16]map[uint16]map[uint16]*ClientVersion),
	}
	client.FieldsCheck = client.CheckFields
	return client
}

// IsOnline 是否上线
func (c *Client) IsOnline() bool {
	currentTime := time.Now().UnixMilli()
	return (c.OnlineAt > 0 && (c.OnlineAt <= currentTime)) && (c.OfflineAt == -1 || c.OfflineAt > currentTime)
}

// IsOffline 是否下线
func (c *Client) IsOffline() bool {
	currentTime := time.Now().UnixMilli()
	return (c.OfflineAt > 0 && (c.OfflineAt <= currentTime)) && (c.OnlineAt == -1 || c.OnlineAt > currentTime)
}

// IsComingOnline 是否即将上线
func (c *Client) IsComingOnline() bool {
	currentTime := time.Now().UnixMilli()
	return c.OnlineAt > currentTime && (c.OfflineAt == -1 || c.OfflineAt < currentTime)
}

// IsComingOffline 是否即将下线
func (c *Client) IsComingOffline() bool {
	currentTime := time.Now().UnixMilli()
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
	if c.Extra["website"] == nil {
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
	if c.Extra["copyrights"] == nil {
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
	if c.Extra["supportUrl"] == nil {
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
	if c.Extra["privacyUrl"] == nil {
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
	if c.Extra["bulletins"] == nil {
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
	if c.Extra["userMaxAccount"] == nil {
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
	if c.Extra["userMaxToken"] == nil {
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

func (c *Client) GetPlatform(platform, area uint16) *ClientPlatform {
	if _, ok := c.Platforms[platform]; !ok {
		return nil
	}
	return c.Platforms[platform][area]
}

func (c *Client) GetLatestCode(platform, area, market uint16) *ClientVersion {
	if _, ok := c.LatestCodes[platform]; !ok {
		return nil
	}
	if _, ok := c.LatestCodes[platform][area]; !ok {
		return nil
	}
	return c.LatestCodes[platform][area][market]
}

const (
	checkClientIPNameLen       = 100
	checkClientOrganizationLen = 100

	checkClientWebsiteLen    = 500
	checkClientCopyrightsNum = 50
	checkClientCopyrightLen  = 100
	checkClientSupportUrlLen = 500
	checkClientPrivacyUrlLen = 500
	checkClientBulletinsNum  = 10000
	checkClientBulletinLen   = 50000
)

// CheckFields 检查字段
func (c *Client) CheckFields() []*tools.CodeError {
	var errors []*tools.CodeError
	if len(c.IPName) <= 0 {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldNil).WithPrefix("IPName"))
	} else if len(c.IPName) > checkClientIPNameLen {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("IPName"))
	}
	if len(c.Organization) <= 0 {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldNil).WithPrefix("Organization"))
	} else if len(c.Organization) > checkClientOrganizationLen {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("Organization"))
	}
	for k, v := range c.Extra {
		switch k {
		case "website":
			if len(v.(string)) > checkClientWebsiteLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("website"))
			}
		case "copyrights":
			if len(v.([]string)) > checkClientCopyrightsNum {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldMax).WithPrefix("copyrights"))
			}
			for kk, vv := range v.([]string) {
				if len(vv) > checkClientCopyrightLen {
					errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix(fmt.Sprintf("copyrights[%d] ", kk)))
				}
			}
		case "supportUrl":
			if len(v.(string)) > checkClientSupportUrlLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("supportUrl"))
			}
		case "privacyUrl":
			if len(v.(string)) > checkClientPrivacyUrlLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("privacyUrl"))
			}
		case "bulletins":
			if len(v.([]string)) > checkClientBulletinsNum {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldMax).WithPrefix("bulletins"))
			}
			for kk, vv := range v.([]string) {
				if len(vv) > checkClientBulletinLen {
					errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix(fmt.Sprintf("bulletins[%d] ", kk)))
				}
			}
		case "userMaxAccount":
		case "userMaxToken":
			continue
		default:
			errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldUnDef).WithPrefix(k))
		}
	}
	return errors
}
