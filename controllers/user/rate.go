package user

import (
	"encoding/json"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/services"
)

type RateController struct {
	controllers.BaseController
}

func (c *RateController) Index() {
	nodes := services.NewSubscribeService().GetUserNodes(c.User, false)
	var nodeList []map[string]any
	for _, v := range nodes {
		nodeList = append(nodeList, map[string]any{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	if len(nodes) == 0 {
		nodeList = append(nodeList, map[string]any{
			"id":   0,
			"name": "暂无节点",
		})
	}
	c.Data["node_list"] = nodeList
	c.TplName = "views/tabler/user/rate.tpl"
}
func (c *RateController) Ajax() {
	event := map[string]any{
		"drawChart": map[string]any{
			// TODO
			// "msg":  node,
			// "data": rates,
		},
	}
	bts, _ := json.Marshal(event)

	eventJsonStr := string(bts)
	c.Ctx.Output.Header("HX-Trigger", eventJsonStr)
	c.Success(&c.Controller, "")
}
