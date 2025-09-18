package main

import (
	aux "ejemplos/auxiliares"
	sort "ejemplos/ordenamientos"
	"fmt"
)

func main() {
	// Merge sort
	// O(n log(n))
	arr1 := aux.CrearArrayRandom(10)
	fmt.Println("Nuevo array es:", arr1)
	fmt.Println("Array ordenado con Merge Sort es:", sort.MergeSort(arr1))

	// Quick sort
	// O(n log(n))
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	fmt.Println(arr2)
	sort.QuickSort(arr2)
	fmt.Println("Array ordenado con Quick Sort es:", arr2)

	// Counting Sort
	// Como sabemos que arreglo va de 0 a 100
	// Podemos ordenar con counting sort ya que es rango acotado

}
