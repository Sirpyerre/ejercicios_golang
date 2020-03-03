package main

import (
	"fmt"
)

func main() {
	producto := struct {
		nombre   string
		precio   float32
		cantidad int
	}{
		nombre:   "Pantalla Plasma HD",
		precio:   5990.90,
		cantidad: 10,
	}

	fmt.Println(producto.nombre, producto.precio, producto.cantidad)
}

