package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type DetectLog struct {
	Id       int   `orm:"decription(检测记录ID)"`
	UserId   int   `orm:"decription(用户ID)"`
	ListId   int   `orm:"decription(规则ID)"`
	Datetime int64 `orm:"decription(检测时间)"`
	NodeId   int   `orm:"decription(节点ID)"`
	Status   int   `orm:"decription(状态)"`

	Op          string `orm:"-"`
	DatetimeStr string `orm:"-"`
}

func NewDetecLog() *DetectLog {
	return &DetectLog{}
}

func (item *DetectLog) Save() (int64, error) {
	return global.DB.Insert(item)
}

// Fetch all order by id
func (item *DetectLog) FetchAll() []DetectLog {
	var items []DetectLog
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}

// Fetch list order by id filter with userId
func (item *DetectLog) FetchAllByUserId(userId int) []DetectLog {
	var items []DetectLog
	global.DB.QueryTable(item).Filter("user_id", userId).OrderBy("-id").All(&items)
	return items
}

// count
func (item *DetectLog) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

// fetch list by limit, page
func (item *DetectLog) FetchList(limit, page int) []DetectLog {
	var items []DetectLog
	global.DB.QueryTable(item).Limit(limit).Offset((page - 1) * limit).OrderBy("-datetime").All(&items)
	return items
}

// parse
func (item *DetectLog) Parse() {
	item.DatetimeStr = time.UnixMilli(item.Datetime).Format("2006-01-02 15:04:05")
}
