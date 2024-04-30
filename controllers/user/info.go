package user

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
	"github.com/viadroid/xpanel-go/utils"
)

type InfoController struct {
	controllers.BaseController
}

func (c *InfoController) Index() {

	// TODO
	// c.Data["user"] = c.User
	c.Data["themes"] = []string{}
	c.Data["methods"] = utils.GetSsMethod("method")
	c.Data["ga_url"] = services.NewMFA().GetGaUrl(c.User.GaToken, c.User.Email)
	c.TplName = "views/tabler/user/edit.tpl"
}

func (c *InfoController) UpdateEmail() {
	newEmail := c.GetString("newemail")

	oldEmail := c.User.Email

	if !global.ConfMap["enable_change_email"].(bool) || c.User.IsShadowBanned {
		c.Error(&c.Controller, "修改失败")
		return
	}

	if newEmail == "" {
		c.Error(&c.Controller, "未填写邮箱")
		return
	}

	if !utils.IsEmail(newEmail) {
		c.Error(&c.Controller, "无效的邮箱")
		return
	}

	if newEmail == oldEmail {
		c.Error(&c.Controller, "新邮箱不能和旧邮箱一样")
		return
	}

	if models.NewUser().IsExist(newEmail) {
		c.Error(&c.Controller, "邮箱已经被使用了")
		return
	}

	if models.NewConfig().Obtain("reg_email_verify").ValueToBool() {
		emailCode := c.GetString("emailcode")

		key := fmt.Sprintf("email_verify:%s", emailCode)
		if _, err := global.Redis.Get(c.Ctx.Request.Context(), key).Result(); err == redis.Nil || err != nil {
			c.Error(&c.Controller, "你的邮箱验证码不正确")
			return
		}

		global.Redis.Del(c.Ctx.Request.Context(), key)
	}

	c.User.Email = newEmail
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}

	c.SuccessWithData(&c.Controller, "修改成功", map[string]any{
		"email": newEmail,
	})

}

func (c *InfoController) UpdateUsername() {
	userName := c.GetString("newusername")

	if c.User.IsShadowBanned {
		c.Error(&c.Controller, "修改失败")
		return
	}

	if userName == "" {
		c.Error(&c.Controller, "未填写用户名")
		return
	}

	c.User.UserName = userName

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}
	c.SuccessWithData(&c.Controller, "修改成功", map[string]any{
		"username": userName,
	})
}

func (c *InfoController) UnbindIM() {
	if _, err := c.User.UnbindIM(); err != nil {
		c.Error(&c.Controller, "解绑失败")
		return
	}
	c.Ctx.Output.Header("HX-Refresh", "true")
	c.Success(&c.Controller, "解绑成功")
}

func (c *InfoController) UpdatePassword() {
	oldPassword := c.GetString("oldpassword")
	newPassword := c.GetString("new_password")
	confirmPassword := c.GetString("confirm_new_password")

	if oldPassword == "" || newPassword == "" || confirmPassword == "" {
		c.Error(&c.Controller, "密码不能为空")
		return
	}

	if utils.CheckPassword(c.User.Pass, oldPassword) {
		c.Error(&c.Controller, "旧密码错误")
		return
	}

	if newPassword != confirmPassword {
		c.Error(&c.Controller, "两次输入的密码不一致")
		return
	}

	if len(newPassword) < 8 {
		c.Error(&c.Controller, "密码太短啦")
		return
	}

	c.User.Pass = utils.PasswordHash(newPassword)

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}

	if models.NewConfig().Obtain("enable_forced_replacement").ValueToBool() {
		c.User.RemoveLink()
	}

	c.Success(&c.Controller, "修改成功")

}

func (c *InfoController) ResetPassword() {
	c.User.Passwd = utils.GenRandomString(16)
	c.Uuid = uuid.NewString()

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "重置失败")
		return
	}
	c.SuccessWithData(&c.Controller, "重置成功", map[string]any{
		"passwd": c.Passwd,
		"uuid":   c.Uuid,
	})
}

func (c *InfoController) ResetApiToken() {
	c.User.ApiToken = utils.GenRandomString(32)

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "重置失败")
		return
	}
	c.Success(&c.Controller, "重置成功")
}

func (c *InfoController) UpdateMethod() {

	method := c.GetString("method")
	if method == "" {
		c.Error(&c.Controller, "非法输入")
		return
	}

	if !slices.Contains(utils.GetSsMethod("method"), method) {
		c.Error(&c.Controller, "加密无效")
		return
	}

	c.User.Method = method

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "重置失败")
		return
	}
	c.Success(&c.Controller, "重置成功")

}

func (c *InfoController) ResetURL() {
	if _, err := c.User.RemoveLink(); err != nil {
		c.Error(&c.Controller, "重置失败")
		return
	}
	c.Success(&c.Controller, "重置成功")
}

func (c *InfoController) UpdateDailyMail() {
	value, err := c.GetInt("mail")
	if err != nil {
		c.Error(&c.Controller, "非法输入")
		return
	}

	if !slices.Contains([]int{0, 1, 2}, value) {
		c.Error(&c.Controller, "非法输入")
		return
	}

	c.User.DailyMailEnable = value

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}
	c.Success(&c.Controller, "修改成功")
}
func (c *InfoController) UpdateContactMethod() {
	value, err := c.GetInt("contact")
	if err != nil {
		c.Error(&c.Controller, "参数错误")
		return
	}

	if !slices.Contains([]int{1, 2}, value) {
		c.Error(&c.Controller, "参数错误")
		return
	}

	c.User.ContactMethod = value

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}
	c.Success(&c.Controller, "修改成功")
}

func (c *InfoController) UpdateTheme() {

	theme := c.GetString("theme")
	if theme == "" {
		c.Error(&c.Controller, "主题不能为空")
		return
	}

	c.User.Theme = theme

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}
	c.Ctx.Output.Header("HX-Refresh", "true")

	c.Success(&c.Controller, "修改成功")
}

func (c *InfoController) SendToGulag() {
	password := c.GetString("password")

	if password != "" || utils.CheckPassword(c.User.Pass, password) {
		c.Error(&c.Controller, "密码错误")
		return
	}

	if global.ConfMap["enable_gulag"].(bool) {
		services.NewAuthService().Logout(&c.Controller)

		c.User.Kill()
		c.Ctx.Output.Header("HX-Refresh", "true")

		c.Success(&c.Controller, "你将被送去古拉格接受劳动改造，再见")
		return
	}
	c.Error(&c.Controller, "自助账号删除未启用")
}
