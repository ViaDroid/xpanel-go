package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type SubController struct {
	controllers.BaseController
}

var sub_update_field = []string{
	"enable_forced_replacement",
	"enable_ss_sub",
	"enable_v2_sub",
	"enable_trojan_sub",
}

func (c *SubController) Index() {
	settings := models.NewConfig().FindByClass("subscribe")

	c.Data["update_field"] = sub_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/sub.tpl"
}

func (c *SubController) Save() {
	for _, item := range sub_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}
