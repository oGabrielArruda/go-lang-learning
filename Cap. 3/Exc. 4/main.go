package main

import (
	"fmt"
)

type meuInt int

var x meuInt

func main() {
	x = 42
	fmt.Printf("%v %T", x, x)
}
