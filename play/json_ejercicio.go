package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Peter",
		Edad:   38,
	}

	u2 := usuario{
		Nombre: "Fer",
		Edad:   29,
	}

	u3 := usuario{
		Nombre: "Angy",
		Edad:   37,
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println("se produjo un error")
	}

	fmt.Print(string(bs))
}
