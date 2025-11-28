package hash

/*
Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un número K y determinar en
O(n) si existe un par de elementos en el arreglo que sumen exactamente K.
*/
func existeSuma(arr []int, k int) bool {
	// Idea iterar cada arreglo, si su complemento esta, entonces existe
	hash := CrearHash[int, int]()
	for _, num := range arr {
		complemento := k - num
		if hash.Pertenece(complemento) {
			return true
		}
		hash.Guardar(num, num)
	}
	// Si llega aca entonces no existe
	return false
}

// COmplejidad O(n). Por cada elemento realizo operaciones con hash de O(1). N * O(1) = O(n)
