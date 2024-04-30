package gateway

import "github.com/beego/beego/v2/server/web"

type EPay struct {
	Base
}

func (a *EPay) Purchase(c *web.Controller) {

}

func (a *EPay) Notify(c *web.Controller) {

}

func (a *EPay) Name() string {
	return "epay"
}

func (a *EPay) ReadableName() string {
	return "EPay"
}

func (e *EPay) Enabled() bool {
	return e.GetActiveGateway("epay")
}

func (a *EPay) GetPurchaseHTML(c *web.Controller) string {
	return ""
}
