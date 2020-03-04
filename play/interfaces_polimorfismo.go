package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (p persona) presentarse() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta")
}

type humano interface {
	presentarse()
}

func bar(h humano) {
	fmt.Println("Fui pasado a la función bar", h)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Pedro",
			apellido: "Rojas",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Mafer",
			apellido: "Bau",
		},
		lpm: false,
	}

	p := persona{

		nombre:   "Carito",
		apellido: "Guz",
	}

	fmt.Println(as1)
	as1.presentarse()
	as2.presentarse()

	bar(as1)
	bar(as2)
	bar(p)
}

