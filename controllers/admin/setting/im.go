package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type ImController struct {
	controllers.BaseController
}

var im_update_field = []string{
        "enable_telegram",
        "telegram_token",
        "telegram_chatid",
        "telegram_bot",
        "telegram_request_token",
        "telegram_add_node",
        "telegram_add_node_text",
        "telegram_update_node",
        "telegram_update_node_text",
        "telegram_delete_node",
        "telegram_delete_node_text",
        "telegram_node_gfwed",
        "telegram_node_gfwed_text",
        "telegram_node_ungfwed",
        "telegram_node_ungfwed_text",
        "telegram_node_offline",
        "telegram_node_offline_text",
        "telegram_node_online",
        "telegram_node_online_text",
        "telegram_daily_job",
        "telegram_daily_job_text",
        "telegram_diary",
        "telegram_diary_text",
        "telegram_unbind_kick_member",
        "telegram_group_bound_user",
        "enable_welcome_message",
        "telegram_group_quiet",
        "allow_to_join_new_groups",
        "group_id_allowed_to_join",
        "help_any_command",
        "user_not_bind_reply",
        "discord_bot_token",
        "discord_client_id",
        "discord_client_secret",
        "discord_guild_id",
        "slack_token",
        "slack_client_id",
        "slack_client_secret",
        "slack_team_id",
}

func (c *ImController) Index() {
	settings := models.NewConfig().FindByClass("im")

	c.Data["update_field"] = im_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/im.tpl"
}


func (c *ImController) Save() {
	for _, item := range im_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}


// TODO Test IM (Slack, Telegram, Discord)