package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type InvoiceController struct {
	controllers.BaseController
}

var invoice_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "账单ID"),
		NewField("OrderId", "订单ID"),
		NewField("Price", "账单金额"),
		NewField("StatusStr", "账单状态"),
		NewField("CreateTimeStr", "创建时间"),
		NewField("UpdateTimeStr", "更新时间"),
		NewField("PayTimeStr", "支付时间"),
	},
}

func (c *InvoiceController) Index() {
	c.Data["details"] = invoice_menu_details
	c.TplName = "views/tabler/user/invoice/index.tpl"
}

func (c *InvoiceController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "Id错误")
		return
	}

	invoice, err := models.NewInvoice().Find(id)
	if err != nil || invoice.UserId != c.User.Id {
		c.Redirect("/user/invoice", 302)
		return
	}

	var paylist models.Paylist
	if invoice.Status == "paid_gateway" {
		global.DB.QueryTable(models.NewPaylist()).
			Filter("invoice_id", invoice.Id).
			Filter("status", 1).One(&paylist)
	}

	invoice.Parse()

	c.Data["invoice"] = invoice
	c.Data["invoice_content"] = invoice.ContentMap
	c.Data["paylist"] = paylist
	c.Data["payments"] = services.NewPayment().GetPaymentsEnabled()

	c.TplName = "views/tabler/user/invoice/view.tpl"

}

func (c *InvoiceController) PayBalance() {
	invoice_id, _ := c.GetInt("invoice_id")

	invoice, err := models.NewInvoice().Find(invoice_id)
	if err != nil {
		c.Error(&c.Controller, "账单不存在")
		return
	}

	if c.User.IsShadowBanned {
		c.Error(&c.Controller, "支付失败，请稍后再试")
		return
	}

	if c.User.Money < invoice.Price {
		c.Error(&c.Controller, "余额不足")
		return
	}

	if invoice.Status != "unpaid" {
		c.Error(&c.Controller, "账单状态错误")
		return
	}

	tx, _ := global.DB.Begin()
	now := time.Now()

	moneyBefore := c.User.Money
	c.User.Money -= invoice.Price

	if _, err := tx.Update(c.User); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "支付失败，请稍后再试")
		return
	}

	// add user money log
	userMoneyLog := models.UserMoneyLog{
		UserId:     c.User.Id,
		Before:     moneyBefore,
		After:      c.User.Money,
		Amount:     -invoice.Price,
		Remark:     fmt.Sprintf("支付账单 #%d", invoice.Id),
		CreateTime: now.UnixMilli(),
	}
	if _, err := tx.Insert(&userMoneyLog); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "支付失败，请稍后再试")
		return
	}

	// update invoice status
	invoice.Status = "paid_gateway"
	invoice.UpdateTime = now.UnixMilli()
	invoice.PayTime = now.UnixMilli()

	if _, err := tx.Update(invoice); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "支付失败，请稍后再试")
		return
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "支付失败，请稍后再试")
		return 
	}
	

	c.Success(&c.Controller, "支付成功")
}

func (c *InvoiceController) Ajax() {
	list := models.NewInvoice().FetchAllByUserId(c.User.Id)

	for i, v := range list {
		list[i].Op = fmt.Sprintf(`<a class="btn btn-blue" href="/user/invoice/%d/view">查看</a>`, v.Id)
		list[i].Parse()
	}

	c.JSONResp(map[string]any{
		"invoices": list,
	})
}
