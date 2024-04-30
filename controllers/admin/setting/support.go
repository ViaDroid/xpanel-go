package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type SupportController struct {
	controllers.BaseController
}

var support_update_field = []string{
	"live_chat",
	"crisp_id",
	"livechat_license",
	// Ticket
	"enable_ticket",
	"mail_ticket",
	"ticket_limit",
}

func (c *SupportController) Index() {
	settings := models.NewConfig().FindByClass("support")

	c.Data["update_field"] = support_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/support.tpl"
}

func (c *SupportController) Save() {
	for _, item := range support_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}
