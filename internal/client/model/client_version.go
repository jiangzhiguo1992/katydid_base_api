package model

import (
	"fmt"
	"katydid_base_api/internal/pkg/database"
	"katydid_base_api/internal/pkg/utils"
	"katydid_base_api/tools"
	"time"
)

// ClientVersion 客户端版本
type ClientVersion struct {
	*database.BaseModel
	CPid   uint64 `json:"cpid"`   // 客户端平台id
	Market uint   `json:"market"` // 市场/渠道/投放
	Code   uint   `json:"code"`   // 版本标识

	Enable    bool  `json:"enable"`    // 是否可用 (是否可下载+可使用+是否对用户可见?)
	BuildAt   int64 `json:"buildAt"`   // 构建时间 (一般是上传时间)
	PublishAt int64 `json:"publishAt"` // 发布时间 (审核通过后发布时间，是否对用户可见?)

	Url    string `json:"url"`    // 升级地址 (安装包地址，或market跳转地址)
	Force  bool   `json:"force"`  // 是否强制升级
	AppKey string `json:"appKey"` // app密钥 (终端使用，版本更替，确定后不可改) TODO:GG 不会返回给客户端，在网关/代理层处理?

	Extra map[string]interface{} `json:"extra" gorm:"serializer:json"` // 额外信息
}

func NewClientVersionDefault(
	CPid uint64, market uint, code uint,
	enable bool,
	url string, force bool, appKey string,
) *ClientVersion {
	return &ClientVersion{
		BaseModel: database.NewBaseModelEmpty(),
		CPid:      CPid, Market: market, Code: code,
		Enable: enable, BuildAt: -1, PublishAt: -1,
		Url: url, Force: force, AppKey: appKey,
		Extra: map[string]interface{}{},
	}
}

func (c *ClientVersion) IsBuild() bool {
	return c.BuildAt > time.Now().UnixMilli()
}

func (c *ClientVersion) IsPublish() bool {
	return c.IsBuild() && (c.PublishAt > time.Now().UnixMilli()) && (len(c.Url) > 0)
}

// SetName 版本名
func (c *ClientVersion) SetName(name *string) {
	if (name != nil) && (len(*name) > 0) {
		c.Extra["name"] = *name
	} else {
		delete(c.Extra, "name")
	}
}

func (c *ClientVersion) GetName() string {
	if c.Extra["name"] == nil {
		return ""
	}
	return c.Extra["name"].(string)
}

// SetSize 安装包大小 (上传pkg的时候统计)
func (c *ClientVersion) SetSize(size *uint64) {
	if (size != nil) && (*size >= 0) {
		c.Extra["size"] = *size
	} else {
		delete(c.Extra, "size")
	}
}

func (c *ClientVersion) GetSize() uint64 {
	if c.Extra["size"] == nil {
		return 0
	}
	return c.Extra["size"].(uint64)
}

// SetIconUrl app图标 (节日/活动/...)
func (c *ClientVersion) SetIconUrl(iconUrl *string) {
	if (iconUrl != nil) && (len(*iconUrl) > 0) {
		c.Extra["iconUrl"] = *iconUrl
	} else {
		delete(c.Extra, "iconUrl")
	}
}

func (c *ClientVersion) GetIconUrl() string {
	if c.Extra["iconUrl"] == nil {
		return ""
	}
	return c.Extra["iconUrl"].(string)
}

// SetCompact app兼容性 (eg: 9.0+)
func (c *ClientVersion) SetCompact(compact *string) {
	if (compact != nil) && (len(*compact) > 0) {
		c.Extra["compact"] = *compact
	} else {
		delete(c.Extra, "compact")
	}
}

func (c *ClientVersion) GetCompact() string {
	if c.Extra["compact"] == nil {
		return ""
	}
	return c.Extra["compact"].(string)
}

// SetLog 升级日志
func (c *ClientVersion) SetLog(log *string) {
	if (log != nil) && (len(*log) > 0) {
		c.Extra["log"] = *log
	} else {
		delete(c.Extra, "log")
	}
}

func (c *ClientVersion) GetLog() string {
	if c.Extra["log"] == nil {
		return ""
	}
	return c.Extra["log"].(string)
}

// SetImgUrls 介绍图片Url
func (c *ClientVersion) SetImgUrls(imgUrls *[]string) {
	if (imgUrls != nil) && (len(*imgUrls) > 0) {
		c.Extra["imgUrls"] = *imgUrls
	} else {
		delete(c.Extra, "imgUrls")
	}
}

func (c *ClientVersion) GetImgUrls() []string {
	if c.Extra["imgUrls"] == nil {
		return []string{}
	}
	return c.Extra["imgUrls"].([]string)
}

// SetVideoUrls 介绍视频Url
func (c *ClientVersion) SetVideoUrls(videoUrls *[]string) {
	if (videoUrls != nil) && (len(*videoUrls) > 0) {
		c.Extra["videoUrls"] = *videoUrls
	} else {
		delete(c.Extra, "videoUrls")
	}
}

func (c *ClientVersion) GetVideoUrls() []string {
	if c.Extra["videoUrls"] == nil {
		return []string{}
	}
	return c.Extra["videoUrls"].([]string)
}

// SetMarketName 设置广告渠道名称 (广告投放)
func (c *ClientVersion) SetMarketName(name string) {
	if c.Market < MarketTypeAdsMin {
		return
	}
	c.Extra["marketName"] = name
}

func (c *ClientVersion) GetMarketName(platform uint16) string {
	if c.Market < MarketTypeAdsMin {
		return platformMarketName(platform, c.Market)
	}
	if c.Extra["marketName"] == nil {
		return ""
	}
	return c.Extra["marketName"].(string)
}

const (
	checkClientVersionUrlLen    = 500
	checkClientVersionAppKeyLen = 100

	checkClientVersionNameLen      = 100
	checkClientVersionIconUrlLen   = 500
	checkClientVersionCompactLen   = 100
	checkClientVersionLogLen       = 10000
	checkClientVersionImgUrlsNum   = 50
	checkClientVersionImgUrlLen    = 500
	checkClientVersionVideoUrlsNum = 50
	checkClientVersionVideoUrlLen  = 500
	checkClientVersionIosIdLen     = 50
)

// CheckFields 检查字段
func (c *ClientVersion) CheckFields() []*tools.CodeError {
	var errors []*tools.CodeError
	if !isMarketTypeOk(c.Market) {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldRange).WithPrefix("Market "))
	}
	if len(c.Url) <= 0 {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldNil).WithPrefix("Url "))
	} else if len(c.Url) > checkClientVersionUrlLen {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("Url "))
	}
	if len(c.AppKey) <= 0 {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldNil).WithPrefix("AppKey "))
	} else if len(c.AppKey) > checkClientVersionAppKeyLen {
		errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("AppKey "))
	}
	for k, v := range c.Extra {
		switch k {
		case "name":
			if len(v.(string)) > checkClientVersionNameLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("name "))
			}
		case "iconUrl":
			if len(v.(string)) > checkClientVersionIconUrlLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("iconUrl "))
			}
		case "compact":
			if len(v.(string)) > checkClientVersionCompactLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("compact "))
			}
		case "log":
			if len(v.(string)) > checkClientVersionLogLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("log "))
			}
		case "imgUrls":
			if len(v.([]string)) > checkClientVersionImgUrlsNum {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldMax).WithPrefix("imgUrls "))
			}
			for kk, vv := range v.([]string) {
				if len(vv) > checkClientVersionImgUrlLen {
					errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix(fmt.Sprintf("imgUrls[%d] ", kk)))
				}
			}
		case "videoUrls":
			if len(v.([]string)) > checkClientVersionVideoUrlsNum {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldMax).WithPrefix("videoUrls "))
			}
			for kk, vv := range v.([]string) {
				if len(vv) > checkClientVersionVideoUrlLen {
					errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix(fmt.Sprintf("videoUrls[%d] ", kk)))
				}
			}
		case "iosId":
			if len(v.(string)) > checkClientVersionIosIdLen {
				errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldLarge).WithPrefix("iosId "))
			}
		default:
			errors = append(errors, utils.MatchErrorByCode(utils.ErrorCodeDBFieldUnDef).WithPrefix(k+" "))
		}
	}
	return errors
}

const (
	MarketTypeLinuxOfficial uint = 1
	MarketTypeLinuxSteam    uint = 2

	MarketTypeWindowsOfficial uint = 1
	MarketTypeWindowsSteam    uint = 2

	MarketTypeMacOsOfficial uint = 1
	MarketTypeMacOsAppStore uint = 2
	MarketTypeMacOsSteam    uint = 3

	MarketTypeWebOfficial uint = 1
	MarketTypeWebChrome   uint = 2
	MarketTypeWebSafari   uint = 3
	MarketTypeWebFirefox  uint = 4
	MarketTypeWebEdge     uint = 5
	MarketTypeWebOpera    uint = 6
	MarketTypeWebIE       uint = 7
	MarketTypeWeb360      uint = 8
	MarketTypeWebQQ       uint = 9
	MarketTypeWebHuoHu    uint = 10
	MarketTypeWebLieBao   uint = 11
	MarketTypeWebSouGou   uint = 12

	MarketTypeAndroidOfficial   uint = 1
	MarketTypeAndroidGooglePlay uint = 2
	MarketTypeAndroidTapTap     uint = 3
	MarketTypeAndroidHuawei     uint = 4
	MarketTypeAndroidXiaomi     uint = 5
	MarketTypeAndroidOppo       uint = 6
	MarketTypeAndroidVivo       uint = 7
	MarketTypeAndroidMeizu      uint = 8
	MarketTypeAndroidOnePlus    uint = 9
	MarketTypeAndroidSamsung    uint = 10
	MarketTypeAndroidLenovo     uint = 11
	MarketTypeAndroidSony       uint = 12
	MarketTypeAndroidLG         uint = 13
	MarketTypeAndroidHTC        uint = 14
	MarketTypeAndroidMotorola   uint = 15
	MarketTypeAndroidNokia      uint = 16
	MarketTypeAndroidTencent    uint = 30
	MarketTypeAndroidBaidu      uint = 31
	MarketTypeAndroid360        uint = 32

	MarketTypeIOSOfficial uint = 1
	MarketTypeIOSAppStore uint = 2

	MarketTypeAppletOfficial    uint = 1
	MarketTypeAppletWeChat      uint = 2
	MarketTypeAppletQQ          uint = 3
	MarketTypeAppletDouYin      uint = 4
	MarketTypeAppletKuaiShou    uint = 5
	MarketTypeAppletXiaoHongShu uint = 6
	MarketTypeAppletBaidu       uint = 7
	MarketTypeAppletZhiFuBao    uint = 8
	MarketTypeAppletTaoBao      uint = 9
	MarketTypeAppletJingDong    uint = 10
	MarketTypeAppletDingDing    uint = 11

	//ItchIo     uint = 2
	//KongReGate uint = 3
	//IndieDb    uint = 4

	MarketTypeAdsMin uint = 1000 // 大于这个的都是投量广告
)

var platformMarketInfos = map[uint16]map[uint]string{
	PlatformTypeLinux: {
		MarketTypeLinuxOfficial: "Linux_官网",
		MarketTypeLinuxSteam:    "Linux_Steam",
	},
	PlatformTypeWindows: {
		MarketTypeWindowsOfficial: "Windows_官网",
		MarketTypeWindowsSteam:    "Windows_Steam",
	},
	PlatformTypeMacOS: {
		MarketTypeMacOsOfficial: "MacOS_官网",
		MarketTypeMacOsAppStore: "MacOS_应用商店",
		MarketTypeMacOsSteam:    "MacOS_Steam",
	},
	PlatformTypeWeb: {
		MarketTypeWebOfficial: "Web_官网",
		MarketTypeWebChrome:   "Web_Chrome",
		MarketTypeWebSafari:   "Web_Safari",
		MarketTypeWebFirefox:  "Web_Firefox",
		MarketTypeWebEdge:     "Web_Edge",
		MarketTypeWebOpera:    "Web_Opera",
		MarketTypeWebIE:       "Web_IE",
		MarketTypeWeb360:      "Web_360",
		MarketTypeWebQQ:       "Web_QQ",
		MarketTypeWebHuoHu:    "Web_火狐",
		MarketTypeWebLieBao:   "Web_猎豹",
		MarketTypeWebSouGou:   "Web_搜狗",
	},
	PlatformTypeAndroid: {
		MarketTypeAndroidOfficial:   "Android_官网",
		MarketTypeAndroidGooglePlay: "Android_谷歌",
		MarketTypeAndroidTapTap:     "Android_TapTap",
		MarketTypeAndroidHuawei:     "Android_华为",
		MarketTypeAndroidXiaomi:     "Android_小米",
		MarketTypeAndroidOppo:       "Android_oppo",
		MarketTypeAndroidVivo:       "Android_vivo",
		MarketTypeAndroidMeizu:      "Android_魅族",
		MarketTypeAndroidOnePlus:    "Android_一加",
		MarketTypeAndroidSamsung:    "Android_三星",
		MarketTypeAndroidLenovo:     "Android_联想",
		MarketTypeAndroidSony:       "Android_索尼",
		MarketTypeAndroidLG:         "Android_LG",
		MarketTypeAndroidHTC:        "Android_HTC",
		MarketTypeAndroidMotorola:   "Android_摩托罗拉",
		MarketTypeAndroidNokia:      "Android_诺基亚",
		MarketTypeAndroidTencent:    "Android_腾讯",
		MarketTypeAndroidBaidu:      "Android_百度",
		MarketTypeAndroid360:        "Android_360",
	},
	PlatformTypeIOS: {
		MarketTypeIOSOfficial: "IOS_官网",
		MarketTypeIOSAppStore: "IOS_AppStore",
	},
	PlatformTypeApplet: {
		MarketTypeAppletOfficial:    "Applet_官网",
		MarketTypeAppletWeChat:      "Applet_微信",
		MarketTypeAppletQQ:          "Applet_QQ",
		MarketTypeAppletDouYin:      "Applet_抖音",
		MarketTypeAppletKuaiShou:    "Applet_快手",
		MarketTypeAppletXiaoHongShu: "Applet_小红书",
		MarketTypeAppletBaidu:       "Applet_百度",
		MarketTypeAppletZhiFuBao:    "Applet_支付宝",
		MarketTypeAppletTaoBao:      "Applet_淘宝",
		MarketTypeAppletJingDong:    "Applet_京东",
		MarketTypeAppletDingDing:    "Applet_钉钉",
	},
}

func isMarketTypeOk(market uint) bool {
	if market < MarketTypeAdsMin {
		for _, platforms := range platformMarketInfos {
			if _, ok := platforms[market]; ok {
				return true
			}
		}
	}
	return true
}

func isPlatformMarketTypeOk(platform uint16, market uint) bool {
	if _, ok := platformMarketInfos[platform]; !ok {
		return false
	}
	if market < MarketTypeAdsMin {
		if _, ok := platformMarketInfos[platform][market]; !ok {
			return false
		}
	}
	return true
}

func platformMarketName(platform uint16, market uint) string {
	if _, ok := platformMarketInfos[platform]; !ok {
		return ""
	}
	if market < MarketTypeAdsMin {
		if _, ok := platformMarketInfos[platform][market]; !ok {
			return ""
		}
		return platformMarketInfos[platform][market]
	}
	return ""
}
