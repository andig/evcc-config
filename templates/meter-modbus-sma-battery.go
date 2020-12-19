package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA Sunny Boy Storage (Battery)",
		Sample: `uri: 192.0.2.2:502
id: 126 # ModBus slave id
model: sma-sunspec
power: Power # default value, optionally override
soc: ChargeState # battery soc (Ladezustand)`,
	}

	registry.Add(template)
}
