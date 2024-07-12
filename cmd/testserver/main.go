package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

func main() {

	carudpAddr, _ := net.ResolveUDPAddr("udp4", "192.168.100.16:6811")
	udpServerCar, err := net.ListenUDP("udp4", carudpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	carInfo, _ := hex.DecodeString("000000320011444941474144523130424d574d4143303031413337303444413635424d5756494e574241575933313033304c353233363131")
	// Read from UDP listener in endless loop
	for {
		var buf [512]byte
		_, addr, err := udpServerCar.ReadFromUDP(buf[0:])
		fmt.Println("got requests")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Write back the message over UPD
		// TODO need to write what was read from udp message
		udpServerCar.WriteToUDP(carInfo, addr)
	}

}
