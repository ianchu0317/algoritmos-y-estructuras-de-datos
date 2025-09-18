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
	// Quick sort
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	fmt.Println(arr2)
	sort.QuickSort(arr2)
	fmt.Println("Array ordenado con Quick Sort es:", arr2)
}
