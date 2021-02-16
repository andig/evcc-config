package templates

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class: "meter",
		Type:  "modbus",
		Name:  "Kostal Inverter (PV Meter)",
		Sample: `uri: 192.0.2.2:1502 # inverter port
id: 71`,
	}

	registry.Add(template)
}
