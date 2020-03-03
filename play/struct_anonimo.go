package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := struct {
		nombre   string
		apellido string
		edad     int
	}{
		nombre:   "Pedro",
		apellido: "Rojas",
		edad:     42,
	}

	fmt.Println(p1)
}

