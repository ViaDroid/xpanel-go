package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type Payback struct {
	Id        int     `orm:"description(记录ID)"`
	Total     float64 `orm:"description(总金额)"`
	UserId    int     `orm:"description(用户ID)"`
	RefBy     int     `orm:"description(推荐人ID)"`
	RefGet    float64 `orm:"description(推荐人获得金额)"`
	InvoiceId int     `orm:"description(账单ID)"`
	Datetime  int64   `orm:"description(创建时间)"`

	DatetimeStr string `orm:"-"`
}

func NewPayback() *Payback {
	return &Payback{}
}

func (item *Payback) FindByRefBy(refBy int) []Payback {
	var list []Payback
	global.DB.QueryTable(item).Filter("ref_by", refBy).OrderBy("-id").All(&list)
	return list
}

func (item *Payback) SumByRefBy(refBy int) float64 {
	var sum float64
	global.DB.Raw("select sum(ref_get) from payback where ref_by=?", refBy).QueryRow(&sum)
	return sum
}

func (item *Payback) UserName() string {
	var user User
	err := global.DB.QueryTable(User{}).Filter("id", item.UserId).One(&user)
	if err != nil {
		return "已注销"
	}
	return user.UserName
}

// fetch all order by id
func (item *Payback) FindAll() []Payback {
	var list []Payback
	global.DB.QueryTable(item).OrderBy("-id").All(&list)
	return list
}

func (item *Payback) Parse() {
	item.DatetimeStr = time.UnixMilli(item.Datetime).Format(time.DateTime)
}
