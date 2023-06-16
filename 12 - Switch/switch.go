package main

import "fmt"

// Primeira maneira
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// Segunda maneira
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
		fallthrough // essa cláusula faz com que a próxima condição seja executada diretamente, sem passar pelo passo da avaliação
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia3 := diaDaSemana(3)
	fmt.Println(dia3)

	dia5 := diaDaSemana(5)
	fmt.Println(dia5)

	dia1 := diaDaSemana(1)
	fmt.Println(dia1)

	dia200 := diaDaSemana(200)
	fmt.Println(dia200)

	fmt.Println()

	diaDaSemana := diaDaSemana2(1)
	fmt.Println(diaDaSemana)

	diaDaSemana = diaDaSemana2(200)
	fmt.Println(diaDaSemana)

	diaDaSemana = diaDaSemana2(7)
	fmt.Println(diaDaSemana)

}
