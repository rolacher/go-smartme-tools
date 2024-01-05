# smecli

smecli is built to show the usage of [rolacher/go-smartme](https://github.com/rolacher/go-smartme) package.
To use the tool, it needs an account on the [Smart-me API](https://api.smart-me.com/swagger/index.html).

## Current features

* List data structure of all devices: `smecli devices`
* Show the data structure of one given device: `smecli devices <device-id>`
* Show the latest values of a device: `smecli values <device-id>`
* Show values from the past of a device: `smecli values <device-id> <date>`
* Show multiple values from of a device: `smecli values <device-id> <start-date> <end-date> <interval>`

**Examples:**

Get information about a device:

```shell
$ smecli devices 40XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXX12 -p
https://api.smart-me.com/Devices/40XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXX12
{
  "id": "40XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXX12",
  "name": "Bilanz VK",
  "serial": 9202922,
  "deviceEnergyType": 1,
  "familyType": 6,
  "activePower": 101.632,
  "activePowerUnit": "kW",
  "counterReading": 2427600.968,
  "counterReadingUnit": "kWh",
  "counterReadingT1": 1174326.392,
  "counterReadingT2": 1253274.576,
  "counterReadingT3": 0,
  "counterReadingT4": 0,
  "counterReadingImport": 2428878.776,
  "counterReadingExport": 1277.808,
  "voltage": 230,
  "voltageL1": 230,
  "voltageL2": 229,
  "voltageL3": 230,
  "currentL1": 166.4,
  "currentL2": 143.28,
  "currentL3": 135.12,
  "valueDate": "2024-01-04T17:25:43.7242222Z",
  "chargeStationState": null
}
```

Getting an interval of measures from the past. The OBIS codes are explained on the [Smart-me Wiki](https://wiki.smart-me.com/schnittstellen/api#h.gwxrm95omkeg) (in German).

```shell
$ smecli values 2cXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXf1 2023-11-16T12:00:00Z 2023-11-16T15:00:00Z -p 
https://api.smart-me.com/ValuesInPastMultiple/2cXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXf1?endDate=2023-11-16T15%3A00%3A00Z&interval=60&startDate=2023-11-16T12%3A00%3A00Z
[
  {
    "deviceId": "2cXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXf1",
    "date": "2023-11-16T12:00:00Z",
    "Values": [
      {
        "obis": "1-0:1.8.0*255",
        "value": 9711213
      },
      {
        "obis": "1-0:2.8.0*255",
        "value": 0
      },
      {
        "obis": "1-0:1.8.1*255",
        "value": 9637442
      },
      {
        "obis": "1-0:1.8.2*255",
        "value": 73771.19
      },
      {
        "obis": "1-0:2.8.1*255",
        "value": 0
      },
      {
        "obis": "1-0:2.8.2*255",
        "value": 0
      },
      {
        "obis": "1-1:5.8.0*255",
        "value": 877158.25
      },
      {
        "obis": "1-1:6.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:7.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:8.8.0*255",
        "value": 5385211.5
      }
    ]
  },
  {
    "deviceId": "2c95f279-899e-4285-8e0d-7dee770561f1",
    "date": "2023-11-16T13:00:00Z",
    "Values": [
      {
        "obis": "1-0:1.8.0*255",
        "value": 9711369
      },
      {
        "obis": "1-0:2.8.0*255",
        "value": 0
      },
      {
        "obis": "1-0:1.8.1*255",
        "value": 9637598
      },
      {
        "obis": "1-0:1.8.2*255",
        "value": 73771.19
      },
      {
        "obis": "1-0:2.8.1*255",
        "value": 0
      },
      {
        "obis": "1-0:2.8.2*255",
        "value": 0
      },
      {
        "obis": "1-1:5.8.0*255",
        "value": 877187.94
      },
      {
        "obis": "1-1:6.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:7.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:8.8.0*255",
        "value": 5385370.5
      }
    ]
  },
  {
    "deviceId": "2c95f279-899e-4285-8e0d-7dee770561f1",
    "date": "2023-11-16T14:00:00Z",
    "Values": [
      {
        "obis": "1-0:1.8.0*255",
        "value": 9711503
      },
      {
        "obis": "1-0:2.8.0*255",
        "value": 0
      },
      {
        "obis": "1-0:1.8.1*255",
        "value": 9637731
      },
      {
        "obis": "1-0:1.8.2*255",
        "value": 73771.19
      },
      {
        "obis": "1-0:2.8.1*255",
        "value": 0
      },
      {
        "obis": "1-0:2.8.2*255",
        "value": 0
      },
      {
        "obis": "1-1:5.8.0*255",
        "value": 877215.6
      },
      {
        "obis": "1-1:6.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:7.8.0*255",
        "value": 0
      },
      {
        "obis": "1-1:8.8.0*255",
        "value": 5385520.5
      }
    ]
  }
]

```

## Configuration

Setup a Json file in your home directory (%USERPROFILE% on Windows) with the following structure:

```json
{
  "Host": "https://smart-me.com:443",
  "Username": "YYYYYYYYY",
  "Password": "XXXXXXXXX",
}
```

You can also use environment variables to configure these values (`SMEAPI_HOST`, `SMEPI_USER`, `SMEPI_PASSWORD`).
