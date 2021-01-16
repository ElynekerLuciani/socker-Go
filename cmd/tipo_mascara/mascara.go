package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprint(os.Stderr, "Uso: %s ip qtd-bits-mascara qtd-bits-ip", os.Args[0])
		os.Exit(1)
	}

	addr := net.ParseIP(os.Args[1])

	//verificar se o endereço de ip informado é nulo
	if addr == nil {
		fmt.Println("Endereço Inválido")
		os.Exit(1)
	}

	//conversão de bits da máscara para inteiro
	ones, _ := strconv.Atoi(os.Args[2])

	//conversão dos bits do ip para saber se é ipv4 ou ipv5
	bits, _ := strconv.Atoi(os.Args[3])

	//criar uma máscara de rede do tipo IPMask
	mask := net.CIDRMask(ones, bits)

	//verificar se um endereço pertence a uma rede ou não

	network := addr.Mask(mask)
	fmt.Println(
		"O endereço é: ", addr.String(),
		"\nMáscara em HEX: ", mask.String(),
		"\nRede é: ", network.String())

	os.Exit(0)

}

//para executar dentro da pasta: go run mask.go (ip local por exemplo, consultar no terminal ip addr)
//
