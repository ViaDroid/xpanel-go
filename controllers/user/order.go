package user

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

type OrderController struct {
	controllers.BaseController
}

func NewField(key, value string) *controllers.Field {
	return &controllers.Field{
		Key:   key,
		Value: value,
	}
}

var order_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "订单ID"),
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
	c.TplName = "views/tabler/user/order/index.tpl"
}

func (c *OrderController) Create() {
	product_id := c.GetString("product_id")

	if product_id == "" {
		c.Redirect("/user/product", 302)
		return
	}

	id, err := strconv.Atoi(product_id)
	if err != nil {
		c.Redirect("/user/product", 302)
		return
	}

	// redir := c.Ctx.GetCookie("redir")

	product, err := models.NewProduct().Find(id)
	if err != nil {
		c.Redirect("/user/product", 302)
		return
	}

	product.Parse()

	c.Data["product"] = product
	c.TplName = "views/tabler/user/order/create.tpl"
}

func (c *OrderController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect("/user/order", 302)
		return
	}

	order, err := models.NewOrder().FindByUserIdAndId(c.User.Id, id)
	if err != nil {
		c.Redirect("/user/order", 302)
		return
	}
	order.Parse()

	invoice, err := models.NewInvoice().FindOneByOrderId(order.Id)

	if err == nil {
		invoice.Parse()
		c.Data["invoice"] = invoice
	}

	c.Data["order"] = order
	c.TplName = "views/tabler/user/order/view.tpl"

}

func (c *OrderController) Process() {
	coupon_raw := c.GetString("coupon")
	product_id, err := c.GetInt("product_id")

	if err != nil {
		c.Error(&c.Controller, "商品ID错误")
		return
	}

	if c.User.IsShadowBanned {
		c.Error(&c.Controller, "您已被封禁")
		return
	}

	product, err := models.NewProduct().Find(product_id)
	if err != nil {
		c.Error(&c.Controller, "商品不存在或库存不足")
		return
	}

	product.Parse()

	buyPrice := product.Price
	var discount float64
	var coupon = new(models.UserCoupon)
	if coupon_raw != "" {
		coupon, err := models.NewUserCoupon().FindByCode(coupon_raw)
		if err != nil {
			c.Error(&c.Controller, "优惠码不存在或已过期")
			return
		}
		coupon.Parse()
		if coupon.IsExpired {
			c.Error(&c.Controller, "优惠码已过期")
			return
		}

		if coupon.Disabled == "是" {
			c.Error(&c.Controller, "优惠码已被禁用")
			return
		}

		if coupon.LimitMap["product_id"] != "" && slices.Contains(strings.Split(coupon.LimitMap["product_id"].(string), ","), strconv.Itoa(product_id)) {
			c.Error(&c.Controller, "优惠码不适用于该商品")
			return
		}

		coupon_use_limit := coupon.LimitMap["use_time"].(float64)
		if coupon_use_limit > 0 {
			use_count, _ := global.DB.QueryTable(models.NewOrder()).Filter("user_id", c.User.Id).Filter("coupon", coupon_raw).Count()
			if float64(use_count) >= coupon_use_limit {
				c.Error(&c.Controller, "优惠码使用次数已达上限")
				return
			}
		}

		coupon_total_use_limit := coupon.LimitMap["total_use_time"].(float64)

		if coupon_total_use_limit > 0 && coupon.UseCount >= int(coupon_total_use_limit) {
			c.Error(&c.Controller, "优惠码已达总使用次数上限")
			return
		}

		if coupon.ContentMap["type"] == "percentage" {
			discount = product.Price * coupon.ContentMap["value"].(float64) / 100
		} else {
			discount = coupon.ContentMap["value"].(float64)
		}

		buyPrice = product.Price - discount
	}

	if c.User.Class < int(product.LimitMap["class_required"].(float64)) {
		c.Error(&c.Controller, "您的账户等级不足，无法购买此商品")
		return
	}

	if c.User.NodeGroup != int(product.LimitMap["node_group_required"].(float64)) {
		c.Error(&c.Controller, "您所在的用户组无法购买此商品")
		return
	}

	if int(product.LimitMap["node_group_required"].(float64)) != 0 {
		exist := global.DB.QueryTable(models.NewOrder()).Filter("user_id", c.User.Id).Exist()
		if exist {
			// c.Error(&c.Controller, "您已购买过此商品")
			c.Error(&c.Controller, "此商品仅限新用户购买")
			return
		}
	}

	// transaction
	tx, _ := global.DB.Begin()
	defer tx.Commit()

	// create order
	now := time.Now()
	order := models.Order{
		UserId:         c.User.Id,
		ProductId:      product_id,
		ProductType:    product.Type,
		ProductName:    product.Name,
		ProductContent: product.Content,
		Coupon:         coupon_raw,
		Price:          buyPrice,
		Status:         "pending_payment",
		CreateTime:     now.UnixMilli(),
		UpdateTime:     now.UnixMilli(),
	}

	if _, err := tx.Insert(&order); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "创建订单失败")
		return
	}

	// create invoice
	content := map[string]any{
		"content_id": 0,
		"name":       product.Name,
		"price":      product.Price,
	}
	if coupon_raw != "" {
		content = map[string]any{
			"content_id": 0,
			"name":       "优惠码 " + coupon_raw,
			"price":      fmt.Sprintf("-%f", discount),
		}
	}

	contentJsonBs, _ := json.Marshal(content)
	invoice := models.Invoice{
		OrderId:    order.Id,
		UserId:     c.User.Id,
		Content:    string(contentJsonBs),
		Price:      buyPrice,
		Status:     "unpaid",
		CreateTime: now.UnixMilli(),
		UpdateTime: now.UnixMilli(),
		PayTime:    0,
	}
	if _, err := tx.Insert(&invoice); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "创建发票失败")
		return
	}

	// update product stock
	if product.Stock > 0 {
		product.Stock--
	}
	product.SaleCount++
	// product.UpdateTime = now.UnixMilli()
	if _, err := tx.Update(product); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "更新商品库存失败")
		return
	}

	// update coupon use count
	if coupon_raw != "" {
		coupon.UseCount++
		if _, err := tx.Update(coupon); err != nil {
			tx.Rollback()
			c.Error(&c.Controller, "更新优惠码使用次数失败")
			return
		}
	}

	c.JSONResp(map[string]any{
		"ret":        1,
		"msg":        "成功创建订单，正在跳转账单页面",
		"invoice_id": invoice.Id,
	})

}

func (c *OrderController) Ajax() {
	list := models.NewOrder().FetchListByUserId(c.User.Id)

	for i, v := range list {

		op := fmt.Sprintf(`<a class="btn btn-blue" href="/user/order/%d/view">查看</a>`, v.Id)

		if v.Status == "pending_payment" {
			invoice, _ := models.NewInvoice().FindOneByOrderId(v.Id)
			op += fmt.Sprintf(`<a class="btn btn-red" style="margin-left:8px" href="/user/invoice/%d/view">支付</a>`, invoice.Id)
		}
		list[i].Op = op
		list[i].Parse()
	}

	c.JSONResp(map[string]any{
		"orders": list,
	})
}
