package main

import (
	"fmt"
)

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(x)
}
//cantidad variable de parametros
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0

	for i, v := range x {
		fmt.Printf("i:%v, v:%v\n", i, v)
		suma += v
	}

	return suma

}

