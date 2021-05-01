package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "sma",
		Name:   "SMA Sunny Home Manager / Energy Meter (Speedwire)",
		Sample: `uri: 192.0.2.2`,
	}

	registry.Add(template)
}
