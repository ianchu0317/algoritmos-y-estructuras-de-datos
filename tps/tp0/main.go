package main

import (
	"fmt"

	"tp0/ejercicios"
)

func main() {
	// Verificación de Swap
	a, b := 5, 6
	fmt.Println("Las variables son a, b=", a, b)
	ejercicios.Swap(&a, &b)
	fmt.Println("Las variables son a, b=", a, b)

}
