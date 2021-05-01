package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "E3DC (PV Meter)",
		Sample: `power:
  type: modbus
  uri: e3dc.fritz.box:502
  id: 1 # ModBus slave id
  register: # manual register configuration
    address: 40067 # (40068 - 1) "Photovoltaikleistung in Watt"
    type: holding
    decode: int32s`,
	}

	registry.Add(template)
}
