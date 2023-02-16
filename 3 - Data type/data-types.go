package main

import "fmt"

func main() {
	//int
	//alias for int32 -> rune
	var number1 int8 = 100
	fmt.Println(number1)
	var numberAliasInt32 rune = 18
	fmt.Println(numberAliasInt32)

	//unsign int = uint
	//alias for uint8 -> byte
	var number2 uint8 = 10
	fmt.Println(number2)

	//real numbers
	//float32, float64
	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	//strings
	//always use double quotes -> ""
	var str string = "test"
	fmt.Println(str)

	//Go not have "char"
	//when I use single quote, Go prints the ascii number for the string informed
	str2 := 'B'
	fmt.Println(str2)
	//The example above prints 66(ascii number) instead character B
}
