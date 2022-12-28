package main

import (
	"fmt"
)

func menuCalculadora() {

	fmt.Println("Bem vindo! Selecione abaixo qual operação deseja ser feita")
	fmt.Println("Digite 1 para soma")
	fmt.Println("Digite 2 para subtração")
	fmt.Println("Digite 3 para multiplicação")
	fmt.Println("Digite 4 para divisão")

}

func main() {
	var num2, num1, numSelecionado int

	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número")
	fmt.Scanln(&num2)

	menuCalculadora()
	fmt.Scanln(&numSelecionado)

	switch numSelecionado {

	case 1:
		fmt.Println(soma(int8(num1), int8(num2)))

	case 2:
		fmt.Println(subtracao(int8(num1), int8(num2)))

	case 3:
		fmt.Println(multiplicacao(int8(num1), int8(num2)))

	case 4:
		fmt.Println(divisao(float32(num1), float32(num2)))

	}

}

func divisao(numero1, numero2 float32) float32 {
	var divisao = numero1 / numero2
	fmt.Println("Divisão total dos números")
	return divisao
}
func multiplicacao(numero1, numero2 int8) int8 {
	var multiplicacao = numero1 * numero2

	fmt.Println("Multiplicação total dos números")
	return multiplicacao
}

func subtracao(numero1, numero2 int8) int8 {
	var subtracaoTotal = numero1 - numero2
	fmt.Println("Subtração total dos números")
	return subtracaoTotal
}

func soma(numero1, numero2 int8) int8 {
	var somaTotal = numero1 + numero2

	fmt.Println("Soma total dos números")
	return somaTotal
}
