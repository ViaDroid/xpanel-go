package admin

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type ProductController struct {
	controllers.BaseController
}

var product_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "商品ID"),
		NewField("TypeText", "类型"),
		NewField("Name", "名称"),
		NewField("Price", "售价"),
		NewField("StatusStr", "销售状态"),
		NewField("CreateTimeStr", "创建时间"),
		NewField("UpdateTimeStr", "更新时间"),
		NewField("SaleCount", "累计销售"),
		NewField("Stock", "库存"),
	},
}

var product_update_fields = []string{
	"type",
	"name",
	"price",
	"status",
	"stock",
	"time",
	"bandwidth",
	"class",
	"class_time",
	"node_group",
	"speed_limit",
	"ip_limit",
	"class_required",
	"node_group_required",
}

var product_invalid_data_msg = "无效商品数据"

func (c *ProductController) Index() {
	c.Data["details"] = product_menu_details
	c.TplName = "views/tabler/admin/product/index.tpl"
}

func (c *ProductController) Create() {
	c.Data["update_field"] = product_update_fields
	c.TplName = "views/tabler/admin/product/create.tpl"
}

func (c *ProductController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "商品不存在")
		return
	}
	product, err := models.NewProduct().Find(id)
	if err != nil {
		c.Error(&c.Controller, "商品不存在")
		return
	}

	var content models.Content
	var limit models.Limit
	json.Unmarshal([]byte(product.Content), &content)
	json.Unmarshal([]byte(product.Limit), &limit)

	c.Data["product"] = product
	c.Data["content"] = content
	c.Data["limit"] = limit
	c.Data["update_field"] = product_update_fields
	c.TplName = "views/tabler/admin/product/edit.tpl"
}

func (c *ProductController) Add() {
	// base product
	typ := c.GetString("type")
	name := c.GetString("name")
	price, _ := c.GetFloat("price", 0)
	status, _ := c.GetInt("status", 1)
	stock, _ := c.GetInt("stock", -1)

	// content
	duation_time, _ := c.GetInt("time", 0)
	bandwidth, _ := c.GetInt("bandwidth", 0)
	class, _ := c.GetInt("class", 0)
	class_time, _ := c.GetInt("class_time", 0)
	node_group, _ := c.GetInt("node_group", 0)
	speed_limit, _ := c.GetInt("speed_limit", 0)
	ip_limit, _ := c.GetInt("ip_limit", 0)
	// limit
	class_required, _ := c.GetInt("class_required", 0)
	node_group_required, _ := c.GetInt("node_group_required", 0)
	new_user_required, _ := c.GetInt("new_user_required", 0)

	var content models.Content
	if price < 0 {
		c.Error(&c.Controller, product_invalid_data_msg)
		return
	}
	switch typ {
	case "tabp":
		if duation_time <= 0 || class_time <= 0 || bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Time:       duation_time,
			Bandwidth:  bandwidth,
			Class:      class,
			ClassTime:  class_time,
			NodeGroup:  node_group,
			SpeedLimit: speed_limit,
			IpLimit:    ip_limit,
		}
	case "time":
		if duation_time <= 0 || class_time <= 0 || bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Time:       duation_time,
			Class:      class,
			ClassTime:  class_time,
			NodeGroup:  node_group,
			SpeedLimit: speed_limit,
			IpLimit:    ip_limit,
		}
	case "bandwidth":
		if bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Bandwidth: bandwidth,
		}

	default:
		c.Error(&c.Controller, product_invalid_data_msg)
		return
	}

	limit := models.Limit{
		ClassRequired:     class_required,
		NodeGroupRequired: node_group_required,
		NewUserRequired:   new_user_required,
	}

	contentJsonBs, _ := json.Marshal(content)
	limitJsonBs, _ := json.Marshal(limit)

	product := models.Product{
		Type:       typ,
		Name:       name,
		Price:      price,
		Status:     status,
		Stock:      stock,
		Content:    string(contentJsonBs),
		Limit:      string(limitJsonBs),
		SaleCount:  0,
		CreateTime: time.Now().UnixMilli(),
		UpdateTime: time.Now().UnixMilli(),
	}

	if _, err := product.Save(); err != nil {
		c.Error(&c.Controller, "添加失败")
		return
	}

	c.Success(&c.Controller, "添加成功")
}

func (c *ProductController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "商品Id错误")
		return
	}
	product, err := models.NewProduct().Find(id)
	if err != nil {
		c.Error(&c.Controller, "商品不存在")
		return
	}

	// base product
	typ := c.GetString("type")
	name := c.GetString("name")
	price, _ := c.GetFloat("price", 0)
	status, _ := c.GetInt("status", 1)
	stock, _ := c.GetInt("stock", -1)

	// content
	duation_time, _ := c.GetInt("time", 0)
	bandwidth, _ := c.GetInt("bandwidth", 0)
	class, _ := c.GetInt("class", 0)
	class_time, _ := c.GetInt("class_time", 0)
	node_group, _ := c.GetInt("node_group", 0)
	speed_limit, _ := c.GetInt("speed_limit", 0)
	ip_limit, _ := c.GetInt("ip_limit", 0)
	// limit
	class_required, _ := c.GetInt("class_required", 0)
	node_group_required, _ := c.GetInt("node_group_required", 0)
	new_user_required, _ := c.GetInt("new_user_required", 0)

	var content models.Content
	if price < 0 {
		c.Error(&c.Controller, product_invalid_data_msg)
		return
	}
	switch typ {
	case "tabp":
		if duation_time <= 0 || class_time <= 0 || bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Time:       duation_time,
			Bandwidth:  bandwidth,
			Class:      class,
			ClassTime:  class_time,
			NodeGroup:  node_group,
			SpeedLimit: speed_limit,
			IpLimit:    ip_limit,
		}
	case "time":
		if duation_time <= 0 || class_time <= 0 || bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Time:       duation_time,
			Class:      class,
			ClassTime:  class_time,
			NodeGroup:  node_group,
			SpeedLimit: speed_limit,
			IpLimit:    ip_limit,
		}
	case "bandwidth":
		if bandwidth <= 0 {
			c.Error(&c.Controller, product_invalid_data_msg)
			return
		}

		content = models.Content{
			Bandwidth: bandwidth,
		}

	default:
		c.Error(&c.Controller, product_invalid_data_msg)
		return
	}

	limit := models.Limit{
		ClassRequired:     class_required,
		NodeGroupRequired: node_group_required,
		NewUserRequired:   new_user_required,
	}

	contentJsonBs, _ := json.Marshal(content)
	limitJsonBs, _ := json.Marshal(limit)

	product.Type = typ
	product.Name = name
	product.Price = price
	product.Content = string(contentJsonBs)
	product.Limit = string(limitJsonBs)
	product.Status = status
	product.Stock = stock
	product.UpdateTime = time.Now().UnixMilli()

	if _, err := product.Update(); err != nil {
		c.Error(&c.Controller, "更新失败")
		return
	}

	c.Success(&c.Controller, "更新成功")
}

func (c *ProductController) Copy() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if !models.NewProduct().IsExist(id) {
		c.Error(&c.Controller, "商品不存在")
		return
	}
	product, _ := models.NewProduct().Find(id)

	product.Id = 0
	product.Name += "_副本"
	product.CreateTime = time.Now().UnixMilli()
	product.UpdateTime = time.Now().UnixMilli()
	product.SaleCount = 0

	if _, err := product.Save(); err != nil {
		c.Error(&c.Controller, "复制失败")
		return
	}
	c.Success(&c.Controller, "复制成功")

}
func (c *ProductController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if !models.NewProduct().IsExist(id) {
		c.Error(&c.Controller, "商品不存在")
		return
	}

	if _, err := models.NewProduct().Delete(id); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}

	c.Success(&c.Controller, "删除成功")
}

func (c *ProductController) Ajax() {
	list := models.NewProduct().FetchAll()
	for i, v := range list {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-product-%d"
             onclick="deleteProduct(%d)">删除</button>
            <button type="button" class="btn btn-orange" id="copy-product-%d"
             onclick="copyProduct(%d)">复制</button>
            <a class="btn btn-blue" href="/admin/product/%d/edit">编辑</a>`, v.Id, v.Id, v.Id, v.Id, v.Id)
		list[i].Parse()

		list[i].Op = op
	}

	c.JSONResp(map[string]any{
		"products": list,
	})
}
