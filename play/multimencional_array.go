package main

import "fmt"

func main() {
	things := [][]string{
		{"Batman", "Jefe", "Vestido de negro"},
		{"Robin", "Ayudante", "Vestido de colores"},
	}

	for i, item := range things {
		//fmt.Println(i, item)
		for j, val := range item {
			fmt.Println(i, j, val)
		}

	}

}
