package admin

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type PaybackController struct {
	controllers.BaseController
}

var payback_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("Total", "原始金额"),
		NewField("UserId", "发起用户ID"),
		// TODO NewField("UserName", "发起用户名"),
		NewField("RefBy", "获利用户ID"),
		// NewField("RefUserName", "获利用户名"),
		NewField("RefGet", "获利金额"),
		NewField("InvoiceId", "账单ID"),
		NewField("DatetimeStr", "时间"),
	},
}

func (c *PaybackController) Index() {
	c.Data["details"] = payback_menu_details
	c.TplName = "views/tabler/admin/log/payback.tpl"
}

func (c *PaybackController) Ajax() {
	list := models.NewPayback().FindAll()

	for i, _ := range list {
		// TODO Payback list
		list[i].Parse()
		// $payback->datetime = Tools::toDateTime((int) $payback->datetime);
		//     $payback->user_name = $payback->getAttributes();
		//     $payback->ref_user_name = $payback->getAttributes();
		// onlines[i].LocationStr = v.Location()
		// list[i].FirstTimeStr = time.UnixMilli(v.FirstTime).Format("2006-01-02 15:04:05")
		// list[i].LastTimeStr = time.UnixMilli(v.LastTime).Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"paybacks": list,
	})
}
