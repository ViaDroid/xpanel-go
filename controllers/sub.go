package controllers

import (
	"fmt"
	"slices"
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type SubController struct {
	BaseController
}

func (c *SubController) Index() {
	errMsg := "订阅链接无效"
	typ := c.Ctx.Input.Param(":type")
	subtype := c.Ctx.Input.Param(":subtype")
	subtype_list := []string{"json", "clash", "sip008", "singbox", "v2rayjson", "sip002", "ss", "v2ray", "trojan"}

	if !global.ConfMap["Subscribe"].(bool) ||
		!slices.Contains(subtype_list, subtype) ||
		global.ConfMap["subUrl"].(string) != fmt.Sprintf("https://%s/sub/%s/%s", c.Ctx.Request.Host, typ, subtype) {
		c.Error(&c.Controller, errMsg)
		return
	}

	token := c.GetString("token")
	rateLimit := services.NewRateLimit()
	if global.ConfMap["enable_rate_limit"].(bool) &&
		rateLimit.CheckRateLimit(services.RateLimitTypes.SUB_IP, c.Ctx.Request.RemoteAddr) ||
		rateLimit.CheckRateLimit(services.RateLimitTypes.SUB_TOKEN, token) {
		c.Error(&c.Controller, errMsg)
		return
	}

	link, err := models.NewLink().FindByToken(token)
	if err != nil || !link.IsValid() {
		c.Error(&c.Controller, errMsg)
		return
	}

	user := models.NewUser().FindById(link.UserId)

	sub_info := services.NewSubscribeService().GetContent(c.User)

	content_type := func(typ string) string {
		switch typ {
		case "clash":
			return "application/yaml"
		case "json", "sip008", "singbox", "v2rayjson":
			return "application/json"
		default:
			return "text/plain"
		}
	}(subtype)

	sub_details := fmt.Sprintf(" upload=%d; download=%d; total=%d; expire=%s",
		c.User.U, c.User.D, c.User.TransferEnable, c.User.ClassExpire)

	if models.NewConfig().Obtain("subscribe_log").ValueToBool() {
		requestIp := c.Ctx.Input.IP()
		ua := c.Ctx.Request.UserAgent()
		if _, err := models.NewSubscribeLog().Add(user, subtype, requestIp, ua); err != nil {
			return
		}

		isExist := global.DB.QueryTable(models.NewSubscribeLog()).
			Filter("user_id", user.Id).
			Filter("request_ip__like", "%"+requestIp+"%").Exist()
		if models.NewConfig().Obtain("notify_new_subscribe").ValueToBool() && !isExist {
			title := fmt.Sprintf("%s-新订阅通知", global.ConfMap["appName"])
			msg := fmt.Sprintf("你的账号于 %s 通过 %s 地址订阅了新的节点", time.Now().Format(time.DateTime), requestIp)
			services.NewNotification().NotifyUser(user, title, msg, "")
		}
	}

	c.Ctx.Output.Header("Subscription-Userinfo", sub_details)
	c.Ctx.Output.Header("Content-Type", content_type)
	c.Ctx.WriteString(sub_info)
}
