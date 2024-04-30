package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type CaptchaController struct {
	controllers.BaseController
}

var captcha_update_field = []string{
	"captcha_provider",
	"enable_reg_captcha",
	"enable_login_captcha",
	"enable_checkin_captcha",
	"enable_reset_password_captcha",
	// Turnstile
	"turnstile_sitekey",
	"turnstile_secret",
	// Geetest
	"geetest_id",
	"geetest_key",
	// hCaptcha
	"hcaptcha_sitekey",
	"hcaptcha_secret",
}

func (c *CaptchaController) Index() {
	settings := models.NewConfig().FindByClass("captcha")

	c.Data["update_field"] = captcha_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/captcha.tpl"
}

func (c *CaptchaController) Save() {

	for _, item := range captcha_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}

	c.Success(&c.Controller, "保存成功")
}
