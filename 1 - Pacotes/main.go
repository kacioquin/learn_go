package main

import (
	"fmt"
	"modulo/secondPackage"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Testing modules")
	secondPackage.Write()

	err := checkmail.ValidateFormat("test@test.com")
	fmt.Println(err)
}
