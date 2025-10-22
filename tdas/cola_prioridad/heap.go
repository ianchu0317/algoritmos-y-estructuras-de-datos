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
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arreglo[0]
}

func (heap *heapArr[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	primerElem := heap.arreglo[0]
	// Swap ultimo elemento y downheap al primer elemnto
	heap.cantidad--
	heap.arreglo[0] = heap.arreglo[heap.cantidad]
	downHeap(0, heap.arreglo[:heap.cantidad+1], heap.cmp)
	// Chequear redimension
	if heap.cantidad < heap.capacidad/4 {
		heap.redimensionar(heap.capacidad / 2)
	}
	return primerElem
}

func (heap heapArr[T]) Cantidad() int {
	return heap.cantidad
}

// *** Funciones de creacion de Heap (Constructores)***

// CrearHeap(cmp) toma funcion de comparacion cmp
// y devuelve una instancia de Heap con la prioridad pasada
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := heapArr[T]{
		cantidad: 0,
		cmp:      funcion_cmp,
	}
	heap.redimensionar(1)
	return &heap
}

// CrearHeapArr toma un arreglo y una funcion de comparacion
// y crea un Heap a partir de la prioridad y los elementos del arreglo
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	aux := make([]T, len(arreglo))
	copy(aux, arreglo)
	heapify(aux, funcion_cmp)
	heap := heapArr[T]{
		arreglo:   aux,
		cantidad:  len(arreglo),
		capacidad: len(arreglo) + 1,
		cmp:       funcion_cmp,
	}
	return &heap
}

// *** HeapSort ***
func HeapSort[T any](arr []T, funcion_cmp func(T, T) int) {
	// Heapify el arreglo
	heapify(arr, funcion_cmp)
	// Swap primer y ultimo elemento
	// Downheap del primer elemento al anteultimo
	for i := 0; i < len(arr); i++ {
		swap(&arr[0], &arr[len(arr)-1-i])
		downHeap(0, arr[:len(arr)-1-i], funcion_cmp)
	}
}

// *** Funciones auxiliares ***

// heap.redimensionar(nuevaCapacidad) toma una capacidad en enteros
// y redimensiona el arreglo interno a la nueva capacidad
func (heap *heapArr[T]) redimensionar(capacidad int) {
	aux := make([]T, capacidad)
	copy(aux, heap.arreglo)
	heap.arreglo = aux
	heap.capacidad = capacidad
}

// swap toma dos elementos y los intercambia de lugar
func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// downHeap(pos, arreglo) toma una posicion en arreglo de un elemento y el arreglo
// y aplica algoritmo down-heap al elemento de esa posicion
func downHeap[T any](pos int, arreglo []T, cmp func(T, T) int) {
	// Si se pasa de este indice entonces el resto son hijos
	if pos > len(arreglo)/2 {
		return
	}
	indxIzq := pos*2 + 1
	indxDer := pos*2 + 2
	mayor := pos
	// Comparar y encontrar el mayor de los hijos si existe
	if indxIzq < len(arreglo) && cmp(arreglo[mayor], arreglo[indxIzq]) < 0 {
		mayor = indxIzq
	}
	if indxDer < len(arreglo) && cmp(arreglo[mayor], arreglo[indxDer]) < 0 {
		mayor = indxDer
	}
	// Swap si el mayor es alguno de los hijos
	if mayor != pos {
		swap(&arreglo[pos], &arreglo[mayor])
		downHeap(mayor, arreglo, cmp)
	}
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
	// Si actual necesita ir antes, swap, sino salir de recursion
	if cmp(actual, padre) > 0 {
		swap(&arreglo[posPadre], &arreglo[pos])
		upHeap(posPadre, arreglo, cmp)
	}
}

// heapify(arreglo, funcion) toma un arreglo y una funcion de comparacion
// y modifica el arreglo en forma de heap segun la funcion pasada
func heapify[T any](arreglo []T, cmp func(T, T) int) {
	for i := len(arreglo) - 1; i >= 0; i-- {
		downHeap(i, arreglo, cmp)
	}
}
