package main

import (
	"bufio"
	"encoding/base64"
	"net"
	"os"
)

func readImage() []byte {
	file, _ := os.Open("go.jpeg")
	defer file.Close()

	fileStat, _ := file.Stat()
	var fileSize = fileStat.Size()
	buffer := make([]byte, fileSize)

	reader := bufio.NewReader(file)
	reader.Read(buffer)

	return buffer
}

func main() {

	binary := readImage()

	enc := base64.StdEncoding.EncodeToString(binary)
	service := ":1200"

	addr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte(enc))

	os.Exit(0)

}
