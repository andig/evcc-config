package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "vzlogger (Push Server/ Websocket)",
		Sample: `power:
  type: ws # use websocket plugin
  uri: ws://192.0.2.2:8082/socket
  jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse response json
  timeout: 30s
  scale: 1`,
	}

	registry.Add(template)
}
