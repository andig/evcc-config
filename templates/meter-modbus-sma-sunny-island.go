package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA Sunny Island (Battery)",
		Sample: `model: sunny-island
uri: 192.0.2.2:502
id: 126
power: Power # default values, optionally override
soc: ChargeState # battery soc (Ladezustand)`,
	}

	registry.Add(template)
}
