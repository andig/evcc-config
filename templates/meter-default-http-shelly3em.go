package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Shelly 3EM (HTTP) / General Purpose Meter",
		Sample: `power: # power reading
- type: http # use http
  uri: http://192.0.2.2/status
  jq: .emeters | map(.power) | add # sum power values of all three (L1, L2, L3) meters
currents: # Optional: List of currents in amperes
- type: http # use http
  uri: http://192.0.2.2/emeter/0
  jq: .current
- type: http # use http
  uri: http://192.0.2.2/emeter/1
  jq: .current
- type: http # use http
  uri: http://192.0.2.2/emeter/2
  jq: .current`,
	}

	registry.Add(template)
}
