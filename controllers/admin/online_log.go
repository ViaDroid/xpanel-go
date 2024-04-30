package admin

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type OnlineLogController struct {
	controllers.BaseController
}

var online_log_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "用户ID"),
		NewField("NodeId", "节点ID"),
		NewField("NodeName", "节点名"),
		NewField("Ip", "IP"),
		NewField("LocationStr", "IP归属地"),
		NewField("FirstTimeStr", "首次连接"),
		NewField("LastTimeStr", "最后连接"),
	},
}

func (c *OnlineLogController) Index() {
	c.Data["details"] = online_log_menu_details
	c.TplName = "views/tabler/admin/log/online.tpl"
}

func (c *OnlineLogController) Ajax() {
	length, _ := c.GetInt("length")
	page, _ := c.GetInt("start")
	draw := c.GetString("draw")

	page = page/length + 1

	onlineLog := models.NewOnlineLog()
	total, _ := onlineLog.Count()

	onlines := onlineLog.FetchList(length, page)

	for i := range onlines {
		onlines[i].Parse()
	}

	c.JSONResp(map[string]any{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": total,
		"onlines":         onlines,
	})
}
