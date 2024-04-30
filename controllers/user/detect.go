package user

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DetectController struct {
	controllers.BaseController
}

func (c *DetectController) Rule() {
	rules := models.NewDetecRule().FetchAll()
	c.Data["rules"] = rules
	c.TplName = "views/tabler/user/detect/index.tpl"
}

func (c *DetectController) Log() {
	if !models.NewConfig().Obtain("display_detect_log").ValueToBool() {
		c.Redirect("/user", 302)
		return
	}

	logs := models.NewDetecLog().FetchAllByUserId(c.User.Id)
	for i, _ := range logs {
		logs[i].Parse()
	}

	c.Data["logs"] = logs
	c.TplName = "views/tabler/user/detect/log.tpl"
}
