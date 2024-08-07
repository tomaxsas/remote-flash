package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/inetaf/tcpproxy"
)

const (
	LOCAL_IP_ADDRESS = "169.254.10.10"
	UDP_PORT         = "6811"
	TCP_PORT         = "6801"
)

var carInfo []byte

func getCarInfo(ip string) error {

	vpnIP := strings.TrimSuffix(strings.TrimSpace(ip), "\n")

	// udpClient, _ := net.Dial("udp4", vpnIP+":20000")
	udpClient, err := net.Dial("udp4", vpnIP+":"+UDP_PORT)
	if err != nil {
		return err
	}
	data, err := hex.DecodeString("000000000011")
	if err != nil {
		panic(err)
	}
	udpClient.SetDeadline(time.Now().Add(time.Millisecond * 1500))
	_, err = udpClient.Write(data)
	if err != nil {
		log.Println("Got error writing to UDP client: ", err)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return fmt.Errorf("timeout sending data to car")
		}
		return err
	}
	carInfoFromUDP := make([]byte, 512)
	_, err = udpClient.Read(carInfoFromUDP)

	if err != nil {
		log.Println("Got error while reading car data: ", err)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return fmt.Errorf("timeout reading car data")
		}
		return err
	}
	carInfo = carInfoFromUDP
	return nil
}
func startProxy(vpnIP string) error {
	// udp server setup
	udpServerAddr, err := net.ResolveUDPAddr("udp4", LOCAL_IP_ADDRESS+":"+UDP_PORT)

	if err != nil {
		log.Println("Got error resolving UDP server address: ", err)
		return err
	}

	udpServer, err := net.ListenUDP("udp4", udpServerAddr)

	if err != nil {
		log.Println("Got error on starting listening on UDP: ", err)
		return err
	}
	log.Println("Starting TCP proxy")
	go func() {
		var p tcpproxy.Proxy
		p.AddRoute(LOCAL_IP_ADDRESS+":"+TCP_PORT, tcpproxy.To(vpnIP+":"+TCP_PORT))
		log.Fatal(p.Run())

	}()

	// Read from UDP listener in endless loop
	log.Println("Starting UDP proxy")
	go func() {
		for {
			var buf [512]byte
			_, addr, err := udpServer.ReadFromUDP(buf[0:])
			if err != nil {
				log.Println("Got error while starting UDP server", err)
			}
			log.Println("Got car info UDP request")

			// Write back the message over UPD
			udpServer.WriteToUDP(carInfo, addr)
		}
	}()
	return nil

}
