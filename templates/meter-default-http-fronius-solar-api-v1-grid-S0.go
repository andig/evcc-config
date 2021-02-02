package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Solar API V1 (Grid S0 meter/ HTTP)",
		Sample: `power: # Grid power reading Fronius Solar API V1 GetPowerFlowRealtimeData.P_Grid
  type: http # use http plugin for grid power (P_Grid)
  uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  jq: if .Body.Data.Site.P_Grid == null then 0 else .Body.Data.Site.P_Grid end # parse GetPowerFlowRealtimeData P_Grid response`,
	}

	registry.Add(template)
}
