package controllers

import "github.com/viadroid/xpanel-go/models"

type CallbackController struct {
	BaseController
}

func (c *CallbackController) Index() {

}

func (c *CallbackController) Telegram() {
	conf := models.NewConfig()
	token := c.GetString("token")
	if conf.Obtain("enable_telegram").ValueToBool() && token == conf.ObtainValue("telegram_request_token") {
		// TODO             Telegram::process($request);

		c.Ctx.Output.Status = 204
		return
	}
	c.Ctx.Output.Status = 400
}
