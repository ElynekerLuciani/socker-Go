package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s host:porta", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	addr, _ := net.ResolveUDPAddr("udp", service)

	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte("Uma frase"))
	var buffer [512]byte
	n, _ := conn.Read(buffer[0:])
	fmt.Println(string(buffer[0:n]))

	os.Exit(0)

}
