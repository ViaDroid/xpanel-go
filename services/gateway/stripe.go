package gateway

import "github.com/beego/beego/v2/server/web"

type Stripe struct {
	Base
}


func (a *Stripe) Purchase(c *web.Controller) {

}

func (a *Stripe) Notify(c *web.Controller) {

}

func (a *Stripe) Name() string {
	return "stripe"
}

func (a *Stripe) ReadableName() string {
	return "Stripe"
}

func (s *Stripe) Enabled() bool {
	return s.GetActiveGateway("stripe")
}


func (a *Stripe) GetPurchaseHTML(c *web.Controller) string {
	return ""
}