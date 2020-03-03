package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}

	fmt.Println(m)

	fmt.Println(m["Batman"])

	if v, ok := m["Robin"]; ok {
		fmt.Printf("desde if, la edad es: %v\n", v)
	}

	for key, val := range m {
		fmt.Println(key, val)
	}
