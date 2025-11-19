package parcialitos

import TDACola "tdas/cola"

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
	colaPar := TDACola.CrearColaEnlazada[T]()
	colaImpar := TDACola.CrearColaEnlazada[T]()
	actual := cola.primero
	contador := 0

	for actual != nil {
		if contador%2 == 0 {
			colaPar.Encolar(actual.valor)
		} else {
			colaImpar.Encolar(actual.valor)
		}
		actual = actual.siguiente
		contador++
	}

	return colaPar, colaImpar
}
