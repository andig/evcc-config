package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus Meter (RTU)",
		Sample: `model: sdm
device: /dev/ttyUSB0 # serial port
baudrate: 9600
comset: 8N1
id: 1
energy: Sum # only required for charge meter usage`,
	}

	registry.Add(template)
}
