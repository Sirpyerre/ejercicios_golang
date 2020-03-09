package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Mi primera expresión función")
	}

	f()

	g := func(x int) {
		fmt.Println("Se descubrió america en", x)
	}

	g(1492)
}

