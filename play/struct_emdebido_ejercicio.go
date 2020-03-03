package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	tractoCamion := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "negro",
		},
		cuatroRuedas: false,
	}

	vocho := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "rojo",
		},
		lujoso: true,
	}

	fmt.Println(tractoCamion.puertas, tractoCamion.color, tractoCamion.cuatroRuedas)
	fmt.Println(vocho.puertas, vocho.color, vocho.lujoso)

}

