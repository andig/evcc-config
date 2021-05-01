package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus Meter (RTU-over-TCP)",
		Sample: `model: sdm
uri: 192.0.2.2:502
rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
id: 1
energy: Sum # this assignment is only required for charge meter usage`,
	}

	registry.Add(template)
}
