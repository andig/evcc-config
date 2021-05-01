package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "discovergy",
		Name:   "Discovergy",
		Sample: `user: demo@discovergy.com 
password: demo # password 
meter: 1ESY1161229886`,
	}

	registry.Add(template)
}
