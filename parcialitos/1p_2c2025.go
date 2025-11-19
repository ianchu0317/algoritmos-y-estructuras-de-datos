package parcialitos

import (
	TDACola "tdas/cola"
)

// Enunciados y resolucion del 1 parcialito 2c2025

/*

*** Ejercicio 1 ***

Implementar una primitiva del TDA Cola que devuelva dos colas:
una con los elementos de las posiciones pares, y otra con los elemntos de posiciones impares.
La colaoriginal debe quedar en el mismo estado inicial. Indicar y justificar la complejidad de la primitiva.
La firma de la primitiva debe ser Dividir() (Cola[T], Cola[T])

*/

type nodoCola[T any] struct {
	valor     T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func (cola colaEnlazada[T]) Dividir() (TDACola.Cola[T], TDACola.Cola[T]) {
	// ****************************************************
	// COmplejidad total O(1) + O(n) = O(n)
	// ****************************************************
	// Declaraciones de estructuras auxiliares y asignación de variables O(1)
	colaPar := TDACola.CrearColaEnlazada[T]()
	colaImpar := TDACola.CrearColaEnlazada[T]()
	actual := cola.primero
	contador := 0

	// Visitar cada n-nodo de la cola y realizar operaciones de O(1) -> n * O(1) = O(n)
	for actual != nil {
		if contador%2 == 0 {
			colaPar.Encolar(actual.valor)
		} else {
			colaImpar.Encolar(actual.valor)
		}
		actual = actual.siguiente
		contador++
	}

	// Return O(1)
	return colaPar, colaImpar
}

/*
*** EJercicio 2 ***

Implementar una función que, dado un arreglo ordenado y sin elementos repetidos de valores enteros no negativos,
obtenga el mínimo valor que no se encuentre en el arreglo.
Indicar y justificar adecuadamente la complejidad del algoritmo.
*/

func MinimoExcluido(arr []int) int {
	return _minimoExcluido(arr, 0, len(arr)-1)
}

func _minimoExcluido(arr []int, ini, fin int) int {
	// ****************************************************
	// COmplejidad total O(log(n)) Por teorema maestro
	// ****************************************************
	if ini == fin {
		if arr[ini] > ini {
			return ini
		} else {
			return ini + 1
		}
	}
	medio := (ini + fin) / 2
	if arr[medio] == medio {
		return _minimoExcluido(arr, medio+1, fin)
	} else {
		return _minimoExcluido(arr, ini, medio-1)
	}
}

/*
*** Ejercicio 3 ***
implementar una función masGrandePosible(digitos []int) int que dado un arreglo de digitos [0-9]
determine cuál es el número más grande que se puede formar con dichos dígitos.
Indicar y justificar la complejidad del algoritmo implementado.
*/

func MasGrandePosible(digitos []int) int {
	// ****************************************************
	// COmplejidad total O(n)
	// ****************************************************
	// O(n) counting sort
	ordenados := ordenarDigitos(digitos, 10)
	num := 0
	// Armar numero O(n) -> visito para cada nodo de arreglo y hago operacion O(1)
	for digito := range ordenados {
		num = num*10 + digito
	}
	return num
}

func ordenarDigitos(arr []int, rango int) []int {
	// Counting sort O(n + 10)
	// Contar frecuencias O(n)
	frecuencias := make([]int, rango)
	for num := range arr {
		frecuencias[num]++
	}
	// Calcular inicios O(k) -> O(10)
	inicios := make([]int, rango)
	for i := 1; i < rango; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}
	// Ordenar elementos O(n)
	ordenados := make([]int, len(arr))
	maxLen := len(arr) - 1
	for num := range arr {
		indx := maxLen - inicios[num]
		ordenados[indx] = num
		inicios[num]++
	}
	return ordenados
}
