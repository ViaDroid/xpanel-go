package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type Paylist struct {
	Id        int       `orm:"description(记录ID)"`
	Userid    int       `orm:"description(用户ID)"`
	Total     float64   `orm:"description(总金额)"`
	Status    int       `orm:"description(状态)"`
	InvoiceId int       `orm:"description(账单ID)"`
	Tradeno   string    `orm:"description(网关单号)"`
	Gateway   string    `orm:"description(支付网关)"`
	Datetime  time.Time `orm:"description(创建时间)"`

	CreateTimeStr string `orm:"-"`
	StatusStr     string `orm:"-"`
}

func NewPaylist() *Paylist {
	return &Paylist{}
}

// fetch all order by id
func (item *Paylist) FindAll() []Paylist {
	var list []Paylist
	global.DB.QueryTable(item).OrderBy("-id").All(&list)
	return list
}

// find by invoice ID
func (item *Paylist) FindByInvoiceId(invoiceId int) (*Paylist, error) {
	err := global.DB.QueryTable(item).Filter("invoice_id", invoiceId).Filter("status", 1).One(item)
	return item, err
}

// Fetch list by invoiceId
func (item *Paylist) FetchListByInvoiceId(invoiceId int) []Paylist {
	var list []Paylist
	global.DB.QueryTable(item).Filter("invoice_id", invoiceId).OrderBy("-id").All(&list)
	return list
}

func (item *Paylist) ParseStatus() string {
	switch item.Status {
	case 0:
		return "未支付"
	case 1:
		return "已支付"
	// case 2:
	// 	return "已取消"
	default:
		return "未知"
	}
}

func (item *Paylist) Parse() {
	item.CreateTimeStr = item.Datetime.Format("2006-01-02 15:04:05")
	item.StatusStr = item.ParseStatus()
}
