package services

import (
	"github.com/viadroid/xpanel-go/models"
)

type Notification struct{}

func NewNotification() *Notification {
	return &Notification{}
}

func (*Notification) NotifyUser(user *models.User, title, msg, template string) {
	if template == "" {
		template = "warn.tpl"
	}
	if user.ContactMethod == 1 && user.ImType == 0 {

		m := map[string]any{
			"user":  user,
			"title": title,
			"msg":   msg,
		}
		models.NewEmailQueue().Add(user.Email, title, template, m)
	} else {
		IM{}.Send(user.ImValue, msg, user.ImType)
	}
}
