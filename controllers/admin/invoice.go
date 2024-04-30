package admin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

type InvoiceController struct {
	controllers.BaseController
}

var invoice_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "账单ID"),
		NewField("UserId", "归属用户"),
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
	c.TplName = "views/tabler/admin/invoice/index.tpl"
}

func (c *InvoiceController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	invoice, err := models.NewInvoice().Find(id)
	if err != nil {
		c.Error(&c.Controller, "查询账单失败")
		return
	}

	var payList *models.Paylist
	if invoice.Status == "paid_gateway" {
		payList, err = models.NewPaylist().FindByInvoiceId(invoice.Id)
		if err == nil {
			payList.ParseStatus()
		}
	}

	invoice.Parse()

	c.Data["invoice"] = invoice
	c.Data["invoice_content"] = invoice.ContentMap
	c.Data["paylist"] = payList

	c.TplName = "views/tabler/admin/invoice/view.tpl"
}

func (c *InvoiceController) MarkPaid() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	invoice, err := models.NewInvoice().Find(id)
	if err != nil {
		c.Error(&c.Controller, "查询账单失败")
		return
	}

	order, _ := models.NewOrder().Find(invoice.OrderId)
	if order.Status == "cancelled" {
		c.Error(&c.Controller, "关联订单已被取消，标记失败")
	}

	tx, _ := global.DB.Begin()

	order.UpdateTime = time.Now().UnixMilli()
	order.Status = "pending_activation"
	if _, err := tx.Update(order); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "更新订单失败")
		return
	}

	invoice.UpdateTime = time.Now().UnixMilli()
	invoice.PayTime = time.Now().UnixMilli()
	invoice.Status = "paid_admin"
	if _, err := tx.Update(invoice); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "更新账单失败")
		return
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "更新账单失败")
		return
	}
	c.Success(&c.Controller, "成功标记账单为已支付（管理员）")
}

func (c *InvoiceController) Ajax() {
	list := models.NewInvoice().FetchAll()

	for i, v := range list {
		list[i].Op = fmt.Sprintf(`<a class="btn btn-blue" href="/admin/invoice/%d/view">查看</a>`, v.Id)
		list[i].Parse()
	}

	c.JSONResp(map[string]any{
		"invoices": list,
	})
}
