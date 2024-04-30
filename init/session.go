package init

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"github.com/viadroid/xpanel-go/global"
)

func init() {
	redisAddr := fmt.Sprintf("%s:%v", global.ConfMap["redis_host"], global.ConfMap["redis_port"])
	web.BConfig.WebConfig.Session.SessionProvider = "redis"
	web.BConfig.WebConfig.Session.SessionProviderConfig = redisAddr
	web.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 30 // 30天
	// web.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
}
