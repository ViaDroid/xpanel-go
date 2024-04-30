package subscribe

import "github.com/viadroid/xpanel-go/models"

type base interface {
	GetContent(*models.User) string
}

func isEnabled(subType string) bool {
	return models.NewConfig().Obtain(subType).ValueIsOne()
}
