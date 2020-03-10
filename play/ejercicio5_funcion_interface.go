package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	longitud float32
}

type circulo struct {
	radio float32
}

func (c circulo) calculaArea() float32 {
	return math.Pi * c.radio * c.radio
}

func (sq cuadrado) calculaArea() float32 {
	return sq.longitud * sq.longitud
}

type forma interface {
	calculaArea() float32
}

func info(f forma) {
	fmt.Printf("El área tipo %T es: %v\n", f, f.calculaArea())
}

func main() {

	sq := cuadrado{
		longitud: 10,
	}

	circle := circulo{
		radio: 5,
	}

	info(sq)
	info(circle)

}

