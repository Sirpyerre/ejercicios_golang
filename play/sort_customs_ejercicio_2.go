package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type PorEdad []usuario
type PorApellido []usuario

func (p PorEdad) Len() int           { return len(p) }
func (p PorEdad) Less(i, j int) bool { return p[i].Edad < p[j].Edad }
func (p PorEdad) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p PorApellido) Len() int           { return len(p) }
func (p PorApellido) Less(i, j int) bool { return p[i].Apellido < p[j].Apellido }
func (p PorApellido) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}
	//	fmt.Println(usuarios)

	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)

		for _, d := range v.Dichos {
			fmt.Println("\t", d)
		}
	}

	fmt.Println("=================")

	sort.Sort(PorEdad(usuarios))
	//	fmt.Println(usuarios)

	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)

		for _, d := range v.Dichos {
			fmt.Println("\t", d)
		}
	}

	fmt.Println("=================")

	sort.Sort(PorApellido(usuarios))
	//	fmt.Println(usuarios)

	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		sort.Strings(v.Dichos)
		for _, d := range v.Dichos {
			fmt.Println("\t", d)
		}
	}
}
