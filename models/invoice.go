package models

import (
	"encoding/json"
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type Invoice struct {
	Id         int     `orm:"description(账单ID)"`
	UserId     int     `orm:"description(归属用户ID)"`
	OrderId    int     `orm:"description(订单ID)"`
	Content    string  `orm:"description(账单内容)"`
	Price      float64 `orm:"description(账单金额)"`
	Status     string  `orm:"description(账单状态)"`
	CreateTime int64   `orm:"description(创建时间)"`
	UpdateTime int64   `orm:"description(更新时间)"`
	PayTime    int64   `orm:"description(支付时间)"`

	Op            string         `orm:"-"`
	StatusStr     string         `orm:"-"`
	CreateTimeStr string         `orm:"-"`
	UpdateTimeStr string         `orm:"-"`
	PayTimeStr    string         `orm:"-"`
	ContentMap    map[string]any `orm:"-"`
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

// save
func (item *Invoice) Save() error {
	_, err := global.DB.Insert(item)
	return err
}

// update
func (item *Invoice) Update() error {
	_, err := global.DB.Update(item)
	return err
}

// delete
func (item *Invoice) Delete() error {
	_, err := global.DB.Delete(item)
	return err
}

// find by id
func (item *Invoice) Find(id int) (*Invoice, error) {
	err := global.DB.QueryTable(item).Filter("Id", id).One(item)
	return item, err
}

// find one by order id
func (item *Invoice) FindOneByOrderId(orderId int) (*Invoice, error) {
	err := global.DB.QueryTable(item).Filter("order_id", orderId).One(item)
	return item, err
}

// fetch all order by id
func (item *Invoice) FetchAll() []Invoice {
	var list []Invoice
	global.DB.QueryTable(item).OrderBy("-id").All(&list)
	return list
}

// fetch list by userId
func (item *Invoice) FetchAllByUserId(userId int) []Invoice {
	var list []Invoice
	global.DB.QueryTable(item).Filter("user_id", userId).OrderBy("-id").All(&list)
	return list
}

// fetch all order by id
func (item *Invoice) FetchAllByPage(page, pageSize int) ([]*Invoice, error) {
	var list []*Invoice
	_, err := global.DB.QueryTable(item).OrderBy("-id").Limit(pageSize, (page-1)*pageSize).All(&list)
	return list, err
}

// count all order
func (item *Invoice) Count() (int64, error) {
	return global.DB.QueryTable(item).Count()
}

func (item *Invoice) ParseStatus() string {
	switch item.Status {
	case "unpaid":
		return "未支付"
	case "paid_gateway":
		return "已支付（支付网关）"
	case "paid_balance":
		return "已支付（账户余额）"
	case "paid_admin":
		return "已支付（管理员）"
	case "cancelled":
		return "已取消"
	case "refunded_balance":
		return "已退款（账户余额）"
	default:
		return "未知"
	}
}

func (item *Invoice) Parse() {
	json.Unmarshal([]byte(item.Content), &item.ContentMap)
	item.StatusStr = item.ParseStatus()
	item.CreateTimeStr = time.UnixMilli(item.CreateTime).Format("2006-01-02 15:04:05")
	item.UpdateTimeStr = time.UnixMilli(item.UpdateTime).Format("2006-01-02 15:04:05")

	if item.PayTime == 0 {
		item.PayTimeStr = "——"
	} else {
		item.PayTimeStr = time.UnixMilli(item.PayTime).Format("2006-01-02 15:04:05")
	}
}
