package admin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type AnnController struct {
	controllers.BaseController
}

var ann_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "公告ID"),
		NewField("Date", "日期"),
		NewField("Content", "公告内容"),
	},
}

var ann_update_fields = []string{
	"email_notify_class",
}

// 后台公告页面
func (c *AnnController) Index() {
	c.Data["details"] = ann_menu_details
	c.TplName = "views/tabler/admin/announcement/index.tpl"
}

// 后台公告创建页面
func (c *AnnController) Create() {
	c.Data["update_field"] = ann_update_fields
	c.TplName = "views/tabler/admin/announcement/create.tpl"
}

// 后台添加公告
func (c *AnnController) Add() {
	// email_notify_class, _ := c.GetInt("email_notify_class")
	email_notify, _ := c.GetBool("email_notify")
	content := c.GetString("content")

	if content != "" {
		ann := models.NewAnn()
		ann.Date = time.Now().Format("2006-01-02 15:04:05")
		ann.Content = content

		if _, err := ann.Save(); err != nil {
			c.Error(&c.Controller, "公告保存失败")
			return
		}

	}

	if email_notify {
		// TODO notify to email

	}

	// TODO notify to telegram
	if models.NewConfig().Obtain("enable_telegram").ValueToBool() {

	}

	if true {
		c.Success(&c.Controller, "公告添加成功，邮件发送成功")
	} else {
		c.Success(&c.Controller, "公告添加成功")
	}
}

// 后台公告编辑页面
func (c *AnnController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, err.Error())
		return
	}

	ann, err := models.NewAnn().FindById(id)
	if err != nil {
		c.Error(&c.Controller, "查询公告失败")
		return
	}

	c.Data["ann"] = ann
	c.TplName = "views/tabler/admin/announcement/edit.tpl"
}

// update
func (c *AnnController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	content := c.GetString("content")

	ann, err := models.NewAnn().FindById(id)
	if err != nil {
		c.Error(&c.Controller, "公告更新失败")
		return
	}

	ann.Date = time.Now().Format("2006-01-02 15:04:05")
	ann.Content = content

	if _, err := ann.Update(); err != nil {
		c.Error(&c.Controller, "公告更新失败")
		return
	}

	// TODO notify to telegram
	if models.NewConfig().Obtain("enable_telegram").ValueToBool() {

	}

	c.Success(&c.Controller, "公告更新成功")
}

// 后台公告删除
func (c *AnnController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "删除失败，参数错误")
		return
	}
	ann, _ := models.NewAnn().FindById(id)
	if _, err = ann.Delete(); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}

	c.Success(&c.Controller, "删除成功")
}

// 后台公告详情页面
func (c *AnnController) Ajax() {
	anns := models.NewAnn().FetchAll()

	for i, v := range anns {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-announcement-%d" 
            onclick="deleteAnn(%d)">删除</button>
            <a class="btn btn-blue" href="/admin/announcement/%d/edit">编辑</a>
		`, v.Id, v.Id, v.Id)
		anns[i].Op = op
	}

	c.JSONResp(map[string]any{
		"anns": anns,
	})
}
