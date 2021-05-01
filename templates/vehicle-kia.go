package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "kia",
		Name:   "Kia (e-Niro, e-Soul, etc)",
		Sample: `title: e-Niro # display name for UI
capacity: 64 # kWh
user: # user
password: # password`,
	}

	registry.Add(template)
}
