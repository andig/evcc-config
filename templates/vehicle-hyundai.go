package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "hyundai",
		Name:   "Hyundai (Kona, Ioniq)",
		Sample: `title: Kona # display name for UI
capacity: 64 # kWh
user: # user
password: # password`,
	}

	registry.Add(template)
}
