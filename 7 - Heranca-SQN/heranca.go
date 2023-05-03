package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type student struct {
	person
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := person{"Cassio", "de Freitas", 33, 169}
	fmt.Println(p1)

	e1 := student{p1, "Sistemas de informação", "FESP-UEMG"}
	fmt.Println(e1.nome)
}
