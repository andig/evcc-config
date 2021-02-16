package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "sma",
		Name:   "SMA Sunny Home Manager 2.0 / SMA Energy Meter 30",
		Sample: `serial: 1234567890 # Serial number of the device`,
	}

	registry.Add(template)
}
