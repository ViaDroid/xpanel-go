package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/server/web"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type AuthService struct{}

// NewAuthService returns a new AuthService instance
func NewAuthService() *AuthService {
	return &AuthService{}
}

// 492809000 // 126944800 //604800
func (s *AuthService) Login(c *web.Controller, userId int, timeGap int) {
	user := models.NewUser().FindById(userId)

	now := time.Now()

	d := now.Add(time.Second * time.Duration(timeGap))
	fmt.Printf("%v", d)

	expire_in := now.Add(time.Second * time.Duration(timeGap)).UnixMilli()
	domain := c.Ctx.Input.Host()
	remoteIp := c.Ctx.Input.IP()
	ua := c.Ctx.Input.UserAgent()

	claims := utils.NewAuthClaims(userId, user.Email, utils.CookieHash(
		user.Pass, expire_in),
		utils.IpHash(remoteIp, userId, expire_in),
		utils.DeviceHash(ua, userId, expire_in),
		expire_in)

	c.SetSession("claims", claims.ToJSON())

	setCookieWithDomain(
		c,
		map[string]any{
			"uid":       fmt.Sprintf("%d", userId),
			"email":     user.Email,
			"key":       utils.CookieHash(user.Pass, expire_in),
			"ip":        utils.IpHash(remoteIp, userId, expire_in),
			"device":    utils.DeviceHash(ua, userId, expire_in),
			"expire_in": fmt.Sprintf("%d", expire_in),
		},
		expire_in, domain)
}

func (s *AuthService) GetUser2(c *context.Context) *models.User {
	claimsStr, ok := c.Input.Session("claims").(string)
	if !ok {
		return nil
	}

	claims := &utils.AuthClaims{}
	err := claims.FromJSON(claimsStr)
	if err != nil {
		web.GlobalSessions.SessionDestroy(c.Input.Context.ResponseWriter, c.Request)
		return nil
	}

	if !claims.IsValid() || claims.IsExpired() {
		web.GlobalSessions.SessionDestroy(c.Input.Context.ResponseWriter, c.Request)
		return nil
	}

	if global.ConfMap["enable_login_bind_ip"].(bool) {
		remoteIp := c.Input.IP()
		_, err := models.NewNode().FindIp4Or6(remoteIp)

		if err != nil && claims.Ip != utils.IpHash(remoteIp, claims.Uid, claims.ExpireIn) {
			web.GlobalSessions.SessionDestroy(c.Input.Context.ResponseWriter, c.Request)
			return nil
		}
	}

	if global.ConfMap["enable_login_bind_device"].(bool) {
		ua := c.Input.UserAgent()
		if claims.Device != utils.DeviceHash(ua, claims.Uid, claims.ExpireIn) {
			web.GlobalSessions.SessionDestroy(c.Input.Context.ResponseWriter, c.Request)
			return nil
		}

	}

	user := models.NewUser().FindById(claims.Uid)

	if user.Id == 0 {
		web.GlobalSessions.SessionDestroy(c.Input.Context.ResponseWriter, c.Request)
		return nil
	}

	return user
}

func (s *AuthService) GetUser(c *context.Context) *models.User {
	uid := c.Input.Cookie("uid")
	email := c.Input.Cookie("email")
	key := c.Input.Cookie("key")
	ipHash := c.Input.Cookie("ip")
	deviceHash := c.Input.Cookie("device")
	expire_in := c.Input.Cookie("expire_in")

	if uid == "" || email == "" || key == "" || ipHash == "" || deviceHash == "" || expire_in == "" {
		return nil
	}

	userId, err := strconv.Atoi(uid)
	if err != nil {
		return nil
	}

	exp, err := strconv.ParseInt(expire_in, 10, 64)
	if err != nil {
		return nil
	}
	if exp < time.Now().UnixMilli() {
		return nil
	}

	if global.ConfMap["enable_login_bind_ip"].(bool) {
		remoteIp := c.Input.IP()
		node, _ := models.NewNode().FindIp4Or6(remoteIp)

		if node == nil && ipHash != utils.IpHash(remoteIp, userId, exp) {
			return nil
		}
	}

	if global.ConfMap["enable_login_bind_device"].(bool) {
		ua := c.Input.UserAgent()
		if deviceHash != utils.DeviceHash(ua, userId, exp) {
			return nil
		}

	}
	user := models.NewUser()

	user.FindById(userId)

	if user.Id == 0 || user.Email != email || key != utils.CookieHash(user.Pass, exp) {
		return nil
	}
	return user
}

func (s *AuthService) Logout(c *web.Controller) {
	domain := c.Ctx.Input.Host()

	c.DestroySession()
	// web.GlobalSessions.SessionDestroy(c.Ctx.Input.Context.ResponseWriter, c.Ctx.Request)

	setCookieWithDomain(
		c,
		map[string]any{
			"uid":       "",
			"email":     "",
			"key":       "",
			"ip":        "",
			"device":    "",
			"expire_in": "",
		},
		0, domain)
}

/*
 * @param int $type 1 = failed, 0 = success
 */
func (s *AuthService) CollectLoginIP(ip string, loginType int, userId int) {
	conf := models.NewConfig()

	item := models.NewLoginIp()

	if conf.Obtain("login_log").ValueToBool() {
		item.Ip = ip
		item.UserId = userId
		item.Type = loginType
		item.Datetime = time.Now().UnixMilli()

		if conf.Obtain("notify_new_login").ValueToBool() && userId != 0 && !item.IsExist(userId) {
			title := fmt.Sprintf("%s-新登录通知", global.ConfMap["appName"])
			msg := fmt.Sprintf("你的账号于 %s 通过 %s 地址登录了用户面板", time.Now().Format(time.DateTime), ip)
			user := models.NewUser().FindById(userId)
			NewNotification().NotifyUser(user, title, msg, "")
		}
	}

	_, err := global.DB.Insert(item)
	if err != nil {
		println(err)
	}
}

func setCookieWithDomain(c *web.Controller, m map[string]any, t int64, domain string) {
	for k, v := range m {
		c.Ctx.SetCookie(k, fmt.Sprintf("%s", v), t, "/", domain, true, true)
		// c.Ctx.Input.Context.SetCookie(k, fmt.Sprintf("%s", v), t, "/", domain, true, true)
	}
}
