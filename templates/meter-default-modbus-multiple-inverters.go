package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Multiple PV inverters combined (PV Meter)",
		Sample: `power:
  type: calc
  add:
  - type: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
  - type: modbus
    model: sunspec
    uri: 192.0.2.3:502
    id: 1`,
	}

	registry.Add(template)
}
