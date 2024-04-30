package webapi

import "github.com/viadroid/xpanel-go/controllers"

type FuncController struct {
	controllers.BaseController
}

func (c *FuncController) Ping() {
	c.Ctx.WriteString("Pong? Pong!")
}

func (c *FuncController) GetDetectRules() {

}
