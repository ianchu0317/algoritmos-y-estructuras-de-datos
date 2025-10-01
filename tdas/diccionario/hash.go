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
// y devuelve la clave en hash dentro del rango pasado.
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func djb2Hash[K any](clave K, largo int) int {
	bytes := convertirABytes(clave)
	hash := 5381
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + int(b) // hash * 33 + b
	}
	return hash % largo
}

// buscarCelda toma una clave del hash y devuelve la celda correspondiente si existe.
// En caso de no existir devuelve nil
func (hash hashAbierto[K, V]) buscarCelda(clave K) *celda[K, V] {
	// Obtener la lista de celdas
	claveHash := djb2Hash(clave, len(hash.arreglo))
	listaHash := hash.arreglo[claveHash]
	// Crear iterador externo de lista e iterar
	iter := listaHash.Iterador()
	for iter.HaySiguiente() {
		celdaActual := iter.VerActual()
		if celdaActual.clave == clave {
			return &celdaActual
		}
		iter.Siguiente()
	}
	return nil
}

// Primitivas de hash abierto

func (hash hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash hashAbierto[K, V]) Guardar(clave K, dato V) {

}
