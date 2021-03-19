package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "fritzdect",
		Name:   "FritzDECT",
		Sample: `uri: https://fritz.box # FRITZ!Box ip address (local)
user: xxxxxxxxxx # FRITZ!Box username (Has to have Smart Home privileges!)
password: yyyyyyyyyy # FRITZ!Box password
ain: 117788992233 # switch actor identification number without blanks (see AIN number on switch sticker)
standbypower: 10 # standbypower threshold in W (depends on embeded vehicle charger)`,
	}

	registry.Add(template)
}
