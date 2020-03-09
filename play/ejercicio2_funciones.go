package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 4, 7, 6, 4}
	total := foo(numeros...)

	fmt.Println(total)
	
	totalBar := bar(numeros)
	fmt.Println(totalBar)
}

func foo(x ...int) int {

	suma := 0
	for _, v := range x {

		suma += v
	}

	return suma
}

func bar(y[] int) int {

	suma := 0
	for _, v := range y {

		suma += v
	}

	return suma
}

