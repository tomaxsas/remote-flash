# Remote flash helper

Currently works with TGflash on BMW via ENET.

Use this tool at your own risk!!!

## Setup

- Install Microsoft KM-TEST Loopback adapter
- Set static ip on this new interface to 169.254.10.10 mask 255.255.0.0
  
## Flashing remotely

- Setup VPN connection between your and client PC
- Start some remotely utility (bimmertool from bimmerutility, remote net etc.) on client PC
- On your pc start some ZGW search tool and note IP address (it should be VPN IP)
- Start remoteflash utility
- Enter ZGW IP IP address
- CLick connect to car and Start
- Open TGFlash and start flashing