package user

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

type InviteController struct {
	controllers.BaseController
}

func (c *InviteController) Index() {

	inviteCode := models.NewInviteCode().FindByUserId(c.User.Id)
	if inviteCode.Id == 0 {
		inviteCode.Add(c.User.Id)
	}
	paybacks := models.NewPayback().FindByRefBy(c.User.Id)
	paybacks_sum := models.NewPayback().SumByRefBy(c.User.Id)

	invite_url := fmt.Sprintf("%s/auth/register?code=%s", global.ConfMap["baseUrl"], inviteCode.Code)
	invite_reward_rate := models.NewConfig().Obtain("invite_reward_rate").ValueToInt() * 100

	c.Data["paybacks"] = paybacks
	c.Data["invite_url"] = invite_url
	c.Data["paybacks_sum"] = paybacks_sum
	c.Data["invite_reward_rate"] = invite_reward_rate

	c.TplName = "views/tabler/user/invite.tpl"
}

func (c *InviteController) Reset() {

}
