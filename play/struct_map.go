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

	m := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for i, v := range m {
		fmt.Printf("k=%v, v=%v\n", i, v.nombre)

		for _, favor := range v.saboresHelado {
			fmt.Printf("\tfavor=%v\n", favor)
		}
		fmt.Println("-------------------")

	}

}

