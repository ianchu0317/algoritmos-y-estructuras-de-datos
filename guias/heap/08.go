package heap

/*
Escribir una función en Go que, dado un arreglo de `n` cadenas y un entero positivo `k` cadenas más largas. Se
espera que el orden del algoritmo sea O(n+klogn). Justificar el orden.
*/
func kCadenasLargas(k int, cadenas []string) []string {
	// Básicamente un heapify, y pasar funcion de comparacion O(n)
	MaxHeap := CrearHeapArr[string](cadenas, func(a, b string) int {
		return len(a) - len(b)
	})
	// Operaciones declaracion de variabels O(1)
	resultado := make([]string, min(len(cadenas), k))
	contador := 0
	// si k<n entonces hacemos k*log(n) -> desencolar heap k veces
	for contador < k && !MaxHeap.EstaVacia() {
		resultado[contador] = MaxHeap.Desencolar()
		contador++
	}
	return resultado
}
