package main

import (
	"fmt"
)

var x int
var g func() = func() {
	fmt.Println("g desde fuera")
}

func main() {
	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("f: %T\n", f)

	g()
	fmt.Printf("g: %T\n", g)
	fmt.Println("Listo")
}

