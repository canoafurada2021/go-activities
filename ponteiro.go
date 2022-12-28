package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiros são referências de memória

	//guarda um inteiro
	var variavel3 int
	//guarda o endereço de memória de um inteiro
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) // processo de desreferenciação

	//mudança do valor da variável 3 para verificar se há ou não mudança no ponteiro
	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
