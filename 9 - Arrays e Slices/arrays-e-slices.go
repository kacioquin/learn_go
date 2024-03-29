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

	fmt.Println("Slices ---------------------")

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "PosicaoAlterada"
	fmt.Println(slice2)

	fmt.Println("Arrays internos ------------------- ")
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println("tamanho - ", len(slice3))
	fmt.Println("capacidade - ", cap(slice3))

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	slice4 = append(slice4, 5)
	fmt.Println("tamanho - ", len(slice4))
	fmt.Println("capacidade - ", cap(slice4))

}
