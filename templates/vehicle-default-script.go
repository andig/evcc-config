package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "default",
		Name:   "Generic (Script)",
		Sample: `title: Auto # display name for UI
capacity: 50 # kWh
charge:
  type: script # use script plugin
  cmd: /bin/sh -c "echo 50" # actual command
  timeout: 3s # kill script after 3 seconds
cache: 5m # cache duration`,
	}

	registry.Add(template)
}
