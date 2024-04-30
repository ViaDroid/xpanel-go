package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
	"github.com/viadroid/xpanel-go/utils"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Login() {
	var captcha map[string]string
	regCaptchaConf := models.Config{}.Obtain("enable_reg_captcha")
	if regCaptchaConf.ValueIsOne() {
		captcha = services.NewCaptcha().Generate()
	}
	c.Data["captcha"] = captcha
	c.TplName = "views/tabler/auth/login.tpl"
}

func Login2(ctx *context.Context) {
	ctx.WriteString("hello, world")
}

func (c *AuthController) LoginHandle() {

	conf := models.NewConfig()
	if conf.Obtain("enable_login_captcha").ValueIsOne() && !services.NewCaptcha().Verify(c.ParamsMap()) {
		c.Error(&c.Controller, "系统无法接受你的验证结果，请刷新页面后重试。")
		// return
	}

	//form
	//email passwd mfa_code remember_me
	email := c.GetString("email")
	password := c.GetString("passwd")
	mfaCode := c.GetString("mfa_code")
	rememberMe, _ := c.GetBool("remember_me", false)

	redir := c.Ctx.GetCookie("redir")
	if redir == "" {
		redir = "/user"
	}

	user := models.NewUser()
	authService := services.NewAuthService()

	remoteIp := c.Ctx.Input.IP()
	if _, err := user.FindByEmail(email); err != nil {
		authService.CollectLoginIP(remoteIp, 1, 0)
		c.Error(&c.Controller, "邮箱或者密码错误")
		return
	}

	if !utils.CheckPassword(user.Pass, password) {
		authService.CollectLoginIP(remoteIp, 1, user.Id)
		c.Error(&c.Controller, "邮箱或者密码错误")
		return
	}

	if user.GaEnable && len(mfaCode) != 6 /*|| */ {
		authService.CollectLoginIP(remoteIp, 1, user.Id)
		c.Error(&c.Controller, "两步验证码错误")
		return
	}

	timeGap := 3600

	if rememberMe {
		d, ok := global.ConfMap["rememberMeDuration"].(int)
		if !ok {
			d = 7
		}
		timeGap = 86400 * d
	}

	authService.Login(&c.Controller, user.Id, timeGap)

	authService.CollectLoginIP(remoteIp, 0, user.Id)

	user.LastLoginTime = time.Now().UnixMilli()

	user.Update()

	c.Ctx.Output.Header("HX-Redirect", redir)

	c.Success(&c.Controller, "登录成功")
	// c.Redirect("/user", 302)
}

func (c *AuthController) Register() {
	c.Data["invite_code"] = c.GetString("code", "")
	var captcha map[string]string
	regCaptchaConf := models.Config{}.Obtain("enable_reg_captcha")
	if regCaptchaConf.ValueIsOne() {
		captcha = services.NewCaptcha().Generate()
	}

	c.Data["captcha"] = captcha
	c.TplName = "views/tabler/auth/register.tpl"
}

func RegisterHelper(c *BaseController, name, email, passwd, inviteCode string, imtype int, imvalue string, money int, is_admin_reg bool) {
	redir := c.Ctx.GetCookie("redir")
	if redir == "" {
		redir = "/user"
	}
	conf := models.NewConfig()
	configs := conf.FindByClass("reg")
	// do reg user
	user := models.NewUser()

	user.UserName = name
	user.Email = email
	user.Remark = ""
	user.Pass = utils.PasswordHash(passwd)
	user.Passwd = utils.GenRandomString(16)
	user.Uuid = uuid.NewString()
	user.ApiToken = uuid.NewString()
	user.Port = utils.GetSsPort()
	user.U = 0
	user.D = 0
	user.Method = configs["reg_method"].(string)
	user.ForbiddenIp = conf.ObtainValue("reg_forbidden_ip")
	user.ForbiddenPort = conf.ObtainValue("reg_forbidden_port")
	user.ImType = imtype
	user.ImValue = imvalue
	user.TransferEnable = utils.ToGB(configs["reg_traffic"].(int))
	user.AutoResetDay = conf.Obtain("free_user_reset_day").ValueToInt()
	user.AutoResetBandwidth = float64(conf.Obtain("free_user_reset_bandwidth").ValueToInt())
	if configs["reg_daily_report"].(bool) {
		user.DailyMailEnable = 1
	}

	if money > 0 {
		user.Money = float64(money)
	} else {
		user.Money = 0
	}

	user.RefBy = 0

	if inviteCode != "" {
		invite := models.NewInviteCode().FindByCode(inviteCode)
		if invite != nil {
			user.RefBy = invite.UserId

		}
	}

	now := time.Now()

	user.GaToken = services.NewMFA().GenerateGaToken()
	user.GaEnable = false
	user.Class = configs["reg_class"].(int)

	gap := configs["reg_class_time"].(int) * 86400
	user.ClassExpire = now.Add(time.Duration(gap)).Format(time.DateTime)
	user.NodeIplimit = configs["reg_ip_limit"].(int)
	speedLimit := configs["reg_speed_limit"]
	user.NodeSpeedlimit = float64(speedLimit.(int))
	user.RegDate = now.Format(time.DateTime)

	realIp := c.Ctx.Request.Header.Get("X-Real-ip")
	realIp = c.Ctx.Input.IP()
	user.RegIp = realIp
	user.Theme = global.ConfMap["theme"].(string)
	user.Locale = global.ConfMap["locale"].(string)
	random_group := conf.Obtain("random_group").Value

	if random_group == "" {
		user.NodeGroup = 0
	} else {
		user.NodeGroup = utils.GetRandomGroup(random_group)
	}

	if _, err := user.Save(); err != nil {
		c.Error(&c.Controller, "未知错误")
		return
	}

	if !is_admin_reg {
		if user.RefBy != 0 {
			services.NewRewardService().IssueRegReward(user.Id, user.RefBy)
		}
		expire_in := now.Add(3600)
		domain := c.Ctx.Input.Host()

		// Auth::login(user.ID, 3600);
		c.Ctx.SetCookie("uid", strconv.Itoa(user.Id), expire_in, "/", domain, true, true)
		c.Ctx.SetCookie("email", user.Email, expire_in, "/", domain, true, true)
		// (new LoginIp())->collectLoginIP(realIp, 0, user.Id);

		c.Ctx.Output.Header("HX-Redirect", redir)
		c.Success(&c.Controller, "注册成功！正在进入登录界面")
		return
	}
}

func (c *AuthController) RegisterHandle() {
	conf := models.NewConfig()

	if conf.ObtainValue("reg_mode") == "colse" {
		c.Error(&c.Controller, "未开放注册。")
		return
	}

	if conf.Obtain("enable_reg_captcha").ValueIsOne() && !services.NewCaptcha().Verify(c.ParamsMap()) {
		c.Error(&c.Controller, "系统无法接受你的验证结果，请刷新页面后重试。")
		return
	}

	tos, _ := c.GetBool("tos")
	email := strings.ToLower(c.GetString("email"))
	name := c.GetString("name")
	passwd := c.GetString("passwd")
	repasswd := c.GetString("repasswd")
	inviteCode := c.GetString("invite_code")

	if !tos {
		c.Error(&c.Controller, "请同意服务条款")
		return
	}
	if inviteCode == "" && conf.ObtainValue("reg_mode") == "invite" {
		c.Error(&c.Controller, "邀请码不能为空")
		return
	}

	if inviteCode != "" {
		invite := models.NewInviteCode().FindByCode(inviteCode)
		if invite == nil {
			c.Error(&c.Controller, "邀请码无效")
			return
		}

		if user := models.NewUser().FindById(invite.UserId); user == nil {
			c.Error(&c.Controller, "邀请码无效")
			return
		}
	}

	var imtype int
	var imvalue string

	// check email format
	checkRes := utils.IsEmailLegal(email)
	if checkRes["ret"] == 0 {
		c.JSONResp(checkRes)
		return
	}
	// check email
	isExist := models.NewUser().IsExist(email)
	if isExist {
		c.Error(&c.Controller, "邮箱已经被注册了")
		return
	}

	// check pwd length
	if len(passwd) < 8 {
		c.Error(&c.Controller, "密码请大于8位")
		return
	}
	// check pwd re
	if passwd != repasswd {
		c.Error(&c.Controller, "两次密码输入不符")
		return
	}

	if conf.Obtain("reg_email_verify").ValueIsOne() {
		email_verify_code := c.GetString("emailcode")
		key := fmt.Sprintf("email_verify:%s", email_verify_code)
		if _, err := global.Redis.Get(c.Ctx.Request.Context(), key).Result(); err == redis.Nil || err != nil {
			c.Error(&c.Controller, "你的邮箱验证码不正确")
			return
		}

		global.Redis.Del(c.Ctx.Request.Context(), key)

	}

	RegisterHelper(&c.BaseController, name, email, passwd, inviteCode, imtype, imvalue, 0, false)

}

func (c *AuthController) SendVerify() {
	conf := models.NewConfig()

	if conf.Obtain("reg_email_verify").ValueIsOne() {
		email := strings.ToLower(c.GetString("email"))

		if email == "" {
			c.Error(&c.Controller, "未填写邮箱")
			return
		}

		// check email format
		checkRes := utils.IsEmailLegal(email)
		if checkRes["ret"] == 0 {
			c.JSONResp(checkRes)
			return
		}

		ip := c.Ctx.Input.IP()
		rateLimit := services.NewRateLimit()
		if !services.NewRateLimit().CheckEmailIpLimit(ip) ||
			!rateLimit.CheckEmailAddressLimit(email) {
			c.Error(&c.Controller, "你的请求过于频繁，请稍后再试")
			return
		}

		_, err := models.NewUser().FindByEmail(email)
		if err == nil {
			c.Error(&c.Controller, "此邮箱已经注册")
			return
		}

		email_code := utils.GenRandomString(6)

		ttl := conf.Obtain("email_verify_code_ttl").ValueToInt()

		global.Redis.SetEx(
			c.Ctx.Request.Context(),
			fmt.Sprintf("email_verify:%s", email_code),
			email,
			time.Duration(ttl),
		)

		subject := fmt.Sprintf("%s - 验证邮件", global.ConfMap["appName"])
		template := "verify_code.tpl"
		body := map[string]any{
			"code":   email_code,
			"expire": time.Now().Add(time.Second * time.Duration(ttl)).Format(time.DateTime),
		}
		err = services.NewMailService().Send(email, subject, template, body)
		if err != nil {
			c.Error(&c.Controller, "邮件发送失败，请联系网站管理员。")
			return
		}

		c.Success(&c.Controller, "验证码发送成功，请查收邮件。")
	}

	c.Error(&c.Controller, "站点未启用邮件验证")
}

func (c *AuthController) Logout() {
	services.NewAuthService().Logout(&c.Controller)
	c.Ctx.Output.Header("Location", "/auth/login")
	c.Redirect("/auth/login", 302)
}
