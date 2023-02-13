package main

import "fmt"

const nomeCompleto = "Cassio de Freitas Quintino"

func main() {
	var idade int = 33
	data()
	fmt.Println(idade)
}

func data() {
	fmt.Println(nomeCompleto)
}
