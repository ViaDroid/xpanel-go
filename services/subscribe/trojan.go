package subscribe

import (
	"encoding/json"
	"fmt"

	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/services"
)

type Trojan struct{}

func (s *Trojan) GetContent(user *models.User) string {

	links := ""
	if isEnabled("enable_trojan_sub") {
		return links
	}

	nodes := services.NewSubscribeService().GetUserNodes(user, false)

	for _, v := range nodes {
		if v.Sort == 14 {
			var customConfig map[string]any
			json.Unmarshal([]byte(v.CustomConfig), &customConfig)
			trojan_port := customConfig["offset_port_user"]
			if trojan_port == "" {
				trojan_port = customConfig["offset_port_node"]
			}
			if trojan_port == "" {
				trojan_port = 443
			}

			host := customConfig["host"]
			allow_insecure := customConfig["allow_insecure"]
			if allow_insecure == "" {
				allow_insecure = "0"
			}
			security := customConfig["security"]
			if security == "" {
				security = "tls"
			}
			mux := customConfig["mux"]
			if mux == "" {
				mux = "0"
			}
			network := customConfig["network"]
			if network == "" {
				network = "tcp"
			}
			transport_plugin := customConfig["transport_plugin"]
			transport_method := customConfig["transport_method"]
			servicename := customConfig["servicename"]
			path := customConfig["path"]

			links += fmt.Sprintf("trojan://%s@%s:%d?peer=%s&sni=%s&obfs=%s&path=%s&mux=%s&allowInsecure=%s&obfsParam=%s&type=%s&security=%s&serviceName=%s#%s\n",
				user.Uuid, v.Server, trojan_port, host, host, transport_plugin, path, mux, allow_insecure, transport_method, network, security, servicename, v.Name)

		}
	}

	return ""
}
