package main

import (
	"net"
	"time"
)

func handleClient(conn *net.UDPConn) {
	var buffer [512]byte
	_, addr, _ := conn.ReadFromUDP(buffer[0:])

	dayTime := time.Now().String()

	conn.WriteToUDP([]byte(dayTime), addr)

}

func main() {
	service := ":1200"
	udpAddr, _ := net.ResolveUDPAddr("udp", service)

	conn, _ := net.ListenUDP("udp", udpAddr)

	for {
		handleClient(conn)
	}

}
