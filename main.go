package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Escolha v0, v1, v2 ou v3")
	}

	opcao := os.Args[1]

	if opcao == "v0" {
		RodarV0()
	} else if opcao == "v1" {
		RodarV1()
	} else if opcao == "v2" {
		RodarV2()
	} else if opcao == "v3" {
		RodarV3()
	} else {
		fmt.Println("Opção inválida")
	}
}
