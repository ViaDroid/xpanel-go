package subscribe

import (
	"encoding/base64"
	"fmt"

	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type SS struct {
}

func (s *SS) GetContent(user *models.User) string {
	links := ""
	if !isEnabled("enable_ss_sub") {
		return links
	}

	noeds := services.NewSubscribeService().GetUserNodes(user, false)

	for _, v := range noeds {
		if v.Sort == 0 {
			item := fmt.Sprintf("%s:%s@%s:%d", user.Method, user.Passwd, v.Server, user.Port)
			links += fmt.Sprintf("%s#%s\n", base64.StdEncoding.EncodeToString([]byte(item)), v.Name)
		}
	}

	return links
}
