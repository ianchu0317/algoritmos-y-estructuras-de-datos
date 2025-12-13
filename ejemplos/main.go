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
	fmt.Println("\n***Merge Sort****")
	fmt.Println("Nuevo array es:", arr1)
	fmt.Println("Array ordenado con Merge Sort es:", sort.MergeSort(arr1))

	// Quick sort
	// O(n log(n))
	arr2 := aux.CrearArrayRandom(10)
	fmt.Println("\n***Quick Sort****")
	fmt.Println("\nArray es:", arr2)
	sort.QuickSort(arr2)
	fmt.Println("Array ordenado con Quick Sort es:", arr2)

	// Counting Sort
	// Como sabemos que arreglo va de 0 a 100
	// Podemos ordenar con counting sort ya que es rango acotado
	arr3 := aux.CrearArrayRandom(10)
	fmt.Println("\n***Counting Sort****")
	fmt.Println("Nuevo array es:", arr3)
	fmt.Println("Array ordenado con Counting Sort es:", sort.CountingSort(arr3))

	// Radix Sort
	// Como sabemos que arreglo va de 0 a 100 -> cada num tiene 2 digitos
	// Podemos ordenar con Radix Sort auxiliar counting Sort
	arr4 := aux.CrearArrayRandom(10)
	fmt.Println("\n***Radix Sort****")
	fmt.Println("Nuevo array es:", arr4)
	fmt.Println("Array ordenado con Radix Sort es:", sort.RadixSort(arr4))

	// Bucket Sort
	arr5 := aux.CrearArrayRandom(20)
	fmt.Println("\n***Bucket Sort****")
	fmt.Println("Nuevo array es:", arr5)
	fmt.Println("Array ordenado con Bucket Sort es:", sort.BucketSort(arr5))
}
