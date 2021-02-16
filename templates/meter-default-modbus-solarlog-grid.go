package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Solarlog (Grid Meter)",
		Sample: `power:
  type: modbus
  uri: 192.0.2.2:502 # IP address of the SolarLog device and ModBus port address
  id: 1
  register:
    address: 3518
    type: input
    decode: uint32s`,
	}

	registry.Add(template)
}
