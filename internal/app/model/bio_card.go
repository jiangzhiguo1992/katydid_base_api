package model

// TODO:GG PGSQL <- BioCard = Users
// TODO:GG PGSQL <- BioAuth = Users

// IBioCard 生物卡接口 (身份证/护照/...) (UIndex = Area + Number)
type IBioCard interface {
	IBase
	GetArea() int      // 区域 (区域政策不一致)
	GetNumber() string // 唯一标识

	GetName() string     // 姓名 (FirstName + LastName)
	GetSex() int         // 性别
	GetYear() int        // 出生年
	GetMonth() int       // 出生月
	GetDay() int         // 出生日
	GetAddress() string  // 住址
	GetFrontUrl() string // 身份证正面
	GetBackUrl() string  // 身份证背面

	IsAdults() bool                  // 是否成年 (和Enable是独立开的)
	IsEnable(clientId int64) bool    // 是否可用 -1为所有都不可用
	GetReason(clientId int64) string // 拒绝原因
}

// ChinaMainlandBioCard 中国大陆身份证
type ChinaMainlandBioCard struct {
	*Base
	Area   int    `json:"area"`   // 区域 (区域政策不一致)
	Number string `json:"number"` // 身份证号

	Name     string `json:"name"`     // 姓名
	Sex      int    `json:"sex"`      // 性别 (女-1，男+1)
	Year     int    `json:"year"`     // 出生年
	Month    int    `json:"month"`    // 出生月
	Day      int    `json:"day"`      // 出生日
	Address  string `json:"address"`  // 住址
	FrontUrl string `json:"frontUrl"` // 身份证正面
	BackUrl  string `json:"backUrl"`  // 身份证背面

	Nation string `json:"nation"` // 民族
	Period string `json:"period"` // 身份证有效期

	Accounts map[int64]int    `json:"accounts"` // 账号数
	Enable   map[int64]bool   `json:"enable"`   // 是否可用
	Reason   map[int64]string `json:"reason"`   // 拒绝原因
}

func NewChinaMainlandBioCard(
	base *Base,
	area int,
	number string,
	name string,
	sex int,
	year int,
	month int,
	day int,
	address string,
	frontUrl string,
	backUrl string,
	nation string,
	period string,
) *ChinaMainlandBioCard {
	return &ChinaMainlandBioCard{
		Base:     base,
		Area:     area,
		Number:   number,
		Name:     name,
		Sex:      sex,
		Year:     year,
		Month:    month,
		Day:      day,
		Address:  address,
		FrontUrl: frontUrl,
		BackUrl:  backUrl,
		Nation:   nation,
		Period:   period,
		Accounts: map[int64]int{},
		Enable:   map[int64]bool{},
		Reason:   map[int64]string{},
	}
}

func (c ChinaMainlandBioCard) GetId() int64 {
	return c.Id
}

func (c ChinaMainlandBioCard) GetCreateAt() int64 {
	return c.CreateAt
}

func (c ChinaMainlandBioCard) GetUpdateAt() int64 {
	return c.UpdateAt
}

func (c ChinaMainlandBioCard) GetArea() int {
	return c.Area
}

func (c ChinaMainlandBioCard) GetNumber() string {
	return c.Number
}

func (c ChinaMainlandBioCard) GetName() string {
	return c.Name
}

func (c ChinaMainlandBioCard) GetSex() int {
	return c.Sex
}

func (c ChinaMainlandBioCard) GetYear() int {
	return c.Year
}

func (c ChinaMainlandBioCard) GetMonth() int {
	return c.Month
}

func (c ChinaMainlandBioCard) GetDay() int {
	return c.Day
}

func (c ChinaMainlandBioCard) GetAddress() string {
	return c.Address
}

func (c ChinaMainlandBioCard) GetFrontUrl() string {
	return c.FrontUrl
}

func (c ChinaMainlandBioCard) GetBackUrl() string {
	return c.BackUrl
}

func (c ChinaMainlandBioCard) IsAdults() bool {
	//TODO implement me
	panic("implement me")
}

func (c ChinaMainlandBioCard) IsEnable(clientId int64) bool {
	if enable, ok := c.Enable[-1]; ok {
		if !enable {
			return false
		}
	}
	if enable, ok := c.Enable[clientId]; ok {
		return enable
	}
	return true
}

func (c ChinaMainlandBioCard) GetReason(clientId int64) string {
	if reason, ok := c.Reason[clientId]; ok {
		if len(reason) > 0 {
			return reason
		}
	}
	if reason, ok := c.Reason[clientId]; ok {
		return reason
	}
	return ""
}
