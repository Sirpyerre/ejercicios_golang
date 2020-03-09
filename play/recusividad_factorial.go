package main

import "fmt"

func main() {

	n1 := factorial(15)
	fmt.Println(n1)

	n2 := cicloFac(16)
	fmt.Println(n2)

}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func cicloFac(n int) int {
	total := 1
	for ; n > 0; n-- {

		total *= n
	}

	return total
}

