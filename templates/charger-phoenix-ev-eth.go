package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "phoenix-ev-eth",
		Name:   "Phoenix EV-ETH Controller (Modbus/TCP)",
		Sample: `uri: 192.168.0.8:502
meter:
  power: true    # charge meter connected to controller
  energy: true   # charge meter connected to controller
  currents: true # charge meter connected to controller`,
	}

	registry.Add(template)
}
