package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {
	fmt.Println("Funções recursivas")

	posicaoFinal := fibonacci(6)

	for i := uint(1); i <= posicaoFinal; i++ {
		fmt.Printf("Sequência %d \n", fibonacci(i))
	}

	fmt.Println("Posição final", posicaoFinal)

}
