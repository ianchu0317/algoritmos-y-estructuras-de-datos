package ordenamientos

// Complejidad O(n+k), k es rango 0-100
// NO In-place
// Estable

func CountingSort(arr []int) []int {
	// Crear array de frecuencias
	contador := make([]int, 101)
	for _, num := range arr {
		contador[num] += 1
	}
	// Crear array de inicios
	inicios := make([]int, 101)
	for i := range inicios {
		inicios[i]
	}
}
