package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//verificar a quantidade de argumentos passados pelo usuário
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	//IPAddr que representa um endereco id, se não for ip valido sera nulo
	addr := net.ParseIP(os.Args[1])

	if addr == nil {
		fmt.Println("Endereco inválido.")
	} else {
		fmt.Println("O endereco é: ", addr.String())
	}

	os.Exit(0)

}

//executar dentro da pasta: go run ip.go (numero de ip, ex: 127.0.0.1)
//a saida é a verificação de um ip válido passado
