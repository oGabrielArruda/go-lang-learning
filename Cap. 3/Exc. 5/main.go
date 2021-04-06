package main

import (
	"fmt"
)

type meuInt int

var x meuInt
var y int

func main() {
	x = 42
	fmt.Printf("x: %v %T \n", x, x)

	y = int(x)
	fmt.Printf("y: %v %T", y, y)
}
