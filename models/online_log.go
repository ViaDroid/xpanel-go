package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type OnlineLog struct {
	Id        int    `orm:"description(记录ID)"`
	UserId    int    `orm:"description(用户ID)"`
	Ip        string `orm:"description(IP地址)"`
	NodeId    int    `orm:"description(节点ID)"`
	FirstTime int64  `orm:"description(首次在线时间)"`
	LastTime  int64  `orm:"description(最后在线时间)"`

	LocationStr  string `orm:"-"`
	FirstTimeStr string `orm:"-"`
	LastTimeStr  string `orm:"-"`
	NodeName     string `orm:"-"`
}

func NewOnlineLog() *OnlineLog {
	return &OnlineLog{}
}

func (item *OnlineLog) FetchListByUserId(userId int) ([]OnlineLog, error) {
	var items []OnlineLog
	_, err := global.DB.QueryTable(item).Filter("user_id", userId).Filter("last_time__gt", time.Now()).OrderBy("-last_time").All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// count
func (item *OnlineLog) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

// fetch list by limit, page
func (item *OnlineLog) FetchList(limit, page int) []OnlineLog {
	var items []OnlineLog
	global.DB.QueryTable(item).Limit(limit).Offset((page - 1) * limit).OrderBy("-last_time").All(&items)
	return items
}

func (item *OnlineLog) GetNodeName() string {
	node := &Node{}
	global.DB.QueryTable(node).Filter("id", item.NodeId).One(&node)
	return node.Name
}

func (item *OnlineLog) Parse() {
	item.FirstTimeStr = time.UnixMilli(item.FirstTime).Format("2006-01-02 15:04:05")
	item.LastTimeStr = time.UnixMilli(item.LastTime).Format("2006-01-02 15:04:05")
	item.NodeName = item.GetNodeName()

	// TODO item.LocationStr = global.GetLocation(item.Ip)
}
