package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "tplink",
		Name:   "TP-LINK Smart Plug",
		Sample: `uri: 192.0.2.2:9999 # TP-LINK Smart Plug ip address + port (local)
standbypower: 15 # standbypower threshold / trickle charge of used charger in W`,
	}

	registry.Add(template)
}
