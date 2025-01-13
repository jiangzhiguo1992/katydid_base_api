package model

import "time"

// ClientVersion 客户端版本
type ClientVersion struct {
	*Base
	Cid      int64 `json:"cid"`      // 客户端id (多对一)
	Area     int   `json:"area"`     // 区域编号 (从conf读取，不一定是国家或洲际)
	Platform int   `json:"platform"` // 平台 (Android/IOS/Web/PC/MACOS/WeChat/DouYin/TaoBao/...)
	Market   int   `json:"market"`   // 渠道
	Code     int   `json:"code"`     // 版本标识

	Name    string   `json:"name"`    // 版本名
	Log     string   `json:"log"`     // 升级日志
	ImgUrls []string `json:"imgUrls"` // 介绍图片
	Url     string   `json:"url"`     // 升级地址
	Force   bool     `json:"force"`   // 是否强制升级
	PageUrl string   `json:"pageUrl"` // 产品页面 (方便控制台跳转)
	Enable  bool     `json:"enable"`  // 是否可用 (和Client.enable一样但不冲突)

	PublishAt int64       `json:"publishAt"` // 发布时间 (不是审核时间)
	StatsList map[int]int `json:"statsList"` // [stats_kind]统计数据 (整点更新)
}

func NewClientVersion(
	base *Base,
	cid int64,
	area int,
	platform int,
	market int,
	code int,
	name string,
	log string,
	imgUrls []string,
	url string,
	force bool,
	pageUrl string,
	enable bool,
) *ClientVersion {
	return &ClientVersion{
		Base:      base,
		Cid:       cid,
		Area:      area,
		Platform:  platform,
		Market:    market,
		Code:      code,
		Name:      name,
		Log:       log,
		ImgUrls:   imgUrls,
		Url:       url,
		Force:     force,
		PageUrl:   pageUrl,
		Enable:    enable,
		PublishAt: -1,
		StatsList: map[int]int{},
	}
}

func (c *ClientVersion) IsPublished() bool {
	return c.PublishAt > time.Now().Unix()
}
