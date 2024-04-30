package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type CronController struct {
	controllers.BaseController
}

var cron_update_field = []string{
	"daily_job_hour",
	"daily_job_minute",
	"enable_daily_finance_mail",
	"enable_weekly_finance_mail",
	"enable_monthly_finance_mail",
	"enable_detect_gfw",
	"enable_detect_ban",
	"enable_detect_inactive_user",
	"detect_inactive_user_checkin_days",
	"detect_inactive_user_login_days",
	"detect_inactive_user_use_days",
	"remove_inactive_user_link_and_invite",
}

func (c *CronController) Index() {
	settings := models.NewConfig().FindByClass("cron")

	c.Data["update_field"] = cron_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/cron.tpl"
}

func (c *CronController) Save() {
	daily_job_hour, _ := c.GetInt("daily_job_hour")
	daily_job_minute, _ := c.GetInt("daily_job_minute")

	if daily_job_hour < 0 || daily_job_hour > 23 {
		c.Error(&c.Controller, "每日任务执行时间的小时数必须在 0-23 之间")
		return
	}
	if daily_job_minute < 0 || daily_job_minute > 59 {
		c.Error(&c.Controller, "每日任务执行时间的分钟数必须在 0-59 之间")
		return
	}

	for _, item := range cron_update_field {
		if item == "daily_job_minute" {
			value := fmt.Sprintf("%d", daily_job_minute-(daily_job_minute%5))
			if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
				c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
				return 
			}
			
			continue
		}
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}

	c.Success(&c.Controller, "保存成功")

}
