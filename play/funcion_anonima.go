package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println("función anónima")
	}()

	func(x int) {
		fmt.Println("el significado de la vida es:", x)
	}(42)
}

func foo() {
	fmt.Println("foo se ejecuto")
}

