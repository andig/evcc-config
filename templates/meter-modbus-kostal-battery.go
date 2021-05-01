package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Hybrid Inverter (Battery Meter)",
		Sample: `uri: 192.0.2.2:1502
id: 71
power: 802:W
soc: 802:SoC`,
	}

	registry.Add(template)
}
