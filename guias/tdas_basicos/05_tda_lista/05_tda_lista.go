package main

import "fmt"

type nodoLista[T any] struct {
	prox *nodoLista[T]
	dato T
}

type ListaEnlazada[T any] struct {
	prim *nodoLista[T]
}

// Agregar un elemento a la lista
func (lista *ListaEnlazada[T]) Enlistar(elemento T) {
	actual := lista.prim
	anterior := lista.prim
	nuevoNodo := nodoLista[T]{prox: nil, dato: elemento}

	// ir al ultimo nodo
	for actual != nil {
		anterior = actual
		actual = actual.prox
	}
	// agregar nuevo nodo
	anterior.prox = &nuevoNodo
}

// Muestra todos los elementos
func (lista ListaEnlazada[T]) Mostrar() {
	actual := lista.prim

	// ir al ultimo nodo
	for actual != nil {
		actual = actual.prox
		fmt.Printf("%v ", actual.dato)
	}
}

func main() {
	miLista := ListaEnlazada[int]{prim: nil}
	vec1 := []int{1, 5, 10, 3, 6, 8}

	for _, e := range vec1 {
		fmt.Printf("%v ", e)
		miLista.Enlistar(e)
	}
	miLista.Mostrar()
}
