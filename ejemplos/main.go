package main

import (
	aux "ejemplos/auxiliares"
	sort "ejemplos/ordenamientos"
	"fmt"
)

func main() {
	// Merge sort
	arr1 := aux.CrearArrayRandom(10)
	fmt.Println("Nuevo array es:", arr1)
	fmt.Println("Array ordenado con Merge Sort es:", sort.MergeSort(arr1))
}
