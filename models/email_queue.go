package models

import (
	"encoding/json"
	"time"

	"github.com/viadroid/xpanel-go/global"
)

type EmailQueue struct {
	Id       int    `orm:"description(记录ID)"`
	ToEmail  string `orm:"description(收件人邮箱)"`
	Subject  string `orm:"description(邮件主题)"`
	Template string `orm:"description(邮件模板)"`
	Array    string `orm:"description(模板内容)"`
	Time     int64  `orm:"description(添加时间)"`
}

func NewEmailQueue() *EmailQueue {
	return &EmailQueue{}
}

func (item *EmailQueue) Add(to, subject, template string, body map[string]any) (int64, error) {
	if template == "" {
		template = "warn.tpl"
	}

	jsonBytes, _ := json.Marshal(body)

	eq := &EmailQueue{
		ToEmail:  to,
		Subject:  subject,
		Template: template,
		Time:     time.Now().UnixMilli(),
		Array:    string(jsonBytes),
	}
	return global.DB.Insert(eq)
}
