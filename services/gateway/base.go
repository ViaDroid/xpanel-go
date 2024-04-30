package gateway

import (
	"encoding/json"
	"slices"

	"github.com/beego/beego/v2/server/web"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

// var Gateways = map[string]Gateway{
// 	"f2f":    &AlipayF2F{},
// 	"stripe": &Stripe{},
// 	"epay":   &EPay{},
// 	"paypay": &PayPal{},
// }

var Gateways = []Gateway{
	&AlipayF2F{},
	&Stripe{},
	&EPay{},
	&PayPal{},
}

type Gateway interface {
	Purchase(c *web.Controller)
	Notify(c *web.Controller)
	GetPurchaseHTML(c *web.Controller) string
	Name() string
	ReadableName() string
	Enabled() bool
	GetActiveGateway(key string) bool
	GetReturnHTML(c *web.Controller)
	GetStatus(c *web.Controller)
}

type Base struct{}

func (b *Base) GetReturnHTML(c *web.Controller) {
	c.Ctx.WriteString("ok")
}

func (b *Base) GetStatus(c *web.Controller) {
	payList := models.NewPaylist()
	global.DB.QueryTable(models.NewPaylist()).Filter("tradeno", "").One(payList)

	c.JSONResp(map[string]interface{}{
		"ret":    1,
		"result": payList.Status,
	})
}

func (b *Base) GetActiveGateway(key string) bool {
	config := models.NewConfig().Obtain("payment_gateway")

	var active_gateways []string
	json.Unmarshal([]byte(config.Value), &active_gateways)

	return slices.Contains(active_gateways, key)
}
