package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	x,s := bar()

	fmt.Println(x, s)
}

func foo() int {
	return 12
}

func bar() (int, string) {
	return 10, "Mega mente"
}

