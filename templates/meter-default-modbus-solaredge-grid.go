package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "SolarEdge (Grid Meter)",
		Sample: `power:
  type: modbus
  uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
  id: 1
  register:
    address: 40206 # Meter 1 Total Real Power (sum of active phases)
    type: holding
    decode: int16
  scale: -1`,
	}

	registry.Add(template)
}
