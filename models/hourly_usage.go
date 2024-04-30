package models

type HourlyUsage struct {
	Id     int    `orm:"description(记录ID)"`
	UserId int    `orm:"description(用户ID)"`
	Date   string `orm:"description(记录日期)"`
	Usage  string `orm:"description(流量用量)"`
}

func NewHourlyUsage() *HourlyUsage {
	return &HourlyUsage{}
}


