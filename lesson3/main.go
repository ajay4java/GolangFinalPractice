package main

import "fmt"

func main() {

	var smallPositiveVal uint8 = 10

	fmt.Println(smallPositiveVal)

	var smallPositiveNegativeVal int8 = -10
	fmt.Println(smallPositiveNegativeVal)

	var myInt int = 2300006
	fmt.Println(myInt)

	myInt = int(smallPositiveVal)
	fmt.Println(myInt)

	var myByte byte = 'A'
	fmt.Println(myByte)
	myByte = 10
	fmt.Println(myByte)

	var myRune rune = 'B'
	fmt.Println(myRune)

	myRune = 'ē'
	fmt.Println(myRune)
}
