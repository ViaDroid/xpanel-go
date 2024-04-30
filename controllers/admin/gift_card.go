package admin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type GiftCardController struct {
	controllers.BaseController
}

var gift_card_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "礼品卡ID"),
		NewField("Card", "卡号"),
		NewField("Balance", "面值"),
		NewField("CreateTimeStr", "创建时间"),
		NewField("StatusStr", "使用状态"),
		NewField("UseTimeStr", "使用时间"),
		NewField("UseUser", "使用用户"),
	},
	"create_dialog": []any{
		map[string]any{
			"id":          "card_number",
			"info":        "创建数量",
			"type":        "input",
			"placeholder": "",
		},
		map[string]any{
			"id":          "card_value",
			"info":        "礼品卡面值",
			"type":        "input",
			"placeholder": "",
		},
		map[string]any{
			"id":   "card_length",
			"info": "礼品卡长度",
			"type": "select",
			"select": map[string]any{
				"12": "12位",
				"18": "18位",
				"24": "24位",
				"30": "30位",
				"36": "36位",
			},
		},
	},
}

func (c *GiftCardController) Index() {
	c.Data["details"] = gift_card_menu_details
	c.TplName = "views/tabler/admin/giftcard.tpl"
}

func (c *GiftCardController) Add() {
	card_number, _ := c.GetInt("card_number", 0)
	card_value, _ := c.GetFloat("card_value", 0)
	card_length, _ := c.GetInt("card_length", 0)

	if card_number <= 0 {
		c.Error(&c.Controller, "生成数量不能为空或小于0")
		return
	}

	if card_value <= 0 {
		c.Error(&c.Controller, "礼品卡面值不能为空或小于0")
		return
	}

	if card_length <= 0 {
		c.Error(&c.Controller, "礼品卡长度不能为空或小于0")
		return
	}

	card_added := ""

	// 开启事务
	tx, _ := global.DB.Begin()
	for i := 0; i < card_number; i++ {
		now := time.Now()
		card := &models.GiftCard{
			Card:       utils.GenRandomString(card_length),
			Balance:    card_value,
			CreateTime: now.UnixMilli(),
			Status:     0,
			UseTime:    0,
			UseUser:    0,
		}

		if _, err := tx.Insert(card); err != nil {
			tx.Rollback()
			c.Error(&c.Controller, "添加失败")
			return
		}

		card_added += card.Card + ","
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "添加失败")
		return
	}

	c.Success(&c.Controller, "添加成功, "+card_added)
}

func (c *GiftCardController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}
	if _, err := models.NewGiftCard().Delete(id); err != nil {
		c.Error(&c.Controller, "删除失败")
		return
	}
	c.Success(&c.Controller, "删除成功")
}

func (c *GiftCardController) Ajax() {
	list := models.NewGiftCard().FetchAll()
	for i, v := range list {
		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-gift-card-%d" 
        onclick="deleteGiftCard(%d)">删除</button>
		`, v.Id, v.Id)
		list[i].Op = op
		list[i].StatusStr = v.ParseStatus()
		list[i].CreateTimeStr = time.UnixMilli(v.CreateTime).Format("2006-01-02 15:04:05")
		if v.UseTime == 0 {
			list[i].UseTimeStr = "-"
		} else {
			list[i].UseTimeStr = time.UnixMilli(v.UseTime).Format("2006-01-02 15:04:05")
		}
	}

	c.JSONResp(map[string]any{
		"giftcards": list,
	})
}
