package admin

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type PaylistController struct {
	controllers.BaseController
}

var paylist_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "用户ID"),
		NewField("Total", "金额"),
		NewField("StatusStr", "状态"),
		NewField("Gateway", "支付网关"),
		NewField("Tradeno", "网关单号"),
		NewField("Datetime", "支付时间"),
		NewField("InvoiceId", "关联账单ID"),
	},
}

func (c *PaylistController) Index() {
	c.Data["details"] = paylist_menu_details
	c.TplName = "views/tabler/admin/log/gateway.tpl"
}

func (c *PaylistController) Ajax() {
	list := models.NewPaylist().FindAll()

	for i, v := range list {
		list[i].StatusStr = v.ParseStatus()
		list[i].CreateTimeStr = v.Datetime.Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"paylists": list,
	})
}
