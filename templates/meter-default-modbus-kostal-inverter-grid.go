package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Kostal Energy Meter via inverter (Grid Meter)",
		Sample: `power:
  type: modbus # use ModBus plugin
  uri: 192.0.2.2:1502 # inverter port
  id: 71
  register: # manual non-sunspec register configuration
    address: 252 # (see ba_kostal_interface_modbus-tcp_sunspec.pdf)
    type: holding
    decode: float32s # swapped float encoding`,
	}

	registry.Add(template)
}
