package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Arrays e Slices -----")

	// primeira maneira de popular um array
	var array1 [5]int
	array1[0] = 5
	array1[1] = 4
	array1[2] = 3
	array1[3] = 2
	array1[4] = 1
	fmt.Println(array1)

	//segunda maneira de popular um array
	array2 := [5]string{"teste1", "teste2", "teste3", "teste3", "teste4"}
	fmt.Println(array2)

	// terceira maneira de poupular um array
	// esta maneira não o faz ter um tamanho dinâmico, mas sim ter o mesmo numero de posições que você popular nele
	array3 := [...]int{1, 3, 4, 5, 6}
	fmt.Println(array3)

	//--- sllices ---

	// slices se parecem com arrays mas não são arrays

	slice := []int{9, 10, 11, 12, 13}
	fmt.Println(slice)

	// maneira de adicionar itens em um slice
	slice = append(slice, 64)
	fmt.Println(slice)

	// ---- Arrays internos ----
	fmt.Println("----------------")

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
}
