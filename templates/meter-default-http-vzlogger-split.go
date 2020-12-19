package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "vzlogger (split import/export channels)",
		Sample: `power:
  type: calc # use calc plugin
  add:
  - type: http # import channel
    uri: http://demo.volkszaehler.org/api/data/<import-uuid>.json?from=now
    jq: .data.tuples[0][1] # parse response json
  - type: http # export channel
    uri: http://demo.volkszaehler.org/api/data/<export-uuid>.json?from=now
    jq: .data.tuples[0][1] # parse response json
    scale: -1 # export must result in negative values`,
	}

	registry.Add(template)
}
