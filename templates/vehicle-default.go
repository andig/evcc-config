package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "default",
		Name:   "Generic",
		Sample: `title: Mein Auto # display name for UI
capacity: 50 # kWh
charge:
  type: ...
  # ...`,
	}

	registry.Add(template)
}
