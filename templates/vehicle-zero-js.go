package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "default",
		Name:   "Zero S (No web api - only bluetooth)",
		Sample: `type: default
title: Zero S # display name for UI
capacity: 10.8 # ZF 7.2 + ZF 3.6kWh
charge:
  type: js
  vm: shared
  script: |
    if (state.chargepower > 0) state.battery++; else state.battery--;
    if (state.battery < 15) state.battery = 15;
    if (state.battery > 100) state.battery = 100;
    state.battery;
cache: 1s`,
	}

	registry.Add(template)
}
