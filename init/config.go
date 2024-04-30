package init

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/config/yaml"
	"github.com/beego/beego/v2/server/web"
	"github.com/viadroid/xpanel-go/conf"
	"github.com/viadroid/xpanel-go/global"
)

func init() {

	c, err := yaml.ReadYmlReader(conf.ConfDir("conf.yaml"))
	if err != nil {
		web.BeeApp.Server.ErrorLog.Fatalln(err)
	}
	global.ConfMap = c

	// Load custom config
	err = config.InitGlobalInstance("yaml", conf.ConfDir("conf.yaml"))
	if err != nil {
		panic(err)
	}

	initRedis(c)
}
