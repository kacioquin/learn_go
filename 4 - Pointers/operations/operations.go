package operations

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func ChangeValue(a *int, b *int) {
	fmt.Printf("Initial memory addresses for a = %v and b = %v\n", a, b)

	temp := *a
	fmt.Println("Value for temp", temp)

	*a = *b
	fmt.Println("Value for a", *a)

	*b = temp
	fmt.Println("Value for b", *b)
}
