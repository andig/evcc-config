package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "keba",
		Name:   "KEBA Connect",
		Sample: `uri: 192.0.2.2 # KEBA ip address
rfid:
  tag: 765765348 # RFID tag, see `+"`"+`evcc charger`+"`"+` to show tag`,
	}

	registry.Add(template)
}
