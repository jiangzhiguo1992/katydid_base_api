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

	ClientStatsKindDuration1   = 10 // 在线时长 (min)
	ClientStatsKindDuration5   = 11 // 在线时长 (min)
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
)

// ClientStats 客户端统计量
type ClientStats struct {
	*Base
	Cid   int64 `json:"cid"`   // 客户端id
	CVid  int64 `json:"cvid"`  // 客户端版本id
	Kind  int   `json:"kind"`  // 类型
	Year  int   `json:"year"`  // 年
	Month int   `json:"month"` // 月
	Week  int   `json:"week"`  // 周 (和月冲突)
	Day   int   `json:"day"`   // 日
	Hour  int   `json:"hour"`  // 时

	Num int `json:"num"` // 数量
}

func NewClientStats(
	base *Base,
	cid int64,
	cvid int64,
	kind int,
	year int,
	month int,
	week int,
	day int,
	hour int,
	num int,
) *ClientStats {
	return &ClientStats{
		Base:  base,
		Cid:   cid,
		CVid:  cvid,
		Kind:  kind,
		Year:  year,
		Month: month,
		Week:  week,
		Day:   day,
		Hour:  hour,
		Num:   num,
	}
}
