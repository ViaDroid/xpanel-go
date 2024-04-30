package models

import (
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/utils"
)

type InviteCode struct {
	Id     int    `orm:"description(记录ID)"`
	Code   string `orm:"description(邀请码)"`
	UserId int    `orm:"description(用户ID)"`
}

func (*InviteCode) TableName() string {
	return "user_invite_code"
}

func NewInviteCode() *InviteCode {
	return &InviteCode{}
}

func (item *InviteCode) FindByCode(code string) *InviteCode {
	global.DB.QueryTable(*item).Filter("code", code).One(item)
	return item
}

func (item *InviteCode) FindByUserId(userId int) *InviteCode {
	global.DB.QueryTable(*item).Filter("user_id", userId).One(item)
	return item
}

func (item *InviteCode) Add(userId int) (int64, error) {
	item.Code = utils.GenRandomString(10)
	item.UserId = userId
	return global.DB.Insert(item)
}
