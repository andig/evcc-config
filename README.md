# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters

- [Discovergy (Grid or PV meter/ HTTP)](#meter-discovergy-grid-or-pv-meter-http)
- [E3DC (Battery)](#meter-e3dc-battery)
- [E3DC (Grid Meter)](#meter-e3dc-grid-meter)
- [E3DC (PV Meter)](#meter-e3dc-pv-meter)
- [Generisch (MQTT)](#meter-generisch-mqtt)
- [Generisch (Script)](#meter-generisch-script)
- [Kostal Inverter (Grid Meter)](#meter-kostal-inverter-grid-meter)
- [Kostal Inverter (PV Meter)](#meter-kostal-inverter-pv-meter)
- [Kostal Smart Energy Meter (Grid Meter)](#meter-kostal-smart-energy-meter-grid-meter)
- [Modbus (Ethernet)](#meter-modbus-ethernet)
- [Modbus (RTU)](#meter-modbus-rtu)
- [Multiple Grid Inverters combined (PV Meter)](#meter-multiple-grid-inverters-combined-pv-meter)
- [SMA Sunny Boy Storage (Battery)](#meter-sma-sunny-boy-storage-battery)
- [SMA Sunny Home Manager 2.0 / SMA Energy Meter 30](#meter-sma-sunny-home-manager-2-0--sma-energy-meter-30)
- [SMA Sunny Island (Battery)](#meter-sma-sunny-island-battery)
- [SMA SunnyBoy / TriPower / other SunSpec PV-inverters (PV Meter)](#meter-sma-sunnyboy--tripower--other-sunspec-pv-inverters-pv-meter)
- [Solarlog (Grid Meter)](#meter-solarlog-grid-meter)
- [Solarlog (PV Meter)](#meter-solarlog-pv-meter)
- [Sonnenbatterie Eco/10 (Battery/ HTTP)](#meter-sonnenbatterie-eco-10-battery-http)
- [Sonnenbatterie Eco/10 (Grid meter/ HTTP)](#meter-sonnenbatterie-eco-10-grid-meter-http)
- [Sonnenbatterie Eco/10 (PV meter/ HTTP)](#meter-sonnenbatterie-eco-10-pv-meter-http)
- [Tesla Powerwall (Battery)](#meter-tesla-powerwall-battery)
- [Tesla Powerwall (Grid meter)](#meter-tesla-powerwall-grid-meter)
- [Tesla Powerwall (PV meter)](#meter-tesla-powerwall-pv-meter)
- [vzlogger (HTTP)](#meter-vzlogger-http)
- [vzlogger (Push Server/ Websocket)](#meter-vzlogger-push-server-websocket)
- [vzlogger (split import/export channels)](#meter-vzlogger-split-import-export-channels)

## Chargers

- [EVSE Wifi](#charger-evse-wifi)
- [Generisch](#charger-generisch)
- [Generisch (MQTT)](#charger-generisch-mqtt)
- [go-eCharger (Cloud)](#charger-go-echarger-cloud)
- [go-eCharger (Lokal)](#charger-go-echarger-lokal)
- [KEBA Connect](#charger-keba-connect)
- [Mobile Charger Connect](#charger-mobile-charger-connect)
- [NRGKick Bluetooth](#charger-nrgkick-bluetooth)
- [NRGKick Connect](#charger-nrgkick-connect)
- [openWB (MQTT)](#charger-openwb-mqtt)
- [OpenWB (remote-controlled using MQTT)](#charger-openwb-remote-controlled-using-mqtt)
- [Phoenix EM-CP Controller (Ethernet/Modbus TCP)](#charger-phoenix-em-cp-controller-ethernet-modbus-tcp)
- [Phoenix EV-CC Controller (Modbus)](#charger-phoenix-ev-cc-controller-modbus)
- [Simple EVSE (Ethernet/Modbus TCP)](#charger-simple-evse-ethernet-modbus-tcp)
- [Simple EVSE (USB)](#charger-simple-evse-usb)
- [Wallbe (Eco, Pro)](#charger-wallbe-eco-pro)
- [Wallbe (pre 2019 EV-CC-AC1 controller)](#charger-wallbe-pre-2019-ev-cc-ac1-controller)

## Vehicles

- [Audi (eTron etc)](#vehicle-audi-etron-etc)
- [BMW (i3)](#vehicle-bmw-i3)
- [Ford (Kuga, Mustang, etc)](#vehicle-ford-kuga-mustang-etc)
- [Generisch](#vehicle-generisch)
- [Generisch (Script)](#vehicle-generisch-script)
- [Hyundai (Kona, Ioniq)](#vehicle-hyundai-kona-ioniq)
- [Kia (e-Niro, e-Soul, etc)](#vehicle-kia-e-niro-e-soul-etc)
- [Nissan (Leaf)](#vehicle-nissan-leaf)
- [Porsche](#vehicle-porsche)
- [Renault (Zoe)](#vehicle-renault-zoe)
- [Tesla](#vehicle-tesla)
- [VW (eUp, ID.3, ID.4, etc)](#vehicle-vw-eup-id-3-id-4-etc)

## Details

### Meters


<a id="meter-discovergy-grid-or-pv-meter-http"></a>
#### Discovergy (Grid or PV meter/ HTTP)

```yaml
- type: default
  power: # power reading
    type: http # use http plugin
    auth:
      type: basic
      user: demo@discovergy.com # Discovergy user name
      password: demo # password 
    uri: https://api.discovergy.com/public/v1/last_reading?meterId=659a3da00324400da66cef81e1cbe3c5 # append meter id
    jq: .values.power
    scale: 0.001
```

<a id="meter-e3dc-battery"></a>
#### E3DC (Battery)

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

<a id="meter-generisch-mqtt"></a>
#### Generisch (MQTT)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt plugin
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-generisch-script"></a>
#### Generisch (Script)

```yaml
- type: default
  power:
    type: script # use script plugin
    cmd: /bin/sh -c "echo 0" # actual command
    timeout: 3s # kill script after 3 seconds
```

<a id="meter-kostal-inverter-grid-meter"></a>
#### Kostal Inverter (Grid Meter)

```yaml
- type: default
  power:
    type: modbus # use ModBus plugin
    model: kostal
    uri: 192.0.2.2:1502 
    id: 71 # Configured Modbus Device ID 
    register: # manual register configuration
      address: 252 # (see https://www.kostal-solar-electric.com/de-de/download/-/media/document-library-folder---kse/2018/08/30/08/53/ba_kostal_interface_modbus-tcp_sunspec.pdf)
      type: holding
      decode: float32s #swapped float encoding
```

<a id="meter-kostal-inverter-pv-meter"></a>
#### Kostal Inverter (PV Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.0.2.2:1502
  id: 71
  power: Power
```

<a id="meter-kostal-smart-energy-meter-grid-meter"></a>
#### Kostal Smart Energy Meter (Grid Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.0.2.2:502
  id: 71
  power: Power
  energy: Energy
```

<a id="meter-modbus-ethernet"></a>
#### Modbus (Ethernet)

```yaml
- type: modbus
  model: sdm
  uri: 192.0.2.2:502
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

<a id="meter-modbus-rtu"></a>
#### Modbus (RTU)

```yaml
- type: modbus
  model: sdm
  uri: 192.0.2.2:502
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default value, optionally override
  energy: Sum # energy value (Zählerstand)
```

<a id="meter-multiple-grid-inverters-combined-pv-meter"></a>
#### Multiple Grid Inverters combined (PV Meter)

```yaml
- type: default
  power:
    type: calc # use the calc plugin
    add: # The add function sums up both string values
    - type: modbus
      model: sunspec
      value: 160:1:DCW # string 1
      uri: 192.0.2.2:1502 
      id: 71 # Configured Modbus Device ID 
    - type: modbus  
      value: 160:2:DCW # string 2
      uri: 192.0.2.2:1502 
      id: 71 # Configured Modbus Device ID 
```

<a id="meter-sma-sunny-boy-storage-battery"></a>
#### SMA Sunny Boy Storage (Battery)

```yaml
- type: modbus
  uri: 192.0.2.2:502
  id: 126 # ModBus slave id
  model: sma-sunspec
  power: Power # default value, optionally override
  soc: ChargeState # battery soc (Ladezustand)
```

<a id="meter-sma-sunny-home-manager-2-0--sma-energy-meter-30"></a>
#### SMA Sunny Home Manager 2.0 / SMA Energy Meter 30

```yaml
- type: sma
  serial: 1234567890 # Serial number of the device
```

<a id="meter-sma-sunny-island-battery"></a>
#### SMA Sunny Island (Battery)

```yaml
- type: modbus
  model: sunny-island
  uri: 192.0.2.2:502
  id: 126
  power: Power # default values, optionally override
  soc: ChargeState # battery soc (Ladezustand)
```

<a id="meter-sma-sunnyboy--tripower--other-sunspec-pv-inverters-pv-meter"></a>
#### SMA SunnyBoy / TriPower / other SunSpec PV-inverters (PV Meter)

```yaml
- type: modbus
  uri:192.0.2.2:502
  id: 126 # ModBus slave id
  model: sma-sunspec
  power: Power # default value, optionally override
  energy: Sum # energy value (Zählerstand)
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

<a id="meter-sonnenbatterie-eco-10-battery-http"></a>
#### Sonnenbatterie Eco/10 (Battery/ HTTP)

```yaml
- type: default
  power: # power reading
    type: http # use http plugin
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Pac_total_W
    scale: -1 # reverse direction
  soc:
    type: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .USOC
```

<a id="meter-sonnenbatterie-eco-10-grid-meter-http"></a>
#### Sonnenbatterie Eco/10 (Grid meter/ HTTP)

```yaml
- type: default
  power: # power reading
    type: http # use http plugin
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .GridFeedIn_W
    scale: -1 # reverse direction
```

<a id="meter-sonnenbatterie-eco-10-pv-meter-http"></a>
#### Sonnenbatterie Eco/10 (PV meter/ HTTP)

```yaml
- type: default
  power: # power reading
    type: http # use http plugin
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Production_W
```

<a id="meter-tesla-powerwall-battery"></a>
#### Tesla Powerwall (Battery)

```yaml
- type: tesla
  uri: http://192.0.2.2/api/meters/aggregates
  usage: battery # grid meter: `site`, pv: `solar`, battery: `battery`
```

<a id="meter-tesla-powerwall-grid-meter"></a>
#### Tesla Powerwall (Grid meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/api/meters/aggregates
  usage: site # grid meter: `site`, pv: `solar`, battery: `battery`
```

<a id="meter-tesla-powerwall-pv-meter"></a>
#### Tesla Powerwall (PV meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/api/meters/aggregates
  usage: solar # grid meter: `site`, pv: `solar`, battery: `battery`
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


<a id="charger-evse-wifi"></a>
#### EVSE Wifi

```yaml
- type: evsewifi
  uri: http://192.0.2.2 # SimpleEVSE-Wifi address
```

<a id="charger-generisch"></a>
#### Generisch

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

<a id="charger-generisch-mqtt"></a>
#### Generisch (MQTT)

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

<a id="charger-go-echarger-cloud"></a>
#### go-eCharger (Cloud)

```yaml
- type: go-e
  token: 4711c # go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

<a id="charger-go-echarger-lokal"></a>
#### go-eCharger (Lokal)

```yaml
- type: go-e
  uri: http://192.0.2.2 # go-e ip address (local)
```

<a id="charger-keba-connect"></a>
#### KEBA Connect

```yaml
- type: keba
  uri: 192.0.2.2 # KEBA ip address
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

<a id="charger-mobile-charger-connect"></a>
#### Mobile Charger Connect

```yaml
- type: mcc
  uri: https://192.0.2.2 # Mobile Charger Connect address
  password: # home user password
```

<a id="charger-nrgkick-bluetooth"></a>
#### NRGKick Bluetooth

```yaml
- type: nrgkick-bluetooth
  macaddress: 00:99:22 # MAC address
  pin: # pin
```

<a id="charger-nrgkick-connect"></a>
#### NRGKick Connect

```yaml
- type: nrgkick-connect
  uri: http://192.0.2.2
  mac: 00:99:22 # MAC address
  password: # password
```

<a id="charger-openwb-mqtt"></a>
#### openWB (MQTT)

```yaml
- type: openwb
  broker: 192.0.2.2 # openWB IP
  id: 1 # loadpoint id
```

<a id="charger-openwb-remote-controlled-using-mqtt"></a>
#### OpenWB (remote-controlled using MQTT)

```yaml
- type: default
  status:
    # with openWB, charging status (A..F) this is split between "plugged" and "charging"
    # the openwb plugin combines both into status (charging=C, plugged=B, otherwise=A)
    type: openwb # use openwb plugin
    plugged:
      type: mqtt
      topic: openWB/lp/1/boolPlugStat
    charging:
      type: mqtt
      topic: openWB/lp/1/boolChargeStat
  enabled:
    type: mqtt
    topic: openWB/lp/1/ChargePointEnabled
    timeout: 30s
  enable:
    type: mqtt
    topic: openWB/set/lp1/ChargePointEnabled
    payload: ${enable:%d} # write payload definition
  maxcurrent:
    type: mqtt
    topic: openWB/set/lp1/DirectChargeAmps
    payload: ${maxCurrent:%d} # write payload definition
```

<a id="charger-phoenix-em-cp-controller-ethernet-modbus-tcp"></a>
#### Phoenix EM-CP Controller (Ethernet/Modbus TCP)

```yaml
- type: phoenix-emcp
  uri: 192.0.2.2:502 # TCP ModBus address
  id: 1
```

<a id="charger-phoenix-ev-cc-controller-modbus"></a>
#### Phoenix EV-CC Controller (Modbus)

```yaml
- type: phoenix-evcc
  device: /dev/ttyUSB0
  baudrate: 9600
  comset: 8N1
  id: 1 # slave id
```

<a id="charger-simple-evse-ethernet-modbus-tcp"></a>
#### Simple EVSE (Ethernet/Modbus TCP)

```yaml
- type: simpleevse
  uri: 192.0.2.2:502 # TCP ModBus address
```

<a id="charger-simple-evse-usb"></a>
#### Simple EVSE (USB)

```yaml
- type: simpleevse
  device: /dev/usb1 # RS485 ModBus device
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

<a id="vehicle-generisch"></a>
#### Generisch

```yaml
- type: default
  title: Mein Auto # display name for UI
  capacity: 50 # kWh
  charge:
    type: ...
    # ...
```

<a id="vehicle-generisch-script"></a>
#### Generisch (Script)

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

<a id="vehicle-vw-eup-id-3-id-4-etc"></a>
#### VW (eUp, ID.3, ID.4, etc)

```yaml
- type: vw
  title: ID.3 # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WVWZZZ... # optional
```

