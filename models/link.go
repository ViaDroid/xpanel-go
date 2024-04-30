package models

import "github.com/viadroid/xpanel-go/global"

type Link struct {
	Id     int    `orm:"description(记录ID)"`
	Token  string `orm:"unique;description(订阅token)"`
	UserId int    `orm:"unique;description(用户ID)"`
}

func NewLink() *Link {
	return &Link{}
}

func (item *Link) FindByUserId(userId int) (*Link, error) {
	err := global.DB.QueryTable(item).Filter("user_id", userId).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
// find by token*) 

func (item *Link) FindByToken(token string) (*Link, error) {
	err := global.DB.QueryTable(item).Filter("token", token).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}


func (item *Link) Save() (int64, error) {
	return global.DB.Insert(item)
}

func (item *Link) IsValid() bool {
	user := NewUser().FindById(item.UserId)
	if user.Id != 0 && !user.IsBanned {
		return true
	}

	return false
}
