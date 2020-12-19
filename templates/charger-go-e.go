package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "go-e",
		Name:   "go-eCharger (Lokal)",
		Sample: `uri: http://192.0.2.2 # go-e ip address (local)`,
	}

	registry.Add(template)
}
