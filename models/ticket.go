package models

import "github.com/viadroid/xpanel-go/global"

type Ticket struct {
	Id       int    `orm:"description(工单ID)"`
	Title    string `orm:"description(工单标题)"`
	Content  string `orm:"type(text);description(工单内容)"`
	UserId   int    `orm:"description(用户ID)"`
	Datetime int64  `orm:"description(创建时间)"`
	Status   string `orm:"description(工单状态)"`
	Type     string `orm:"description(工单类型)"`

	Op          string `orm:"-"`
	DatetimeStr string `orm:"-"`
}

func NewTicket() *Ticket {
	return &Ticket{}
}

// Fetch by user id
func (t *Ticket) FetchByUserId(userId int) []Ticket {
	var tickets []Ticket
	global.DB.QueryTable(t).Filter("user_id", userId).OrderBy("-datetime").All(&tickets)
	return tickets
}

// Fetch all Tickets
func (t *Ticket) FetchAll() []Ticket {
	var tickets []Ticket
	global.DB.QueryTable(t).OrderBy("-id").All(&tickets)
	return tickets
}

// Save a ticket
func (t *Ticket) Save() error {
	_, err := global.DB.Insert(t)
	return err
}

// Update a ticket
func (t *Ticket) Update() error {
	_, err := global.DB.Update(t)
	return err
}

// Find a ticket by id
func (t *Ticket) FindById(id int) error {
	return global.DB.QueryTable(t).Filter("id", id).One(t)
}

// Find a ticket by UserId
func (t *Ticket) FindByUserId(userId int) error {
	return global.DB.QueryTable(t).Filter("user_id", userId).One(t)
}

// Delete by id
func (t *Ticket) DeleteById(id int) (int64, error) {
	return global.DB.QueryTable(t).Filter("id", id).Delete()
}

// 工单类型
func (t *Ticket) ParseType() string {
	switch t.Type {
	case "howto":
		return "使用"
	case "billing":
		return "财务"
	case "account":
		return "账户"
	default:
		return "其它"
	}
}

func (t *Ticket) ParseStatus() string {
	switch t.Status {
	case "closed":
		return "已结单"
	case "open_wait_user":
		return "等待用户回复"
	case "open_wait_admin":
		return "进行中"
	default:
		return "未知"
	}
}
