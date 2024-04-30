package admin

import (
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type MoneyLogController struct {
	controllers.BaseController
}

var money_menu_details = map[string]any{
	"field": []any{
		NewField("Id", "事件ID"),
		NewField("UserId", "发起用户ID"),
		NewField("Before", "变动前余额"),
		NewField("After", "变动后余额"),
		NewField("Amount", "变动金额"),
		NewField("Remark", "备注"),
		NewField("CreateTimeStr", "变动时间"),
	},
}

func (c *MoneyLogController) Index() {
	c.Data["details"] = money_menu_details
	c.TplName = "views/tabler/admin/log/money.tpl"
}

func (c *MoneyLogController) Ajax() {
	list := models.NewUserMoneyLog().FindAll()

	for i, v := range list {
		list[i].CreateTimeStr = time.UnixMilli(v.CreateTime).Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"money_logs": list,
	})
}
