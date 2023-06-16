package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	// fmt.Println("Semelhante a while....")

	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("Semelhante a um for em outra linguagem....")

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	fmt.Println("Loops com o uso do 'range'. O range é usado quando precisamos iterar sobre alguma estrutura")
	fmt.Println("Iterando sobre um array ou slice")

	nomes := []string{"Cassio", "João", "Maria"}
	fmt.Println("slice ->", nomes)
	for index, value := range nomes {
		fmt.Println(index, value)
	}

	fmt.Println("Iterando sobre uma string")
	fmt.Println("string -> PALAVRA")

	for index, value := range "PALAVRA" {
		fmt.Println(index, string(value))
	}

	fmt.Println("Iterando sobre um map")
	dados := map[string]string{
		"nome":      "Cassio",
		"sobrenome": "Freitas",
	}

	for index, value := range dados {
		fmt.Println(index, value)
	}

	// Loop infinito
	// for {
	// 	fmt.Print("Ao infinito e além")
	// }

}
