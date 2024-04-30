package user

import (
	"slices"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
)

type ClientController struct {
	controllers.BaseController
}

func (c *ClientController) GetClients() {
	name := c.Ctx.Input.Param(":name")

	if name == "" || !global.ConfMap["enable_r2_client_download"].(bool) {
		c.Abort("404")
		return
	}

	clients := []string{
		"Clash.Verge.exe",
		"Clash.Verge_aarch64.dmg",
		"Clash.Verge.AppImage.tar.gz",
		"CMFA.apk",
		"SFA.apk",
		"SFM.zip",
	}

	if !slices.Contains(clients, name) {
		c.Abort("404")
		return
	}
	// TODO genR2PresignedUrl
	presignedUrl := ""
	c.Ctx.Output.Header("Location", presignedUrl)
	c.Redirect(presignedUrl, 302)
}
