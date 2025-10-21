package cola_prioridad

// *** Estructura Interna Heap ***

type heapArr[T any] struct {
	arr       []T
	capacidad int
	cantidad  int
	cmp       func(T, T) int
}

// *** Primitivas Heap ***

// *** HeapSort ***

// *** Funciones auxiliares ***

// heap.redimensionar(nuevaCapacidad) toma una capacidad en enteros
// y redimensiona el arreglo interno a la nueva capacidad
func (heap *heapArr[T]) redimensionar(capacidad int) {

}

// downHeap(pos, arreglo) toma una posicion en arreglo de un elemento y el arreglo
// y aplica algoritmo down-heap al elemento de esa posicion
func downHeap[T any](pos int, arreglo []T) {

}

// upHeap(pos, arreglo) toma una posicion en arreglo de un elemento y el arreglo
// y aplica algoritmo up-heap al elemento de esa posicion
func upHeap[T any](pos int, arreglo []T) {

}

// heapify(arreglo, funcion) toma un arreglo y una funcion de comparacion
// y modifica el arreglo en forma de heap segun la funcion pasada
func heapify[T any](arreglo []T, cmp func()) {

}
