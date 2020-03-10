package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hola, en main")

}

func foo() {
	defer func() {
		fmt.Println("foo diferida corrió")
	}()

	fmt.Println("foo corrió")
}

