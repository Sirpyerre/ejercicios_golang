package main

import (
	"fmt"
)

func main() {
	i := 1982
	for {

		if i > 2020 {
			break
		}
		fmt.Println(i)
		i++
	}
}
