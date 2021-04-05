package main

import "fmt"

type meutipo int

func main() {
	fmt.Println("Hello world!")
	nmrDeBytes, erros := fmt.Println("Eaeee!", "Fala galera", 100)
	fmt.Println("Numero de bytes: ", nmrDeBytes, "Erros: ", erros)

	x := 16
	var y string = "Fala galera" // y := "Fala galera"
	z := true

	fmt.Printf("X: %v, %T", x, x) // X: 16, int
	fmt.Printf("Y: %v, %T", y, y) // Y: Fala galera, string
	fmt.Printf("Z: %v, %T", z, z) // Z: true, bool
}
