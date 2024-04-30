package admin

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type OrderController struct {
	controllers.BaseController
}

var order_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "订单ID"),
		NewField("UserId", "提交用户"),
		NewField("ProductId", "商品ID"),
		NewField("ProductTypeStr", "商品类型"),
		NewField("ProductName", "商品名称"),
		NewField("Coupon", "优惠码"),
		NewField("Price", "金额"),
		NewField("StatusStr", "状态"),
		NewField("CreateTimeStr", "创建时间"),
		NewField("UpdateTimeStr", "更新时间"),
	},
}

func (c *OrderController) Index() {
	c.Data["details"] = order_menu_details
	c.TplName = "views/tabler/admin/order/index.tpl"
}

func (c *OrderController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	order, err := models.NewOrder().Find(id)
	if err != nil {
		c.Redirect("/admin/order/index", 302)
		return
	}

	order.Parse()
	invoice, _ := models.NewInvoice().FindOneByOrderId(order.Id)
	invoice.Parse()

	c.Data["order"] = order
	c.Data["invoice"] = invoice
	c.TplName = "views/tabler/admin/order/view.tpl"
}

func (c *OrderController) Cancel() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	order, err := models.NewOrder().Find(id)
	if err != nil {
		c.Error(&c.Controller, "订单不存在!")
		return
	}

	if slices.Contains([]string{"activated", "expired", "cancelled"}, order.Status) {
		c.Error(&c.Controller, "不能取消"+order.Status+"的产品!")
		return
	}

	order.UpdateTime = time.Now().UnixMilli()
	order.Status = "cancelled"

	if _, err := order.Update(); err != nil {
		c.Error(&c.Controller, "取消失败!")
		return
	}

	invoice, err := models.NewInvoice().FindOneByOrderId(order.Id)
	if err != nil {
		c.Success(&c.Controller, "订单取消成功，但关联账单不存在!")
		return
	}

	if slices.Contains([]string{"paid_gateway", "paid_balance", "paid_admin"}, invoice.Status) {
		// TODO refundToBalance
		c.Success(&c.Controller, "订单取消成功，关联账单已退款至余额")
	}

	invoice.UpdateTime = time.Now().UnixMilli()
	invoice.Status = "cancelled"

	if err := invoice.Update(); err != nil {
		c.Error(&c.Controller, "取消失败!")
		return
	}
	c.Success(&c.Controller, "取消成功！")
}

func (c *OrderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if _, err := models.NewOrder().Delete(id); err != nil {
		c.Error(&c.Controller, "删除失败!")
		return
	}

	c.Success(&c.Controller, "删除成功！")
}

func (c *OrderController) Ajax() {

	list := models.NewOrder().FetchAll()
	var order_status = []string{"pending_payment", "pending_activation"}

	for i, v := range list {

		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-order-%d" onclick="deleteOrder(%d)">删除</button>`, v.Id, v.Id)

		if slices.Contains(order_status, v.Status) {
			op += fmt.Sprintf(`<button type="button" class="btn btn-orange" id="cancel-order-%d" onclick="cancelOrder(%d)" style="margin-left: 6px;margin-right: 6px;">取消</button>`, v.Id, v.Id)
		}

		op += fmt.Sprintf(`<a class="btn btn-blue" href="/admin/order/%d/view">查看</a>`, v.Id)

		list[i].Op = op
		list[i].Parse()
	}

	c.JSONResp(map[string]any{
		"orders": list,
	})
}
