package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type LoginIp struct {
	Id       int    `orm:"description(记录ID)"`
	UserId   int    `orm:"description(用户ID)"`
	Ip       string `orm:"description(登录IP)"`
	Datetime int64  `orm:"description(登录时间)"`
	Type     int    `orm:"description(登录类型)"`

	LocationStr string `orm:"-"`
	TypeStr     string `orm:"-"`
	DatetimeStr string `orm:"-"`
}

func NewLoginIp() *LoginIp {
	return &LoginIp{}
}

func (item *LoginIp) Save() (int64, error) {
	return global.DB.Insert(item)
}

// count
func (item *LoginIp) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

// fetch all
func (item *LoginIp) FetchAll() ([]LoginIp, error) {
	var items []LoginIp
	_, err := global.DB.QueryTable(item).All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// fetch list by limit, page
func (item *LoginIp) FetchList(limit, page int) []LoginIp {
	var items []LoginIp
	global.DB.QueryTable(item).Limit(limit).Offset((page - 1) * limit).OrderBy("-datetime").All(&items)
	return items
}

// find by user_id
func (item *LoginIp) FindByUserId(userId int) error {
	return global.DB.QueryTable(item).Filter("user_id", userId).One(item)
}

func (item *LoginIp) IsExist(userId int) bool {
	return global.DB.QueryTable(item).Filter("user_id", userId).Exist()
}

func (item *LoginIp) FetchListByUserId(userId int) ([]LoginIp, error) {
	var items []LoginIp
	_, err := global.DB.QueryTable(item).Filter("user_id", userId).Filter("type", 0).OrderBy("-datetime").Limit(10).All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (item *LoginIp) DateTimeStr() string {
	return time.UnixMilli(item.Datetime).Format(time.DateTime)
}

func (item *LoginIp) Location() string {
	return "TODO"
}

func (item *LoginIp) ParseTypeStr() string {
	if item.Type == 0 {
		return "成功"
	}
	return "失败"
}
