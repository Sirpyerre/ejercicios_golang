package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := persona{
		nombre:   "Pedro",
		apellido: "Rojas",
		edad:     36,
	}

	p2 := persona{
		nombre:   "Fer",
		apellido: "Bautista",
		edad:     26,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nombre, p1.apellido, p1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}

