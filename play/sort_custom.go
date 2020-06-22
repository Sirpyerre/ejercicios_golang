package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type PorNombre []Persona

func (a PorNombre) Len() int           { return len(a) }
func (a PorNombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorNombre) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {
	p1 := Persona{"Peter", 38}
	p2 := Persona{"Ardd", 19}
	p3 := Persona{"Alice", 52}
	p4 := Persona{"Fer", 26}

	personas := []Persona{p1, p2, p3, p4}

	fmt.Println(personas)

	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

	sort.Sort(PorNombre(personas))

	fmt.Println(personas)

}
