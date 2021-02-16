package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Inverter (PV Meter)",
		Sample: `model: kostal
uri: 192.0.2.2:1502
id: 71
power: Power`,
	}

	registry.Add(template)
}
