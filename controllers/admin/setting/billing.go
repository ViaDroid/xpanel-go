package setting

import (
	"encoding/json"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type BillingController struct {
	controllers.BaseController
}

var billing_update_field = []string{
	// 支付宝当面付
	"f2f_pay_app_id",
	"f2f_pay_pid",
	"f2f_pay_public_key",
	"f2f_pay_private_key",
	"f2f_pay_notify_url",
	// Stripe
	"stripe_api_key",
	"stripe_endpoint_secret",
	"stripe_currency",
	"stripe_card",
	"stripe_alipay",
	"stripe_wechat",
	"stripe_min_recharge",
	"stripe_max_recharge",
	// EPay
	"epay_url",
	"epay_pid",
	"epay_key",
	"epay_sign_type",
	"epay_alipay",
	"epay_wechat",
	"epay_qq",
	"epay_usdt",
	// PayPal
	"paypal_mode",
	"paypal_client_id",
	"paypal_client_secret",
	"paypal_currency",
	"paypal_locale",
}

func (c *BillingController) Index() {

	settings := models.NewConfig().FindByClass("billing")
	c.Data["update_field"] = billing_update_field
	c.Data["payment_gateways"] = returnGatewaysList()
	c.Data["active_payment_gateway"] = returnActiveGateways()
	c.Data["settings"] = settings

	c.TplName = "views/tabler/admin/setting/billing.tpl"
}

func (c *BillingController) Save() {
	gateway_in_use := []string{}
	for k := range returnGatewaysList() {
		payment_enable := c.GetString(k)

		if payment_enable == "true" {
			gateway_in_use = append(gateway_in_use, k)
		}
	}

	conf := models.NewConfig()
	gateway := conf.Obtain("payment_gateway")

	gateway_in_use_json_bs, _ := json.Marshal(gateway_in_use)
	gateway.Value = string(gateway_in_use_json_bs)

	tx, _ := global.DB.Begin()
	if _, err := tx.Update(&gateway); err != nil {
		tx.Rollback()
		c.Error(&c.Controller, "保存支付网关时出错")
		return
	}

	for _, v := range billing_update_field {
		var item models.Config
		global.DB.QueryTable(item).Filter("item", v).One(&item)
		item.Value = c.GetString(v)
		if _, err := tx.Update(&item); err != nil {
			tx.Rollback()
			c.Error(&c.Controller, "保存配置时出错")
			return
		}
	}
	if err := tx.Commit(); err != nil {
		c.Error(&c.Controller, "保存配置时出错")
		return
	}

	c.Success(&c.Controller, "保存成功")
}

func (c *BillingController) SetStripeWebhook() {
	// stripe_api_key := c.GetString("stripe_api_key")
	// Stripe::setApiKey($stripe_api_key);

}

func returnGatewaysList() map[string]string {
	list := map[string]string{
		// 支付宝当面付
		// "f2f": "支付宝当面付",
		// // Stripe
		// "stripe": "Stripe",
		// // EPay
		// "epay": "EPay",
		// // PayPal
		// "paypal": "PayPal",
	}
	for _, v := range services.NewPayment().GetAllPaymentMap() {
		list[v.Name()] = v.ReadableName()
	}
	return list
}

func returnActiveGateways() []string {
	value := models.NewConfig().Obtain("payment_gateway").Value

	var payment_gateways []string
	json.Unmarshal([]byte(value), &payment_gateways)
	return payment_gateways
}
