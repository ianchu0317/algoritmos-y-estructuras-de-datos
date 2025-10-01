package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

// *** Estructura de hash abierto ***

type celda[K any, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K any, V any] struct {
	arreglo  []TDALista.Lista[celda[K, V]]
	cantidad int
}

// Funciones auxiliares

// convertirABytes toma la clave y la convierte a bytes
func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2Hash toma una clave y el numero maximo que puede valer,
// y devuelve la clave en hash dentro del rango
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func Djb2Hash[K any](clave K, largo int) int {
	bytes := convertirABytes(clave)
	hash := 5381
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + int(b) // hash * 33 + b
	}
	return hash % largo
}

// Primitivas de hash abierto
