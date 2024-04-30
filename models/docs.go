package models

import "github.com/viadroid/xpanel-go/global"

type Docs struct {
	Id      int    `orm:"description(文档ID)"`
	Date    string `orm:"description(文档日期)"`
	Title   string `orm:"description(文档标题)"`
	Content string `orm:"description(文档内容)"`

	Op string `orm:"-"`
}

func NewDocs() *Docs {
	return &Docs{}
}

func (item *Docs) FindDocsById(id int) Docs {
	global.DB.QueryTable(item).Filter("id", id).One(item)
	return *item
}
func (item *Docs) FetchDocs() []Docs {
	var list []Docs
	global.DB.QueryTable(item).OrderBy("-id").All(&list)
	return list
}

// Save
func (item *Docs) Save() (int64, error) {
	return global.DB.Insert(item)
}

// Find by id
func (item *Docs) FindById(id int) (*Docs, error) {
	return item, global.DB.QueryTable(item).Filter("id", id).One(item)
}

// Update
func (item *Docs) Update() (int64, error) {
	return global.DB.Update(item)
}

// Delete
func (item *Docs) Delete() (int64, error) {
	return global.DB.Delete(item)
}

// Fetch all order by id desc
func (item *Docs) FetchAll() []Docs {
	var items []Docs
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}
