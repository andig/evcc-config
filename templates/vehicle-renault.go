package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "renault",
		Name:   "Renault (Zoe)",
		Sample: `title: Zoe # display name for UI
capacity: 60 # kWh
user: # user
password: # password
vin: WREN... # optional`,
	}

	registry.Add(template)
}
