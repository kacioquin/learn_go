package main

import (
	"examples/pointers/operations"
	"fmt"
)

func main() {
	var x int = 10
	var y int = 20

	// fmt.Printf("Values for x = %v and y = %v\n", x, y)
	// fmt.Printf("Pointers for: x = %v and y = %v\n", &x, &y)

	// //z receive the memory address given to x
	// z := &x

	// //show the pointer created for z variable
	// fmt.Println("Pointer for z =", &z)

	// //show the value in z. In this case will be the memory address given to x
	// fmt.Println("Value for z =", z)

	// //show the value stored in memory address saved in z that in this case is the memory address given to x
	// fmt.Println("Value stored in memory address saved in z =", *z)

	fmt.Printf("Initial values: x = %v and y = %v\n", x, y)

	fmt.Printf("Initial memory addresses for x = %v and y = %v\n", &x, &y)

	operations.ChangeValue(&x, &y)

	fmt.Printf("Value after 'ChangeValue function': x = %v, y = %v\n", x, y)
	fmt.Printf("Memory address after 'ChangeValue function': x = %v, y = %v\n", &x, &y)
}
