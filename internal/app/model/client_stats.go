package model

// ClientStats 客户端统计量
type ClientStats struct {
	*Base
	Cid   int64 `json:"cid"`   // 客户端id
	CVid  int64 `json:"cvid"`  // 客户端版本id
	Kind  int   `json:"kind"`  // 类型 (1:查看Watched 2:下载Download 3:启动Opener 4:评分Score)
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
