package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Symo GEN24 Plus (Battery Meter)",
		Sample: `power:
  type: calc
  add:
  - type: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: 160:3:DCW # mppt 3 charge
    scale: -1
  - type: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: 160:4:DCW # mppt 4 discharge
soc:
  type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 1
  value: ChargeState`,
	}

	registry.Add(template)
}
