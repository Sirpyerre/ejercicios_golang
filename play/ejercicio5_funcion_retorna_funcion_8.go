package main

import (
	"fmt"
)

func main() {
	f := bar()

	fmt.Printf("Cual es el significado de la vida? %v\n", f())
}

func bar() func() int {
	return func() int {
		return 42
	}
}

