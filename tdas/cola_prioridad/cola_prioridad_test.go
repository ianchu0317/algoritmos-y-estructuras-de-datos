package cola_prioridad_test

import (
	"math/rand"
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

// testOrdenDesencolar(t, arr, heap) toma un arreglo y un heap.
// Itera el arreglo y compara si coincide con orden de desencolado
func testOrdenDesencolar[T any](t *testing.T, arr []T, heap TDAColaPrioridad.ColaPrioridad[T]) {
	for _, elem := range arr {
		require.Equal(t, elem, heap.Desencolar(), "El orden del heap deberia coincidir con el arreglo esperado")
	}
}

// Tests especificos
func TestCasosEspecificos(t *testing.T) {
	arr := []int{1, 6, 3, 9, 5, 2, 4, 8, 7, 0}
	arrMax := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arrMin := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Crear heaps de maximos y minimos y chequear vacios
	heapMax := TDAColaPrioridad.CrearHeap[int](cmpIntMax)
	heapMin := TDAColaPrioridad.CrearHeap[int](cmpIntMin)
	testsHeapVacio(t, heapMax)
	testsHeapVacio(t, heapMin)
	// Cargar heaps de maximos y minimos con arr
	cargarHeapArr(heapMax, arr)
	cargarHeapArr(heapMin, arr)
	// CHequear orden de desencolado
	testOrdenDesencolar(t, arrMax, heapMax)
	testOrdenDesencolar(t, arrMin, heapMin)
	// chequear nuevamente heap vacio
	testsHeapVacio(t, heapMax)
	testsHeapVacio(t, heapMin)
}

// Tests para heaps creados desde el arreglo
func TestCrearHeapArr(t *testing.T) {
	arr := []int{1, 6, 3, 9, 5, 2, 4, 8, 7, 0}
	arrMax := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arrMin := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Crear heaps de maximos y minimos y chequear vacios
	heapMax := TDAColaPrioridad.CrearHeapArr[int](arr, cmpIntMax)
	heapMin := TDAColaPrioridad.CrearHeapArr[int](arr, cmpIntMin)
	// CHequear orden de desencolado
	testOrdenDesencolar(t, arrMax, heapMax)
	testOrdenDesencolar(t, arrMin, heapMin)
	// chequear nuevamente heap vacio
	testsHeapVacio(t, heapMax)
	testsHeapVacio(t, heapMin)
}

// crearArregloDesordenado toma cantidad y rango en enteros positivos.
// Devuelve un arreglo de largo 'cantidad' desordenado de numeros aleatorios dentro del rango.
func crearArregloDesordenado(cantidad, rango int) []int {
	nuevoArr := make([]int, cantidad)
	for i := range nuevoArr {
		numRandom := rand.Intn(rango)
		nuevoArr[i] = numRandom
	}
	return nuevoArr
}

// countinSort ordena un arreglo de numeros enteros dentro del rango
func countingSort(arr []int, rango int) []int {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[num]++
	}
	inicios := make([]int, rango)
	for i := 1; i < rango; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}
	ordenado := make([]int, len(arr))
	for _, num := range arr {
		indxInicio := inicios[num]
		ordenado[indxInicio] = num
		inicios[num]++
	}
	return ordenado
}

// testDosArreglosIguales chequea si dos arreglos son iguales
func testDosArreglosIguales[T any](t *testing.T, arr1, arr2 []T) {
	require.Equal(t, len(arr1), len(arr2), "El largo de los dos arreglos debe ser igual")
	for i := 0; i < len(arr1); i++ {
		require.Equal(t, arr1[i], arr2[i], "El elemento en la misma posicion de los dos arreglos debe ser igual")
	}
}

// Test heapsort enteros
func TestHeapSort(t *testing.T) {
	arr1 := crearArregloDesordenado(10, 100)
	ordenado1 := countingSort(arr1, 100)
	TDAColaPrioridad.HeapSort(arr1, cmpIntMax)
	testDosArreglosIguales(t, arr1, ordenado1)
}

// Test volumen
