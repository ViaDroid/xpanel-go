package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	c.TplName = "views/tabler/index.tpl"
}

func (c *HomeController) Tos() {
	c.TplName = "views/tabler/tos.tpl"

}

func (c *HomeController) Staff() {
	if c.User != nil {
		c.TplName = "views/tabler/staff.tpl"
	} else {
		c.TplName = "views/tabler/404.tpl"
	}
}
