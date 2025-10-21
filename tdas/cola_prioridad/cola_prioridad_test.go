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
	heap := TDAColaPrioridad.CrearHeap[string]()
	require.Equal(t, 0, heap.Cantidad(), "Heap recien creado")
}

// Test heap un elemento (probar todas las primitivas)

// Test volumen

// Test heapify

// Test heapsort enteros
