package user

import "github.com/viadroid/xpanel-go/controllers"

type SubLogController struct {
	controllers.BaseController
}

func (c *SubLogController) Index() {
	c.TplName = "user/sub_log.html"
}
