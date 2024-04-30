package models

import "github.com/viadroid/xpanel-go/global"

type DetectBanLog struct {
	Id              int   `orm:"description(封禁记录ID)"`
	UserId          int   `orm:"index;description(用户ID)"`
	DetectNumber    int   `orm:"description(本次违规次数)"`
	BanTime         int64 `orm:"description(封禁时长)"`
	StartTime       int64 `orm:"description(封禁开始时间)"`
	EndTime         int64 `orm:"description(封禁结束时间)"`
	AllDetectNumber int   `orm:"description(累计违规次数)"`

	StartTimeStr string `orm:"-"`
	EndTimeStr   string `orm:"-"`
}

func NewDetecBanLog() *DetectBanLog {
	return &DetectBanLog{}
}

func (item *DetectBanLog) Save() (int64, error) {
	return global.DB.Insert(item)
}

// Fetch all order by id
func (item *DetectBanLog) FetchAll() []DetectBanLog {
	var items []DetectBanLog
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}

// count
func (item *DetectBanLog) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

// fetch list by limit, page
func (item *DetectBanLog) FetchList(limit, page int) []DetectBanLog {
	var items []DetectBanLog
	global.DB.QueryTable(item).Limit(limit).Offset((page - 1) * limit).OrderBy("-id").All(&items)
	return items
}
