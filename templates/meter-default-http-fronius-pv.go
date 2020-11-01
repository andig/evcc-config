package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Symo, Symo Hybrid (PV Meter/ HTTP)",
		Sample: `power:
  type: calc
  add:
  - type: http # first symo wr hybrid model
    uri: http://192.168.0.5/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    method: GET # default HTTP method
    headers:
    - content-type: application/json
    insecure: true
    jq: .Body.Data.Site.P_PV
  - type: http # second symo wr
    uri: http://192.168.0.6/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    method: GET # default HTTP method
    headers:
    - content-type: application/json
    insecure: true
    jq: .Body.Data.Site.P_PV`,
	}

	registry.Add(template)
}
