package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis_rate/v10"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

var RateLimitTypes = struct {
	SUB_IP                int
	SUB_TOKEN             int
	WEBAPI_IP             int
	WEBAPI_KEY            int
	USER_API_IP           int
	USER_API_KEY          int
	ADMIN_API_IP          int
	ADMIN_API_KEY         int
	NODE_API_IP           int
	NODE_API_KEY          int
	EMAIL_REQUEST_IP      int
	EMAIL_REQUEST_ADDRESS int
	TICKET                int
}{
	SUB_IP:    0,
	SUB_TOKEN: 1,
	WEBAPI_IP: 2,
	WEBAPI_KEY: 3,
	USER_API_IP: 4,
	USER_API_KEY: 5,
	ADMIN_API_IP: 6,
	ADMIN_API_KEY: 7,
	NODE_API_IP: 8,
	NODE_API_KEY: 9,
	EMAIL_REQUEST_IP: 10,
	EMAIL_REQUEST_ADDRESS: 11,
	TICKET: 12,
}

var ratesMap = []map[string]string{
	{
		"rate":       "rate_limit_sub_ip",
		"prefix-key": "xpanel_sub_ip:",
	},
	{
		"rate":       "rate_limit_sub",
		"prefix-key": "xpanel_sub_token:",
	},
	{
		"rate":       "rate_limit_webapi_ip",
		"prefix-key": "xpanel_webapi_ip:",
	},
	{
		"rate":       "rate_limit_webapi",
		"prefix-key": "xpanel_webapi_key:",
	},
	{
		"rate":       "rate_limit_user_api_ip",
		"prefix-key": "xpanel_user_api_ip:",
	},
	{
		"rate":       "rate_limit_user_api",
		"prefix-key": "xpanel_user_api_key:",
	},
	{
		"rate":       "rate_limit_admin_api_ip",
		"prefix-key": "xpanel_admin_api_ip:",
	},
	{
		"rate":       "rate_limit_admin_api",
		"prefix-key": "xpanel_admin_api_key:",
	},
	{
		"rate":       "rate_limit_node_api_ip",
		"prefix-key": "xpanel_node_api_ip:",
	},
	{
		"rate":       "rate_limit_node_api",
		"prefix-key": "xpanel_node_api_key:",
	},
	{
		"rate":       "email_request_ip_limit",
		"prefix-key": "xpanel_email_request_ip:",
	},
	{
		"rate":       "email_request_address_limit",
		"prefix-key": "xpanel_email_request_address:",
	},
	{
		"rate":       "rate_limit_ticket",
		"prefix-key": "xpanel_ticket:",
	},
}

type RateLimit struct {}

func NewRateLimit() *RateLimit {
	return &RateLimit{}
}

func (s *RateLimit) CheckRateLimit(limit_type int, value string) bool {
	
	if limit_type < 0 || limit_type >= len(ratesMap) {
		return false
	}

	limiter := redis_rate.NewLimiter(global.Redis)

	rate := models.NewConfig().Obtain(ratesMap[limit_type]["rate"]).ValueToInt()
	limit := redis_rate.PerMinute(rate)

	_, err := limiter.Allow(context.Background(), fmt.Sprintf("%s%s", ratesMap[limit_type]["prefix-key"], value), limit)

	return err == nil
}

func (s *RateLimit) CheckEmailIpLimit(request_ip string) bool {
	// TODO
	return false
}

func (s *RateLimit) CheckEmailAddressLimit(request_address string) bool {
	// TODO
	return false
}

func (s *RateLimit) CheckTicketLimit(user_id int) bool {
	// TODO
	return false
}
