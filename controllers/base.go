package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
	"github.com/viadroid/xpanel-go/utils"
)

type BaseController struct {
	web.Controller
	utils.ResponseHelper
	*models.User
}

func (c *BaseController) Prepare() {
	c.Data["public_setting"] = global.SettingMap

	c.Data["config"] = global.ConfMap

	c.User = services.NewAuthService().GetUser2(c.Ctx)

	c.Data["isLogin"] = c.User != nil
	c.Data["user"] = c.User

}

func (c *BaseController) ParamsMap() map[string]any {
	params := map[string]any{}
	for k, v := range c.Ctx.Request.Form {
		params[k] = v[0]
	}
	return params
}
