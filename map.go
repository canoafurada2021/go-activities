package main

import (
	"fmt"
)

func main() {

	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "antonio",
	}
	fmt.Println(usuario["nome"])

  // forma de utilizar um map dentro de outro map
	usuario2 := map[string]map[string]string{
		"Nome": {
			"primeiro": "Patricia",
			"ultimo":   "carneiro",
		},
		"curso": {
			"nome":  "educação física",
			"local": "USP",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "Nome")

	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}
	fmt.Println(usuario2)
}
