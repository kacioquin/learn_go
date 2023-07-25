package main

import (
	"fmt"
	"math/rand"
	"time"
)

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Ok! Houve um pânico, mas não entre em pânico. A execução foi recuperada!")
	}
}

func fakePanic(n1 int) int {
	fmt.Println("Atenção, vamos gerar um PANIC!!!", n1)

	if n1 == 10 {
		panic("PANIC, PANIC, PANIC")
	}
	return n1
}

func main() {
	fmt.Println("main")
	defer recoverExecution()

	for i := 1; i <= 50; i++ {
		time.Sleep(time.Second)
		fakePanic(rand.Intn(20))
	}
}
