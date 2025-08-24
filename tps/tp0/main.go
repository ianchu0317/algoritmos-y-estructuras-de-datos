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
	vector := []int{1, 2, 3, 4}
	fmt.Println("Vector es:", vector)
	fmt.Println("Elementos x elemento: ")
	for i := 0; i <= 3; i++ {
		fmt.Println(vector[i])
	}
	/*
		// Probar ordenamiento de seleccion
		vec1 := []int{43, 25, 13, 94, 5, 10, 29, 48}
		fmt.Println("Vector es:", vec1)
		fmt.Println("Ordenado es:", ejercicios.Seleccion(vec1))
	*/

	// Probar Suma de array
	fmt.Println("Vector es:", vector)
	fmt.Println("Suma de elementos es:", ejercicios.Suma(vector))
}
