package user

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/services"
)

type MFAController struct {
	controllers.BaseController
}

func (c *MFAController) CheckGa() {
	code := c.GetString("code")
	if code == "" {
		c.Error(&c.Controller, "二维码不能为空")
		return
	}

	if !services.NewMFA().ValidateGaToken(c.User.GaToken, code) {
		c.Error(&c.Controller, "验证码错误")
		return
	}
	c.Success(&c.Controller, "验证成功")
}

func (c *MFAController) SetGa() {
	enable := c.GetString("enable")
	if enable == "" {
		c.Error(&c.Controller, "选项无效")
		return
	}

	c.User.GaEnable = enable == "1"
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "设置失败")
		return
	}

	c.Success(&c.Controller, "设置成功")
}

func (c *MFAController) ResetGa() {
	c.User.GaToken = services.NewMFA().GenerateGaToken()
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "重置失败")
		return
	}

	c.SuccessWithData(&c.Controller, "重置成功", map[string]any{
		"ga-token": c.User.GaToken,
		"ga-url":   services.NewMFA().GetGaUrl(c.User.GaToken, c.User.Email),
	})
}
