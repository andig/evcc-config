package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "vzlogger (HTTP)",
		Sample: `power: # power reading
  type: http # use http plugin
  uri: http://demo.volkszaehler.org/api/data/<uuid>.json?from=now
  jq: .data.tuples[0][1] # parse response json`,
	}

	registry.Add(template)
}
