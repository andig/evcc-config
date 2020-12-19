package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "tesla",
		Name:   "Tesla Powerwall (PV meter)",
		Sample: `uri: http://192.0.2.2/api/meters/aggregates
usage: solar # grid meter: `+"`"+`site`+"`"+`, pv: `+"`"+`solar`+"`"+`, battery: `+"`"+`battery`+"`"+``,
	}

	registry.Add(template)
}
