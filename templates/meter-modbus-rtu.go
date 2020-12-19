package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Modbus (RTU)",
		Sample: `model: sdm
uri: 192.0.2.2:502
rtu: true # rs485 device connected using ethernet adapter
id: 2
power: Power # default value, optionally override
energy: Sum # energy value (Zählerstand)`,
	}

	registry.Add(template)
}
