package user

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type TicketController struct {
	controllers.BaseController
}

func (c *TicketController) Index() {

	if !models.NewConfig().Obtain("enable_ticket").ValueToBool() {
		c.Redirect("/user", 302)
		return
	}

	tickets := models.NewTicket().FetchByUserId(c.User.Id)

	for i, v := range tickets {
		tickets[i].DatetimeStr = time.UnixMilli(v.Datetime).Format("2006-01-02 15:04:05")
	}

	c.Data["tickets"] = tickets
	c.TplName = "views/tabler/user/ticket/index.tpl"
}

func (c *TicketController) Add() {
	title := c.GetString("title")
	comment := c.GetString("comment")
	typeStr := c.GetString("type")

	if !models.NewConfig().Obtain("enable_ticket").ValueToBool() || c.User.IsShadowBanned ||
		// !services.NewRateLimit().CheckTicketLimit(c.User.Id) ||
		title == "" || comment == "" || typeStr == "" {
		c.Error(&c.Controller, "请求失败")
		return
	}

	contents := []any{
		map[string]any{
			"comment_id":     int(0),
			"commenter_name": c.User.UserName,
			"comment":        comment,
			"datetime":       time.Now().UnixMilli(),
		},
	}

	bs, _ := json.Marshal(contents)

	ticket := models.Ticket{
		UserId:   c.User.Id,
		Title:    title,
		Type:     typeStr,
		Status:   "open_wait_admin",
		Content:  string(bs),
		Datetime: time.Now().UnixMilli(),
	}

	if err := ticket.Save(); err != nil {
		c.Error(&c.Controller, "请求失败")
		return
	}

	// Notify admin
	if models.NewConfig().Obtain("mail_ticket").ValueToBool() {
		// TODO
	}
	c.Success(&c.Controller, "提交成功")
}

func (c *TicketController) Create() {

}

func (c *TicketController) Detail() {
	if !models.NewConfig().Obtain("enable_ticket").ValueToBool() {
		c.Redirect("/user", 302)
		return
	}
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "请求失败")
		return
	}
	ticket := models.NewTicket()
	if err := ticket.FindById(id); err != nil {
		c.Redirect("/user/ticket", 302)
		return
	}

	comments := make([]map[string]any, 0)
	json.Unmarshal([]byte(ticket.Content), &comments)

	for i, v := range comments {
		comments[i]["datetime"] = time.UnixMilli(int64(v["datetime"].(float64))).Format("2006-01-02 15:04:05")
	}

	ticket.DatetimeStr = time.UnixMilli(ticket.Datetime).Format("2006-01-02 15:04:05")
	ticket.Type = ticket.ParseType()
	// ticket.Status = ticket.ParseStatus()

	c.Data["ticket"] = ticket
	c.Data["comments"] = comments
	c.TplName = "views/tabler/user/ticket/view.tpl"
}

func (c *TicketController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "请求失败")
		return
	}

	comment := c.GetString("comment")

	if !models.NewConfig().Obtain("enable_ticket").ValueToBool() || c.User.IsShadowBanned ||
		comment == "" {
		c.Error(&c.Controller, "请求失败")
		return
	}

	ticket := models.NewTicket()
	if err := ticket.FindById(id); err != nil {
		c.Error(&c.Controller, "请求失败")
		return
	}

	comments := make([]map[string]any, 0)
	json.Unmarshal([]byte(ticket.Content), &comments)

	contentItem := map[string]any{
		"comment_id":     comments[len(comments)-1]["comment_id"].(float64) + 1,
		"commenter_name": c.User.UserName,
		"comment":        comment,
		"datetime":       time.Now().UnixMilli(),
	}

	bs, _ := json.Marshal(append(comments, contentItem))

	ticket.Content = string(bs)
	ticket.Status = "open_wait_admin"

	if err := ticket.Update(); err != nil {
		c.Error(&c.Controller, "请求失败")
		return
	}
	// Notify admin
	if models.NewConfig().Obtain("mail_ticket").ValueToBool() {
		// TODO
	}
	c.Success(&c.Controller, "提交成功")

}
