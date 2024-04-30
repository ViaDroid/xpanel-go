package controllers

import (
	"fmt"
	"time"

	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {
	conf := models.NewConfig()
	var captcha map[string]string

	class_expire_days := 0
	if c.User.Class > 0 {
		expireTime, _ := time.Parse(time.DateTime, c.User.ClassExpire)
		class_expire_days = int(time.Until(expireTime).Milliseconds() / 86400)
	}

	if conf.Obtain("enable_reg_captcha").ValueIsOne() {
		captcha = services.NewCaptcha().Generate()
	}
	ann, _ := models.NewAnn().FetchOne()
	c.Data["ann"] = ann
	c.Data["captcha"] = captcha
	c.Data["class_expire_days"] = class_expire_days
	c.Data["UniversalSub"] = services.NewSubscribeService().GetUniversalSubLink(c.User)
	c.TplName = "views/tabler/user/index.tpl"
}

func (c *UserController) Profile() {

	loginIpList, err := models.NewLoginIp().FetchListByUserId(c.User.Id)

	if err != nil {
		loginIpList = []models.LoginIp{}
	}

	logList, err := models.NewOnlineLog().FetchListByUserId(c.User.Id)

	if err != nil {
		logList = []models.OnlineLog{}
	}

	c.Data["logins"] = loginIpList
	c.Data["ips"] = logList

	c.TplName = "views/tabler/user/profile.tpl"
}

func (c *UserController) Announcement() {
	anns := models.NewAnn().FetchList()
	c.Data["anns"] = anns
	c.TplName = "views/tabler/user/announcement.tpl"
}

func (c *UserController) CheckIn() {
	conf := models.NewConfig()

	if !conf.Obtain("enable_checkin").ValueIsOne() || !c.User.IsAbleToCheckin() {
		c.Error(&c.Controller, "暂时还不能签到")
		return
	}
	if conf.Obtain("enable_checkin_captcha").ValueIsOne() {
		if !services.NewCaptcha().Verify(c.ParamsMap()) {
			c.Error(&c.Controller, "系统无法接受你的验证结果，请刷新页面后重试")
			return
		}
	}

	traffic, err := services.NewRewardService().IssueCheckinReward(c.User.Id)
	if err != nil {
		c.Error(&c.Controller, "签到失败")
		return
	}

	res := map[string]any{
		"ret":  1,
		"msg":  fmt.Sprintf("获得了 %d MB 流量", traffic),
		"data": map[string]any{"last-checkin-time": time.Now().Format(time.DateTime)},
	}
	c.JSONResp(res)
}

func (c *UserController) SwitchThemeMode() {
	if c.User.IsDarkMode == 1 {
		c.User.IsDarkMode = 0
	} else {
		c.User.IsDarkMode = 1
	}

	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "切换失败")
		return
	}

	c.Ctx.Output.Header("HX-Refresh", "true")
	c.Success(&c.Controller, "切换成功")
}

func (c *UserController) Banned() {

	c.Data["banned_reason"] = c.User.BannedReason
	c.TplName = "view/tabler/user/banned.tpl"

}

func (c *UserController) Logout() {
	services.NewAuthService().Logout(&c.Controller)

	c.Redirect("/", 302)
}
