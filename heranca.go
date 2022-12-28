package main

import "fmt"

// Utilização de structs como classes (Analogia parecida com a POO)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	// herança de pessoa
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("---Herança---")

	pessoa1 := pessoa{"Patricia", "Cordeiro", 17, 165}
	estudante1 := estudante{pessoa1, "Filosofia", "UFSC"}
	fmt.Println(estudante1.idade)
}
