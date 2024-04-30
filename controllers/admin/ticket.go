package admin

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type TicketController struct {
	controllers.BaseController
}

var ticket_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "工单ID"),
		NewField("Title", "主题"),
		NewField("Status", "工单状态"),
		NewField("Type", "工单类型"),
		NewField("UserId", "提交用户"),
		NewField("DatetimeStr", "创建时间"),
	},
}

var err_msg = "请求失败"

func (c *TicketController) Index() {
	c.Data["details"] = ticket_menu_details
	c.TplName = "views/tabler/admin/ticket/index.tpl"
}

func (c *TicketController) Update() {
	comment := c.GetString("comment")
	if comment == "" {
		c.Error(&c.Controller, err_msg)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ticket := models.NewTicket()
	if err := ticket.FindById(id); err != nil {
		c.Error(&c.Controller, err_msg)
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
	// Notify user
	// TODO
	// user := models.NewUser()
	// if err := user.FindById(ticket.UserId); err != nil {
	// 	c.Error(&c.Controller, "请求失败")
	// 	return
	// }
	//  Notification::notifyUser(
	// 	// $user,
	// 	$_ENV['appName'] . '-工单被回复',
	// 	'你好，有人回复了<a href="' . $_ENV['baseUrl'] . '/user/ticket/' . $ticket->id . '/view">工单</a>，请你查看。'
	// );

	c.Success(&c.Controller, "提交成功")
}
func (c *TicketController) Add() {
	c.TplName = "admin/ticket/add.tpl"
}

func (c *TicketController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ticket := models.NewTicket()
	if err := ticket.FindById(id); err != nil {
		c.Redirect("/admin/ticket", 302)
		return
	}

	comments := make([]map[string]any, 0)
	json.Unmarshal([]byte(ticket.Content), &comments)

	for i, v := range comments {
		comments[i]["datetime"] = time.UnixMilli(int64(v["datetime"].(float64))).Format("2006-01-02 15:04:05")
	}

	c.Data["ticket"] = ticket
	c.Data["comments"] = comments
	c.TplName = "views/tabler/admin/ticket/view.tpl"
}

func (c *TicketController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if _, err := models.NewTicket().DeleteById(id); err != nil {
		c.Error(&c.Controller, err_msg)
		return
	}

	c.Success(&c.Controller, "删除成功")
}

func (c *TicketController) Close() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ticket := models.NewTicket()
	if err := ticket.FindById(id); err != nil {
		c.Error(&c.Controller, err_msg)
		return
	}

	if ticket.Status == "closed" {
		c.Error(&c.Controller, "工单已关闭")
		return
	}

	ticket.Status = "closed"
	if err := ticket.Update(); err != nil {
		c.Error(&c.Controller, err_msg)
		return
	}

	c.Success(&c.Controller, "关闭成功")
}

func (c *TicketController) UpdateAi() {
	// TODO Update AI
	c.TplName = "admin/"
}

func (c *TicketController) Ajax() {

	tickets := models.NewTicket().FetchAll()

	for i, v := range tickets {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-ticket" 
            onclick="deleteTicket(%d)">删除</button>
		`, v.Id)

		if v.Status != "closed" {
			op += fmt.Sprintf(`
			<button type="button" class="btn btn-green" id="close-ticket" 
            onclick="closeTicket(%d)">关闭</button>
			`, v.Id)
		}

		op += fmt.Sprintf(`
		<a class="btn btn-blue" href="/admin/ticket/%d/view">查看</a>`, v.Id)

		tickets[i].Op = op
		tickets[i].Type = v.ParseType()
		tickets[i].Status = v.ParseStatus()
		tickets[i].DatetimeStr = time.UnixMilli(v.Datetime).Format("2006-01-02 15:04:05")
	}

	c.JSONResp(map[string]any{
		"tickets": tickets,
	})
}
