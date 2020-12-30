package main

import (
	"encoding/base64"
	"io/ioutil"
	"net"
)

func handleClient(conn *net.UDPConn) {
	var netBuffer [8000]byte
	var fileBuffer [8000]byte

	conn.ReadFromUDP(netBuffer[0:])
	//fmt.Println("Dado codificado: ", string(netBuffer[:]))
	//dec, _ := base64.StdEncoding.DecodeString(string(netBuffer[:]))
	//fmt.Println("Dado decodificado: ", dec)
	base64.StdEncoding.Decode(fileBuffer[:], netBuffer[:])

	ioutil.WriteFile("go.jpeg", fileBuffer[:], 0444)

}

func main() {
	service := ":1200"
	addr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.ListenUDP("udp", addr)

	for {
		handleClient(conn)
	}

}
