package admin

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type SystemController struct {
	controllers.BaseController
}

func (c *SystemController) Index() {
	conf := models.NewConfig()
	last_daily_job_time := conf.Obtain("last_daily_job_time").Value
	db_version := conf.Obtain("db_version").Value
	version := conf.Obtain("version").Value

	c.Data["last_daily_job_time"] = last_daily_job_time
	c.Data["db_version"] = db_version
	c.Data["version"] = version
	c.TplName = "views/tabler/admin/system.tpl"
}

func (c *SystemController) CheckUpdate() {
	c.TplName = "admin/system/info.tpl"
}
