package model

const (
	ClientStatsKindWatched  = 1 // 查看Watched
	ClientStatsKindDownload = 2 // 下载Download
	ClientStatsKindOpener   = 3 // 启动Opener
	ClientStatsKindScore    = 4 // 评分Score
)

// ClientStats 客户端统计量 (UIndex = 除Num外的所有字段)
type ClientStats struct {
	*Base
	Cid   int64 `json:"cid"`   // 客户端id
	CVid  int64 `json:"cvid"`  // 客户端版本id
	Kind  int   `json:"kind"`  // 类型
	Year  int   `json:"year"`  // 年
	Month int   `json:"month"` // 月
	Day   int   `json:"day"`   // 日
	Hour  int   `json:"hour"`  // 时
	Num   int   `json:"num"`   // 数量
}

func NewClientStats(
	base *Base,
	cid int64,
	cvid int64,
	kind int,
	year int,
	month int,
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
		Day:   day,
		Hour:  hour,
		Num:   num,
	}
}
