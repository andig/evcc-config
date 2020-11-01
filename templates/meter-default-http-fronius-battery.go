package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Symo Hybrid (Battery Meter/ HTTP)",
		Sample: `power:
  type: http
  uri: http://192.168.11.71/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  method: GET # default HTTP method
  headers:
  - content-type: application/json
  insecure: true
  jq: .Body.Data.Site.P_Akku
soc:
  type: http
  uri: http://192.168.11.71/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  method: GET # default HTTP method
  headers:
  - content-type: application/json
  insecure: true
  jq: .Body.Data.Inverters."1".SOC`,
	}

	registry.Add(template)
}
