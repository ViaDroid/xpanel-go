package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/viadroid/xpanel-go/init"
	_ "github.com/viadroid/xpanel-go/routers"
)

func main() {
	web.Run()
}
