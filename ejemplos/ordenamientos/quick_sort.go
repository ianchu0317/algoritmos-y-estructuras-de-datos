package ordenamientos

// Particiona el array y devuelve index de pivot
func particion(arr []int) int {
	// ultimo elemento como pivot
	pivot := arr[len(arr)-1]
	// index del pivot?
	indxPivot := 0

	// Para cada elemento menos el ultimo
	// Si el elemento es menor al pivot,
	// Avanzar contador del pivot
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[i], arr[indxPivot] = arr[indxPivot], arr[i]
			indxPivot++
		}
	}
	// Poner pivot en su indice
	arr[indxPivot], arr[len(arr)-1] = arr[len(arr)-1], arr[indxPivot]
	return indxPivot
}

// Ordena array con algoritmo Quick Sort
func QuickSort(arr []int) {
	// Caso base si solo hay un elemento
	if len(arr) <= 1 {
		return
	}
	// Pivot como primer elemento
	indxPivot := particion(arr)
	QuickSort(arr[:indxPivot])
	QuickSort(arr[indxPivot+1:])
}
