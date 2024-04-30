package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/viadroid/xpanel-go/services"
)

func UserAuthHandler(ctx *context.Context) {
	user := services.NewAuthService().GetUser2(ctx)

	if user == nil {
		ctx.Redirect(302, "/auth/login")
		return
	}

	// ctx.Output.Session("user", user)

	// _, ok := ctx.Input.Session("uid").(int)
	// if !ok && ctx.Request.RequestURI != "/login" {
	//     ctx.Redirect(302, "/login")
	// }
}

func AdminAuthHandler(ctx *context.Context) {
	user := services.NewAuthService().GetUser2(ctx)
	if user == nil {
		ctx.Redirect(302, "/auth/login")
		return
	}

	if !user.IsAdmin {
		ctx.Redirect(302, "/user")
		return
	}

}

func GuestHandler(ctx *context.Context) {
	user := services.NewAuthService().GetUser2(ctx)

	if user != nil {
		// ctx.Output.Header("Location", "/user")
		ctx.Redirect(302, "/user")
		return
	}
}
