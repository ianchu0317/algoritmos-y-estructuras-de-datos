package hash

/*Implementar una función de orden O(n) que dado un arreglo de n números enteros
devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces.
Justificar el orden de la solución. */

func hayNumeroMasMitad(arr []int) bool {
	hash := CrearHash[int, int](func(a, b int) bool {
		return a == b
	})
	// Para cada numero guardo su frecuencia. Si supera len(arr)/2 entonces true, sino false
	// Realizo n veces operaciones de O(1)
	for _, num := range arr {
		frecuencia := 0
		if hash.Pertenece(num) {
			frecuencia = hash.Obtener(num)
		}
		hash.Guardar(num, frecuencia+1)
		if hash.Obtener(num) > len(arr)/2 {
			return true
		}
	}
	return false
}
