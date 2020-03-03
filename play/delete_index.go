package main

import "fmt"

func main() {
	/*
		favoriteThings1 := []string{
			"computers", "mounting", "beach",
		}

	*/
	favoriteGame := map[string][]string{
		"peter_rojas": {
			"computers", "mounting", "beach",
		},
		"carlos hernandez": {
			"read", "buy", "music",
		},
		"juan perez": {
			"ice cream", "painting", "dance",
		},
	}

	favoriteGame["lalo peña"] = []string{
		`swiming`, `pool dancing`, `speak foreing languages`,
	}

	for i, item := range favoriteGame {
		fmt.Println(i)
		for j, v := range item {
			fmt.Printf("\ti=%v, favorite=%v\n", j, v)
		}
	}

	delete(favoriteGame, "pedro_rojas")

	fmt.Println(favoriteGame)

}
