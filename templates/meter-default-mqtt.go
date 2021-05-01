package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Generic (MQTT)",
		Sample: `power: # power reading
  type: mqtt # use mqtt plugin
  topic: mbmd/sdm1-1/Power # mqtt topic
  timeout: 10s # don't use older values
currents: # optional: List of currents in Amperes
- type: mqtt # use mqtt plugin
  topic: mbmd/sdm1-1/CurrentL1 # Current of L1
- type: mqtt # use mqtt plugin
  topic: mbmd/sdm1-1/CurrentL2 # Current of L2
- type: mqtt # use mqtt plugin
  topic: mbmd/sdm1-1/CurrentL3 # Current of L3`,
	}

	registry.Add(template)
}
