package setting

import (
	"fmt"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type EmailController struct {
	controllers.BaseController
}

var email_update_field = []string{
	"email_driver",
	"email_verify_code_ttl",
	"email_password_reset_ttl",
	"email_request_ip_limit",
	"email_request_address_limit",
	// SMTP
	"smtp_host",
	"smtp_username",
	"smtp_password",
	"smtp_port",
	"smtp_name",
	"smtp_sender",
	"smtp_ssl",
	"smtp_bbc",
	// Mailgun
	"mailgun_key",
	"mailgun_domain",
	"mailgun_sender",
	"mailgun_sender_name",
	// Sendgrid
	"sendgrid_key",
	"sendgrid_sender",
	"sendgrid_name",
	// AWS SES
	"aws_ses_access_key_id",
	"aws_ses_access_key_secret",
	"aws_ses_region",
	"aws_ses_sender",
	// Postal
	"postal_host",
	"postal_key",
	"postal_sender",
	"postal_name",
	// Mailchimp
	"mailchimp_key",
	"mailchimp_from_email",
	"mailchimp_from_name",
	// Alibaba Cloud
	"alibabacloud_dm_access_key_id",
	"alibabacloud_dm_access_key_secret",
	"alibabacloud_dm_endpoint",
	"alibabacloud_dm_account_name",
	"alibabacloud_dm_from_alias",
}

func (c *EmailController) Index() {
	settings := models.NewConfig().FindByClass("email")

	c.Data["update_field"] = email_update_field
	c.Data["settings"] = settings
	c.TplName = "views/tabler/admin/setting/email.tpl"
}

func (c *EmailController) Save() {
	for _, item := range email_update_field {
		value := c.GetString(item)
		if _, err := models.NewConfig().UpdateByItem(item, value); err != nil {
			c.Error(&c.Controller, fmt.Sprintf("保存 %s 时出错", item))
			return
		}
	}
	c.Success(&c.Controller, "保存成功")
}

func (c *EmailController) TestEmail() {
	// TODO Test Send Mail
}
