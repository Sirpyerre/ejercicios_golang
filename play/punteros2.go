package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Printf("antes: valor x=%v, referencia& '%v'\n", x, &x)
	foo(&x)
	fmt.Printf("after: valor x='%v', direccion& '%v'\n", x, &x)
}

func foo(y *int) {
	fmt.Printf("before: direción y='%v', apunta al valor* %v\n", y, *y)
	*y = 42
	fmt.Printf("after: direción y='%v', apunta al valor* %v\n", y, *y)
}

