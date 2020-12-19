package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "phoenix-emcp",
		Name:   "Phoenix EM-CP Controller (Ethernet/Modbus TCP)",
		Sample: `uri: 192.0.2.2:502 # TCP ModBus address
id: 1`,
	}

	registry.Add(template)
}
