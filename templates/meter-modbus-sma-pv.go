package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA SunnyBoy / TriPower / other SunSpec PV-inverters (PV Meter)",
		Sample: `uri:192.0.2.2:502
id: 126 # ModBus slave id
model: sma-sunspec
power: Power # default value, optionally override
energy: Sum # energy value (Zählerstand)`,
	}

	registry.Add(template)
}
