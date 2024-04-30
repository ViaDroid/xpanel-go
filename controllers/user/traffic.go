package user

import "github.com/viadroid/xpanel-go/controllers"

type TrafficLogController struct {
	controllers.BaseController
}

func (c *TrafficLogController) Index() {
	c.TplName = "user/traffic_log.tpl"
}
