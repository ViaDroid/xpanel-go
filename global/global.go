package global

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/session"
	"github.com/redis/go-redis/v9"
)

var (
	DB         orm.Ormer
	ConfMap    map[string]any
	SettingMap map[string]any

	Redis   *redis.Client
	Session *session.Manager
)
