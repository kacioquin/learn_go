package main

import "fmt"

func main() {
	//if else
	numero := -1
	if numero > 0 {
		fmt.Println("número maior que 0")
	} else {
		fmt.Println("número menor ou igual a 0")
	}

	//if init
	//a variável criada no if init vale apenas dentro desse contexto
	if outroNumero := numero; outroNumero < 0 {
		fmt.Println("outroNumero menor que 0")
	}

	// a atribuição abaixo geraria um erro, pois a variável outroNumero não existe fora do contexto do if que foi declarada, no caso, o if acima.
	// ultimoNumero := outroNumero

}
