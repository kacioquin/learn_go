package main

import "fmt"

func main() {
	usuario1 := map[string]string{
		"nome":      "Cassio",
		"sobrenome": " de Freitas",
	}

	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])

	//nested maps
	usuario2 := map[string]map[string]string{
		"dados": {
			"nome":      "John",
			"sobrenome": "Doe",
		},
		"curso": {
			"nome":      "Sistemas de Informação",
			"faculdade": "FESP-UEMG",
		},
	}

	fmt.Println(usuario2)

	//delete properties
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	//add properties
	usuario2["curso"] = map[string]string{
		"nome": "Sistemas de Informação",
	}
	fmt.Println(usuario2)

}
