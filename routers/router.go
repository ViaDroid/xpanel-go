package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/controllers/admin"
	"github.com/viadroid/xpanel-go/controllers/admin/setting"
	"github.com/viadroid/xpanel-go/controllers/user"
	"github.com/viadroid/xpanel-go/middleware"
)

func init() {
	// set static resourses route
	web.SetStaticPath("/images/", "static/images")
	web.SetStaticPath("/assets/", "static/assets")
	web.SetStaticPath("/js/", "static/js")

	web.Router("/beego", &controllers.MainController{})

	// Home
	web.CtrlGet("/", (*controllers.HomeController).Index)
	web.CtrlGet("/tos", (*controllers.HomeController).Tos)
	web.CtrlGet("/staff", (*controllers.HomeController).Staff)

	// Error Page
	// web.ErrorController(&controllers.ErrorController{})

	// Bot Callback
	web.CtrlGet("/callback/:type", (*controllers.HomeController).Index)
	// OAuth
	web.CtrlAny("/oauth/:type", (*controllers.OAuthController).Index)
	web.CtrlGet("/oauth/:type", (*controllers.OAuthController).Index)

	// 通用订阅
	web.CtrlGet("/sub/:type/:subtype", (*controllers.SubController).Index)

	// User
	userNs := web.NewNamespace("/user/",
		web.NSBefore(middleware.UserAuthHandler),
		web.NSCtrlGet("", (*controllers.UserController).Index),

		// 签到
		web.NSCtrlPost("checkin", (*controllers.UserController).CheckIn),
		// 公告
		web.NSCtrlGet("announcement", (*controllers.UserController).Announcement),

		// 文档
		web.NSCtrlGet("docs", (*user.DocsController).Index),
		web.NSCtrlGet("docs/:id:int/view", (*user.DocsController).Detail),

		// 个人资料
		web.NSCtrlGet("profile", (*controllers.UserController).Profile),
		// Invite
		web.NSCtrlGet("invite", (*user.InviteController).Index),
		web.NSCtrlGet("invite/reset", (*user.InviteController).Reset),
		// 封禁
		web.NSCtrlGet("banned", (*controllers.UserController).Banned),
		// 节点
		web.NSCtrlGet("server", (*user.ServerController).Index),
		// 动态倍率
		web.NSCtrlGet("rate", (*user.RateController).Index),
		web.NSCtrlPost("rate", (*user.RateController).Ajax),
		// 审计
		web.NSCtrlGet("detect", (*user.DetectController).Rule),
		web.NSCtrlGet("detect/log", (*user.DetectController).Log),
		// 工单
		web.NSCtrlGet("ticket", (*user.TicketController).Index),
		web.NSCtrlGet("ticket/create", (*user.TicketController).Create),
		web.NSCtrlPost("ticket", (*user.TicketController).Add),
		web.NSCtrlGet("ticket/:id:int/view", (*user.TicketController).Detail),
		web.NSCtrlPut("ticket/:id:int", (*user.TicketController).Update),
		// 资料编辑
		web.NSCtrlGet("edit", (*user.InfoController).Index),
		web.NSCtrlPost("email", (*user.InfoController).UpdateEmail),
		web.NSCtrlPost("username", (*user.InfoController).UpdateUsername),
		web.NSCtrlPost("unbind_im", (*user.InfoController).UnbindIM),
		web.NSCtrlPost("password", (*user.InfoController).UpdatePassword),
		web.NSCtrlPost("passwd_reset", (*user.InfoController).ResetPassword),
		web.NSCtrlPost("apitoken_reset", (*user.InfoController).ResetApiToken),
		web.NSCtrlPost("method", (*user.InfoController).UpdateMethod),
		web.NSCtrlPost("url_reset", (*user.InfoController).ResetURL),
		web.NSCtrlPost("daily_mail", (*user.InfoController).UpdateDailyMail),
		web.NSCtrlPost("contact_method", (*user.InfoController).UpdateContactMethod),
		web.NSCtrlPost("theme", (*user.InfoController).UpdateTheme),
		web.NSCtrlPost("kill", (*user.InfoController).SendToGulag),
		// 发送验证邮件
		web.NSCtrlGet("send", (*controllers.AuthController).SendVerify),
        // MFA
		web.NSCtrlPost("ga_check", (*user.MFAController).CheckGa),
		web.NSCtrlPost("ga_set", (*user.MFAController).SetGa),
		web.NSCtrlPost("ga_reset", (*user.MFAController).ResetGa),
		// 深色模式切换
		web.NSCtrlPost("switch_theme_mode", (*controllers.UserController).SwitchThemeMode),
		// 订阅记录
		web.NSCtrlGet("/subscribe", (*user.SubLogController).Index),
		// 流量记录
		web.NSCtrlGet("/traffic", (*user.TrafficLogController).Index),
		// 账户余额
		web.NSCtrlGet("/money", (*user.MoneyController).Index),
		web.NSCtrlPost("/giftcard", (*user.MoneyController).Giftcard),
		// 产品页面
		web.NSCtrlGet("/product", (*user.ProductController).Index),
		// 订单页面
		web.NSCtrlGet("/order", (*user.OrderController).Index),
		web.NSCtrlGet("/order/create", (*user.OrderController).Create),
		web.NSCtrlPost("/order/create", (*user.OrderController).Process),
		web.NSCtrlGet("/order/:id:int/view", (*user.OrderController).Detail),
		web.NSCtrlPost("/order/ajax", (*user.OrderController).Ajax),
		// 账单页面
		web.NSCtrlGet("/invoice", (*user.InvoiceController).Index),
		web.NSCtrlGet("/invoice/:id:int/view", (*user.InvoiceController).Detail),
		web.NSCtrlPost("/invoice/pay_balance", (*user.InvoiceController).PayBalance),
		web.NSCtrlPost("/invoice/ajax", (*user.InvoiceController).Ajax),
		// 新优惠码系统
		web.NSCtrlPost("/coupon", (*user.CouponController).Check),
		// 支付
		web.NSCtrlPost("/payment/purchase/:type", (*controllers.PaymentController).Purchase),
		web.NSCtrlGet("/payment/purchase/:type", (*controllers.PaymentController).Purchase),
		web.NSCtrlPost("/payment/return/:type", (*controllers.PaymentController).ReturnHTML),

		// Get Clients
		web.NSCtrlGet("/clients/:name", (*user.ClientController).GetClients),
		// 登出
		web.NSCtrlGet("logout", (*controllers.UserController).Logout),
	)
	// web.InsertFilter("/user/*", )
	web.AddNamespace(userNs)

	// Auth
	authNs := web.NewNamespace("/auth/",
		web.NSBefore(middleware.GuestHandler),
		web.NSCtrlGet("login", (*controllers.AuthController).Login),
		web.NSCtrlPost("login", (*controllers.AuthController).LoginHandle),
		web.NSCtrlGet("register", (*controllers.AuthController).Register),
		web.NSCtrlPost("register", (*controllers.AuthController).RegisterHandle),
		web.NSCtrlPost("send", (*controllers.AuthController).SendVerify),
		web.NSCtrlGet("logout", (*controllers.AuthController).Logout),
	)
	web.AddNamespace(authNs)

	// Password
	passwordNs := web.NewNamespace("/password/",
		web.NSBefore(middleware.GuestHandler),
		web.NSCtrlGet("reset", (*controllers.PasswordController).Reset),
		web.NSCtrlPost("reset", (*controllers.PasswordController).HandleReset),
		web.NSCtrlGet("token/:token", (*controllers.PasswordController).Token),
		web.NSCtrlPost("token/:token", (*controllers.PasswordController).HandleToken),
	)
	web.AddNamespace(passwordNs)

	// Admin
	adminNs := web.NewNamespace("/admin/",
		web.NSBefore(middleware.AdminAuthHandler),
		web.NSCtrlGet("", (*controllers.AdminController).Index),
		// Node
		web.NSCtrlGet("node", (*admin.NodeController).Index),
		web.NSCtrlGet("node/create", (*admin.NodeController).Create),
		web.NSCtrlPost("node", (*admin.NodeController).Add),
		web.NSCtrlGet("node/:id:int/edit", (*admin.NodeController).Edit),
		web.NSCtrlPost("node/:id:int/reset_password", (*admin.NodeController).ResetPassword),
		web.NSCtrlPost("node/:id:int/reset_bandwidth", (*admin.NodeController).ResetBandwidth),
		web.NSCtrlPost("node/:id:int/copy", (*admin.NodeController).Copy),
		web.NSCtrlPut("node/:id:int", (*admin.NodeController).Update),
		web.NSCtrlDelete("node/:id:int", (*admin.NodeController).Delete),
		web.NSCtrlPost("node/ajax", (*admin.NodeController).Ajax),
		// Ticket
		web.NSCtrlGet("ticket", (*admin.TicketController).Index),
		// web.NSCtrlPost("ticket", (*admin.TicketController).Add),
		web.NSCtrlGet("ticket/:id:int/view", (*admin.TicketController).Detail),
		web.NSCtrlPut("ticket/:id:int/close", (*admin.TicketController).Close),
		web.NSCtrlPut("ticket/:id:int", (*admin.TicketController).Update),
		web.NSCtrlPut("ticket/:id:int/ai", (*admin.TicketController).UpdateAi),
		web.NSCtrlDelete("ticket/:id:int", (*admin.TicketController).Delete),
		web.NSCtrlPost("ticket/ajax", (*admin.TicketController).Ajax),
		// Ann
		web.NSCtrlGet("announcement", (*admin.AnnController).Index),
		web.NSCtrlGet("announcement/create", (*admin.AnnController).Create),
		web.NSCtrlPost("announcement", (*admin.AnnController).Add),
		web.NSCtrlGet("announcement/:id:int/edit", (*admin.AnnController).Edit),
		web.NSCtrlPut("announcement/:id:int", (*admin.AnnController).Update),
		web.NSCtrlDelete("announcement/:id:int", (*admin.AnnController).Delete),
		web.NSCtrlPost("announcement/ajax", (*admin.AnnController).Ajax),
		// Docs
		web.NSCtrlGet("docs", (*admin.DocsController).Index),
		web.NSCtrlGet("docs/create", (*admin.DocsController).Create),
		web.NSCtrlPost("docs", (*admin.DocsController).Add),
		web.NSCtrlPost("docs/generate", (*admin.DocsController).Generate),
		web.NSCtrlGet("docs/:id:int/edit", (*admin.DocsController).Edit),
		web.NSCtrlPut("docs/:id:int", (*admin.DocsController).Update),
		web.NSCtrlDelete("docs/:id:int", (*admin.DocsController).Delete),
		web.NSCtrlPost("docs/ajax", (*admin.DocsController).Ajax),
		// 审计规则
		web.NSCtrlGet("detect", (*admin.DetectRuleController).Index),
		// web.NSCtrlGet("detect/create", (*admin.DetectRuleController).Create),
		web.NSCtrlPost("detect/add", (*admin.DetectRuleController).Add),
		web.NSCtrlDelete("detect/:id:int", (*admin.DetectRuleController).Delete),
		web.NSCtrlPost("detect/ajax", (*admin.DetectRuleController).Ajax),
		// 审计触发日志
		web.NSCtrlGet("detect/log", (*admin.DetectLogController).Index),
		web.NSCtrlPost("detect/log/ajax", (*admin.DetectLogController).Ajax),
		// 审计封禁日志
		web.NSCtrlGet("detect/ban", (*admin.DetectBanLogController).Index),
		web.NSCtrlPost("detect/ban/ajax", (*admin.DetectBanLogController).Ajax),
		// User
		web.NSCtrlGet("user", (*admin.UserController).Index),
		web.NSCtrlGet("user/:id:int/edit", (*admin.UserController).Edit),
		web.NSCtrlPut("user/:id:int", (*admin.UserController).Update),
		web.NSCtrlPost("user/create", (*admin.UserController).Create),
		web.NSCtrlDelete("user/:id:int", (*admin.UserController).Delete),
		web.NSCtrlPost("user/ajax", (*admin.UserController).Ajax),
		// Coupon
		web.NSCtrlGet("coupon", (*admin.CouponController).Index),
		web.NSCtrlPost("coupon", (*admin.CouponController).Add),
		web.NSCtrlPost("coupon/ajax", (*admin.CouponController).Ajax),
		web.NSCtrlDelete("coupon/:id:int", (*admin.CouponController).Delete),
		web.NSCtrlPost("coupon/:id:int/disable", (*admin.CouponController).Disable),
		// 登录日志
		web.NSCtrlGet("login", (*admin.LoginLogController).Index),
		web.NSCtrlPost("login/ajax", (*admin.LoginLogController).Ajax),
		// 在线IP日志
		web.NSCtrlGet("online", (*admin.OnlineLogController).Index),
		web.NSCtrlPost("online/ajax", (*admin.OnlineLogController).Ajax),
		// 订阅日志
		web.NSCtrlGet("subscribe", (*admin.SubLogController).Index),
		web.NSCtrlPost("subscribe/ajax", (*admin.SubLogController).Ajax),
		// 返利日志
		web.NSCtrlGet("payback", (*admin.PaybackController).Index),
		web.NSCtrlPost("payback/ajax", (*admin.PaybackController).Ajax),
		// 用户余额日志
		web.NSCtrlGet("money", (*admin.MoneyLogController).Index),
		web.NSCtrlPost("money/ajax", (*admin.MoneyLogController).Ajax),
		// 支付网关日志
		web.NSCtrlGet("gateway", (*admin.PaylistController).Index),
		web.NSCtrlPost("gateway/ajax", (*admin.PaylistController).Ajax),
		// 系统状态
		web.NSCtrlGet("system", (*admin.SystemController).Index),
		web.NSCtrlPost("system/check_update", (*admin.SystemController).CheckUpdate),
		// 设置中心
		web.NSCtrlGet("/setting/billing", (*setting.BillingController).Index),
		web.NSCtrlPost("/setting/billing", (*setting.BillingController).Save),
		web.NSCtrlPost("setting/billing/set_stripe_webhook", (*setting.BillingController).SetStripeWebhook),
		web.NSCtrlGet("setting/captcha", (*setting.CaptchaController).Index),
		web.NSCtrlPost("setting/captcha", (*setting.CaptchaController).Save),
		web.NSCtrlGet("setting/cron", (*setting.CronController).Index),
		web.NSCtrlPost("setting/cron", (*setting.CronController).Save),
		web.NSCtrlGet("setting/email", (*setting.EmailController).Index),
		web.NSCtrlPost("setting/email", (*setting.EmailController).Save),
		web.NSCtrlGet("setting/feature", (*setting.FeatureController).Index),
		web.NSCtrlPost("setting/feature", (*setting.FeatureController).Save),
		web.NSCtrlGet("setting/im", (*setting.ImController).Index),
		web.NSCtrlPost("setting/im", (*setting.ImController).Save),
		web.NSCtrlGet("setting/ref", (*setting.RefController).Index),
		web.NSCtrlPost("setting/ref", (*setting.RefController).Save),
		web.NSCtrlGet("setting/reg", (*setting.RegController).Index),
		web.NSCtrlPost("setting/reg", (*setting.RegController).Save),
		web.NSCtrlGet("setting/sub", (*setting.SubController).Index),
		web.NSCtrlPost("setting/sub", (*setting.SubController).Save),
		web.NSCtrlGet("setting/support", (*setting.SupportController).Index),
		web.NSCtrlPost("setting/support", (*setting.SupportController).Save),
		
		// 设置测试

		// 礼品卡
		web.NSCtrlGet("giftcard", (*admin.GiftCardController).Index),
		web.NSCtrlPost("giftcard", (*admin.GiftCardController).Add),
		web.NSCtrlPost("giftcard/ajax", (*admin.GiftCardController).Ajax),
		web.NSCtrlDelete("giftcard/:id:int", (*admin.GiftCardController).Delete),
		// 商品
		web.NSCtrlGet("product", (*admin.ProductController).Index),
		web.NSCtrlGet("product/create", (*admin.ProductController).Create),
		web.NSCtrlPost("product", (*admin.ProductController).Add),
		web.NSCtrlGet("product/:id:int/edit", (*admin.ProductController).Edit),
		web.NSCtrlPost("product/:id:int/copy", (*admin.ProductController).Copy),
		web.NSCtrlPut("product/:id:int", (*admin.ProductController).Update),
		web.NSCtrlDelete("product/:id:int", (*admin.ProductController).Delete),
		web.NSCtrlPost("product/ajax", (*admin.ProductController).Ajax),
		// 订单
		web.NSCtrlGet("order", (*admin.OrderController).Index),
		web.NSCtrlGet("order/:id:int/view", (*admin.OrderController).Detail),
		web.NSCtrlPost("order/:id:int/cancel", (*admin.OrderController).Cancel),
		web.NSCtrlDelete("order/:id:int", (*admin.OrderController).Delete),
		web.NSCtrlPost("order/ajax", (*admin.OrderController).Ajax),
		// 账单
		web.NSCtrlGet("invoice", (*admin.InvoiceController).Index),
		web.NSCtrlGet("invoice/:id:int/view", (*admin.InvoiceController).Detail),
		web.NSCtrlPost("invoice/:id:int/mark_paid", (*admin.InvoiceController).MarkPaid),
		web.NSCtrlPost("invoice/ajax", (*admin.InvoiceController).Ajax),
	)
	web.AddNamespace(adminNs)

	// WebAPI

}
