package main

import "fmt"

func calculo(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := calculo(10, 20, 3, 4, 500)
	fmt.Println(totalDaSoma)
}
