package main

import (
	"fmt"
)

func main() {
	x := []int{5, 3, 76, 37, 2}

	for i, v := range x {
		fmt.Println(i, v)
	}

}
