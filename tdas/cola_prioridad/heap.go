package cola_prioridad

// *** Estructura Interna Heap ***

type heapArr[T any] struct {
	arr       []T
	capacidad int
	cantidad  int
	cmp       func(T, T) int
}

// *** Primitivas Heap ***

func (heap heapArr[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heapArr[T]) Encolar(T) {

}

func (heap heapArr[T]) VerMax() T {
	return heap.arr[0]
}

func (heap heapArr[T]) Desencolar() T {
	return heap.arr[0]
}

func (heap heapArr[T]) Cantidad() int {
	return heap.cantidad
}

// *** Funciones de creacion de Heap (Constructores)***
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := heapArr[T]{cantidad: 0, cmp: funcion_cmp}
	heap.redimensionar(0)
	return &heap
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heapify(arreglo, funcion_cmp)
	heap := heapArr[T]{cantidad: 0, cmp: funcion_cmp}
	return &heap
}

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
func heapify[T any](arreglo []T, cmp func(T, T) int) {

}
