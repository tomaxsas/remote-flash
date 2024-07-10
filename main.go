package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

const (
	LOCAL_IP_ADDRESS = "169.254.10.10"
	UDP_PORT         = "6811"
	TCP_PORT         = "6801"
)

func main() {
	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)

	// Print the instruction to the reader in the console
	fmt.Println("Please enter your IP address:")

	// Call the reader to read user's input
	vpnIP, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("Your VPN IP address is:", vpnIP)

	udpClient, _ := net.Dial("udp4", vpnIP+UDP_PORT)
	data, err := hex.DecodeString("11")
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x", data)

	udpClient.Write(data)
	carInfoFromUDP, _, err := bufio.NewReader(udpClient).ReadLine()
	fmt.Println("Received from vpnIP", string(carInfoFromUDP))
	if err != nil {
		panic(err)
	}
	// udp server setup
	udpAddr, err := net.ResolveUDPAddr("udp4", LOCAL_IP_ADDRESS+UDP_PORT)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	udpServer, err := net.ListenUDP("udp4", udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from UDP listener in endless loop
	for {
		var buf [512]byte
		_, addr, err := udpServer.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		// Write back the message over UPD
		// TODO need to write what was read from udp message
		udpServer.WriteToUDP(carInfoFromUDP, addr)
	}

}
