package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Miss","Apellido":"Moneypenny","Edad":32}]
	`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", personas)

	for i, v := range personas {
		fmt.Println("\nPersona Número", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
