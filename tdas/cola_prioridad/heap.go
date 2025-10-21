package cola_prioridad

// *** Estructura Interna Heap ***

type heapArr[T any] struct {
	arreglo   []T
	capacidad int
	cantidad  int
	cmp       func(T, T) int
}

// *** Primitivas Heap ***

func (heap heapArr[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heapArr[T]) Encolar(elemento T) {
	// Agregar elemento al ultimo
	heap.arreglo[heap.cantidad] = elemento
	upHeap(heap.cantidad, heap.arreglo, heap.cmp)
	heap.cantidad++
	// Chequeo redimension
	if heap.cantidad == heap.capacidad {
		heap.redimensionar(heap.capacidad * 2)
	}
}

func (heap heapArr[T]) VerMax() T {
	return heap.arreglo[0]
}

func (heap heapArr[T]) Desencolar() T {
	return heap.arreglo[0]
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

// swap toma dos elementos y los intercambia de lugar
func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// downHeap(pos, arreglo) toma una posicion en arreglo de un elemento y el arreglo
// y aplica algoritmo down-heap al elemento de esa posicion
func downHeap[T any](pos int, arreglo []T) {

}

// upHeap(pos, arreglo) toma una posicion en arreglo de un elemento y el arreglo
// y aplica algoritmo up-heap al elemento de esa posicion
func upHeap[T any](pos int, arreglo []T, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	// Hallar posicion padre dependiendo si es hijo izquierdo o derecho
	posPadre := (pos - 1) / 2
	if pos%2 == 0 {
		posPadre = (pos - 2) / 2
	}
	padre := arreglo[posPadre]
	actual := arreglo[pos]

	// Comparar padre con hijo
	if cmp(actual, padre) > 0 {
		swap(&actual, &padre)
	} else {
		return
	}
	// Llamada recursiva
	upHeap(posPadre, arreglo, cmp)
}

// heapify(arreglo, funcion) toma un arreglo y una funcion de comparacion
// y modifica el arreglo en forma de heap segun la funcion pasada
func heapify[T any](arreglo []T, cmp func(T, T) int) {

}
