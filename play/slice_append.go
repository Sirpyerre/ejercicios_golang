package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	y := []int{333, 444, 666}
	fmt.Println(x)

	x = append(x, 44, 55, 34)
	fmt.Println(x)

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}

