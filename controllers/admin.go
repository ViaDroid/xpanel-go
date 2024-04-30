package controllers

import "github.com/viadroid/xpanel-go/services"

type AdminController struct {
	BaseController
}

func (c *AdminController) Index() {

	service := services.NewAnalyticsService()
	todayIncome := service.GetIncome("today")
	yesterday_income := service.GetIncome("yesterday_income")
	this_month_income := service.GetIncome("this month")
	total_income := service.GetIncome("total")
	total_user := service.GetTotalUser()
	checkin_user := service.GetCheckinUser()
	today_checkin_user := service.GetTodayCheckinUser()
	inactive_user := service.GetInactiveUser()
	active_user := service.GetActiveUser()
	total_node := service.GetTotalNode()
	alive_node := service.GetAliveNode()
	raw_today_traffic := service.GetRawGbTodayTrafficUsage()
	raw_last_traffic := service.GetRawGbLastTrafficUsage()
	raw_unused_traffic := service.GetRawGbUnusedTrafficUsage()
	today_traffic := service.GetTodayTrafficUsage()
	last_traffic := service.GetLastTrafficUsage()
	unused_traffic := service.GetUnusedTrafficUsage()

	// 流水
	c.Data["today_income"] = todayIncome
	c.Data["yesterday_income"] = yesterday_income
	c.Data["this_month_income"] = this_month_income
	c.Data["total_income"] = total_income
	// 用户的签到情况
	c.Data["total_user"] = total_user
	c.Data["checkin_user"] = checkin_user
	c.Data["today_checkin_user"] = today_checkin_user
	// 闲置账户
	c.Data["inactive_user"] = inactive_user
	c.Data["active_user"] = active_user
	// 服务器的在线情况
	c.Data["total_node"] = total_node
	c.Data["alive_node"] = alive_node
	// 流量用量
	c.Data["raw_today_traffic"] = int(raw_today_traffic)
	c.Data["raw_last_traffic"] = int(raw_last_traffic)
	c.Data["raw_unused_traffic"] = int(raw_unused_traffic)
	c.Data["today_traffic"] = today_traffic
	c.Data["last_traffic"] = last_traffic
	c.Data["unused_traffic"] = unused_traffic

	c.TplName = "views/tabler/admin/index.tpl"
}
