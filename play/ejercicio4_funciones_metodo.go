package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Printf("Hola soy %v y tengo %v años\n", p.nombre, p.edad)

}

func main() {
	alumno := persona{
		nombre:   "Pedro",
		apellido: "RR",
		edad:     45,
	}

	alumno.presentar()

}

