package main

import (
	"fmt"
	"io/ioutil"
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
		fmt.Fprintf(os.Stderr, "Uso: %s\n host:porta", os.Args[0])
		os.Exit(1)
	}

	//armazenar o endereço do serviço conectado
	service := os.Args[1]

	//transformar o endereco para estabelecer conexao
	//TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	//usar o endereco tcp para criar um socket
	//retorna TCPConn
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//ler os bytes da resposta
	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result))
	os.Exit(0)

}
