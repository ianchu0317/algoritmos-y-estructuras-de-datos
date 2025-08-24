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

	// Verificar exceso de item en array
	vector := [4]int{1, 2, 3, 4}
	fmt.Println("Vector es:", vector)
	fmt.Println("Elementos x elemento: ")
	for i := 0; i <= 4; i++ {
		fmt.Println(vector[i])
	}

}
