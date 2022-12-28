package main

import "fmt"

// Utilização de structs como classes (Analogia parecida com a POO)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {

	enderecoEx := endereco{"Rua Florianópolis", 370}

	user := usuario{"Samuel", 16, enderecoEx}

	fmt.Println(user)
}
