package diccionario

import "fmt"

// *** Estructura de ABB ***

type nodo[K any, V any] struct {
	clave K
	dato  V
	izq   *nodo[K, V]
	der   *nodo[K, V]
}

type arbolBinario[K any, V any] struct {
	raiz     *nodo[K, V]
	comparar func(K, K) int
}

// Funciones auxiliares

// crearNodo crea un nodo para el abb
func crearNodo[K any, V any](clave K, dato V) *nodo[K, V] {
	return &nodo[K, V]{clave: clave, dato: dato, izq: nil, der: nil}
}

// convertirBytes toma la clave y la convierte a bytes
func convertirBytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2HashInt toma una clave y devuelve un entero unico con respecto a la clave.
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func djb2HashInt[K any](clave K) int {
	bytes := convertirBytes(clave)
	hash := uint32(5381)
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + uint32(b) // hash*33 + b
	}
	return int(hash)
}

// Primitivas de ABB
