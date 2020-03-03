package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {
	ag1 := agenteSecreto{
		persona: persona{
			nombre:   "Pedro",
			apellido: "Rojas",
			edad:     36,
		},
		lpm: true,
	}

	p2 := persona{
		nombre:   "Fer",
		apellido: "Bautista",
		edad:     26,
	}

	fmt.Println(p2)

	fmt.Println(ag1.nombre, ag1.apellido, ag1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}

