package admin

import (
	"fmt"
	"strconv"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DetectRuleController struct {
	controllers.BaseController
}

var detect_rule_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "规则ID"),
		NewField("Name", "规则名称"),
		NewField("Text", "规则介绍"),
		NewField("Regex", "正则表达式"),
		NewField("TypeStr", "规则类型"),
	},
	"add_dialog": []any{
		map[string]any{
			"id":          "name",
			"info":        "规则名称",
			"type":        "input",
			"placeholder": "审计规则名称",
		},
		map[string]any{
			"id":          "text",
			"info":        "规则介绍",
			"type":        "input",
			"placeholder": "简洁明了地描述审计规则",
		},
		map[string]any{
			"id":          "regex",
			"info":        "正则表达式",
			"type":        "input",
			"placeholder": "用以匹配审计内容的正则表达式",
		},
		map[string]any{
			"id":   "type",
			"info": "规则类型",
			"type": "select",
			"select": map[string]any{
				"1": "数据包明文匹配",
				"0": "数据包十六进制匹配",
			},
		},
	},
}

func (c *DetectRuleController) Index() {
	c.Data["details"] = detect_rule_menu_details
	c.TplName = "views/tabler/admin/detect.tpl"
}

func (c *DetectRuleController) Create() {
}

func (c *DetectRuleController) Add() {
	name := c.GetString("name")
	text := c.GetString("text")
	regex := c.GetString("regex")
	type_, _ := c.GetInt("type")

	if name == "" || text == "" {
		c.Error(&c.Controller, "参数错误")
		return
	}

	detectRule := &models.DetectRule{
		Name:  name,
		Text:  text,
		Regex: regex,
		Type:  type_,
	}

	if _, err := detectRule.Save(); err != nil {
		c.Error(&c.Controller, "添加失败")
		return
	}

	// TODO notify 有新的审计规则：

	c.Success(&c.Controller, "添加成功")
}

func (c *DetectRuleController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}
	if _, err := models.NewDetecRule().Delete(id); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}
	c.Success(&c.Controller, "删除成功")
}

func (c *DetectRuleController) Ajax() {
	list := models.NewDetecRule().FetchAll()
	for i, v := range list {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-rule-%d" onclick="deleteRule(%d)">
		删除</button>`, v.Id, v.Id)
		list[i].Op = op
		list[i].TypeStr = v.ParseType()
	}

	c.JSONResp(map[string]any{
		"rules": list,
	})
}
