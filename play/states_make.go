package main

import (
	"fmt"
)

func main() {
	estados := make([]string, 33, 33)
	estados = []string{
		"Aguascalientes",
		"Baja California",
		"Baja California Sur",
		"Campeche",
		"Chiapas",
		"Chihuahua",
		"Ciudad de México",
		"Coahuila de Zaragoza",
		"Colima",
		"Durango",
		"Estado de México",
		"Guanajuato",
		"Guerrero",
		"Hidalgo",
		"Jalisco",
		"Michoacán de Ocampo",
		"Morelos",
		"Nayarit",
		"Nuevo León",
		"Oaxaca",
		"Puebla",
		"Querétaro",
		"Quintana Roo",
		"San Luis Potosí",
		"Sin Localidad",
		"Sinaloa",
		"Sonora",
		"Tabasco",
		"Tamaulipas",
		"Tlaxcala",
		"Veracruz de Ignacio de la Llave",
		"Yucatán",
		"Zacatecas",
	}

	for i := 0; i < len(estados); i++ {

		fmt.Printf("%v - %v\n", i, estados[i])
	}

	fmt.Println(len(estados))
	fmt.Println(cap(estados))

}

