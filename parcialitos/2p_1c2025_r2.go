package parcialitos

/*
Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para
los que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente
en el heap. Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir
operando. Justificar la complejidad de la primitiva implementada.*/

type Heap[T any] struct {
	arreglo  []T
	cantidad int
}

func (heap *Heap[T]) Filter(f func(T) bool) {
	// Inicializacion de variables O(1)
	aux := make([]T, len(heap.arreglo))
	cantidadAux := 0

	// para los n-elementos del heap, agregar a arreglo auxiliar si pasa filtro
	for i := 0; i < heap.cantidad; i++ {
		elem := heap.arreglo[i]
		if f(elem) {
			aux[cantidadAux] = elem
			cantidadAux++
		}
	}
	// Actualizar variabels O(1)
	heap.arreglo = aux
	heap.cantidad = cantidadAux

	// O(n) se considera ya implementada
	heapify(heap.arreglo, heap.cantidad, heap.cmp) // Le paso arreglo, cantidad de arreglo, funcion comparacion
	heap.chequearRedimension()                     // Chequear redimension del arreglo si hace falta / Sino se chequea en proximo encolar/desencolar
}
