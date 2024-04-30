package controllers

import (
	"github.com/viadroid/xpanel-go/services"
)

type PaymentController struct {
	BaseController
}

func (c *PaymentController) Purchase() {
	typ := c.Ctx.Input.Param(":type")
	paymentGayteway := services.NewPayment().GetPaymentByName(typ)

	if paymentGayteway == nil {
		c.Abort("404")
		return
	}

	paymentGayteway.Purchase(&c.Controller)
}
func (c *PaymentController) ReturnHTML() {
	typ := c.Ctx.Input.Param(":type")
	paymentGayteway := services.NewPayment().GetPaymentByName(typ)

	if paymentGayteway == nil {
		c.Abort("404")
		return
	}

	paymentGayteway.GetReturnHTML(&c.Controller)
}

func (c *PaymentController) Notify() {
	typ := c.Ctx.Input.Param(":type")
	paymentGayteway := services.NewPayment().GetPaymentByName(typ)

	if paymentGayteway == nil {
		c.Abort("404")
		return
	}

	paymentGayteway.Notify(&c.Controller)
}

func (c *PaymentController) GetStatus() {
	typ := c.Ctx.Input.Param(":type")
	paymentGayteway := services.NewPayment().GetPaymentByName(typ)

	if paymentGayteway == nil {
		c.Abort("404")
		return
	}
	paymentGayteway.GetStatus(&c.Controller)
}
