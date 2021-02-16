package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus RTU Meter",
		Sample: `model: sdm
device: /dev/ttyUSB0 # serial port
id: 2
energy: Sum # this assignment is only required for charge meter usage`,
	}

	registry.Add(template)
}
