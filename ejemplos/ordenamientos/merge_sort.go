package ordenamientos

// Insertar dos arreglos ordenados en una sola
func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	nuevoArr := make([]int, 0)

	// Insertar ordenadamente
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			nuevoArr = append(nuevoArr, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			nuevoArr = append(nuevoArr, arr2[j])
			j++
		} else {
			nuevoArr = append(nuevoArr, arr1[i])
			nuevoArr = append(nuevoArr, arr2[j])
			i++
			j++
		}
	}

	// Insertar los elementos restantes
	for i < len(arr1) {
		nuevoArr = append(nuevoArr, arr1[i])
		i++
	}
	for j < len(arr2) {
		nuevoArr = append(nuevoArr, arr2[j])
		j++
	}

	// Devolver array ordenado
	return nuevoArr
}

// Merge sort ordena el arreglo utilizando algoritmo merge sort
func MergeSort(arr []int) []int {
	// Caso base si hay un solo elemento, está ordenado
	if len(arr) <= 1 {
		return arr
	}

	// Dividir a la mitad y ordenar
	mitad := len(arr) / 2
	izqOrd := MergeSort(arr[:mitad])
	derOrd := MergeSort(arr[mitad:])

	// Mergear las mitades ordenadas
	nuevoArr := merge(izqOrd, derOrd)
	return nuevoArr
}
