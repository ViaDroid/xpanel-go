package admin

import (
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DetectBanLogController struct {
	controllers.BaseController
}

var detect_ban_log_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "用户ID"),
		NewField("DetectNumber", "违规次数"),
		NewField("BanTime", "封禁时长(分钟)"),
		NewField("StartTime", "统计开始时间"),
		NewField("EndTime", "统计结束&封禁开始时间"),
		NewField("BanEndTime", "封禁结束时间"),
		NewField("AllDetectNumber", "累计违规次数"),
	},
}

func (c *DetectBanLogController) Index() {
	c.Data["details"] = detect_ban_log_menu_details
	c.TplName = "views/tabler/admin/log/detect_ban.tpl"
}

func (c *DetectBanLogController) Ajax() {
	length, _ := c.GetInt("length")
	page, _ := c.GetInt("start")
	draw := c.GetString("draw")

	page = page/length + 1

	subLog := models.NewDetecBanLog()
	total, _ := subLog.Count()

	list := subLog.FetchList(length, page)

	for i, v := range list {
		list[i].StartTimeStr = time.UnixMilli(v.StartTime).Format("2006-01-02 15:04:05")
		list[i].EndTimeStr = time.UnixMilli(v.EndTime).Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": total,
		"logs":            list,
	})
}
