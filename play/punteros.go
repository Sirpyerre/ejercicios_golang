package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	//imprimir la dirección de memoria
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// v apunta a la dirección de a
	var v *int = &a
	fmt.Println(v)
	// Imprimir el valor en la dirección de a
	fmt.Println(*v)

	*v = 43
	fmt.Println(a)
}

