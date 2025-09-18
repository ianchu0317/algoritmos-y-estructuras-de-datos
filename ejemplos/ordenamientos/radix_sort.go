package ordenamientos

import "fmt"

// Los numeros de 0-100 tienen 2 dígitos
// Se puede ordenar con Radix Sort
// 	- Primero se ordena por unidad
// 	- Segundo se ordena por centenas (dígito más significativo)

const RANGO_N = 10

// Counting Sort como función auxiliar
// Ordena según unidad, decena, centena
func CountingSortN(arr []int, n int) []int {
	// 1) Crear array de frecuencias
	contador := make([]int, RANGO_N)
	for _, num := range arr {
		contador[(num/n)%10] += 1
	}
	// 2) Crear array de inicios
	inicios := make([]int, RANGO_N)
	for i := 1; i < RANGO_N; i++ {
		inicios[i] = inicios[i-1] + contador[i-1]
	}
	// 3) Devolver en nuevo array
	nuevoArr := make([]int, len(arr))
	for _, num := range arr {
		digito := (num / n) % 10
		posicion := inicios[digito] // Obtener la posicion de inicio
		nuevoArr[posicion] = num    // Colocar en su posicion de inicio
		inicios[digito]++           // Cambiar su posicion de inicio una vez colocado
	}
	return nuevoArr
}

// Ordenar arreglo de enteros por algoritmo RadixSort
func RadixSort(arr []int) []int {
	// Ordenar primero por unidad y luego por decena
	arr1 := CountingSortN(arr, 1)
	fmt.Println("Count Sort unidad:", arr1)
	arr2 := CountingSortN(arr1, 10)
	fmt.Println("Count Sort decena:", arr2)
	return arr2
}
