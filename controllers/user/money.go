package user

import (
	"fmt"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

type MoneyController struct {
	controllers.BaseController
}

func (c *MoneyController) Index() {
	list := models.NewUserMoneyLog().FindListByUserId(c.User.Id)
	for i := range list {
		list[i].Parse()

	}
	c.Data["moneylogs"] = list
	c.Data["moneylog_count"] = len(list)
	c.TplName = "views/tabler/user/money.tpl"
}

func (c *MoneyController) Giftcard() {
	giftcard := c.GetString("giftcard")
	card, err := models.NewGiftCard().FindByCard(giftcard)
	if err != nil || card.Status != 0 {
		c.Error(&c.Controller, "礼品卡无效")
		return
	}

	if c.User.IsShadowBanned {
		c.Error(&c.Controller, "礼品卡无效")
		return
	}

	tx, _ := global.DB.Begin()

	now := time.Now()
	card.Status = 1
	card.UseTime = now.UnixMilli()
	card.UseUser = c.User.Id
	if _, err := tx.Update(card); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "使用礼品卡失败")
		return
	}

	moneyBefore := c.User.Money
	c.User.Money += card.Balance
	if _, err := tx.Update(c.User); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "使用礼品卡失败")
		return
	}

	log := models.UserMoneyLog{
		UserId:     c.User.Id,
		Before:     moneyBefore,
		After:      c.User.Money,
		Amount:     card.Balance,
		Remark:     fmt.Sprintf("礼品卡充值 %s", card.Card),
		CreateTime: now.UnixMilli(),
	}
	if _, err := tx.Insert(&log); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "使用礼品卡失败")
		return
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "使用礼品卡失败")
		return
	}
	c.Success(&c.Controller, "充值成功")

}
