package main

import (
	"fmt"
)

func main() {
	numbers := [5]int{5, 3, 76, 37, 2}

	for i, v := range numbers {
		fmt.Printf("i=%d ===> %v\n", i, v)
	}

	fmt.Printf("El arreglo  es de tipo%T", numbers)
}
