package model

// ClientVersion 客户端版本
type ClientVersion struct {
	*Base
	Area     int    `json:"area"`     // 区域编号 (从conf读取，不一定是国家或洲际)
	Market   int    `json:"market"`   // 渠道
	Cid      int64  `json:"cid"`      // 客户端id (多对一)
	Cname    string `json:"cname"`    // 客户端名称 (可能需要替换Client.Name)
	CIconUrl string `json:"cIconUrl"` // 客户端图标
	CDesc    string `json:"cDesc"`    // 客户端描述
	CKind    string `json:"cKind"`    // 客户端类型 (Market里的分类)

	PageUrl string   `json:"pageUrl"` // 产品页面 (方便控制台跳转)
	Code    int      `json:"code"`    // 版本标识
	Name    string   `json:"name"`    // 版本名
	Log     string   `json:"log"`     // 升级日志
	ImgUrls []string `json:"imgUrls"` // 介绍图片
	Url     string   `json:"url"`     // 升级地址
	Force   bool     `json:"force"`   // 是否强制升级
	Enable  bool     `json:"enable"`  // 是否可用 (和Client.enable一样但不冲突)

	PublishAt int64 `json:"publishAt"` // 发布时间 (不是审核时间)
	Watched   int   `json:"watched"`   // 总查看数量 (整点更新)
	Download  int   `json:"download"`  // 总下载数量 (整点更新)
	Opener    int   `json:"opener"`    // 总启动数量 (整点更新)
	Score     int   `json:"score"`     // 当前总评分 (整点更新)
	Comments  int   `json:"comments"`  // 当前总评数 (整点更新)
}

func NewClientVersion(
	base *Base,
	area int,
	market int,
	cid int64,
	cname string,
	ciconUrl string,
	cdesc string,
	ckind string,
	pageUrl string,
	code int,
	name string,
	log string,
	imgUrls []string,
	url string,
	force bool,
	enable bool,
) *ClientVersion {
	return &ClientVersion{
		Base:      base,
		Area:      area,
		Market:    market,
		Cid:       cid,
		Cname:     cname,
		CIconUrl:  ciconUrl,
		CDesc:     cdesc,
		CKind:     ckind,
		PageUrl:   pageUrl,
		Code:      code,
		Name:      name,
		Log:       log,
		ImgUrls:   imgUrls,
		Url:       url,
		Force:     force,
		Enable:    enable,
		PublishAt: -1,
		Watched:   -1,
		Download:  -1,
		Opener:    -1,
		Score:     -1,
		Comments:  -1,
	}
}
