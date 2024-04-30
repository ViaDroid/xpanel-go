package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type FeatureController struct {
	controllers.BaseController
}

var feature_update_field = []string{
	"display_detect_log",
	"display_docs",
	"display_docs_only_for_paid_user",
	"traffic_log",
	"traffic_log_retention_days",
	"subscribe_log",
	"subscribe_log_retention_days",
	"notify_new_subscribe",
	"login_log",
	"notify_new_login",
	"enable_checkin",
	"checkin_min",
	"checkin_max",
}

func (c *FeatureController) Index() {
	settings := models.NewConfig().FindByClass("feature")

	c.Data["update_field"] = feature_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/feature.tpl"
}

func (c *FeatureController) Save() {
	for _, item := range feature_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}
