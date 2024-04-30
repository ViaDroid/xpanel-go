package admin

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

var node_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "节点ID"),
		NewField("Name", "名称"),
		NewField("Server", "地址"),
		NewField("TypeStr", "状态"),
		NewField("SortStr", "类型"),
		NewField("TrafficRate", "倍率"),
		NewField("IsDynamicRateStr", "动态倍率"),
		NewField("DynamicRateType", "动态倍率计算方式"),
		NewField("NodeClass", "等级"),
		NewField("NodeGroup", "组别"),
		NewField("NodeBandwidthLimit", "流量限制/GB"),
		NewField("NodeBandwidth", "已用流量/GB"),
		NewField("BandwidthlimitResetday", "重置日"),
	},
}

var node_update_field = []string{
	"name",
	"server",
	"traffic_rate",
	"is_dynamic_rate",
	"dynamic_rate_type",
	"max_rate",
	"max_rate_time",
	"min_rate",
	"min_rate_time",
	"node_group",
	"node_speedlimit",
	"sort",
	"node_class",
	"node_bandwidth_limit",
	"bandwidthlimit_resetday",
}

type NodeController struct {
	controllers.BaseController
}

func (c *NodeController) Index() {
	c.Data["details"] = node_menu_details
	c.TplName = "views/tabler/admin/node/index.tpl"
}
func (c *NodeController) Create() {
	c.Data["update_field"] = node_update_field
	c.TplName = "views/tabler/admin/node/create.tpl"
}

func (c *NodeController) Add() {

	node := models.NewNode()
	node.Name = c.GetString("name")
	node.NodeGroup, _ = c.GetInt("node_group", 0)
	node.Server = c.GetString("server")
	node.TrafficRate, _ = c.GetFloat("traffic_rate", 1)
	node.IsDynamicRate, _ = c.GetBool("is_dynamic_rate", false)
	node.DynamicRateType, _ = c.GetInt("dynamic_rate_type", 0)

	jsonBytes, _ := json.Marshal(map[string]any{
		"max_rate":      c.GetString("max_rate", "1"),
		"max_rate_time": c.GetString("max_rate_time", "22"),
		"min_rate":      c.GetString("min_rate", "1"),
		"min_rate_time": c.GetString("min_rate_time", "3"),
	})

	node.DynamicRateConfig = string(jsonBytes)

	custom_config := c.GetString("custom_config")
	if custom_config == "" {
		custom_config = "{}"
	}
	node.CustomConfig = custom_config

	node.NodeSpeedlimit, _ = c.GetFloat("node_speedlimit", 0)
	node.Type, _ = c.GetBool("type", false)
	node.Sort, _ = c.GetInt("sort", 0)
	node.NodeClass, _ = c.GetInt("node_class", 0)
	node_bandwidth_limit, _ := c.GetInt("node_bandwidth_limit", 0)
	node.NodeBandwidthLimit = utils.ToGB(node_bandwidth_limit)
	node.BandwidthlimitResetday, _ = c.GetInt("bandwidthlimit_resetday", 1)
	node.Password = utils.GenRandomString(32)

	if _, err := node.Insert(); err != nil {
		c.JSONResp(map[string]interface{}{
			"ret": 0,
			"msg": "添加失败",
		})
		return
	}

	// notify
	if models.NewConfig().Obtain("telegram_add_node").ValueToBool() {
		// TODO
	}

	c.JSONResp(map[string]interface{}{
		"ret":     1,
		"msg":     "添加成功",
		"node_id": node.Id,
	})
}

func (c *NodeController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.JSONResp(map[string]interface{}{
			"ret": 0,
			"msg": "编辑失败，节点不存在",
		})
		return
	}

	var dynamic_rate_config map[string]string
	err = json.Unmarshal([]byte(node.DynamicRateConfig), &dynamic_rate_config)
	if err != nil {
		c.Error(&c.Controller, "编辑失败")
		return
	}
	node.MaxRate = dynamic_rate_config["max_rate"]
	node.MaxRateTime = dynamic_rate_config["max_rate_time"]
	node.MinRate = dynamic_rate_config["min_rate"]
	node.MinRateTime = dynamic_rate_config["min_rate_time"]

	node.NodeBandwidth = int64(utils.FlowToGB(float64(node.NodeBandwidth)))
	node.NodeBandwidthLimit = int64(utils.FlowToGB(float64(node.NodeBandwidthLimit)))

	c.Data["node"] = node
	c.Data["update_field"] = node_update_field
	c.TplName = "views/tabler/admin/node/edit.tpl"
}

func (c *NodeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.JSONResp(map[string]interface{}{
			"ret": 0,
			"msg": "删除失败",
		})
		return
	}

	if _, err := node.Delete(id); err != nil {
		c.JSONResp(map[string]interface{}{
			"ret": 0,
			"msg": "删除失败",
		})
	}

	// notify after delete node
	if models.NewConfig().Obtain("telegram_delete_node").ValueToBool() {
		// TODO
	}

	c.JSONResp(map[string]interface{}{
		"ret": 1,
		"msg": "删除成功",
	})
}

func (c *NodeController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.Error(&c.Controller, "编辑失败，节点不存在")
		return
	}

	node.Name = c.GetString("name")
	node.NodeGroup, _ = c.GetInt("node_group", 0)
	node.Server = strings.Trim(c.GetString("server"), "")
	node.TrafficRate, _ = c.GetFloat("traffic_rate", 1)
	node.IsDynamicRate, _ = c.GetBool("is_dynamic_rate", false)
	node.DynamicRateType, _ = c.GetInt("dynamic_rate_type", 0)

	jsonBytes, _ := json.Marshal(map[string]any{
		"max_rate":      c.GetString("max_rate", "1"),
		"max_rate_time": c.GetString("max_rate_time", "22"),
		"min_rate":      c.GetString("min_rate", "1"),
		"min_rate_time": c.GetString("min_rate_time", "3"),
	})

	node.DynamicRateConfig = string(jsonBytes)

	custom_config := c.GetString("custom_config")
	if custom_config == "" {
		custom_config = "{}"
	}
	node.CustomConfig = custom_config

	node.NodeSpeedlimit, _ = c.GetFloat("node_speedlimit", 0)
	node.Type, _ = c.GetBool("type", false)
	node.Sort, _ = c.GetInt("sort", 0)
	node.NodeClass, _ = c.GetInt("node_class", 0)
	node_bandwidth_limit, _ := c.GetInt("node_bandwidth_limit", 0)
	node.NodeBandwidthLimit = utils.ToGB(node_bandwidth_limit)
	node.BandwidthlimitResetday, _ = c.GetInt("bandwidthlimit_resetday", 1)

	if _, err := node.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}

	// notify after update node
	if models.NewConfig().Obtain("telegram_update_node").ValueToBool() {
		// TODO
	}

	c.Success(&c.Controller, "修改成功")

}

func (c *NodeController) ResetPassword() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.Error(&c.Controller, "重置节点通讯密钥失败")
		return
	}

	node.Password = utils.GenRandomString(32)
	if _, err := node.Update(); err != nil {
		c.Error(&c.Controller, "重置节点通讯密钥失败")
		return
	}
	c.Success(&c.Controller, "重置节点通讯密钥成功")
}

func (c *NodeController) ResetBandwidth() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.Error(&c.Controller, "重置节点流量失败")
		return
	}

	node.NodeBandwidth = 0
	if _, err := node.Update(); err != nil {
		c.Error(&c.Controller, "重置节点流量失败")
		return
	}
	c.Success(&c.Controller, "重置节点流量成功")
}
func (c *NodeController) Copy() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	node, err := models.NewNode().Find(id)
	if err != nil {
		c.Error(&c.Controller, "复制失败")
		return
	}

	node.Id = 0
	node.Name += " (副本)"
	node.NodeBandwidth = 0
	node.Password = utils.GenRandomString(32)

	if _, err := node.Insert(); err != nil {
		c.Error(&c.Controller, "复制失败")
		return
	}
	c.Success(&c.Controller, "复制成功")
}
func (c *NodeController) Ajax() {
	nodes := models.NewNode().FetchAll()

	for i, v := range nodes {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-node-%d" 
            onclick="deleteNode(%d)">删除</button>
            <button type="button" class="btn btn-orange" id="copy-node-%d" 
            onclick="copyNode(%d)">复制</button>
            <a class="btn btn-blue" href="/admin/node/%d/edit">编辑</a>
		`, v.Id, v.Id, v.Id, v.Id, v.Id)

		nodes[i].Op = op
		nodes[i].Parse()
	}

	c.JSONResp(map[string]interface{}{
		"nodes": nodes,
	})

}
