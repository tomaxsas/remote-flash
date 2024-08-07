# Remote flash helper

Utility designed to make programs work with BMW ENET over VPN conenction. Most tools only accepts direct connections. This utility will overcome this limitation.

Currently works with:

- TGflash
- BitBox
- BMW UFO


>[!IMPORTANT]
>Use this tool at your own risk!!!
>
>Require stable Internet and VPN connection.
>
>Not responsible for damaged ECU.

## Setup

- Install Microsoft KM-TEST Loopback adapter
- Set static ip on this new interface to 169.254.10.10 mask 255.255.0.0

WiFi is preferred connections. Sometimes local ethernet adapter must be disabled for some apps to look at KM-TEST adapter
  
## Connecting remotely

- Setup VPN connection between your and client PC
- Start some remotely utility (Remoteutiliy from bimmerutility is preferred, others might not work.) on client PC
- On your pc start some ZGW search tool or Remoteutiliy and note IP address (it should be VPN IP). ZGW must be found by automatic tools.
- Start remoteflash utility
- Enter ZGW IP address
- CLick connect to car and Start
- Open target program and start flashing/coding/programing

## TODO

- Test Quick flash from bimmertuning tools
