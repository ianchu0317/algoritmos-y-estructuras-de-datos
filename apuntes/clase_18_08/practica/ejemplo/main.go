package main

import "fmt"

func main() {
	// Variables
	var a int = 5
	var b int = 6
	fmt.Println(a, b)

	// Variables sin tipo
	c := 7
	d := 8
	fmt.Println(c, d)

	// Variables múltiples
	var e, f int = 9, 10
	fmt.Println(e, f)

	// Intercambio de variables
	a, b = b, a // swap
	fmt.Println(a, b)
}
