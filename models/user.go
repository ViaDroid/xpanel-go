package models

import (
	"fmt"
	"math"
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/utils"
	"golang.org/x/crypto/sha3"
)

type User struct {
	Id                 int     `orm:"description(用户ID)"`
	UserName           string  `orm:"description(用户名)"`
	Email              string  `orm:"unique;description(E-Mail)"`
	Pass               string  `orm:"description(登录密码)"`
	Passwd             string  `orm:"description(节点密码)"`
	Uuid               string  `orm:"unique;description(UUID)"`
	U                  int64   `orm:"description(账户当前上传流量)"`
	D                  int64   `orm:"description(账户当前下载流量)"`
	TransferToday      int64   `orm:"description(账户今日所用流量)"`
	TransferTotal      int64   `orm:"description(账户累计使用流量)"`
	TransferEnable     int64   `orm:"description(账户当前可用流量)"`
	Port               int     `orm:"description(端口)"`
	LastDetectBanTime  string  `orm:"description(最后一次被封禁的时间)"`
	AllDetectNumber    int     `orm:"description(累计违规次数)"`
	LastUseTime        int64   `orm:"description(最后使用时间)"`
	LastCheckInTime    int64   `orm:"description(最后签到时间)"`
	LastLoginTime      int64   `orm:"description(最后登录时间)"`
	RegDate            string  `orm:"description(注册时间)"`
	Money              float64 `orm:"description(账户余额)"`
	RefBy              int     `orm:"description(邀请人ID)"`
	Method             string  `orm:"description(Shadowsocks加密方式)"`
	RegIp              string  `orm:"description(注册IP)"`
	NodeSpeedlimit     float64 `orm:"description(用户限速)"`
	NodeIplimit        int     `orm:"description(同时可连接IP数)"`
	IsAdmin            bool    `orm:"description(是否管理员)"`
	ImType             int     `orm:"description(联系方式类型)"`
	ImValue            string  `orm:"description(联系方式)"`
	ContactMethod      int     `orm:"description(偏好的联系方式)"`
	DailyMailEnable    int     `orm:"description(每日报告开关)"`
	Class              int     `orm:"description(等级)"`
	ClassExpire        string  `orm:"description(等级过期时间)"`
	Theme              string  `orm:"description(网站主题)"`
	GaToken            string  `orm:"unique;description(GA密钥)"`
	GaEnable           bool    `orm:"description(GA开关)"`
	Remark             string  `orm:"description(备注)"`
	NodeGroup          int     `orm:"description(节点分组)"`
	IsBanned           bool    `orm:"description(是否封禁)"`
	BannedReason       string  `orm:"description(封禁理由)"`
	IsShadowBanned     bool    `orm:"description(是否处于账户异常状态)"`
	ExpireNotified     int     `orm:"description(过期提醒)"`
	TrafficNotified    int     `orm:"description(流量提醒)"`
	ForbiddenIp        string  `orm:"description(禁止访问IP)"`
	ForbiddenPort      string  `orm:"description(禁止访问端口)"`
	AutoResetDay       int     `orm:"description(自动重置流量日)"`
	AutoResetBandwidth float64 `orm:"description(自动重置流量)"`
	ApiToken           string  `orm:"unique;description(API 密钥)"`
	IsDarkMode         int     `orm:"description(是否启用暗黑模式)"`
	IsInactive         bool    `orm:"description(是否处于闲置状态)"`
	Locale             string  `orm:"description(显示语言)"`

	Op                string `orm:"-"`
	TransferUsed      string `orm:"-"`
	TransferEnableStr string `orm:"-"`
	IsAdminStr        string `orm:"-"`
	IsBannedStr       string `orm:"-"`
	IsInactiveStr     string `orm:"-"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Save() (int64, error) {
	return global.DB.Insert(u)
}

func (u *User) Update() (int64, error) {
	return global.DB.Update(u)
}

func (u *User) FindById(id int) *User {
	err := global.DB.QueryTable(u).Filter("id", id).One(u)
	if err != nil {
		return nil
	}
	return u
}

func (u *User) FindByEmail(email string) (*User, error) {
	err := global.DB.QueryTable(u).Filter("email", email).One(u)
	return u, err
}

func (u *User) IsExist(email string) bool {
	return global.DB.QueryTable(u).Filter("email", email).Exist()
}

func (u *User) Count() int64 {
	count, _ := global.DB.QueryTable(u).Count()
	return count
}

func (u *User) PortArray() []int {
	list := []User{}
	global.DB.QueryTable(u).All(&list)

	arr := []int{}

	for _, v := range list {
		arr = append(arr, v.Port)
	}
	return arr
}

func (u *User) DiceBear() string {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(u.Email))
	str := fmt.Sprintf("%x\n", h)
	return fmt.Sprintf("https://api.dicebear.com/7.x/identicon/svg?seed=%s", str)
}

func (u *User) SumByField(field string) float64 {
	var sum float64
	sql := fmt.Sprintf("select sum(%s) from user", field)
	err := global.DB.Raw(sql).QueryRow(&sum)
	if err != nil {
		return 0
	}
	return sum
}

func (u *User) FetchAll() []User {
	list := []User{}
	global.DB.QueryTable(u).OrderBy("-id").All(&list)
	return list
}

func (u *User) IsBindIm(imType int, imValue string) bool {
	return global.DB.QueryTable(u).Filter("im_type", imType).Filter("im_value", imValue).Exist()
}

// 联系方式类型
func (u *User) GetImType() string {
	var res string
	switch u.ImType {
	case 1:
		res = "Slack"
	case 2:
		res = "Discord"
	default:
		res = "Telegram"
	}
	return res
}

// 今天使用的流量[自动单位]
func (u *User) TodayUsedTraffic() string {
	return utils.AutoBytes(u.TransferToday)
}

/*
 * 今天使用的流量占总流量的百分比
 */
func (u *User) TodayUsedTrafficPercent() float64 {
	if u.TransferEnable == 0 {
		return 0
	}
	return math.Round(float64(u.TransferToday/u.TransferEnable)) * 100

}

// 今天之前已使用的流量[自动单位]
func (u *User) LastusedTraffic() string {
	return utils.AutoBytes(u.U + u.D - u.TransferToday)
}

// 今天之前已使用的流量占总流量的百分比
func (u *User) LastusedTrafficPercent() float64 {
	if u.TransferEnable == 0 {
		return 0
	}
	return math.Round(float64(u.U+u.D-u.TransferToday/u.TransferEnable)) * 100
}

// 最后使用时间
func (u *User) GetLastUseTime() string {
	res := "从未使用"
	if u.LastUseTime != 0 {
		res = time.UnixMilli(u.LastUseTime).Format(time.DateTime)
	}
	return res
}

// 最后签到时间
func (u *User) GetLastCheckInTime() string {
	res := "从未签到"
	if u.LastCheckInTime != 0 {
		res = time.UnixMilli(u.LastCheckInTime).Format(time.DateTime)
	}
	return res
}

/*
 * 总流量[自动单位]
 */
func (u *User) EnableTraffic() string {
	return utils.AutoBytes(u.TransferEnable)
}

/*
 * 当期用量[自动单位]
 */
func (u *User) UsedTraffic() string {
	return utils.AutoBytes(u.U + u.D)
}

/*
 * 累计用量[自动单位]
 */
func (u *User) TotalTraffic() string {
	return utils.AutoBytes(u.TransferTotal)
}

/*
 * 剩余流量[自动单位]
 */
func (u *User) UnusedTraffic() string {
	return utils.AutoBytes(u.TransferEnable - (u.U + u.D))
}

// 是否可以签到
func (u *User) IsAbleToCheckin() bool {
	return time.UnixMilli(u.LastCheckInTime).Format(time.DateOnly) != time.Now().Format(time.DateOnly) && !u.IsShadowBanned
}

// 删除用户的订阅链接
func (u *User) RemoveLink() (int64, error) {
	return global.DB.QueryTable(Link{}).Filter("user_id", u.Id).Delete()
}

// 删除用户的邀请码
func (u *User) RemoveInvite() {
	global.DB.QueryTable(InviteCode{}).Filter("user_id", u.Id).Delete()
}

// 销户
func (u *User) Kill() (int64, error) {
	global.DB.QueryTable(DetectBanLog{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(DetectLog{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(InviteCode{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(OnlineLog{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(Link{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(LoginIp{}).Filter("user_id", u.Id).Delete()
	global.DB.QueryTable(SubscribeLog{}).Filter("user_id", u.Id).Delete()
	return global.DB.QueryTable(User{}).Filter("id", u.Id).Delete()
}

func (u *User) UnbindIM() (int64, error) {
	u.ImType = 0
	u.ImValue = ""
	return u.Update()
}

func (u *User) SendTrafficNotification() {

}

// 发送每日流量报告
func (u *User) SendDailyNotification() {

}

func (u *User) MethodStr() string {
	return u.Method
}

func (u *User) Parse() {
	u.TransferUsed = u.UsedTraffic()
	u.TransferEnableStr = u.EnableTraffic()

	f := func(v bool) string {
		if v {
			return "是"
		} else {
			return "否"
		}
	}

	u.IsAdminStr = f(u.IsAdmin)
	u.IsBannedStr = f(u.IsBanned)
	u.IsInactiveStr = f(u.IsInactive)

}
