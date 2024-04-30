package admin

import (
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type LoginLogController struct {
	controllers.BaseController
}

var login_log_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "用户ID"),
		NewField("Ip", "登录IP"),
		NewField("LocationStr", "IP归属地"),
		NewField("DatetimeStr", "时间"),
		NewField("TypeStr", "类型"),
	},
}

func (c *LoginLogController) Index() {
	c.Data["details"] = login_log_menu_details
	c.TplName = "views/tabler/admin/log/login.tpl"
}

func (c *LoginLogController) Ajax() {
	length, _ := c.GetInt("length")
	page, _ := c.GetInt("start")
	draw := c.GetString("draw")

	page = page/length + 1

	loginIp := models.NewLoginIp()
	total, _ := loginIp.Count()

	logins := loginIp.FetchList(length, page)

	for i, v := range logins {
		logins[i].DatetimeStr = time.UnixMilli(v.Datetime).Format("2006-01-02 15:04:05")
		logins[i].LocationStr = v.Location()
		logins[i].TypeStr = v.ParseTypeStr()
	}

	c.JSONResp(map[string]any{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": total,
		"logins":          logins,
	})
}
