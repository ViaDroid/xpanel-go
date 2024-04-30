package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "views/tabler/404.tpl"
}

func (c *ErrorController) Error405() {
	c.Data["content"] = "server error"
	c.TplName = "views/tabler/405.tpl"
}


func (c *ErrorController) Error500() {
	c.Data["content"] = "internal error"
	c.TplName = "views/tabler/500.tpl"
}