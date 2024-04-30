package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type RefController struct {
	controllers.BaseController
}

var ref_update_field = []string{
	"invite_reg_money_reward",
	"invite_reg_traffic_reward",
	"invite_mode",
	"invite_reward_mode",
	"invite_reward_rate",
	"invite_reward_count_limit",
	"invite_reward_total_limit",
}

func (c *RefController) Index() {
	settings := models.NewConfig().FindByClass("ref")

	c.Data["update_field"] = ref_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/ref.tpl"
}

func (c *RefController) Save() {
	for _, item := range ref_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}
