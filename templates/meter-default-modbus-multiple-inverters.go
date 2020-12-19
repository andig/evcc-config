package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Multiple Grid Inverters combined (PV Meter)",
		Sample: `power:
  type: calc # use the calc plugin
  add: # The add function sums up both string values
  - type: modbus
    model: sunspec
    value: 160:1:DCW # string 1
    uri: 192.0.2.2:1502 
    id: 71 # Configured Modbus Device ID 
  - type: modbus  
    value: 160:2:DCW # string 2
    uri: 192.0.2.2:1502 
    id: 71 # Configured Modbus Device ID `,
	}

	registry.Add(template)
}
