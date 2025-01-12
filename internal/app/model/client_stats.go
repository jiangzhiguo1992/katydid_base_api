package model

const (
	ClientStatsKindWatched  = 1 // 查看Watched (web端的查看，和应用市场的详情页进入)
	ClientStatsKindDownload = 2 // 下载Download (小程序的话，关注也是这个)
	ClientStatsKindOpener   = 3 // 启动Opener
	ClientStatsKindRegister = 4 // 注册Register (转化etc)
	ClientStatsKindActive   = 5 // 激活Active (日/周/月/年活, 留存/回归率)
	ClientStatsKindReturn   = 5 // 回归Active (日/周/月/年活, 留存/回归率)
	ClientStatsKindPay      = 6 // 付费数Payer (日/周/月/年付费, 会员率/续费率/转化率/付费率/平均每付费用户收入/每用户平均收入)
	ClientStatsKindScore    = 7 // 评分Score
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
