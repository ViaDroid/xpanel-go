package user

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/services"
	"github.com/viadroid/xpanel-go/utils"
)

type ServerController struct {
	controllers.BaseController
}

func (c *ServerController) Index() {

	nodes := services.NewSubscribeService().GetUserNodes(c.User, true)

	var node_list []map[string]any
	for _, node := range nodes {
		node.Parse()
		node_bandwidth_limit := "无限制"
		if node.NodeBandwidthLimit == 0 {
			node_bandwidth_limit = utils.AutoBytes(node.NodeBandwidthLimit)
		}
		node_list = append(node_list, map[string]any{
			"id":                   node.Id,
			"name":                 node.Name,
			"class":                node.NodeClass,
			"color":                node.Color(),
			"connection_type":      node.ConnectionType(),
			"sort":                 node.SortStr,
			"online_user":          node.OnlineUser,
			"online":               node.GetNodeOnlineStatus(),
			"traffic_rate":         node.TrafficRate,
			"is_dynamic_rate":      node.IsDynamicRate,
			"node_bandwidth":       utils.AutoBytes(node.NodeBandwidth),
			"node_bandwidth_limit": node_bandwidth_limit,
		})

	}
	c.Data["servers"] = node_list

	c.TplName = "views/tabler/user/server.tpl"
}
