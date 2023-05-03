package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")
	fmt.Println("Arrays ---------------------")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"nome", "idade", "altura", "sexo", "teste"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

}
