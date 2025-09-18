package ordenamientos

// Complejidad O(n+k), k es rango 0-100
// NO In-place
// Estable

const RANGO = 101

func CountingSort(arr []int) []int {
	// 1) Crear array de frecuencias
	contador := make([]int, RANGO)
	for _, num := range arr {
		contador[num] += 1
	}
	//fmt.Println("Contador:", contador)
	// 2) Crear array de inicios
	inicios := make([]int, RANGO)
	for i := 1; i < RANGO; i++ {
		inicios[i] = inicios[i-1] + contador[i-1]
	}
	//fmt.Println("Inicios:", inicios)
	// 3) Devolver en nuevo array
	nuevoArr := make([]int, len(arr))
	for i := range arr {
		num := arr[i]
		posicion := inicios[num] // Obtener la posicion de inicio
		nuevoArr[posicion] = num // Colocar en su posicion de inicio
		inicios[num]++           // Cambiar su posicion de inicio una vez colocado
	}
	return nuevoArr
}
