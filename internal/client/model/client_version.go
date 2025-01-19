package model

import (
	"time"
)

const (
	MarketTypeGooglePlay = 1
	MarketTypeAppStore   = 2
	MarketTypeHuawei     = 3
	MarketTypeXiaomi     = 4
	MarketTypeOppo       = 5
	MarketTypeVivo       = 6
	MarketTypeTencent    = 7
	MarketTypeBaidu      = 8
	MarketType360        = 9
	MarketTypeMeizu      = 10
	MarketTypeCoolpad    = 11
	MarketTypeLenovo     = 12
	MarketTypeZTE        = 13
	MarketTypeSamsung    = 14
	MarketTypeSony       = 15
	MarketTypeLG         = 16
	MarketTypeHTC        = 17
	MarketTypeNokia      = 18
	MarketTypeMotorola   = 19
	MarketTypeOnePlus    = 20
	// TODO:GG 还会有渠道投放，所以可以从conf读取？
)

// ClientVersion 客户端版本
type ClientVersion struct {
	*Base
	CPid   int64 `json:"cpid"`   // 客户端平台id
	Market int   `json:"market"` // 市场/渠道 (从conf读)
	Code   int   `json:"code"`   // 版本标识

	AppName   string `json:"appName"`   // app名称
	AppIcon   string `json:"appIcon"`   // app图标
	AppCompat string `json:"appCompat"` // app兼容性 (eg: 9.0+)
	AppKey    string `json:"appKey"`    // app密钥 (终端使用)

	Name  string   `json:"name"`  // 版本名
	Size  int64    `json:"size"`  // 安装包大小 (上传pkg的时候统计)
	Log   string   `json:"log"`   // 升级日志
	Imgs  []string `json:"imgs"`  // 介绍图片Url
	Url   string   `json:"url"`   // 升级地址
	Force bool     `json:"force"` // 是否强制升级

	BuildAt   int64 `json:"buildAt"`   // 构建时间 (不是发布时间)
	PublishAt int64 `json:"publishAt"` // 发布时间 (不是审核时间)
	Enable    bool  `json:"enable"`    // 是否可用 (没有reason)

	Extra map[string]interface{} `json:"extra"` // 额外信息
}

func NewClientVersion(
	base *Base,
	CPid int64, market int, code int,
	appName string, appIcon string, appCompat string, appKey string,
	name string, size int64, log string, imgs []string, url string, force bool,
	buildAt int64, publishAt int64, enable bool,
	extra map[string]interface{},
) *ClientVersion {
	return &ClientVersion{
		Base: base,
		CPid: CPid, Market: market, Code: code,
		AppName: appName, AppIcon: appIcon, AppCompat: appCompat, AppKey: appKey,
		Name: name, Size: size, Log: log, Imgs: imgs, Url: url, Force: force,
		BuildAt: buildAt, PublishAt: publishAt, Enable: enable,
		Extra: extra,
	}
}

func (c *ClientVersion) IsBuild() bool {
	return (c.BuildAt > time.Now().Unix()) && (c.Size > 0)
}

func (c *ClientVersion) IsPublish() bool {
	return c.PublishAt > time.Now().Unix()
}

func (c *ClientVersion) SetIosId(iosId string) {
	if c.Extra == nil {
		c.Extra = map[string]interface{}{}
	}
	c.Extra["ios_id"] = iosId
}

func (c *ClientVersion) GetIosId() string {
	if c.Extra == nil || c.Extra["ios_id"] == nil {
		return ""
	}
	return c.Extra["ios_id"].(string)
}
