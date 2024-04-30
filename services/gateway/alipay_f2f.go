package gateway

import "github.com/beego/beego/v2/server/web"

type AlipayF2F struct {
	Base
}

// implement Gateway interface

func (a *AlipayF2F) Purchase(c *web.Controller) {

}

func (a *AlipayF2F) Notify(c *web.Controller) {

}

func (a *AlipayF2F) Name() string {
	return "f2f"
}

func (a *AlipayF2F) ReadableName() string {
	return "Alipay F2F"
}

func (a *AlipayF2F) Enabled() bool {
	return a.GetActiveGateway("f2f")
}


func (a *AlipayF2F) GetPurchaseHTML(c *web.Controller) string {
	return "gateway/f2f.tpl"
}
