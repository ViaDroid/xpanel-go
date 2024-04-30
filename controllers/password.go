package controllers

type PasswordController struct {
	BaseController
}

func (c *PasswordController) Reset() {
	c.TplName = "views/tabler/password/reset.tpl"
}

func (c *PasswordController) HandleReset() {

}

func (c *PasswordController) Token() {

}

func (c *PasswordController) HandleToken() {

}
