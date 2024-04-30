package models

import (
	"encoding/json"
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type Product struct {
	Id         int     `orm:"description(商品ID)"`
	Type       string  `orm:"description(类型)"`
	Name       string  `orm:"description(名称)"`
	Price      float64 `orm:"description(售价)"`
	Content    string  `orm:"description(内容)"`
	Limit      string  `orm:"description(购买限制)"`
	Status     int     `orm:"description(销售状态)"`
	CreateTime int64   `orm:"description(创建时间)"`
	UpdateTime int64   `orm:"description(更新时间)"`
	SaleCount  int     `orm:"description(累计销量)"`
	Stock      int     `orm:"description(库存)"`

	Op            string         `orm:"-"`
	StatusStr     string         `orm:"-"`
	CreateTimeStr string         `orm:"-"`
	UpdateTimeStr string         `orm:"-"`
	ContentMap    map[string]any `orm:"-"`
	LimitMap      map[string]any `orm:"-"`
	TypeText      string         `orm:"-"`
}

type Content struct {
	Time       int `json:"time"`
	Bandwidth  int `json:"bandwidth"`
	Class      int `json:"class"`
	ClassTime  int `json:"class_time"`
	NodeGroup  int `json:"node_group"`
	SpeedLimit int `json:"speed_limit"`
	IpLimit    int `json:"ip_limit"`
}

type Limit struct {
	ClassRequired     int `json:"class_required"`
	NodeGroupRequired int `json:"node_group_required"`
	NewUserRequired   int `json:"new_user_required"`
}

func NewProduct() *Product {
	return &Product{}
}

// save
func (item *Product) Save() (int64, error) {
	return global.DB.Insert(item)

}

// update
func (item *Product) Update() (int64, error) {
	return global.DB.Update(item)
}

// delete
func (item *Product) Delete(id int) (int64, error) {
	item.Id = id
	return global.DB.Delete(item)
}

// is exist
func (item *Product) IsExist(id int) bool {
	return global.DB.QueryTable(item).Filter("Id", id).Exist()
}

// find by id
func (item *Product) Find(id int) (*Product, error) {
	err := global.DB.QueryTable(item).Filter("Id", id).One(item)
	return item, err
}

// fetch all order by id
func (item *Product) FetchAll() []Product {
	var products []Product
	global.DB.QueryTable(item).OrderBy("-id").All(&products)
	return products
}

// fetch list by type and status
func (item *Product) FetchList(typeStr string, status int) []Product {
	var products []Product
	global.DB.QueryTable(item).Filter("Type", typeStr).Filter("status", status).OrderBy("-id").All(&products)
	return products
}

func (item *Product) ParseType() string {
	switch item.Type {
	case "tabp":
		item.TypeText = "时间流量包"
	case "time":
		item.TypeText = "时间包"
	case "bandwidth":
		item.TypeText = "流量包"
	default:
		item.TypeText = "其他"
	}
	return item.TypeText
}

func (item *Product) Parse() {
	item.ParseType()

	if item.Status == 1 {
		item.StatusStr = "正常"
	} else {
		item.StatusStr = "下架"
	}
	item.CreateTimeStr = time.UnixMilli(item.CreateTime).Format(time.DateTime)
	item.UpdateTimeStr = time.UnixMilli(item.UpdateTime).Format(time.DateTime)

	json.Unmarshal([]byte(item.Content), &item.ContentMap)
	json.Unmarshal([]byte(item.Limit), &item.LimitMap)
}
