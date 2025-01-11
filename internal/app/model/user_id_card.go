package model

import "time"

// UserIDCard 用户身份证/护照 (UIndex = Area + Number)
type UserIDCard struct {
	*Base
	Uid    int64  `json:"uid"`    // 用户Id (一对多，双国籍?)
	Area   int    `json:"area"`   // 区域 (区域政策不一致)
	Number string `json:"number"` // 身份证号 (唯一标识)

	Name     string                 `json:"name"`     // 姓名
	Sex      int                    `json:"sex"`      // 性别 (女-1，男+1)
	Nation   int                    `json:"nation"`   // 民族
	Year     int                    `json:"year"`     // 出生年
	Month    int                    `json:"month"`    // 出生月
	Day      int                    `json:"day"`      // 出生日
	Address  string                 `json:"address"`  // 住址
	Period   int64                  `json:"period"`   // 有效期
	FrontUrl string                 `json:"frontUrl"` // 正面照片
	BackUrl  string                 `json:"backUrl"`  // 背面照片
	Extra    map[string]interface{} `json:"extra"`    // 额外信息

	BioAuth *UserBioAuth `json:"bioAuth"` // 生物验证
}

func NewUserIDCard(
	base *Base,
	uid int64,
	area int,
	number string,
	name string,
	sex int,
	nation int,
	year int,
	month int,
	day int,
	address string,
	period int64,
	frontUrl string,
	backUrl string,
	extra map[string]interface{},
) *UserIDCard {
	return &UserIDCard{
		Base:     base,
		Uid:      uid,
		Area:     area,
		Number:   number,
		Name:     name,
		Sex:      sex,
		Nation:   nation,
		Year:     year,
		Month:    month,
		Day:      day,
		Address:  address,
		Period:   period,
		FrontUrl: frontUrl,
		BackUrl:  backUrl,
		Extra:    extra,
	}
}

func (c UserIDCard) IsAdults() bool {
	currentYear, currentMonth, currentDay := time.Now().Date()
	age := currentYear - c.Year
	if currentMonth < time.Month(c.Month) || (currentMonth == time.Month(c.Month) && currentDay < c.Day) {
		age--
	}
	return age >= 18
}
