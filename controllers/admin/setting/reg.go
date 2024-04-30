package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type RegController struct {
	controllers.BaseController
}

var reg_update_field = []string{
	"reg_mode",
	"reg_email_verify",
	"reg_daily_report",
	"random_group",
	"min_port",
	"max_port",
	"reg_traffic",
	"free_user_reset_day",
	"free_user_reset_bandwidth",
	"reg_class",
	"reg_class_time",
	"reg_method",
	"reg_ip_limit",
	"reg_speed_limit",
}

func (c *RegController) Index() {
	settings := models.NewConfig().FindByClass("reg")

	c.Data["update_field"] = reg_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/reg.tpl"
}

func (c *RegController) Save() {
	for _, item := range reg_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}
