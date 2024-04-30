package gateway

import "github.com/beego/beego/v2/server/web"

type PayPal struct {
	Base
}


func (a *PayPal) Purchase(c *web.Controller) {

}

func (a *PayPal) Notify(c *web.Controller) {

}

func (a *PayPal) Name() string {
	return "payPal"
}

func (a *PayPal) ReadableName() string {
	return "PayPal"
}

func (p *PayPal) Enabled() bool {
	return p.GetActiveGateway("payPal")
}


func (a *PayPal) GetPurchaseHTML(c *web.Controller) string {
	return ""
}