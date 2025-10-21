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
func TestHeapVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](cmpString)
	require.Equal(t, 0, heap.Cantidad(), "Heap recien creado debe tener cantidad 0")
	require.True(t, heap.EstaVacia(), "Heap recien creado debe estar vacio")
	require.Panics(t, func() { heap.VerMax() }, "Heap recien creado no puede ver elemento de mayor prioridad")
	require.Panics(t, func() { heap.Desencolar() }, "Heap recien creado no puede desencolar")
}

// Test heap un elemento (probar todas las primitivas)

// Test volumen

// Test heapify

// Test heapsort enteros
