package admin

import (
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DetectLogController struct {
	controllers.BaseController
}

var detect_log_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("UserId", "用户ID"),
		NewField("NodeId", "节点ID"),
		NewField("NodeName", "节点名"),
		NewField("ListId", "规则ID"),
		NewField("RuleName", "规则名"),
		NewField("DatetimeStr", "时间"),
	},
}

func (c *DetectLogController) Index() {
	c.Data["details"] = detect_log_menu_details
	c.TplName = "views/tabler/admin/log/detect.tpl"
}

func (c *DetectLogController) Ajax() {
	length, _ := c.GetInt("length")
	page, _ := c.GetInt("start")
	draw := c.GetString("draw")

	page = page/length + 1

	subLog := models.NewDetecLog()
	total, _ := subLog.Count()

	list := subLog.FetchList(length, page)

	for i, v := range list {
		list[i].DatetimeStr = time.UnixMilli(v.Datetime).Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": total,
		"logs":            list,
	})
}
