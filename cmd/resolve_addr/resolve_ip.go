package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	//passando um nome qualquer para o programa, vamos usar uma função para
	//traduzir um nome para um ip

	addr, err := net.ResolveIPAddr("ip", os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Endereço Resolvido: ", addr.String())
	os.Exit(0)
}

//para executar: na pasta executar go run resolve_ip (endereço wwww)
