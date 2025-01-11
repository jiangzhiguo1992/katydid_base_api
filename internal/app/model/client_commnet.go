package model

// ClientComment 客户端评论
type ClientComment struct {
	Cid       int64    `json:"cid"`      // 客户端id
	CVid      int64    `json:"cvid"`     // 客户端版本id
	Timestamp int64    `json:"time"`     // 时间戳
	Username  string   `json:"username"` // 用户名称
	Title     string   `json:"title"`    // 标题
	Tags      []string `json:"tags"`     // 标签
	Body      string   `json:"body"`     // 内容
	ImgUrls   []string `json:"imgUrls"`  // 介绍图片
	Points    int      `json:"points"`   // 点赞数
	Score     float64  `json:"score"`    // 评分 (不影响ClientScore，只是展示，-1为无)
}

func NewClientComment(
	cid int64,
	cvid int64,
	timestamp int64,
	username string,
	title string,
	tags []string,
	body string,
	imgUrls []string,
	points int,
	score float64,
) *ClientComment {
	return &ClientComment{
		Cid:       cid,
		CVid:      cvid,
		Timestamp: timestamp,
		Username:  username,
		Title:     title,
		Tags:      tags,
		Body:      body,
		ImgUrls:   imgUrls,
		Points:    points,
		Score:     score,
	}
}
