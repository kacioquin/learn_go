package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 150
	ponteiro = &variavel3

	var variavel4 int = *ponteiro

	fmt.Println(variavel3, ponteiro, variavel4)

}
