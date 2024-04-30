package init

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/viadroid/xpanel-go/command"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterModel(
		new(models.Ann),
		new(models.Config),
		new(models.DetectBanLog),
		new(models.DetectLog),
		new(models.DetectRule),
		new(models.Docs),
		new(models.EmailQueue),
		new(models.GiftCard),
		new(models.HourlyUsage),
		new(models.InviteCode),
		new(models.Invoice),
		new(models.Link),
		new(models.LoginIp),
		new(models.Node),
		new(models.OnlineLog),
		new(models.Order),
		new(models.Payback),
		new(models.Paylist),
		new(models.Product),
		new(models.SubscribeLog),
		new(models.Ticket),
		new(models.User),
		new(models.UserCoupon),
		new(models.UserMoneyLog),
	)

	db_username := global.ConfMap["db_username"]
	db_password := global.ConfMap["db_password"]
	db_port := global.ConfMap["db_port"]
	db_host := global.ConfMap["db_host"]
	db_database := global.ConfMap["db_database"]

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_username, db_password, db_host, db_port, db_database)
	orm.RegisterDataBase("default", "mysql", dbSource)
	orm.RunSyncdb("default", false, true)
	global.DB = orm.NewOrm()

	command.ImportSetting()
	// command.ExportSetting()

	global.SettingMap = models.NewConfig().GetPublicConfig()

}
