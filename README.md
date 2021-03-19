# Configuration examples for EVCC

[![Build Status](https://github.com/andig/evcc-config/workflows/Build/badge.svg)](https://github.com/andig/evcc-config/actions?query=workflow%3ABuild)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

[EVCC](https://github.com/andig/evcc) supports a growing list of [chargers](#chargers), [meters](#meters) and [vehicles](#vehicles). See below for detailed configuration.
Additional devices can be configured using the `generic` device type and related [plugins](#https://github.com/andig/evcc#plugins).

## Contributing

If you want to contribute configurations to this repository please open a Pull Request ("PR"). PRs should contain:

- updated or new `yaml` configuration (see https://github.com/andig/evcc-config/tree/master/yaml)
- run `make` to generate compilable Go code (configurations are used for testing) and update the README

## Chargers

- [EVSE DIN (Modbus RTU)](#charger-evse-din-modbus-rtu)
- [EVSE DIN (Modbus/TCP)](#charger-evse-din-modbus-tcp)
- [EVSE-Wifi](#charger-evse-wifi)
- [FritzDECT](#charger-fritzdect)
- [Generic](#charger-generic)
- [Generic (MQTT)](#charger-generic-mqtt)
- [go-eCharger](#charger-go-echarger)
- [go-eCharger (Cloud)](#charger-go-echarger-cloud)
- [i-CHARGE CION (Modbus RTU-over-TCP)](#charger-i-charge-cion-modbus-rtu-over-tcp)
- [KEBA Connect](#charger-keba-connect)
- [Mobile Charger Connect (Audi, Bentley, Porsche)](#charger-mobile-charger-connect-audi-bentley-porsche)
- [NRGKick BT (Bluetooth)](#charger-nrgkick-bt-bluetooth)
- [NRGKick Connect](#charger-nrgkick-connect)
- [openWB (MQTT)](#charger-openwb-mqtt)
- [Phoenix EM-CP-PP-ETH Controller (Modbus/TCP)](#charger-phoenix-em-cp-pp-eth-controller-modbus-tcp)
- [Phoenix EV-ETH Controller (Modbus/TCP)](#charger-phoenix-ev-eth-controller-modbus-tcp)
- [Phoenix EV-SER Controller (Modbus RTU)](#charger-phoenix-ev-ser-controller-modbus-rtu)
- [Wallbe (Eco, Pro)](#charger-wallbe-eco-pro)
- [Wallbe (pre 2019 EV-CC-AC1 controller)](#charger-wallbe-pre-2019-ev-cc-ac1-controller)

## Meters

- [Discovergy](#meter-discovergy)
- [E3DC (Battery Meter)](#meter-e3dc-battery-meter)
- [E3DC (Grid Meter)](#meter-e3dc-grid-meter)
- [E3DC (PV Meter)](#meter-e3dc-pv-meter)
- [Eastron SDM Modbus Meter (RTU)](#meter-eastron-sdm-modbus-meter-rtu)
- [Eastron SDM Modbus Meter (RTU-over-TCP)](#meter-eastron-sdm-modbus-meter-rtu-over-tcp)
- [Fronius Solar API V1 (Battery Meter/ HTTP)](#meter-fronius-solar-api-v1-battery-meter-http)
- [Fronius Solar API V1 (Grid Meter/ HTTP)](#meter-fronius-solar-api-v1-grid-meter-http)
- [Fronius Solar API V1 (PV Meter/ HTTP)](#meter-fronius-solar-api-v1-pv-meter-http)
- [Fronius Symo GEN24 Plus (Battery Meter)](#meter-fronius-symo-gen24-plus-battery-meter)
- [Fronius Symo GEN24 Plus (Grid Meter)](#meter-fronius-symo-gen24-plus-grid-meter)
- [Fronius Symo GEN24 Plus (PV Meter)](#meter-fronius-symo-gen24-plus-pv-meter)
- [Generic (MQTT)](#meter-generic-mqtt)
- [Generic (Script)](#meter-generic-script)
- [Generic SunSpec 3-phase meter via inverter (Grid Meter)](#meter-generic-sunspec-3-phase-meter-via-inverter-grid-meter)
- [Generic SunSpec battery inverter (Battery Meter)](#meter-generic-sunspec-battery-inverter-battery-meter)
- [Generic SunSpec PV inverter (PV Meter)](#meter-generic-sunspec-pv-inverter-pv-meter)
- [Kostal Energy Meter via inverter (Grid Meter)](#meter-kostal-energy-meter-via-inverter-grid-meter)
- [Kostal Inverter (PV Meter)](#meter-kostal-inverter-pv-meter)
- [Kostal Smart Energy Meter (Grid Meter)](#meter-kostal-smart-energy-meter-grid-meter)
- [Multiple DC MPP strings combined (PV Meter)](#meter-multiple-dc-mpp-strings-combined-pv-meter)
- [Multiple PV inverters combined (PV Meter)](#meter-multiple-pv-inverters-combined-pv-meter)
- [SMA Sunny Home Manager / Energy Meter (Speedwire)](#meter-sma-sunny-home-manager--energy-meter-speedwire)
- [SMA Sunny Island / Sunny Boy Storage (Battery Meter)](#meter-sma-sunny-island--sunny-boy-storage-battery-meter)
- [SMA SunnyBoy / TriPower / other PV-inverter (PV Meter)](#meter-sma-sunnyboy--tripower--other-pv-inverter-pv-meter)
- [SolarEdge (Grid Meter)](#meter-solaredge-grid-meter)
- [SolarEdge StorEdge (Battery Meter)](#meter-solaredge-storedge-battery-meter)
- [Solarlog (Grid Meter)](#meter-solarlog-grid-meter)
- [Solarlog (PV Meter)](#meter-solarlog-pv-meter)
- [Sonnenbatterie Eco/10 (Battery Meter/ HTTP)](#meter-sonnenbatterie-eco-10-battery-meter-http)
- [Sonnenbatterie Eco/10 (Grid Meter/ HTTP)](#meter-sonnenbatterie-eco-10-grid-meter-http)
- [Sonnenbatterie Eco/10 (PV Meter/ HTTP)](#meter-sonnenbatterie-eco-10-pv-meter-http)
- [Tesla Powerwall (Battery Meter)](#meter-tesla-powerwall-battery-meter)
- [Tesla Powerwall (Grid Meter)](#meter-tesla-powerwall-grid-meter)
- [Tesla Powerwall (PV Meter)](#meter-tesla-powerwall-pv-meter)
- [vzlogger (HTTP)](#meter-vzlogger-http)
- [vzlogger (Push Server/ Websocket)](#meter-vzlogger-push-server-websocket)
- [vzlogger (split import/export channels)](#meter-vzlogger-split-import-export-channels)

## Vehicles

- [Audi (eTron etc)](#vehicle-audi-etron-etc)
- [BMW (i3)](#vehicle-bmw-i3)
- [Ford (Kuga, Mustang, etc)](#vehicle-ford-kuga-mustang-etc)
- [Generic](#vehicle-generic)
- [Generic (Script)](#vehicle-generic-script)
- [Generic EV without SoC (Javascript)](#vehicle-generic-ev-without-soc-javascript)
- [Hyundai (Kona, Ioniq)](#vehicle-hyundai-kona-ioniq)
- [Kia (e-Niro, e-Soul, etc)](#vehicle-kia-e-niro-e-soul-etc)
- [Nissan (Leaf)](#vehicle-nissan-leaf)
- [Porsche](#vehicle-porsche)
- [Renault (Zoe)](#vehicle-renault-zoe)
- [Tesla](#vehicle-tesla)
- [VW (e-Up, e-Golf, etc)](#vehicle-vw-e-up-e-golf-etc)
- [VW ID (ID.3, ID.4, but also e-Golf, e-Up)](#vehicle-vw-id-id-3-id-4-but-also-e-golf-e-up)

## Details

### Meters


<a id="meter-discovergy"></a>
#### Discovergy

```yaml
- type: discovergy
  user: demo@discovergy.com 
  password: demo # password 
  meter: 1ESY1161229886
```

<a id="meter-e3dc-battery-meter"></a>
#### E3DC (Battery Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40069
      type: holding
      decode: int32s
    scale: -1 # reverse direction
  soc:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40082
      type: holding
      decode: uint16
```

<a id="meter-e3dc-grid-meter"></a>
#### E3DC (Grid Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40073
      type: holding
      decode: int32s
```

<a id="meter-e3dc-pv-meter"></a>
#### E3DC (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40067 # (40068 - 1) "Photovoltaikleistung in Watt"
      type: holding
      decode: int32s
```

<a id="meter-eastron-sdm-modbus-meter-rtu"></a>
#### Eastron SDM Modbus Meter (RTU)

```yaml
- type: modbus
  model: sdm
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600
  comset: 8N1
  id: 1
  energy: Sum # only required for charge meter usage
```

<a id="meter-eastron-sdm-modbus-meter-rtu-over-tcp"></a>
#### Eastron SDM Modbus Meter (RTU-over-TCP)

```yaml
- type: modbus
  model: sdm
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  id: 1
  energy: Sum # this assignment is only required for charge meter usage
```

<a id="meter-fronius-solar-api-v1-battery-meter-http"></a>
#### Fronius Solar API V1 (Battery Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_Akku == null then 0 else .Body.Data.Site.P_Akku end
  soc:
    type: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: .Body.Data.Inverters."1".SOC
```

<a id="meter-fronius-solar-api-v1-grid-meter-http"></a>
#### Fronius Solar API V1 (Grid Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_Grid == null then 0 else .Body.Data.Site.P_Grid end
```

<a id="meter-fronius-solar-api-v1-pv-meter-http"></a>
#### Fronius Solar API V1 (PV Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_PV == null then 0 else .Body.Data.Site.P_PV end
```

<a id="meter-fronius-symo-gen24-plus-battery-meter"></a>
#### Fronius Symo GEN24 Plus (Battery Meter)

```yaml
- type: default
  power:
    type: calc
    add:
    - type: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:3:DCW # mppt 3 charge
      scale: -1
    - type: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:4:DCW # mppt 4 discharge
  soc:
    type: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: ChargeState
```

<a id="meter-fronius-symo-gen24-plus-grid-meter"></a>
#### Fronius Symo GEN24 Plus (Grid Meter)

```yaml
- type: modbus
  model: 213 # sunspec meter
  uri: 192.0.2.2:502
  id: 200
```

<a id="meter-fronius-symo-gen24-plus-pv-meter"></a>
#### Fronius Symo GEN24 Plus (PV Meter)

```yaml
- type: default
  power:
    type: calc
    add:
    - type: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:1:DCW # mpp 1 pv
    - type: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:2:DCW # mpp 2 pv
```

<a id="meter-generic-mqtt"></a>
#### Generic (MQTT)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt plugin
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-generic-script"></a>
#### Generic (Script)

```yaml
- type: default
  power:
    type: script # use script plugin
    cmd: /bin/sh -c "echo 0" # actual command
    timeout: 3s # kill script after 3 seconds
```

<a id="meter-generic-sunspec-3-phase-meter-via-inverter-grid-meter"></a>
#### Generic SunSpec 3-phase meter via inverter (Grid Meter)

```yaml
- type: modbus
  model: 203 # sunspec meter
  uri: 192.0.2.2:502
  id: 1
```

<a id="meter-generic-sunspec-battery-inverter-battery-meter"></a>
#### Generic SunSpec battery inverter (Battery Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 1
  soc: ChargeState
```

<a id="meter-generic-sunspec-pv-inverter-pv-meter"></a>
#### Generic SunSpec PV inverter (PV Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 1
```

<a id="meter-kostal-energy-meter-via-inverter-grid-meter"></a>
#### Kostal Energy Meter via inverter (Grid Meter)

```yaml
- type: default
  power:
    type: modbus # use ModBus plugin
    uri: 192.0.2.2:1502 # inverter port
    id: 71
    register: # manual non-sunspec register configuration
      address: 252 # (see ba_kostal_interface_modbus-tcp_sunspec.pdf)
      type: holding
      decode: float32s # swapped float encoding
```

<a id="meter-kostal-inverter-pv-meter"></a>
#### Kostal Inverter (PV Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:1502
  id: 71
```

<a id="meter-kostal-smart-energy-meter-grid-meter"></a>
#### Kostal Smart Energy Meter (Grid Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 71
```

<a id="meter-multiple-dc-mpp-strings-combined-pv-meter"></a>
#### Multiple DC MPP strings combined (PV Meter)

```yaml
- type: default
  power:
    type: calc
    add:
    - type: modbus
      model: sunspec
      value: 160:1:DCW # SunSpec Model 160 MPP string 1 DCW
      uri: 192.0.2.2:502
      id: 1
    - type: modbus
      model: sunspec
      value: 160:2:DCW # SunSpec Model 160 MPP string 2 DCW
      uri: 192.0.2.2:502
      id: 1
```

<a id="meter-multiple-pv-inverters-combined-pv-meter"></a>
#### Multiple PV inverters combined (PV Meter)

```yaml
- type: default
  power:
    type: calc
    add:
    - type: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
    - type: modbus
      model: sunspec
      uri: 192.0.2.3:502
      id: 1
```

<a id="meter-sma-sunny-home-manager--energy-meter-speedwire"></a>
#### SMA Sunny Home Manager / Energy Meter (Speedwire)

```yaml
- type: sma
  uri: 192.0.2.2
```

<a id="meter-sma-sunny-island--sunny-boy-storage-battery-meter"></a>
#### SMA Sunny Island / Sunny Boy Storage (Battery Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 126
  soc: ChargeState
```

<a id="meter-sma-sunnyboy--tripower--other-pv-inverter-pv-meter"></a>
#### SMA SunnyBoy / TriPower / other PV-inverter (PV Meter)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 126
```

<a id="meter-solaredge-grid-meter"></a>
#### SolarEdge (Grid Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.0.2.2:502
    id: 1
    register:
      address: 40207
      type: holding
      decode: int16
    scale: -1
```

<a id="meter-solaredge-storedge-battery-meter"></a>
#### SolarEdge StorEdge (Battery Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.0.2.2:502
    id: 1
    register:
      address: 62836
      type: holding
      decode: float32s
    scale: -1
  soc:
    type: modbus
    uri: 192.0.2.2:502
    id: 1
    register:
      address: 62852
      type: holding
      decode: float32s
```

<a id="meter-solarlog-grid-meter"></a>
#### Solarlog (Grid Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.0.2.2:502 # IP address of the SolarLog device and ModBus port address
    id: 1
    register:
      address: 3518
      type: input
      decode: uint32s
```

<a id="meter-solarlog-pv-meter"></a>
#### Solarlog (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.0.2.2:502 # IP address of the SolarLog  device and ModBus port address
    id: 1
    register:
      address: 3502
      type: input
      decode: uint32s
```

<a id="meter-sonnenbatterie-eco-10-battery-meter-http"></a>
#### Sonnenbatterie Eco/10 (Battery Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Pac_total_W
  soc:
    type: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .USOC
```

<a id="meter-sonnenbatterie-eco-10-grid-meter-http"></a>
#### Sonnenbatterie Eco/10 (Grid Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .GridFeedIn_W
    scale: -1 # reverse direction
```

<a id="meter-sonnenbatterie-eco-10-pv-meter-http"></a>
#### Sonnenbatterie Eco/10 (PV Meter/ HTTP)

```yaml
- type: default
  power:
    type: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Production_W
```

<a id="meter-tesla-powerwall-battery-meter"></a>
#### Tesla Powerwall (Battery Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: battery
```

<a id="meter-tesla-powerwall-grid-meter"></a>
#### Tesla Powerwall (Grid Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: grid
```

<a id="meter-tesla-powerwall-pv-meter"></a>
#### Tesla Powerwall (PV Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: pv
```

<a id="meter-vzlogger-http"></a>
#### vzlogger (HTTP)

```yaml
- type: default
  power: # power reading
    type: http # use http plugin
    uri: http://demo.volkszaehler.org/api/data/<uuid>.json?from=now
    jq: .data.tuples[0][1] # parse response json
```

<a id="meter-vzlogger-push-server-websocket"></a>
#### vzlogger (Push Server/ Websocket)

```yaml
- type: default
  power:
    type: ws # use websocket plugin
    uri: ws://192.0.2.2:8082/socket
    jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse response json
    timeout: 30s
    scale: 1
```

<a id="meter-vzlogger-split-import-export-channels"></a>
#### vzlogger (split import/export channels)

```yaml
- type: default
  power:
    type: calc # use calc plugin
    add:
    - type: http # import channel
      uri: http://demo.volkszaehler.org/api/data/<import-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
    - type: http # export channel
      uri: http://demo.volkszaehler.org/api/data/<export-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
      scale: -1 # export must result in negative values
```


### Chargers


<a id="charger-evse-din-modbus-rtu"></a>
#### EVSE DIN (Modbus RTU)

```yaml
- type: simpleevse
  device: /dev/ttyUSB0 # serial RS485 interface
```

<a id="charger-evse-din-modbus-tcp"></a>
#### EVSE DIN (Modbus/TCP)

```yaml
- type: simpleevse
  uri: 192.0.2.2:502 # Modbus/TCP converter adress
```

<a id="charger-evse-wifi"></a>
#### EVSE-Wifi

```yaml
- type: evsewifi
  uri: http://192.0.2.2
```

<a id="charger-fritzdect"></a>
#### FritzDECT

```yaml
- type: fritzdect
  uri: https://fritz.box # FRITZ!Box ip address (local)
  user: xxxxxxxxxx # FRITZ!Box username (Has to have Smart Home privileges!)
  password: yyyyyyyyyy # FRITZ!Box password
  ain: 117788992233 # switch actor identification number without blanks (see AIN number on switch sticker)
  standbypower: 10 # standbypower threshold in W (depends on embeded vehicle charger)
```

<a id="charger-generic"></a>
#### Generic

```yaml
- type: default
  status: # charger status A..F
    type: ...
    # ...
  enabled: # charger enabled state (true/false or 0/1)
    type: ...
    # ...
  enable: # set charger enabled state
    type: ...
    # ...
  maxcurrent: # set charger max current
    type: ...
    # ...
```

<a id="charger-generic-mqtt"></a>
#### Generic (MQTT)

```yaml
- type: default
  status: # charger status A..F
    type: mqtt
    topic: some/topic1
  enabled: # charger enabled state (true/false or 0/1)
    type: mqtt
    topic: some/topic2
  enable: # set charger enabled state
    type: script
    cmd: /bin/sh -c "echo ${enable}"
  maxcurrent: # set charger max current
    type: script
    cmd: /bin/sh -c "echo ${maxcurrent}"
```

<a id="charger-go-echarger"></a>
#### go-eCharger

```yaml
- type: go-e
  uri: http://192.0.2.2 # go-e ip address (local)
```

<a id="charger-go-echarger-cloud"></a>
#### go-eCharger (Cloud)

```yaml
- type: go-e
  token: 4711c # go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

<a id="charger-i-charge-cion-modbus-rtu-over-tcp"></a>
#### i-CHARGE CION (Modbus RTU-over-TCP)

```yaml
- type: default
  status:
    type: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
        address: 139 # CP-Status
        type: holding
        decode: uint16
  enabled:
    type: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1 
    register: # manual register configuration
      address: 100 # Zustand
      type: holding
      decode: uint16
  enable:
    type: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
      address: 100 # ein / aus
      type: writesingle
      decode: uint16
  maxcurrent:
    type: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
      address: 127 # Strom max
      type: writesingle
      decode: uint16
```

<a id="charger-keba-connect"></a>
#### KEBA Connect

```yaml
- type: keba
  uri: 192.0.2.2
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

<a id="charger-mobile-charger-connect-audi-bentley-porsche"></a>
#### Mobile Charger Connect (Audi, Bentley, Porsche)

```yaml
- type: mcc
  uri: https://192.0.2.2
  password: # home user password
```

<a id="charger-nrgkick-bt-bluetooth"></a>
#### NRGKick BT (Bluetooth)

```yaml
- type: nrgkick-bluetooth
  macaddress: 00:99:22 # BT device MAC address
  pin: # App PIN number (write protection)
```

<a id="charger-nrgkick-connect"></a>
#### NRGKick Connect

```yaml
- type: nrgkick-connect
  uri: http://192.0.2.2
  mac: 00:99:22 # BT device MAC address
  password: # password
```

<a id="charger-openwb-mqtt"></a>
#### openWB (MQTT)

```yaml
- type: openwb
  broker: 192.0.2.2 # openWB IP
  id: 1 # loadpoint id
```

<a id="charger-phoenix-em-cp-pp-eth-controller-modbus-tcp"></a>
#### Phoenix EM-CP-PP-ETH Controller (Modbus/TCP)

```yaml
- type: phoenix-em-eth
  uri: 192.168.0.8:502
  meter:
    power: true    # charge meter connected to controller
    energy: true   # charge meter connected to controller
    currents: true # charge meter connected to controller
```

<a id="charger-phoenix-ev-eth-controller-modbus-tcp"></a>
#### Phoenix EV-ETH Controller (Modbus/TCP)

```yaml
- type: phoenix-ev-eth
  uri: 192.168.0.8:502
  meter:
    power: true    # charge meter connected to controller
    energy: true   # charge meter connected to controller
    currents: true # charge meter connected to controller
```

<a id="charger-phoenix-ev-ser-controller-modbus-rtu"></a>
#### Phoenix EV-SER Controller (Modbus RTU)

```yaml
- type: phoenix-ev-ser
  device: /dev/ttyUSB0
  baudrate: 9600 # configurable (S2/DIP 1)
  comset: 8N1
  id: 1 # configurable (S2/DIP 2–6)
```

<a id="charger-wallbe-eco-pro"></a>
#### Wallbe (Eco, Pro)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # TCP ModBus address
```

<a id="charger-wallbe-pre-2019-ev-cc-ac1-controller"></a>
#### Wallbe (pre 2019 EV-CC-AC1 controller)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # TCP ModBus address
  legacy: true # enable for older Wallbes with Phoenix EV-CC-AC1-M3-CBC-RCM controller
```


### Vehicles


<a id="vehicle-audi-etron-etc"></a>
#### Audi (eTron etc)

```yaml
- type: audi
  title: eTron # display name for UI
  capacity: 14 # kWh
  user: # user
  password: # password
  vin: WAUZZZ... # optional
```

<a id="vehicle-bmw-i3"></a>
#### BMW (i3)

```yaml
- type: bmw
  title: i3 # display name for UI
  capacity: 65 # kWh
  user: # user
  password: # password
  vin: WBMW... # optional
```

<a id="vehicle-ford-kuga-mustang-etc"></a>
#### Ford (Kuga, Mustang, etc)

```yaml
- type: ford
  title: Kuga # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WF0FXX... # optional
```

<a id="vehicle-generic"></a>
#### Generic

```yaml
- type: default
  title: Mein Auto # display name for UI
  capacity: 50 # kWh
  charge:
    type: ...
    # ...
```

<a id="vehicle-generic-script"></a>
#### Generic (Script)

```yaml
- type: default
  title: Auto # display name for UI
  capacity: 50 # kWh
  charge:
    type: script # use script plugin
    cmd: /bin/sh -c "echo 50" # actual command
    timeout: 3s # kill script after 3 seconds
  cache: 5m # cache duration
```

<a id="vehicle-generic-ev-without-soc-javascript"></a>
#### Generic EV without SoC (Javascript)

```yaml
- type: default
  title: My electric vehicle # display name for UI
  capacity: 10 # kWh
  charge:
    type: js
    script: 95 // vehicle SoC in %
```

<a id="vehicle-hyundai-kona-ioniq"></a>
#### Hyundai (Kona, Ioniq)

```yaml
- type: hyundai
  title: Kona # display name for UI
  capacity: 64 # kWh
  user: # user
  password: # password
```

<a id="vehicle-kia-e-niro-e-soul-etc"></a>
#### Kia (e-Niro, e-Soul, etc)

```yaml
- type: kia
  title: e-Niro # display name for UI
  capacity: 64 # kWh
  user: # user
  password: # password
```

<a id="vehicle-nissan-leaf"></a>
#### Nissan (Leaf)

```yaml
- type: nissan
  title: Leaf # display name for UI
  capacity: 60 # kWh
  user: # user
  password: # password
```

<a id="vehicle-porsche"></a>
#### Porsche

```yaml
- type: porsche
  title: Taycan # display name for UI
  capacity: 83 # kWh
  user: # user
  password: # password
  vin: WP...
```

<a id="vehicle-renault-zoe"></a>
#### Renault (Zoe)

```yaml
- type: renault
  title: Zoe # display name for UI
  capacity: 60 # kWh
  user: # user
  password: # password
  vin: WREN... # optional
```

<a id="vehicle-tesla"></a>
#### Tesla

```yaml
- type: tesla
  title: Model S # display name for UI
  capacity: 90 # kWh
  user: # email
  password: # password
  vin: WTSLA...
```

<a id="vehicle-vw-e-up-e-golf-etc"></a>
#### VW (e-Up, e-Golf, etc)

```yaml
- type: vw
  title: Golf # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WVWZZZ... # optional
```

<a id="vehicle-vw-id-id-3-id-4-but-also-e-golf-e-up"></a>
#### VW ID (ID.3, ID.4, but also e-Golf, e-Up)

```yaml
- type: id
  title: ID.3 # display name for UI
  capacity: 50 # kWh
  user: # user
  password: # password
  vin: WVWZZZ... # optional
```

