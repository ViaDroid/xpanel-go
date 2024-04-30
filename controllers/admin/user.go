package admin

import (
	"fmt"
	"strconv"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

func NewField(key, value string) *controllers.Field {
	return &controllers.Field{
		Key:   key,
		Value: value,
	}
}

// 用户详细信息字段
type UserDetailsField struct {
	Op             string // 操作类型（例如“添加”、“编辑”）
	Id             string // 用户 ID
	UserName       string // 昵称
	Email          string // 电子邮件地址
	Money          string // 余额
	RefBy          string // 推荐人 ID
	TransferEnable string // 流量限制
	TransferUsed   string // 当前使用量
	Class          string // 用户等级
	IsAdmin        string // 是否管理员
	IsBanned       string // 是否被封禁
	IsInactive     string // 是否闲置
	RegDate        string // 注册日期
	ClassExpire    string // 等级到期时间
}

// 创建用户对话框字段
type CreateDialogField struct {
	Id          string // 字段 ID
	Info        string // 字段描述
	Type        string // 字段类型（例如“input”）
	Placeholder string // 占位符文本
}

// var detailMap = map[string]string{
// 	"Op":             "操作",
// 	"Id":             "用户ID",
// 	"UserName":       "昵称",
// 	"Email":          "邮箱",
// 	"Money":          "余额",
// 	"RefBy":          "邀请人",
// 	"TransferEnable": "流量限制",
// 	"TransferUsed":   "当期用量",
// 	"Class":          "等级",
// 	"IsAdmin":        "是否管理员",
// 	"IsBanned":       "是否封禁",
// 	"IsInactive":     "是否闲置",
// 	"RegDate":        "注册时间",
// 	"ClassExpire":    "等级过期",
// }

var create_dialog = []any{
	map[string]any{
		"Id":          "email",
		"Info":        "登录邮箱",
		"Type":        "input",
		"Placeholder": "",
	},
	map[string]any{
		"Id":          "password",
		"Info":        "登录密码",
		"Type":        "input",
		"Placeholder": "留空则随机生成",
	},
	map[string]any{
		"Id":          "ref_by",
		"Info":        "邀请人",
		"Type":        "input",
		"Placeholder": "邀请人的用户id，可留空",
	},
	map[string]any{
		"Id":          "balance",
		"Info":        "账户余额",
		"Type":        "input",
		"Placeholder": "-1为按默认设置，其他为指定值",
	},
}

var detailListMap = []any{
	NewField("Op", "操作"),
	NewField("Id", "用户ID"),
	NewField("UserName", "昵称"),
	NewField("Email", "邮箱"),
	NewField("Money", "余额"),
	NewField("RefBy", "邀请人"),
	NewField("TransferEnableStr", "流量限制"),
	NewField("TransferUsed", "当期用量"),
	NewField("Class", "等级"),
	NewField("IsAdminStr", "是否管理员"),
	NewField("IsBannedStr", "是否封禁"),
	NewField("IsInactiveStr", "是否闲置"),
	NewField("RegDate", "注册时间"),
	NewField("ClassExpire", "等级过期"),
}

// var createDialogList = []any{
// 	CreateDialogField{
// 		Id:          "email",
// 		Info:        "登录邮箱",
// 		Type:        "input",
// 		Placeholder: "",
// 	},
// 	CreateDialogField{
// 		Id:          "password",
// 		Info:        "登录密码",
// 		Type:        "input",
// 		Placeholder: "留空则随机生成",
// 	},
// 	CreateDialogField{
// 		Id:          "ref_by",
// 		Info:        "邀请人",
// 		Type:        "input",
// 		Placeholder: "邀请人的用户id，可留空",
// 	},
// 	CreateDialogField{
// 		Id:          "balance",
// 		Info:        "账户余额",
// 		Type:        "input",
// 		Placeholder: "-1为按默认设置，其他为指定值",
// 	},
// }

var details = map[string]any{
	"field":         detailListMap,
	"create_dialog": create_dialog,
}
var user_update_field = []string{
	"email",
	"user_name",
	"remark",
	"pass",
	"money",
	"is_admin",
	"ga_enable",
	"is_banned",
	"banned_reason",
	"is_shadow_banned",
	"transfer_enable",
	"ref_by",
	"class_expire",
	"node_group",
	"class",
	"auto_reset_day",
	"auto_reset_bandwidth",
	"node_speedlimit",
	"node_iplimit",
	"port",
	"passwd",
	"method",
}

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Index() {
	// jsonBytes, _ := json.Marshal(details2)
	// jsonStr := string(jsonBytes)
	c.Data["details"] = details
	c.TplName = "views/tabler/admin/user/index.tpl"
}

func (c *UserController) Create() {
	email := c.GetString("email")
	ref_by := c.GetString("ref_by")
	password := c.GetString("password")
	balanceStr := c.GetString("balance")

	balance := 0
	if balanceStr != "" {
		b, err := strconv.Atoi(balanceStr)
		if err != nil {
			c.Error(&c.Controller, "余额必须是数字")
			return
		}
		balance = b
	}

	refById := 0
	if ref_by != "" {
		refById, err := strconv.Atoi(ref_by)
		if err != nil {
			c.Error(&c.Controller, "邀请人必须是数字")
			return
		}

		// 邀请人不存在
		refBy := models.NewUser().FindById(refById)
		if refBy == nil {
			c.Error(&c.Controller, "邀请人不存在")
			return
		}
	}

	if email == "" {
		c.Error(&c.Controller, "邮箱不能为空")
		return
	}

	// 邮箱已存在
	user := models.NewUser()
	if user.IsExist(email) {
		c.Error(&c.Controller, "邮箱已存在")
		return
	}

	if password == "" {
		password = utils.GenRandomString(16)
	}

	controllers.RegisterHelper(&c.BaseController, "user", email, password, "", 0, "", balance, true)

	user.FindByEmail(email)

	if ref_by != "" {
		user.RefBy = refById
		user.Update()
	}

	c.Success(&c.Controller, fmt.Sprintf("添加成功，用户邮箱: %s，密码：%s", email, password))
}

func (c *UserController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.Error(&c.Controller, err.Error())
		return
	}
	user := models.NewUser().FindById(id)
	c.Data["update_field"] = user_update_field
	c.Data["edit_user"] = user
	c.TplName = "views/tabler/admin/user/edit.tpl"
}

func (c *UserController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	user := models.NewUser().FindById(id)

	// 用户不存在
	if user == nil {
		c.Error(&c.Controller, "用户不存在")
		return
	}

	pass := c.GetString("pass")
	if pass != "" {
		user.Pass = utils.PasswordHash(pass)
		if models.NewConfig().Obtain("enable_forced_replacement").ValueToBool() {
			user.RemoveLink()
		}
	}

	money, err := c.GetFloat("money")
	if err != nil {
		c.Error(&c.Controller, "余额必须是数字")
		return
	}

	if user.Money != money {
		diff := money - user.Money

		remark := "管理员扣除余额"
		if diff > 0 {
			remark = "管理员增加余额"
		}

		models.NewUserMoneyLog().Add(user.Id, user.Money, money, diff, remark)

		user.Money = money
	}

	user.Email = c.GetString("email")
	user.UserName = c.GetString("user_name")
	user.Remark = c.GetString("remark")
	user.IsAdmin, _ = c.GetBool("is_admin", false)
	user.GaEnable, _ = c.GetBool("ga_enable")
	user.IsBanned, _ = c.GetBool("is_banned")
	user.BannedReason = c.GetString("banned_reason")
	user.IsShadowBanned, _ = c.GetBool("is_shadow_banned")
	user.TransferEnable = int64(utils.AutoBytesR(c.GetString("transfer_enable")))
	user.RefBy, _ = c.GetInt("ref_by")
	user.ClassExpire = c.GetString("class_expire")
	user.NodeGroup, _ = c.GetInt("node_group")
	user.Class, _ = c.GetInt("class")
	user.AutoResetDay, _ = c.GetInt("auto_reset_day")
	user.AutoResetBandwidth, _ = c.GetFloat("auto_reset_bandwidth")
	user.NodeSpeedlimit, _ = c.GetFloat("node_speedlimit")
	user.NodeIplimit, _ = c.GetInt("node_iplimit")
	user.Port, _ = c.GetInt("port")
	user.Method = c.GetString("method")

	if _, err := user.Update(); err != nil {
		c.Error(&c.Controller, "修改失败")
		return
	}

	c.Success(&c.Controller, "修改成功")

}

func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, err.Error())
		return
	}

	user := models.User{Id: id}
	if _, err := user.Kill(); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}

	c.Success(&c.Controller, "删除成功")
}

func (c *UserController) Ajax() {
	users := models.NewUser().FetchAll()

	for i, user := range users {
		op := fmt.Sprintf(`<button type="button" class="btn btn-red" id="delete-user-%d"
		onclick="deleteUser(%d)">删除</button>
		<a class="btn btn-blue" href="/admin/user/%d/edit">编辑</a>`, user.Id, user.Id, user.Id)

		users[i].Op = op
		users[i].Parse()
	}

	c.JSONResp(map[string]any{
		"users": users,
	})
}
