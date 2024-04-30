package services

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
)

type SubscribeService struct{}

func NewSubscribeService() *SubscribeService {
	return &SubscribeService{}
}

func (s *SubscribeService) GetUniversalSubLink(user *models.User) string {

	link, err := models.NewLink().FindByUserId(user.Id)
	if err != nil {
		link = models.NewLink()
		link.Token = ""
		link.UserId = user.Id
		link.Save()
	}

	subUrl := global.ConfMap["subUrl"]
	return fmt.Sprintf("%s/sub/%s", subUrl, link.Token)
}

func (s *SubscribeService) GetUserNodes(user *models.User, showAll bool) []models.Node {
	var list []models.Node

	querySetter := global.DB.QueryTable(models.Node{}).Filter("type", 1)

	if !showAll {
		querySetter.Filter("node_class__lte", user.Class)
	}

	if !user.IsAdmin {
		group := []int{0}
		if user.NodeGroup != 0 {
			group = []int{user.NodeGroup, 0}
		}

		querySetter.Filter("node_group__in", group)
	}

	cond := orm.NewCondition()
	cond.And("node_bandwidth_limit", 0) //.Or("node_bandwidth__lt", "`node_bandwidth_limit`")
	querySetter.SetCond(cond)

	// querySetter.GetCond()
	querySetter.OrderBy("node_class").OrderBy("name").All(&list)
	return list

}

func (s *SubscribeService) GetContent(user *models.User) string {
	return ""
}


func (s *SubscribeService) GetClient() {

}