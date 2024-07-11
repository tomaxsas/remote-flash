package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/inetaf/tcpproxy"
)

const (
	LOCAL_IP_ADDRESS = "169.254.10.10"
	UDP_PORT         = "6811"
	TCP_PORT         = "6801"
)

func main() {

	readerIP := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your IP address:")

	// Call the reader to read user's input
	vpnIP, err := readerIP.ReadString('\n')
	if err != nil {
		panic(err)
	}
	vpnIP = strings.TrimSuffix(strings.TrimSpace(vpnIP), "\n")
	fmt.Println("Your VPN IP address is:", vpnIP)

	// udpClient, _ := net.Dial("udp4", vpnIP+":20000")
	udpClient, err := net.Dial("udp4", vpnIP+":"+UDP_PORT)
	if err != nil {
		fmt.Println("Error connecting to UDP server in net Dial")
		panic(err)
	}
	data, err := hex.DecodeString("000000000011")
	if err != nil {
		panic(err)
	}

	_, err = udpClient.Write(data)
	if err != nil {
		panic(err)
	}
	carInfoFromUDP := make([]byte, 512)
	_, err = udpClient.Read(carInfoFromUDP)

	fmt.Println("Received car Info from vpnIP", string(carInfoFromUDP))
	if err != nil {
		panic(err)
	}
	// udp server setup
	udpServerAddr, err := net.ResolveUDPAddr("udp4", LOCAL_IP_ADDRESS+":"+UDP_PORT)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	udpServer, err := net.ListenUDP("udp4", udpServerAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Starting TCP proxy")
	go func() {
		var p tcpproxy.Proxy
		p.AddRoute(LOCAL_IP_ADDRESS+":"+TCP_PORT, tcpproxy.To(vpnIP+":"+TCP_PORT))
		log.Fatal(p.Run())

	}()

	// Read from UDP listener in endless loop
	fmt.Println("Starting UDP proxy")
	for {
		var buf [512]byte
		_, addr, err := udpServer.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Got UDP request from tgflash")

		// Write back the message over UPD
		// TODO need to write what was read from udp message
		udpServer.WriteToUDP(carInfoFromUDP, addr)
	}

}
