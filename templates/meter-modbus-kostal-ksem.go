package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Smart Energy Meter (Grid Meter)",
		Sample: `model: kostal
uri: 192.0.2.2:502
id: 71
power: Power
energy: Energy`,
	}

	registry.Add(template)
}
