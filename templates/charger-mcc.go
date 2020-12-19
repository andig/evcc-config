package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "mcc",
		Name:   "Mobile Charger Connect",
		Sample: `uri: https://192.0.2.2 # Mobile Charger Connect address
password: # home user password`,
	}

	registry.Add(template)
}
