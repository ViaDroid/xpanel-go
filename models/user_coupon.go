package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type UserCoupon struct {
	Id         int    `orm:"description(优惠码ID)"`
	Code       string `orm:"description(优惠码)"`
	Content    string `orm:"description(优惠码内容)"`
	Limit      string `orm:"description(优惠码限制)"`
	UseCount   int    `orm:"description(累计使用次数)"`
	CreateTime int64  `orm:"description(创建时间)"`
	ExpireTime int64  `orm:"description(过期时间)"`

	Op            string         `orm:"-"`
	Type          string         `orm:"-"`
	Value         string         `orm:"-"`
	ProductId     string         `orm:"-"`
	UseTime       string         `orm:"-"`
	TotalUseTime  string         `orm:"-"`
	NewUser       string         `orm:"-"`
	Disabled      string         `orm:"-"`
	CreateTimeStr string         `orm:"-"`
	ExpireTimeStr string         `orm:"-"`
	IsExpired     bool           `orm:"-"`
	ContentMap    map[string]any `orm:"-"`
	LimitMap      map[string]any `orm:"-"`
}

func NewUserCoupon() *UserCoupon {
	return &UserCoupon{}
}

// Save
func (m *UserCoupon) Save() (int64, error) {
	return global.DB.Insert(m)
}

// Update
func (m *UserCoupon) Update() (int64, error) {
	return global.DB.Update(m)
}

// Delete
func (m *UserCoupon) Delete(id int) (int64, error) {
	m.Id = id
	return global.DB.Delete(m)
}

// Fetch all order by id desc
func (m *UserCoupon) FetchAll() []UserCoupon {
	var list []UserCoupon
	global.DB.QueryTable(m).OrderBy("-id").All(&list)
	return list
}

// 优惠码类型
func (m *UserCoupon) GetCouponType() string {
	var content map[string]string
	json.Unmarshal([]byte(m.Content), &content)

	switch content["type"] {
	case "percentage":
		return "百分比"
	case "fixed":
		return "固定金额"
	default:
		return "未知"
	}
}

// 优惠码状态
func (m *UserCoupon) GetCouponStatus() string {
	if m.ExpireTime < time.Now().UnixMilli() {
		return "已过期"
	} else {
		return "激活"
	}
}

// Find by id
func (m *UserCoupon) FindById(id int) (*UserCoupon, error) {
	m.Id = id
	return m, global.DB.Read(m)
}

// Find by code
func (m *UserCoupon) FindByCode(code string) (*UserCoupon, error) {
	m.Code = code
	return m, global.DB.Read(m)
}

// count by code
func (m *UserCoupon) CountByCode(code string) int64 {
	count, _ := global.DB.QueryTable(m).Filter("Code", code).Count()
	return count
}

func (m *UserCoupon) Parse() {
	var content, limit map[string]any
	json.Unmarshal([]byte(m.Content), &content)
	json.Unmarshal([]byte(m.Limit), &limit)

	m.ContentMap = content
	m.LimitMap = limit

	m.Type = m.GetCouponType()
	m.Value = content["value"].(string)
	m.ProductId = limit["product_id"].(string)

	if limit["use_time"].(float64) == 0 {
		m.UseTime = "不限次数"
	} else {
		m.UseTime = fmt.Sprintf("%.f次", limit["use_time"])
	}

	if limit["total_use_time"].(float64) == 0 {
		m.TotalUseTime = "不限次数"
	} else {
		m.TotalUseTime = fmt.Sprintf("%.f次", limit["total_use_time"])
	}

	if limit["new_user"] == 1 {
		m.NewUser = "是"
	} else {
		m.NewUser = "否"
	}

	if limit["disabled"].(float64) == 1 {
		m.Disabled = "是"
	} else {
		m.Disabled = "否"
	}

	m.CreateTimeStr = time.UnixMilli(m.CreateTime).Format("2006-01-02 15:04")

	if m.ExpireTime == 0 {
		m.ExpireTimeStr = "永久有效"
	} else {
		m.ExpireTimeStr = time.UnixMilli(m.ExpireTime).Format("2006-01-02 15:04")
	}

	// is expire
	if m.ExpireTime > 0 && m.ExpireTime < time.Now().UnixMilli() {
		m.IsExpired = true
	} else {
		m.IsExpired = false
	}
}
