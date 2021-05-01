package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "SolarEdge StorEdge (Battery Meter)",
		Sample: `power:
  type: modbus
  uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
  id: 1
  register:
    address: 62836 # Battery 1 Instantaneous Power
    type: holding
    decode: float32s
  scale: -1
soc:
  type: modbus
  uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
  id: 1
  register:
    address: 62852 # Battery 1 State of Energy (SOE)
    type: holding
    decode: float32s`,
	}

	registry.Add(template)
}
