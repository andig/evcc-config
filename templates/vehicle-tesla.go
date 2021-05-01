package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "tesla",
		Name:   "Tesla",
		Sample: `title: Model S # display name for UI
capacity: 90 # kWh
user: # email
password: # password
vin: WTSLA...`,
	}

	registry.Add(template)
}
