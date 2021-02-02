package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Solar API V1 (PV meter/ HTTP)",
		Sample: `power: # pv power reading Fronius Solar API V1 GetPowerFlowRealtimeData.P_PV
        type: http # use http plugin for pv power (P_PV)
        uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
        jq: if .Body.Data.Site.P_PV == null then 0 else .Body.Data.Site.P_PV end # parse GetPowerFlowRealtimeData P_PV response`,
	}

	registry.Add(template)
}
