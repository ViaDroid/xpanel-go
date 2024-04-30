package models

import (
	"encoding/json"
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type Order struct {
	Id             int     `orm:"description(订单ID)"`
	UserId         int     `orm:"description(提交用户ID)"`
	ProductId      int     `orm:"description(商品ID)"`
	ProductType    string  `orm:"description(商品类型)"`
	ProductName    string  `orm:"description(商品名称)"`
	ProductContent string  `orm:"description(商品内容)"`
	Coupon         string  `orm:"description(订单优惠码)"`
	Price          float64 `orm:"description(订单金额)"`
	Status         string  `orm:"description(订单状态)"`
	CreateTime     int64   `orm:"description(创建时间)"`
	UpdateTime     int64   `orm:"description(更新时间)"`

	Op             string         `orm:"-"`
	ProductTypeStr string         `orm:"-"`
	StatusStr      string         `orm:"-"`
	CreateTimeStr  string         `orm:"-"`
	UpdateTimeStr  string         `orm:"-"`
	ContentMap     map[string]any `orm:"-"`
}

func NewOrder() *Order {
	return &Order{}
}

// save
func (item *Order) Save() (int64, error) {
	return global.DB.Insert(item)

}

// update
func (item *Order) Update() (int64, error) {
	return global.DB.Update(item)
}

// delete
func (item *Order) Delete(id int) (int64, error) {
	item.Id = id
	return global.DB.Delete(item)
}

// is exist
func (item *Order) IsExist(id int) bool {
	return global.DB.QueryTable(item).Filter("Id", id).Exist()
}

// find by id
func (item *Order) Find(id int) (*Order, error) {
	err := global.DB.QueryTable(item).Filter("Id", id).One(item)
	return item, err
}

// find by userId and id
func (item *Order) FindByUserIdAndId(userId, id int) (*Order, error) {
	err := global.DB.QueryTable(item).Filter("UserId", userId).Filter("Id", id).One(item)
	return item, err
}

// fetch all order by id
func (item *Order) FetchAll() []Order {
	var products []Order
	global.DB.QueryTable(item).OrderBy("-id").All(&products)
	return products
}

// Fetch list by userId order by id
func (item *Order) FetchListByUserId(userId int) []Order {
	var products []Order
	global.DB.QueryTable(item).Filter("user_id", userId).OrderBy("-id").All(&products)
	return products
}

func (item *Order) ParseStatus() string {
	switch item.Status {
	case "pending_payment":
		return "等待中"
	case "pending_activation":
		return "待激活"
	case "activated":
		return "已激活"
	case "expired":
		return "已过期"
	case "cancelled":
		return "已取消"
	}
	return "未知"
}

func (item *Order) ParseProductType() string {
	switch item.ProductType {
	case "tabp":
		return "时间流量包"
	case "time":
		return "时间包"
	case "bandwidth":
		return "流量包"
	default:
		return "其它"
	}
}

func (item *Order) Parse() {
	json.Unmarshal([]byte(item.ProductContent), &item.ContentMap)
	item.ProductTypeStr = item.ParseProductType()
	item.StatusStr = item.ParseStatus()
	item.CreateTimeStr = time.UnixMilli(item.CreateTime).Format("2006-01-02 15:04:05")
	item.UpdateTimeStr = time.UnixMilli(item.UpdateTime).Format("2006-01-02 15:04:05")
}
