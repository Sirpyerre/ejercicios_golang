package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {

		if i%4 == 0 {
			fmt.Printf("%d modulo 4, %d\n", i, i%4)
		} else if i%4 == 1 {
			fmt.Printf("%d modulo 4, %d\n", i, i%4)
		} else {
			fmt.Printf("lo demas\n")
		}
	}
}

