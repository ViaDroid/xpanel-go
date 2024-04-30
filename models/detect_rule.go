package models

import "github.com/viadroid/xpanel-go/global"

type DetectRule struct {
	Id    int    `orm:"description(审计规则ID)"`
	Name  string `orm:"description(规则名称)"`
	Text  string `orm:"description(规则介绍)"`
	Regex string `orm:"description(规则内容)"`
	Type  int    `orm:"description(规则类型)"`

	Op      string `orm:"-"`
	TypeStr string `orm:"-"`
}

func NewDetecRule() *DetectRule {
	return &DetectRule{}
}

func (item *DetectRule) Save() (int64, error) {
	return global.DB.Insert(item)
}

// Delete
func (item *DetectRule) Delete(id int) (int64, error) {
	item.Id = id
	return global.DB.Delete(item)
}

// Fetch all order by id
func (item *DetectRule) FetchAll() []DetectRule {
	var items []DetectRule
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}

// Fetch one by id
func (item *DetectRule) FetchOne(id int) (*DetectRule, error) {
	item.Id = id
	err := global.DB.Read(item)
	return item, err
}

// parse type
func (item *DetectRule) ParseType() string {
	if item.Type == 1 {
		return "数据包明文匹配"
	} else {
		return "数据包 hex 匹配"
	}
}
