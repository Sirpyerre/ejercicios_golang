package main

import "fmt"

type persona struct {
	nombre        string
	apellido      string
	saboresHelado []string
}

func main() {

	p1 := persona{
		nombre:        "Peter",
		apellido:      "Rojas",
		saboresHelado: []string{"fresa", "tamarindo", "arroz con leche"},
	}

	p2 := persona{
		nombre:        "Fer",
		apellido:      "Bautista",
		saboresHelado: []string{"fresa", "chamoi", "cafe"},
	}

	fmt.Println(p1.nombre, p1.apellido)
	for i, sabor := range p1.saboresHelado {
		fmt.Printf("\t%d - %v\n", i, sabor)
	}

	fmt.Println(p2.nombre, p2.apellido)
	for j, sabor := range p2.saboresHelado {
		fmt.Printf("\t%d - %v\n", j, sabor)
	}
}

