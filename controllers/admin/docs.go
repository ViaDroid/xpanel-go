package admin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DocsController struct {
	controllers.BaseController
}

var docs_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "文档ID"),
		NewField("Date", "日期"),
		NewField("Title", "标题"),
	},
}

func (c *DocsController) Index() {
	c.Data["details"] = docs_menu_details
	c.TplName = "views/tabler/admin/docs/index.tpl"
}

func (c *DocsController) Create() {
	c.TplName = "views/tabler/admin/docs/create.tpl"
}

func (c *DocsController) Add() {
	title := c.GetString("title")
	content := c.GetString("content")

	if title == "" {
		c.Error(&c.Controller, "标题不能为空")
		return
	}

	if content == "" {
		c.Error(&c.Controller, "内容不能为空")
		return
	}

	doc := models.NewDocs()
	doc.Date = time.Now().Format("2006-01-02 15:04:05")
	doc.Content = content
	doc.Title = title

	if _, err := doc.Save(); err != nil {
		c.Error(&c.Controller, "文档添加失败")
		return
	}

	c.Success(&c.Controller, "文档添加成功")
}

// 使用LLM生成文档
func (c *DocsController) Generate() {
	question := c.GetString("question")

	c.JSONResp(map[string]any{
		"ret":      1,
		"msg":      "文档生成成功",
		"question": question,
	})
}

// 文档编辑页面
func (c *DocsController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, err.Error())
		return
	}

	doc, err := models.NewDocs().FindById(id)
	if err != nil {
		c.Error(&c.Controller, "文档更新失败")
		return
	}

	c.Data["doc"] = doc
	c.TplName = "views/tabler/admin/docs/edit.tpl"
}

// 后台编辑文档提交
func (c *DocsController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	title := c.GetString("title")
	content := c.GetString("content")

	doc, err := models.NewDocs().FindById(id)
	if err != nil {
		c.Error(&c.Controller, "公告更新失败")
		return
	}

	doc.Title = title
	doc.Content = content
	doc.Date = time.Now().Format("2006-01-02 15:04:05")

	if _, err := doc.Update(); err != nil {
		c.Error(&c.Controller, "文档更新失败")
		return
	}

	// TODO notify to telegram
	if models.NewConfig().Obtain("enable_telegram").ValueToBool() {

	}

	c.Success(&c.Controller, "文档更新成功")
}

// 后台删除文档
func (c *DocsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "删除失败，参数错误")
		return
	}
	doc, _ := models.NewDocs().FindById(id)
	if _, err = doc.Delete(); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}

	c.Success(&c.Controller, "删除成功")
}

// 后台文档页面 AJAX
func (c *DocsController) Ajax() {
	docs := models.NewDocs().FetchAll()

	for i, v := range docs {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-doc-%d" 
            onclick="deleteDoc(%d)">删除</button>
            <a class="btn btn-blue" href="/admin/docs/%d/edit">编辑</a>
		`, v.Id, v.Id, v.Id)
		docs[i].Op = op
	}

	c.JSONResp(map[string]any{
		"docs": docs,
	})
}
