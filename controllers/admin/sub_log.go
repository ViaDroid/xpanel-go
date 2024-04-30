package admin

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type SubLogController struct {
	controllers.BaseController
}

var sub_log_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "用户ID"),
		NewField("Type", "获取的订阅类型"),
		NewField("RequestIp", "请求IP"),
		NewField("Location", "IP归属地"),
		NewField("RequestTime", "请求时间"),
		NewField("RequestUserAgent", "客户端标识符"),
	},
}

func (c *SubLogController) Index() {
	c.Data["details"] = sub_log_menu_details
	c.TplName = "views/tabler/admin/log/sub.tpl"
}

func (c *SubLogController) Ajax() {
	length, _ := c.GetInt("length")
	page, _ := c.GetInt("start")
	draw := c.GetString("draw")

	page = page/length + 1

	subLog := models.NewSubscribeLog()
	total, _ := subLog.Count()

	list := subLog.FetchList(length, page)

	// for i, v := range list {
	// onlines[i].LocationStr = v.Location()
	// list[i].FirstTimeStr = time.UnixMilli(v.FirstTime).Format("2006-01-02 15:04:05")
	// list[i].LastTimeStr = time.UnixMilli(v.LastTime).Format("2006-01-02 15:04:05")
	// list[i].NodeName = v.GetNodeName()
	// }

	c.JSONResp(map[string]any{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": total,
		"subscribes":      list,
	})
}
