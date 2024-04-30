package models

import (
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/utils"
)

type Node struct {
	Id                     int     `orm:"description(节点ID)"`
	Name                   string  `orm:"description(节点名称)"`
	Type                   bool    `orm:"description(节点启用)"`
	Server                 string  `orm:"description(节点地址)"`
	CustomConfig           string  `orm:"type(json);description(自定义配置)"`
	Sort                   int     `orm:"description(节点类型)"`
	TrafficRate            float64 `orm:"description(流量倍率)"`
	IsDynamicRate          bool    `orm:"description(是否启用动态流量倍率)"`
	DynamicRateType        int     `orm:"description(动态流量倍率计算方式)"`
	DynamicRateConfig      string  `orm:"description(动态流量倍率配置)"`
	NodeClass              int     `orm:"description(节点等级)"`
	NodeSpeedlimit         float64 `orm:"description(节点限速)"`
	NodeBandwidth          int64   `orm:"description(节点流量)"`
	NodeBandwidthLimit     int64   `orm:"description(节点流量限制)"`
	BandwidthlimitResetday int     `orm:"description(流量重置日)"`
	NodeHeartbeat          int     `orm:"description(节点心跳)"`
	OnlineUser             int     `orm:"description(节点在线用户)"`
	Ipv4                   string  `orm:"description(IPv4地址)"`
	Ipv6                   string  `orm:"description(IPv6地址)"`
	NodeGroup              int     `orm:"description(节点群组)"`
	Online                 int     `orm:"description(在线状态)"`
	GfwBlock               int     `orm:"description(是否被GFW封锁)"`
	Password               string  `orm:"unique;description(后端连接密码)"`

	Op               string `orm:"-"`
	MaxRate          string `orm:"-"`
	MaxRateTime      string `orm:"-"`
	MinRate          string `orm:"-"`
	MinRateTime      string `orm:"-"`
	TypeStr          string `orm:"-"`
	SortStr          string `orm:"-"`
	IsDynamicRateStr string `orm:"-"`
}

func NewNode() *Node {
	return &Node{}
}

func (item *Node) Insert() (int64, error) {
	return global.DB.Insert(item)
}

func (item *Node) Update() (int64, error) {
	return global.DB.Update(item)
}

// Find By Id
func (item *Node) Find(id int) (*Node, error) {
	return item, global.DB.QueryTable(item).Filter("id", id).One(item)
}

// Delete by Id
func (item *Node) Delete(id int) (int64, error) {
	item.Id = id
	return global.DB.Delete(item)
}

func (item *Node) FetchAll() []Node {
	var nodes []Node
	global.DB.QueryTable(item).OrderBy("-id").All(&nodes)

	return nodes
}

func (item *Node) FindIp4Or6(ip string) (*Node, error) {
	err := global.DB.QueryTable(item).Filter("ipv4", ip).Filter("ipv6", ip).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Node) Color() string {
	var color = "red"
	switch item.GetNodeOnlineStatus() {
	case 0:
		color = "orange"
	case 1:
		color = "green"
	default:
		color = "red"
	}

	return color
}

// 获取节点在线状态
// @return int 0 = new node, -1 = offline, 1 = online
func (item *Node) GetNodeOnlineStatus() int {
	if item.NodeHeartbeat == 0 {
		return 0
	} else {
		if item.NodeHeartbeat+600 > int(time.Now().UnixMilli()) {
			return 1
		} else {
			return -1
		}
	}
}

// 0 = IPv4, 1 = IPv6, 2 = DualStack
func (item *Node) ConnectionType() int {
	if item.Ipv6 != "::1" && item.Ipv4 != "127.0.0.1" {
		return 2
	} else {
		if item.Ipv4 != "127.0.0.1" {
			return 0
		} else {
			return 1
		}
	}

}

func (item *Node) ParseSort() string {
	switch item.Sort {
	case 0:
		return "Shadowsocks"
	case 1:
		return "Shadowsocks2022"
	case 2:
		return "TUIC"
	case 3:
		return "WireGuard"
	case 11:
		return "Vmess"
	case 14:
		return "Trojan"
	default:
		return "未知"
	}
}

func (item *Node) DynamicRateTypeStr() string {
	switch item.DynamicRateType {
	case 0:
		return "Logistic"
	case 2:
		return "Linear"
	default:
		return "未知"
	}
}

func (item *Node) Parse() {
	if item.Type {
		item.TypeStr = "显示"
	} else {
		item.TypeStr = "隐藏"
	}

	item.SortStr = item.ParseSort()
	
	item.IsDynamicRateStr = func(v bool) string {
		if v {
			return "是"
		} else {
			return "否"
		}
	}(item.IsDynamicRate)

	item.NodeBandwidth = int64(utils.FlowToGB(float64(item.NodeBandwidth)))
	item.NodeBandwidthLimit = int64(utils.FlowToGB(float64(item.NodeBandwidthLimit)))
}
