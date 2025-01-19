package model

const (
	ClientStatsKindWatched  = 1 // 查看 (web端的查看，和应用市场的详情页进入)
	ClientStatsKindDownload = 2 // 下载 (小程序的话，关注也是这个)
	ClientStatsKindStart    = 3 // 首次启动
	ClientStatsKindOpener   = 4 // 启动(包括首次)
	ClientStatsKindOpenerU  = 5 // 启动(包括首次+去重)
	ClientStatsKindRegister = 6 // 首次注册
	ClientStatsKindActive   = 7 // 登录(包括首次)
	ClientStatsKindActiveU  = 8 // 登录(包括首次+去重)
	ClientStatsKindScore    = 9 // 评分 (评论放链接里点击)

	ClientStatsKindDuration0   = 10 // 在线时长 (min)
	ClientStatsKindDuration3   = 11 // 在线时长 (min)
	ClientStatsKindDuration10  = 12 // 在线时长 (min)
	ClientStatsKindDuration30  = 13 // 在线时长 (min)
	ClientStatsKindDuration60  = 14 // 在线时长 (min)
	ClientStatsKindDuration180 = 15 // 在线时长 (min)
	ClientStatsKindReturn7     = 16 // 回归周期 (day)
	ClientStatsKindReturn15    = 17 // 回归周期 (day)
	ClientStatsKindReturn30    = 18 // 回归周期 (day)
	ClientStatsKindReturn90    = 19 // 回归周期 (day)

	ClientStatsKindPayer       = 20 // 付费用户数(去重)
	ClientStatsKindPayFirst    = 21 // 首充次数
	ClientStatsKindPayCount    = 22 // 付费次数(包括首充)
	ClientStatsKindPayMoney1   = 23 // 付费金额 (元)
	ClientStatsKindPayMoney6   = 24 // 付费金额 (元)
	ClientStatsKindPayMoney30  = 25 // 付费金额 (元)
	ClientStatsKindPayMoney68  = 26 // 付费金额 (元)
	ClientStatsKindPayMoney128 = 27 // 付费金额 (元)
	ClientStatsKindPayMoney328 = 28 // 付费金额 (元)
	ClientStatsKindPayMoney648 = 29 // 付费金额 (元)

	// TODO:GG 还有很多可以加？都要放进来吗?
)

// ClientStats 客户端统计量
type ClientStats struct {
	*Base
	IP       int `json:"IP"`       // 系列 (eg:大富翁IP)
	Part     int `json:"part"`     // 类型 (eg:单机版)
	Area     int `json:"area"`     // 区域编号 (从conf读取，不一定是国家或洲际)
	Platform int `json:"platform"` // 平台
	Market   int `json:"market"`   // 渠道
	Code     int `json:"code"`     // 版本标识
	Kind     int `json:"kind"`     // 类型
	Year     int `json:"year"`     // 年
	Month    int `json:"month"`    // 月
	Week     int `json:"week"`     // 周 (和月冲突)
	Day      int `json:"day"`      // 日
	Hour     int `json:"hour"`     // 时

	Num int `json:"num"` // 数量
}

func NewClientStats(
	base *Base,
	IP int, part int, area int, platform int, market int, code int,
	kind int,
	year int, month int, week int, day int, hour int,
	num int,
) *ClientStats {
	return &ClientStats{
		Base: base,
		IP:   IP, Part: part, Area: area, Platform: platform, Market: market, Code: code,
		Kind: kind,
		Year: year, Month: month, Week: week, Day: day, Hour: hour,
		Num: num,
	}
}
