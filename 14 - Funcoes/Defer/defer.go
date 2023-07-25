package main

import (
	"fmt"
	"time"
)

func closeFakeConnection() {
	fmt.Println("Ok, sua conexão com o banco foi fechada com sucesso!")
}

func getFakeData() {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Second)
		fmt.Println("Aguarde, os dados já serão retornados", i)
	}
	time.Sleep(time.Second)
	fmt.Println("Seus dados estão prontos!")
}

func main() {
	fmt.Println("func main")
	defer closeFakeConnection()
	getFakeData()
}
