package models

import (
	"github.com/viadroid/xpanel-go/global"
)

type Ann struct {
	Id      int    `orm:"description(公告ID)"`
	Date    string `orm:"description(公告日期)"`
	Content string `orm:"description(公告内容)"`

	Op string `orm:"-"`
}

func (*Ann) TableName() string {
	return "announcement"
}

func NewAnn() *Ann {
	return &Ann{}
}

func (item *Ann) FetchOne() (*Ann, error) {
	err := global.DB.QueryTable(item).OrderBy("-date").One(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Ann) FetchList() []Ann {
	var items []Ann
	global.DB.QueryTable(item).OrderBy("-date").All(&items)
	return items
}

// Find by id
func (item *Ann) FindById(id int) (*Ann, error) {
	global.DB.QueryTable(item).Filter("id", id).One(item)
	return item, nil
}

// Fetch all order by id desc
func (item *Ann) FetchAll() []Ann {
	var items []Ann
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}

func (item *Ann) Save() (int64, error) {
	return global.DB.Insert(item)
}

func (item *Ann) Update() (int64, error) {
	return global.DB.Update(item)
}

func (item *Ann) Delete() (int64, error) {
	return global.DB.Delete(item)
}
