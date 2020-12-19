package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "go-e",
		Name:   "go-eCharger (Cloud)",
		Sample: `token: 4711c # go-e cloud API token
cache: 10s # go-e cloud API cache duration`,
	}

	registry.Add(template)
}
