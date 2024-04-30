package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type UserMoneyLog struct {
	Id         int     `orm:"description(记录ID)"`
	UserId     int     `orm:"description(用户ID)"`
	Before     float64 `orm:"description(用户变动前账户余额)"`
	After      float64 `orm:"description(用户变动后账户余额)"`
	Amount     float64 `orm:"description(变动总额)"`
	Remark     string  `orm:"description(备注)"`
	CreateTime int64   `orm:"description(创建时间)"`

	CreateTimeStr string `orm:"-"`
}

func NewUserMoneyLog() *UserMoneyLog {
	return &UserMoneyLog{}
}

func (item *UserMoneyLog) Add(userId int, before, after, amount float64, remark string) (int64, error) {
	log := &UserMoneyLog{
		UserId:     userId,
		Before:     before,
		After:      after,
		Amount:     amount,
		Remark:     remark,
		CreateTime: time.Now().UnixMilli(),
	}
	return global.DB.Insert(log)
}

// fetch all order by id
func (item *UserMoneyLog) FindAll() []UserMoneyLog {
	var list []UserMoneyLog
	global.DB.QueryTable(item).OrderBy("-id").All(&list)
	return list
}

// fetch list by userId
func (item *UserMoneyLog) FindListByUserId(userId int) []UserMoneyLog {
	var list []UserMoneyLog
	global.DB.QueryTable(item).Filter("UserId", userId).OrderBy("-id").All(&list)
	return list
}

func (item *UserMoneyLog) Parse() {
	item.CreateTimeStr = time.UnixMilli(item.CreateTime).Format("2006-01-02 15:04:05")
}
