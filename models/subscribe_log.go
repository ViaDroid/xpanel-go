package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type SubscribeLog struct {
	Id               int    `orm:"description(记录ID)"`
	UserId           int    `orm:"description(用户ID)"`
	Type             string `orm:"description(获取的订阅类型)"`
	RequestIp        string `orm:"description(请求IP)"`
	RequestUserAgent string `orm:"description(请求UA)"`
	RequestTime      int64  `orm:"description(请求时间)"`
}

func NewSubscribeLog() *SubscribeLog {
	return &SubscribeLog{}
}

// count
func (item *SubscribeLog) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

// fetch list by limit, page
func (item *SubscribeLog) FetchList(limit, page int) []SubscribeLog {
	var items []SubscribeLog
	global.DB.QueryTable(item).Limit(limit).Offset((page - 1) * limit).OrderBy("-request_time").All(&items)
	return items
}

func (item *SubscribeLog) Add(user *User, typ, requestIp, ua string) (int64, error) {
	log := &SubscribeLog{
		UserId:           user.Id,
		Type:             typ,
		RequestIp:        requestIp,
		RequestUserAgent: ua,
		RequestTime:      time.Now().UnixMilli(),
	}

	return global.DB.Insert(log)
}
