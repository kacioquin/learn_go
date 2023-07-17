package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Retorno da função -> %s", texto)
	}("funcao anonima com parametro")

	fmt.Println(retorno)
}
