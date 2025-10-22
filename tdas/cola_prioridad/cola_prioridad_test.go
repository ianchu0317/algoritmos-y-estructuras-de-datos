package cola_prioridad_test

import (
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

// Funciones de comparacion

func cmpString(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func cmpIntMax(a, b int) int {
	return a - b
}

func cmpIntMin(a, b int) int {
	return b - a
}

// Test heap vacio
// testsHeapVacio corre los tests de heap vacio
func testsHeapVacio[T any](t *testing.T, heap TDAColaPrioridad.ColaPrioridad[T]) {
	require.Equal(t, 0, heap.Cantidad(), "Heap vacio debe tener cantidad 0")
	require.True(t, heap.EstaVacia(), "Heap vacio debe estar vacio")
	require.Panics(t, func() { heap.VerMax() }, "Heap vacio no puede ver elemento de mayor prioridad")
	require.Panics(t, func() { heap.Desencolar() }, "Heap vacio  no puede desencolar")
}

func TestHeapVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](cmpString)
	testsHeapVacio(t, heap)
}

// Test heap un elemento (probar todas las primitivas)
func TestHeapUnElemento(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](cmpString)
	heap.Encolar("A")
	require.Equal(t, 1, heap.Cantidad(), "Heap con un elemento debe tener cantidad 1")
	require.False(t, heap.EstaVacia(), "Heap con un elemento no puede estar vacio")
	require.Equal(t, "A", heap.VerMax(), "Heap con un elemento debe devolver ese unico elemento de mayor prioridad")
	require.Equal(t, "A", heap.Desencolar(), "Heap con un elemento debe desencolar ese unico elemento")
	// Luego de desencolar cumple con heap vacio
	testsHeapVacio(t, heap)
}

// cargarHeapArr(heap, arr) toma un heap y un array y carga los elementos del arreglo al heap
func cargarHeapArr[T any](heap TDAColaPrioridad.ColaPrioridad[T], arr []T) {
	for _, elem := range arr {
		heap.Encolar(elem)
	}
}

// Tests especificos
func TestCasosEspecificos(t *testing.T) {
	//arr := []int{1, 6, 3, 9, 5, 2, 4, 8, 7, 0}
	//arrMax := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//arrMin := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Crear heaps de maximos y minimos y chequear vacios
	heapMax := TDAColaPrioridad.CrearHeap[int](cmpIntMax)
	heapMin := TDAColaPrioridad.CrearHeap[int](cmpIntMin)
	testsHeapVacio(t, heapMax)
	testsHeapVacio(t, heapMin)

}

// Test volumen

// Test heapify

// Test heapsort enteros
